package main

import "github.com/xeus2001/interview-accountapi/pkg/f3"

func main() {
	client := f3.NewClient().WithEndPoint("http://localhost:8080/v1")
	println(client.IsHealthy())
}
