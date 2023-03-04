package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/endyApina/exercise-admin-computer/config"
	"github.com/endyApina/exercise-admin-computer/db/postgres"
	httpServer "github.com/endyApina/exercise-admin-computer/server/http"
)

func main() {
	fmt.Println("starting records service...")
	//get dependencies
	fmt.Println("getting dependencies")
	config, err := config.LoadSecrets(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	//get data store
	store, err := postgres.New(&config)
	if err != nil {
		log.Fatal("error connecting to database")
	}

	httpRouter := httpServer.MountServer(store)

	interruptHandler := make(chan os.Signal, 1)
	signal.Notify(interruptHandler, syscall.SIGTERM, syscall.SIGINT)

	//start httpServer
	go func() {
		httpAddr := fmt.Sprintf(":%s", config.ServicePort)
		log.Println("http service running on", httpAddr)
		if err := http.ListenAndServe(httpAddr, httpRouter); err != nil {
			log.Println("failed to start http server")
		}
	}()

	<-interruptHandler
}
