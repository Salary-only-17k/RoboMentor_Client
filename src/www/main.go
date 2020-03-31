package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"www/framework"
	"www/framework/router"
)

func init() {

	var AppID = flag.String("a", "", "AppID Error")
	var AppSecret = flag.String("s", "", "AppSecret Error")

	flag.Parse()

	if  *AppID == "" || *AppSecret == "" {
		log.Println("\033[31m[Error]\033[0m", "RoboMentorClient AppID AppSecret Error")
	}

	RoboMentor.Init(*AppID, *AppSecret)
}

func main() {

	log.Println("[info]", "RoboMentorClient Start")

	router := Router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8888),
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {}
	}()

	log.Println("[info]", "RoboMentorClient Start Success")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {}

	log.Println("[info]", "RoboMentorClient Stop")
}
