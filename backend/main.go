package main

import "github.com/lucabockmann/go_boilerplate/api"

func main() {
	api.ConnectDatabase()
	api.CreateRoutes()
}
