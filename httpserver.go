package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/healthz", SimpleServer)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
func SimpleServer(w http.ResponseWriter, req *http.Request) {

	fmt.Printf(req.URL.String())
	//记录ip
	fmt.Printf("client ip = %s\n", req.RemoteAddr)
	w.WriteHeader(200)
	//返回页面200
	fmt.Fprintf(w, "200")

}
