package main

import (
	"fmt"
	"net/http"
	"time"
)

var counter int

func main() {

	var x int
	x = 1

	hello := "Hello, World!"

	fmt.Println(hello, x) // Hello, World! 1

	if x == 1 {
		fmt.Println("x is 1") // x is 1
	} else {
		fmt.Println("x is not 1") // x is not 1
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i) // 0 1 2 3 4 5 6 7 8 9
	}

	a := 5
	b := 10

	a, b = b, a

	fmt.Println(a, b)

	counter = 0

	server()

	fmt.Println(add(1, 2))

}

func add(x int, y int) int {
	return x + y
}

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Golang Microservice!")
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		handleUser(w, r)
	})

	// Start the server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}

}

func handleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Processing request #", counter)
	localCounter := counter

	counter++
	time.Sleep(5 * time.Second)

	// Respond to the client
	fmt.Println("Response sent to client #", localCounter)

	fmt.Fprintf(w, "Goodbye, User!")

}
