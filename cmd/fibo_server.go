package main

import (
	"encoding/json"
	"fmt"
	"github.com/ali-mohit/simple-fibonacci-server/internal/cache_fibo"
	"github.com/ali-mohit/simple-fibonacci-server/internal/process_fibo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/vmihailenco/msgpack/v5"
	"log"
	"net/http"
)

func getServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "start server",
		Run: func(cmd *cobra.Command, args []string) {
			httpServer := New(cmd, args)
			if err := httpServer.serve(); err != nil {
				logrus.WithError(err).Fatal("Failed to serve.")
			}
		},
	}
}

type Server interface {
	serve() error
}

type impl struct {
	inMemoryCache    cache_fibo.InMemoryCache
	fibonacciHandler process_fibo.FibonacciProcessHandler
}

func New(cmd *cobra.Command, args []string) Server {
	inMemoryCache := cache_fibo.New()
	fiboProcess := process_fibo.New(inMemoryCache)

	return &impl{
		inMemoryCache:    inMemoryCache,
		fibonacciHandler: fiboProcess,
	}
}

func (s *impl) serve() error {

	http.HandleFunc("/fibonacci", s.handleFibonacci)
	http.HandleFunc("/fibonacci-v2", s.handleFibonacciV2)

	fmt.Println("Server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (s *impl) handleFibonacci(w http.ResponseWriter, r *http.Request) {
	var req process_fibo.FibonacciRequest

	// Determine the request content type
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "application/x-msgpack":
		if err := msgpack.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Unsupported content type", http.StatusUnsupportedMediaType)
		return
	}

	resp, err := s.fibonacciHandler.ProcessFibonacciNumber(&req)
	if err != nil {
		log.Printf("Error ProcessFibonacciNumber: %v", err)
		http.Error(w, "Unknown Error ProcessFibonacciNumber", http.StatusInternalServerError)
		return
	}

	// Determine the response content type
	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	case "application/x-msgpack":
		w.Header().Set("Content-Type", "application/x-msgpack")
		if err := msgpack.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	default:
		http.Error(w, "Unsupported accept type", http.StatusNotAcceptable)
	}
}

func (s *impl) handleFibonacciV2(w http.ResponseWriter, r *http.Request) {
	var req process_fibo.FibonacciRequestV2

	// Determine the request content type
	switch r.Header.Get("Content-Type") {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "application/x-msgpack":
		if err := msgpack.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "Unsupported content type", http.StatusUnsupportedMediaType)
		return
	}

	resp, err := s.fibonacciHandler.ProcessFibonacciNumberV2(&req)
	if err != nil {
		log.Printf("Error ProcessFibonacciNumber: %v", err)
		http.Error(w, "Unknown Error ProcessFibonacciNumber", http.StatusInternalServerError)
		return
	}

	// Determine the response content type
	switch r.Header.Get("Accept") {
	case "application/json":
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	case "application/x-msgpack":
		w.Header().Set("Content-Type", "application/x-msgpack")
		if err := msgpack.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Error encoding response: %v", err)
		}
	default:
		http.Error(w, "Unsupported accept type", http.StatusNotAcceptable)
	}
}
