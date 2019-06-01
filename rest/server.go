package rest

import (
	"context"
	"net/http"

	//nolint:golint
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/studtool/common/logs"
)

// Server is expected to be embedded
type Server struct {
	server *http.Server

	structLogger  logs.Logger
	reflectLogger logs.Logger
	requestLogger logs.Logger

	apiClassifier APIClassifier
}

type ServerParams struct {
	Address string
	Handler http.Handler

	StructLogger  logs.Logger
	ReflectLogger logs.Logger
	RequestLogger logs.Logger

	APIClassifier APIClassifier
}

func NewServer(params ServerParams) *Server {
	return &Server{
		server: &http.Server{
			Addr:    params.Address,
			Handler: params.Handler,
		},

		structLogger:  params.StructLogger,
		reflectLogger: params.ReflectLogger,
		requestLogger: params.RequestLogger,

		apiClassifier: params.APIClassifier,
	}
}

func (srv *Server) MetricsHandler() http.Handler {
	return promhttp.Handler()
}

func (srv *Server) Run() error {
	srv.structLogger.Infof("started on %s", srv.server.Addr)
	go func() {
		if err := srv.server.ListenAndServe(); err != nil {
			srv.structLogger.Fatal(err)
		}
	}()
	return nil
}

func (srv *Server) Shutdown() error {
	srv.structLogger.Infof("stopped")
	return srv.server.Shutdown(context.TODO())
}
