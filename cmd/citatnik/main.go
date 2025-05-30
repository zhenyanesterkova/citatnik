package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/zhenyanesterkova/citatnik/internal/handler"
	"github.com/zhenyanesterkova/citatnik/internal/storage"
)

func main() {
	run()
}

func run() {
	store := storage.New()

	RHandler := handler.NewRepositorieHandler(store)
	router := mux.NewRouter()
	RHandler.InitRouter(router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("Commencing server shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown failed:", err)
	}

	wg.Wait()
	log.Println("Server was gracefully shut down.")
}
