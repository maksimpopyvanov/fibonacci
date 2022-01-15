package fibonacci

import (
	"net/http"
	"time"
)

type HTTPServer struct {
	httpServer *http.Server
}

func (s *HTTPServer) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		WriteTimeout:   20 * time.Second,
		ReadTimeout:    10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
