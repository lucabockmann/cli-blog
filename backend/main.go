package main

import "github.com/lucabockmann/cliblog/api"

func main() {
	api.ConnectDatabase()
	api.CreateRoutes()
}
