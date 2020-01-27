package app

import "damingerdai/address/pkg/routes"

import "fmt"

import "os"

const address = "localhost:9999"

func Run() {
	r := routes.NewRouter()
	err := r.Run(address)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %v\n", err)
	} else {
		fmt.Fprintln(os.Stdout, "run successfully")
	}
}
