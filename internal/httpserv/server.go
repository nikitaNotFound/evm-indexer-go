package httpserv

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/nikitaNotFound/evm-indexer-go/internal/apigen"
	"github.com/nikitaNotFound/evm-indexer-go/internal/storages/postgres"
	"github.com/rs/zerolog/log"
)

type HTTPServer struct {
	tcpAddress string
	app        *echo.Echo
	storage    *postgres.Storage
}

// NewHTTPServer creates a new AppAPI instance. Accepts a TCP address and storage instance.
func NewHTTPServer(tcpAddress string, storage *postgres.Storage) *HTTPServer {
	e := echo.New()

	e.Use(echomiddleware.Recover())

	// CORS middleware
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{"*"},
	}))

	e.Use(echomiddleware.Logger())
	RegisterSwagger(e.Group(""))

	app := &HTTPServer{
		tcpAddress: tcpAddress,
		app:        e,
		storage:    storage,
	}

	apigen.RegisterHandlers(e, app)

	return app
}

// start-stop
func (a *HTTPServer) Listen() error {
	if a.tcpAddress == "" {
		return errors.New("endpoint is not set")
	}
	log.Info().Msgf("Starting server... on %s", a.tcpAddress)

	err := a.app.Start(a.tcpAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Server starting error")
	}
	return nil
}

func (a *HTTPServer) Shutdown() error {
	log.Info().Msg("Shutting down server...")
	// Use the non-context Close method for backward compatibility
	return a.app.Close()
}

// ShutdownWithContext performs a graceful shutdown with the provided context
func (a *HTTPServer) ShutdownWithContext(ctx context.Context) error {
	log.Info().Msg("Gracefully shutting down server with timeout...")
	return a.app.Shutdown(ctx)
}
