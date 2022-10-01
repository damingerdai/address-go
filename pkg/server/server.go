package server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/damingerdai/address-go/pkg/routes"

	"github.com/joho/godotenv"
)

var port = flag.Int("port", 9999, "port")

func Run() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	flag.Parse()
	address := fmt.Sprintf(":%d", *port)

	r := routes.NewRouter()
	s := &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	go func() {
		err = s.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stdout, "error: %v\n", err)
		} else {
			fmt.Fprintln(os.Stdout, "run successfully")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Fprintln(os.Stdin, "shuting the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		fmt.Fprintln(os.Stderr, "Server force to shutdown:", err)
	}
	fmt.Fprintln(os.Stdin, "Server is existing")
}
