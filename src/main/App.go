package main

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	. "github.com/hsedjame/products-api/src/main/api"
	. "github.com/hsedjame/products-api/src/main/repository/impl"
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
		subRouter.Use(controller.Middleware)
		controller.AddRoutes(subRouter)
	}

	/* Configure route to Documentation */
	redocOpts := middleware.RedocOpts{SpecURL: "../../swagger.yaml"}
	redoc := middleware.Redoc(redocOpts, nil)
	router.Handle("/docs", redoc).Methods(http.MethodGet)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

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

	productRepository := NewSimpleProductRepository(log.New(os.Stdout, "[[products-api]]", log.LstdFlags))
	productHandler := NewProductHandler(productRepository)

	App{controllers: []RestController{
		NewProductController(*productHandler),
	}}.start()
}
