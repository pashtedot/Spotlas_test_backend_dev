package internal

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/misnaged/annales/logger"
	version "github.com/misnaged/annales/versioner"

	pb "close-spots-service/protocols/spots"

	"close-spots-service/config"
	"close-spots-service/internal/metrics"
	"close-spots-service/internal/repository"
	"close-spots-service/internal/server"
	"close-spots-service/internal/service"
)

// App is main microservice application instance that
// have all necessary dependencies inside structure
type App struct {
	// application configuration
	config *config.Scheme

	version *version.Version

	connection *sql.DB
	metric     metrics.Metrics
	srv        server.ISpotsService
	grpc       server.IGrpc
}

// NewApplication create new App instance
func NewApplication() (app *App, err error) {
	ver, err := version.NewVersion()
	if err != nil {
		return nil, fmt.Errorf("init app version: %w", err)
	}

	return &App{
		config:  &config.Scheme{},
		version: ver,
	}, nil
}

// Init initialize application and all necessary instances
func (app *App) Init() error {
	app.initMetrics()

	if err := app.initDb(app.config.Db); err != nil {
		return fmt.Errorf("application database initialisation: %w", err)
	}

	if err := app.initService(); err != nil {
		return fmt.Errorf("initi service instance: %w", err)
	}

	if err := app.initGrpc(app.config.Grpc); err != nil {
		return fmt.Errorf("application grpc server initialisation: %w", err)
	}

	return nil
}

func (app *App) initMetrics() {
	app.metric = metrics.NewPrometheus()
}

// initDb initialize Application database instance
func (app *App) initDb(cfg *config.Db) (err error) {
	logger.Log().Infof("Connecting to Postgresql database %s", cfg.Address)

	db, err := sql.Open("postgres", cfg.Address)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	app.connection = db

	logger.Log().Infof("Connected to Postgresql database %s", cfg.Address)

	return nil
}

// initService initialize Application service layer
func (app *App) initService() (err error) {
	app.srv, err = service.NewSpotsService(
		app.metric,
		repository.NewRepository(app.connection, app.metric),
	)
	if err != nil {
		return fmt.Errorf("initialize application service layer instance: %w", err)
	}

	return nil
}

// initGrpc initialize Application GRPC server instance
func (app *App) initGrpc(cfg *config.Grpc) (err error) {
	logger.Log().Info("App GRPC server initializing...")

	app.grpc, err = server.NewServer(cfg)
	if err != nil {
		return fmt.Errorf("create new GRPC server instance: %w", err)
	}

	api := server.NewHandlers(app.srv)

	pb.RegisterCloseSpotsServiceServer(app.grpc.GetServer(), api)

	logger.Log().Info("App GRPC server initialized")
	return nil
}

// Serve start serving Application service
func (app *App) Serve() error {
	go app.grpc.Listen()

	// Gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit

	return nil
}

// Stop shutdown the application
func (app *App) Stop() error {
	if app.grpc != nil {
		app.grpc.Close()
	}
	if app.srv != nil {
		app.srv.Close()
	}
	var err error
	if app.connection != nil {
		err = app.connection.Close()
	}

	return err
}

// Config return App config Scheme
func (app *App) Config() *config.Scheme {
	return app.config
}

// Version return application current version
func (app *App) Version() string {
	return app.version.String()
}
