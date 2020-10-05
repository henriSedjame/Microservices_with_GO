package main

import (
	"context"
	"github.com/gorilla/mux"
	. "github.com/hsedjame/products-api/src/api"
	"github.com/hsedjame/products-api/src/repository/impl"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	controllers []RestController
}

func (app App) start() {

	router := mux.NewRouter()

	for _, controller := range app.controllers {
		subRouter := router.PathPrefix(controller.Path()).Subrouter()
		controller.AddRoutes(subRouter)
	}

	server := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: time.Second,
		WriteTimeout:      time.Second,
	}

	runAndStopServerGracefully(server)
}

func runAndStopServerGracefully(server *http.Server) {
	logger := log.New(os.Stdout, "[[products-api]] ", log.LstdFlags)

	go func() {
		logger.Fatal(server.ListenAndServe())
	}()

	// Créer un chanel qui reçoit des signaux de type os.Signal
	signalChannel := make(chan os.Signal)

	// Envoyer un signal au channel lors :
	// * d'une interruption de la machine
	// * d'un arrêt de la machine
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)

	// Récupérer le signal envoyé

	interruptOrKillSignal := <-signalChannel

	logger.Println("Interruption du server ==> ", interruptOrKillSignal)

	deadline, _ := context.WithTimeout(context.Background(), 30*time.Second)

	logger.Fatal(server.Shutdown(deadline))
}

func main() {
	productRepository := impl.NewSimpleProductRepository(log.New(os.Stdout, "[[products-api]]", log.LstdFlags))
	productHandler := NewProductHandler(productRepository)

	App{controllers: []RestController{
		NewProductController(*productHandler),
	}}.start()
}
