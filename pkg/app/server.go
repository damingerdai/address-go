package app

import (
	"damingerdai/address/pkg/routes"
	"flag"
	"fmt"
	"os"
)

var port = flag.Int("port", 9999, "port")

func Run() {
	flag.Parse()
	address := fmt.Sprintf(":%d", *port)
	fmt.Println(address)
	r := routes.NewRouter()
	err := r.Run(address)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
	} else {
		fmt.Fprintln(os.Stdout, "run successfully")
	}
}
