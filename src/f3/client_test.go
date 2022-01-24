package f3_test

import (
	"github.com/xeus2001/interview-accountapi/src/f3"
	"github.com/xeus2001/interview-accountapi/src/iso"
	"github.com/xeus2001/interview-accountapi/src/iso/countryCode"
	"github.com/xeus2001/interview-accountapi/src/iso/currencyCode"
	"os"
	"strconv"
	"testing"
	"time"
)

// hostname of the account API service.
var hostname *string

// port of the account API service.
var port int32 = 8080

func init() {
	theHostName := os.Getenv("F3_HOST")
	if len(theHostName) == 0 {
		theHostName = "localhost"
	}
	hostname = &theHostName
	thePort := os.Getenv("F3_PORT")
	if len(thePort) != 0 {
		parsedInt, err := strconv.ParseInt(thePort, 10, 32)
		if err == nil {
			port = int32(parsedInt)
		}
	}
}

func TestCurrency(t *testing.T) {
	c, ok := iso.CurrencyByCode[currencyCode.Euro]
	if !ok {
		t.Fatal("EUR currency is wrong")
	}
	if c.Name != "Euro" {
		t.Fatal("EUR currency is wrong")
	}
	c, ok = iso.CurrencyByName["Euro"]
	if !ok {
		t.Fatal("USD currency is wrong")
	}
	if c.Code != currencyCode.Euro {
		t.Fatal("USD currency is wrong")
	}
}

func TestCountry(t *testing.T) {
	c, ok := iso.CountryByCode[countryCode.Germany]
	if !ok {
		t.Fatal("Germany is not in the country map")
	}
	if c.Code != "DE" {
		t.Fatalf("Germany has invalid alpha-2 code, expected DE, found: %s", c.Code)
	}
	if c.Alpha3 != "DEU" {
		t.Fatalf("Germany has invalid alpha-3 code, expected DEU, found: %s", c.Alpha3)
	}
	if c.Name != "Germany" {
		t.Fatalf("Germany has invalid name, expected Germany, found: %s", c.Name)
	}
}

func TestNewId(t *testing.T) {
	uuid, err := f3.NewUuid()
	if err != nil {
		t.Fatal("Failed to create unique Id for new account")
	}
	if uuid == nil {
		t.Fatal("Failed to create UUID")
	}
}

func TestNewAccount(t *testing.T) {
	account, e := f3.NewAccount(
		nil,
		countryCode.UnitedKingdom,
		"400300",
		[]string{"Alexander", "Lowey-Weber"},
		"41426819",
		"")
	if e != nil {
		t.Fatal("Failed to create account")
	}
	if account == nil {
		t.Fatal("Failed to create account")
	}
}

func TestClient_IsHealthy(t *testing.T) {
	client, err := f3.NewClient(*hostname, port, time.Second*5)
	if err != nil {
		t.Fatalf("Creating the F3 client failed, cause: %s", err.Error())
	}
	if !client.IsHealthy() {
		t.Fatalf("Health check for service failed")
	}
}

func TestClient_CreateAccount(t *testing.T) {
	client, e := f3.NewClient(*hostname, port, time.Second*5)
	if e != nil {
		t.Fatalf("Creating the F3 client failed, cause: %s", e.Error())
	}
	account, e := f3.NewAccount(
		nil,
		countryCode.UnitedKingdom,
		"400300",
		[]string{"Alexander", "Lowey-Weber"},
		"41426819",
		"")
	if e != nil {
		t.Fatal("Failed to create account")
	}
	if account == nil {
		t.Fatal("Failed to create account")
	}
	created, e := client.CreateAccount(account)
	if e != nil {
		t.Fatalf("Failed to create test account: %s", e.Error())
	}
	if created == nil {
		t.Fatal("Received nil as created account")
	}
}
