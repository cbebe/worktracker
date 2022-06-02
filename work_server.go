package worktracker

import (
	"fmt"
	"net/http"
	"os"
)

type WorkServer struct {
	service WorkService
	http.Handler
}

func NewWorkServer(store WorkStore) *WorkServer {
	s := new(WorkServer)

	s.service = WorkService{store}

	router := http.NewServeMux()
	router.Handle("/all", http.HandlerFunc(s.getWorkHandler))
	router.Handle("/start", http.HandlerFunc(s.startWorkHandler))
	router.Handle("/stop", http.HandlerFunc(s.stopWorkHandler))
	s.Handler = router

	return s
}

func handleError(w http.ResponseWriter, err error) {
	fmt.Fprint(os.Stdout, err)
	w.WriteHeader(http.StatusInternalServerError)
}

func (s *WorkServer) startWorkHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.service.StartWork(); err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *WorkServer) getWorkHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	works, err := s.service.GetWork()
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	PrintWork(w, works)
}

func (s *WorkServer) stopWorkHandler(w http.ResponseWriter, r *http.Request) {
	if err := s.service.StopWork(); err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
