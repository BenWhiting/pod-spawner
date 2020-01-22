package apiserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type (
	// Server object
	Server struct {
		log
		router *mux.Router
		pods   podCollection
	}

	// Config for server
	Config struct {
		Port int
	}

	log interface {
		Trace(...interface{})
		Debug(...interface{})
		Info(...interface{})
		Warn(...interface{})
		Error(...interface{})
		Fatal(...interface{})
		Panic(...interface{})
	}

	podCollection interface {
		Get() string
		Add() string
		Remove() string
		Update() string
	}
)

// Generate server config
func Generate() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT environment required")
	}
	portI, err := strconv.Atoi(port)
	if err != nil {
		return nil, err
	}

	return &Config{
		Port: portI,
	}, nil
}

// New server instance
func New(log log, podProvider podCollection) (*Server, error) {
	return &Server{log,
		mux.NewRouter(),
		podProvider,
	}, nil
}

// Start the server
func (s Server) Start(c *Config) {
	// load routes
	s.Routes()

	// Set up server
	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", c.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.router,
	}

	// Start the server on a nonblocking call
	go func() {
		s.Info(fmt.Sprintf("Listening on %s...\n", srv.Addr))
		if err := srv.ListenAndServe(); err != nil {
			s.Fatal(err)
		}
	}()

	// setup a channel to listen for os signals on
	sigChan := make(chan os.Signal, 1)

	// Signal on interrupt signal
	signal.Notify(sigChan, os.Interrupt)

	// Block until we receive our signal.
	<-sigChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := srv.Shutdown(ctx)
	if err != nil {
		s.Error("Error while shutting down server:%v", err)
	}

	s.Info("shutting down")
}
