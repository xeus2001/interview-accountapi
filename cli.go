/*
	My Package comment. This file basically is just for me to quickly learn Go.

INSTALL GO
	curl -LO https://go.dev/dl/go1.18beta1.linux-amd64.tar.gz
	sudo tar -C /usr/local -xzf go1.18beta1.linux-amd64.tar.gz
	mkdir ~/go
	export GOROOT=/usr/local/go
	export GOPATH=$HOME/go
	export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

INTERVIEW
	https://github.com/form3tech-oss
		https://github.com/orgs/form3tech-oss/repositories?type=all
	https://github.com/form3tech-oss/interview-accountapi
		https://github.com/form3tech-oss/go-form3   <-- reference implementation
	https://github.com/form3tech-oss/candidate-pack

OTHER
	https://github.com/form3tech-oss/rust-club

BOOKS
	https://go.dev/doc/effective_go 				!!!
	https://github.com/dariubs/GoBooks
	http://www.golangbootcamp.com/book/types

LIBRARY
	https://go.dev/doc/tutorial/create-module												!!!
		https://pkg.go.dev/go/build		<-- compile package?
	https://medium.com/mindorks/how-to-create-a-package-in-go-ae4e79b95241
	https://www.digitalocean.com/community/tutorials/how-to-write-packages-in-go			!!!
	https://www.geeksforgeeks.org/how-to-create-your-own-package-in-golang/
	https://vsoch.github.io/2019/go/
		https://github.com/sci-f/scif-go

TEST
	https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package	!!!
	https://quii.gitbook.io/learn-go-with-tests/
	https://blog.alexellis.io/golang-writing-unit-tests/
	https://pkg.go.dev/testing

SECURITY
	https://github.com/securego/gosec
	https://staticcheck.io/

NETWORKING
	https://iximiuz.com/en/posts/go-net-http-setsockopt-example/
	https://go.dev/doc/database/cancel-operations
	https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/
	https://github.com/SrijanSriv/gonetworking

MEMORY / GC
	reflect/value.go
	builtin/builtin.go
	https://mdlayher.com/blog/unsafe-string-interning-in-go/
	https://go101.org/article/unsafe.html
	Sick: No weak references in Go!
		see: https://github.com/golang/go/issues/43615						!!!!! Implement a weak map
	https://golangbyexample.com/understanding-uintptr-golang/
	https://pkg.go.dev/sync/atomic
	https://github.com/go4org/intern/blob/main/intern.go

GENERICS
	https://www.infoq.com/news/2021/12/go-generics-beta/
	https://www.infoq.com/news/2021/06/go-fuzzing-beta/

JSON
	https://blog.alexellis.io/golang-json-api-client/
	https://www.geeksforgeeks.org/how-to-parse-json-in-golang/
	https://github.com/buger/jsonparser

CURL
	curl -v 'http://172.19.0.4:8080/v1/organisation/accounts' -H 'accept: application/vnd.api+json'

*/
package main

import (
	"fmt"
)

var test = "Hello World"

type SpeedControl interface {
	Speed() float64
	addSpeed(delta float64)
}

type Vehicle struct {
	Name string
}

type Car struct {
	vehicle Vehicle
	age     int32
	speed   float64
}

/*
	Hello World
*/
func testPrint(vehicle *Vehicle) {
	println(vehicle.Name)
}

func (car *Car) Hello(greeter string) {
	fmt.Printf("%s %s\n", greeter, car.vehicle.Name)
}

func (car *Car) Speed() float64 {
	return car.speed
}

func (car *Car) addSpeed(delta float64) {
	car.speed += delta
}

func printSpeed(control SpeedControl) {
	println(control.Speed())
}

/*
	API design:

	Options
		protocol (enum http, https, http2, http3, ...)
		hostname
		port
		version
		connectTimeout
		requestTimeout
		// Add credentials here
	Request
	Response
	Ping
	PingResponse
	ListAccounts
	ListAccountsResponse
	CreateAccount
	CreateAccountResponse
	DeleteAccount
	DeleteAccountResponse

	AccountClient
	client.SetOptions()
	client.Options()
	client.Ping()
	client.Do(Request) Response
*/

//func old() {
//	url := "http://localhost:8080/v1/organisation/accounts"
//	client := http.Client{
//		Timeout: time.Second * 2,
//	}
//	req, err := http.NewRequest(http.MethodGet, url, nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	req.Header.Set("User-Agent", "FooLib")
//	req.Header.Set("Accept", "application/json")
//	res, getErr := client.Do(req)
//	if getErr != nil {
//		log.Fatal(getErr)
//	}
//	if res.Body != nil {
//		defer res.Body.Close()
//	}
//	body, readErr := ioutil.ReadAll(res.Body)
//	if readErr != nil {
//		log.Fatal(readErr)
//	}
//	var data []main2.Account
//	jsonErr := json.Unmarshal(body, &data)
//	if jsonErr != nil {
//		log.Fatal(jsonErr)
//	}
//
//	for d := range data {
//		account := data[d]
//		println(account.Id)
//	}
//
//	myCar := Car{Vehicle{Name: "Hello"}, 32, 0}
//	println(myCar.vehicle.Name)
//	testPrint(&myCar.vehicle)
//	var buh int32 = 0
//	defer fmt.Printf("%d\n", buh)
//	buh++
//	defer fmt.Printf("%d\n", buh)
//	buh++
//	defer fmt.Printf("%d\n", buh)
//	buh++
//	defer fmt.Printf("%d\n", buh)
//	buh++
//	defer fmt.Printf("We expect that after this text 3, 2, 1 and 0 follow\n")
//	if myCar.age+5 < 100 {
//		buh++
//	}
//	fmt.Printf("Hello %s %d\n", myCar.vehicle.Name, buh)
//	myCar.Hello("Cool")
//	myCar.addSpeed(10)
//	var xxx SpeedControl = &myCar
//	printSpeed(xxx)
//
//	s := []string{0: "no error", 4: "Eio", 7: "invalid argument"}
//	for key, value := range s {
//		fmt.Printf("%d = %s, ", key, value)
//	}
//	println()
//	d := map[int32]string{0: "no error", 4: "Eio", 7: "invalid argument"}
//	for key, value := range d {
//		fmt.Printf("%d = %s, ", key, value)
//	}
//	println()
//	f := [...]float64{7.0, 8.5, 9.1}
//	f2 := f[2:3]
//	fmt.Printf("len(f) = %d, %d\n", len(f), len(f2))
//	fmt.Printf("f[0]=%f, f2[0]=%f\n", f[0], f2[0])
//
//	myCar.vehicle.Name = "first"
//	myCar.age = 1
//	pMyCar := &myCar
//	pMyCar.age = 2
//	pMyCar.vehicle.Name = "second"
//	pMyCar2 := &myCar
//	pMyCar2.age = 3
//	pMyCar2.vehicle.Name = "third"
//	println(myCar.vehicle.Name)
//	println(myCar.age)
//	println(pMyCar.vehicle.Name)
//	println(pMyCar.age)
//	println(pMyCar2.vehicle.Name)
//	println(pMyCar2.age)
//}

type TestA struct {
	v string
}

type TestB struct {
	v TestA
}

func main() {
	//old()
	a := TestA{"hello"}
	pA := &a.v
	b := TestA{"hello"}
	pB := &b.v
	c := TestB{TestA{"hello"}}
	println(a == b)
	println(a == c.v)
	println(pA == pB)
	println()
}
