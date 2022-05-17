package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port string
	flag.StringVar(&port, "P", "8080", "端口号")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s", "正在访问"+port)
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
