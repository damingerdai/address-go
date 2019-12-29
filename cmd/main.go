package main

import "damingerdai/address/routes"

import "os"

import "fmt"

var address string

func init() {
	ip := os.Getenv("Ip")
	port := os.Getenv("Port")

	if len(port) == 0 {
		port = "9999"
	}
	address = fmt.Sprintf("%s:%s", ip, port)

	if len(address) == 0 {
		panic("fail to init app")
	}
}

func main() {
	r := routes.NewRouter()
	r.Run(address)
}
