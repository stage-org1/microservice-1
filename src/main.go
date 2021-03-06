package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", doRequest)
	err:= http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func doRequest(w http.ResponseWriter, r *http.Request) {
	subset := r.Header.Get("subset")
	w.Header().Set("subset", subset)
	resp, err := http.Get("http://microservice-2-service")
	statuscode := resp.StatusCode
	if (err == nil && resp.Body != nil) {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if (err == nil) {

			fmt.Fprintf(w, "request was a success, statuscode: " + strconv.Itoa(statuscode) + ", body: " + string(body))
		} else {
			http.Error(w, "an error occured:" + err.Error(), http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "an error occured:" + err.Error() , http.StatusInternalServerError)
	} 

}

