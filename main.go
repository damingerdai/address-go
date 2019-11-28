package main

import "damingerdai/address/routes"

func main() {
	r := routes.NewRouter()
	r.Run(":9999")
}
