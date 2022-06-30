package app

import (
	"flag"
	"fmt"
	"github.com/damingerdai/address-go/pkg/routes"
	"log"
	"os"

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
	fmt.Println(address)
	r := routes.NewRouter()
	err = r.Run(address)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
	} else {
		fmt.Fprintln(os.Stdout, "run successfully")
	}
}
