package main

import (
	"fmt"

	"github.com/viandikafauzi/rmq-ses/rmq"
)

func main() {
	fmt.Println("Server is starting...")

	rmq.Run()
}
