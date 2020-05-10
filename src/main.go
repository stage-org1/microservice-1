package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/unstableGreet", unstableGreet)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func unstableGreet(w http.ResponseWriter, r *http.Request) {
	chance := rand.Intn(100)

	if chance < 75 {
		fmt.Fprintf(w, "hello world")
	} else {
		fmt.Fprintf(w, "dlrow olleh")
	}
}
