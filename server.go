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
		Handler:        handler, //передаем все маршруты
		MaxHeaderBytes: 1 << 20, //1 Mb - лимит заголовков
		ReadTimeout:    10 * time.Second, //время ожидания запроса
		WriteTimeout:   10 * time.Second, //время ожидания ответа
	}
	return s.httpServer.ListenAndServe() //под капотом запускает бесконечный цикл for
}

func (s *Server) Shutdown(ctx context.Context) error { //для остановки сервера
	return s.httpServer.Shutdown(ctx)
}
