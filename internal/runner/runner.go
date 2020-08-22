package runner

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/waypoint/internal/factory"
	"github.com/hashicorp/waypoint/internal/plugin"
	"github.com/hashicorp/waypoint/internal/server"
	pb "github.com/hashicorp/waypoint/internal/server/gen"
	"github.com/hashicorp/waypoint/sdk/component"
	"github.com/hashicorp/waypoint/sdk/terminal"
)

var ErrClosed = errors.New("runner is closed")

// Runners in Waypoint execute operations. These can be local (the CLI)
// or they can be remote (triggered by some webhook). In either case, they
// share this same underlying implementation.
//
// To use a runner:
//
//   1. Initialize it with New. This will setup some initial state but
//      will not register with the server or run jobs.
//
//   2. Start the runner with "Start". This will register the runner and
//      kick off some management goroutines. This will not execute any jobs.
//
//   3. Run a single job with "Accept". This is named to be similar to a
//      network listener "accepting" a connection. This will request a single
//      job from the Waypoint server, block until one is available, and execute
//      it. Repeat this call for however many jobs you want to execute.
//
//   4. Clean up with "Close". This will gracefully exit the runner, waiting
//      for any running jobs to finish.
//
type Runner struct {
	id          string
	logger      hclog.Logger
	client      pb.WaypointClient
	ctx         context.Context
	cleanupFunc func()
	runner      *pb.Runner
	factories   map[component.Type]*factory.Factory
	ui          terminal.UI
	local       bool
	tempDir     string

	closedVal int32
	acceptWg  sync.WaitGroup

	// config is the current runner config.
	config      *pb.RunnerConfig
	originalEnv []*pb.ConfigVar

	// noopCh is used in tests only. This will cause any noop operations
	// to block until this channel is closed.
	noopCh <-chan struct{}
}

// New initializes a new runner.
//
// You must call Start to start the runner and register with the Waypoint
// server. See the Runner struct docs for more details.
func New(opts ...Option) (*Runner, error) {
	// Create our ID
	id, err := server.Id()
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			"failed to generate unique ID: %s", err)
	}

	// Our default runner
	runner := &Runner{
		id:     id,
		logger: hclog.L(),
		ctx:    context.Background(),
		runner: &pb.Runner{Id: id},
		factories: map[component.Type]*factory.Factory{
			component.BuilderType:        plugin.Builders,
			component.RegistryType:       plugin.Registries,
			component.PlatformType:       plugin.Platforms,
			component.ReleaseManagerType: plugin.Releasers,
		},
	}

	// Build our config
	var cfg config
	for _, o := range opts {
		err := o(runner, &cfg)
		if err != nil {
			return nil, err
		}
	}

	// Setup our runner components list
	for t, f := range runner.factories {
		for _, n := range f.Registered() {
			runner.runner.Components = append(runner.runner.Components, &pb.Component{
				Type: pb.Component_Type(t),
				Name: n,
			})
		}
	}

	return runner, nil
}

// Id returns the runner ID.
func (r *Runner) Id() string {
	return r.id
}

// Start starts the runner by registering the runner with the Waypoint
// server. This will spawn goroutines for management. This will return after
// registration so this should not be executed in a goroutine.
func (r *Runner) Start() error {
	if r.closed() {
		return ErrClosed
	}

	log := r.logger

	// Register
	log.Debug("registering runner")
	client, err := r.client.RunnerConfig(r.ctx)
	if err != nil {
		return err
	}
	r.cleanup(func() { client.CloseSend() })

	// Send request
	if err := client.Send(&pb.RunnerConfigRequest{
		Event: &pb.RunnerConfigRequest_Open_{
			Open: &pb.RunnerConfigRequest_Open{
				Runner: r.runner,
			},
		},
	}); err != nil {
		return err
	}

	// Wait for an initial config as confirmation we're registered.
	log.Trace("runner connected, waiting for initial config")
	resp, err := client.Recv()
	if err != nil {
		return err
	}

	// Handle the first config so our initial setup is done
	r.handleConfig(resp.Config)

	// Start the watcher
	ch := make(chan *pb.RunnerConfig)
	go r.watchConfig(ch)

	// Start the goroutine that waits for all other configs
	go r.recvConfig(r.ctx, client, ch)

	log.Info("runner registered with server")
	return nil
}

// Close gracefully exits the runner. This will wait for any pending
// job executions to complete and then deregister the runner. After
// this is called, Start and Accept will no longer function and will
// return errors immediately.
func (r *Runner) Close() error {
	// If we can't swap, we're already closed.
	if !atomic.CompareAndSwapInt32(&r.closedVal, 0, 1) {
		return nil
	}

	// Wait for our jobs to complete
	r.acceptWg.Wait()

	// Run any cleanup necessary
	if f := r.cleanupFunc; f != nil {
		f()
	}

	return nil
}

func (r *Runner) closed() bool {
	return atomic.LoadInt32(&r.closedVal) > 0
}

type config struct{}

type Option func(*Runner, *config) error

// WithClient sets the client directly. In this case, the runner won't
// attempt any connection at all regardless of other configuration (env
// vars or waypoint config file). This will be used.
func WithClient(client pb.WaypointClient) Option {
	return func(r *Runner, cfg *config) error {
		r.client = client
		return nil
	}
}

// WithComponentFactory sets a factory for a component type. If this isn't set for
// a component type, then the builtins will be used.
func WithComponentFactory(t component.Type, f *factory.Factory) Option {
	return func(r *Runner, cfg *config) error {
		r.factories[t] = f
		return nil
	}
}

// WithLogger sets the logger that the runner will use. If this isn't
// set it uses hclog.L().
func WithLogger(logger hclog.Logger) Option {
	return func(r *Runner, cfg *config) error {
		r.logger = logger
		return nil
	}
}

// WithLocal sets the runner to local mode. This only changes the UI
// behavior to use the given UI. If ui is nil then the normal streamed
// UI will be used.
func WithLocal(ui terminal.UI) Option {
	return func(r *Runner, cfg *config) error {
		r.local = true
		r.ui = ui
		return nil
	}
}

// ByIdOnly sets it so that only jobs that target this runner by specific
// ID may be assigned.
func ByIdOnly() Option {
	return func(r *Runner, cfg *config) error {
		r.runner.ByIdOnly = true
		return nil
	}
}
