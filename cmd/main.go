package main

import (
	"waypoint/pkg/web"
)

func main() {
	e := web.SetupServer()
	e.Logger.Fatal(e.Start(":12345"))
}
