// объявление структуры сервера для использования в http
package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1 Mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe() //под капотом запускает бесконечный цикл for
}

func (s *Server) Shutdown(ctx context.Context) error { //будем использовать при выходе из приложения
	return s.httpServer.Shutdown(ctx)
}
