package main

import (
	"github.com/morelmiles/go-redis-caching/internals/routes"
	"github.com/morelmiles/go-redis-caching/pkg/database"
)

func main() {
	database.Config()
	routes.Routes()
}
