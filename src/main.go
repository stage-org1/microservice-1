package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/unstableGreet", unstableGreet)
	http.HandleFunc("/doRequest", doRequest)
	err:= http.ListenAndServe(":80", nil)
	if err != nil {
		os.Exit(1)
	}
}

func doRequest(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://microservice-2-vservice")
	if (err == nil) {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if (err == nil) {
			fmt.Fprintf(w, "request was a success, "+string(body))
		}
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)

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
