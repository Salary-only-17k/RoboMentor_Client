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
	"www/framework/config"
	"www/framework/service/message"
)

func init() {

	var AppID = flag.String("a", "", "AppID Error")
	var AppSecret = flag.String("s", "", "AppSecret Error")

	flag.Parse()

	if  *AppID == "" || *AppSecret == "" {
		log.Println("\033[31m[Error]\033[0m", "RoboMentorSDK AppID AppSecret Error")
	}

	Config.Init(*AppID, *AppSecret)

	MessageService.Messages(*AppID, *AppSecret)

	RoboMentor.Init()
}

func main() {

	log.Println("[info]", "RoboMentorSDK Start")

	router := RoboMentor.InitRouter()

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

	log.Println("[info]", "RoboMentorSDK Start Success")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)

	defer cancel()

	if err := s.Shutdown(ctx); err != nil {}

	log.Println("[info]", "RoboMentorSDK Stop")
}
