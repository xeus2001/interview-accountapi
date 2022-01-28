package main

import (
	"fmt"
	"github.com/xeus2001/interview-accountapi/pkg/f3"
)

func main() {
	fmt.Printf("Try health-check at endpoint: %s\n", f3.DefaultEndPoint)
	client := f3.NewClient()
	println(client.IsHealthy())
}
