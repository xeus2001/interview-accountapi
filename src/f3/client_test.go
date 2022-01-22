package f3_test

import (
	"github.com/xeus2001/interview-accountapi/src/f3"
	"testing"
	"time"
)

func TestHealthy(t *testing.T) {
	client, err := f3.NewClient("localhost", 8080, time.Second*5)
	if err != nil {
		t.Fatalf("Creating the F3 client failed, reason: %s", err.Error())
	}
	if !client.IsHealthy() {
		t.Fatalf("Health check for service failed")
	}
}
