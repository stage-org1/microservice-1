package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", doRequest)
	err:= http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func doRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("doRequest method called")
	resp, err := http.Get("http://microservice-2-service")
	if (err == nil) {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if (err == nil) {
			fmt.Println(err.Error())
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
