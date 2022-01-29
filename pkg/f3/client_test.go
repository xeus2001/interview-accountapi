package f3_test

import (
	"github.com/xeus2001/interview-accountapi/pkg/f3"
	"testing"
)

//
// No tests here, because except for NewClient all other methods require integration tests or mock-ups and
// I tend to agree to the author of this:
//
// https://medium.com/@thrawn01/why-you-should-never-test-private-methods-f822358e010
//

func TestNewClient(t *testing.T) {
	client := f3.NewClient()
	httpClient := client.HttpClient()
	if httpClient.Transport != f3.DefaultTransport {
		t.Fatalf("Wrong transport in default HTTP client")
	}
}
