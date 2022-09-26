package server

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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
	err = s.ListenAndServe()
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
	} else {
		fmt.Fprintln(os.Stdout, "run successfully")
	}
}
