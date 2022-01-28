# Documentation

A simple example of how to use the client to create an account:

```go
package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/xeus2001/interview-accountapi/pkg/f3"
	"time"
)

func main() {
	fmt.Printf("Use endpoint: %s\n", f3.DefaultEndPoint)
	client := f3.NewClient()
	orgId := uuid.New().String()
	account := f3.NewAccount(
		&orgId,     // organization id
		"GB",       // country, iso-3166
		"100000",   // bank-id
		"GBDSC",    // bank-id-code
		"Foo",      // account holder name
		"12345678", // account number
		"GBP",      // currency, iso-4217
		"tests")    // customer-id (optional)
	created, err := client.CreateAccount(account)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	} else {
		fmt.Printf("Account %s created at %s\n", created.Id, created.CreatedOn.Format(time.UnixDate))
	}
}
```

The source code is committed in [cmd/cmd.go](cmd/cmd.go). To build and run it, just do `make clean && make && bin/f3`.

Do make more sophisticated accounts, please either directly create a new fresh account or simply modify the generated 
template. For more information see the documentation of the [f3 client library](./f3.md). 