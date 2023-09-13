package main

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
)

//https://gist.github.com/denji/12b3a568f092ab951456

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":8080", "/home/admin/conf/web/us.example.com/ssl/us.example.com.crt", "/home/admin/conf/web/us.example.com/ssl/us.example.com.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
