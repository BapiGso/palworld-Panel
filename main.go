package main

import (
	"palworld-panel/core"
)

func main() {
	p := core.New()

	p.LoadMiddlewareRoutes()
	p.Listen(":8080")
}
