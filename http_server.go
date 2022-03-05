package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello, world!\n")
}

func main() {
    host := flag.String("host", "127.0.0.1", "Binding host.")
    port := flag.Int("port", 8000, "Binding port.")
    flag.Parse()
    bindstr := fmt.Sprintf("%s:%d", *host, *port)

    mux := http.NewServeMux()
    mux.HandleFunc("/", helloHandler)
    log.Printf("Start HTTP server on %s ...\n", bindstr)
    if err := http.ListenAndServe(bindstr, mux); err != nil {
        log.Fatalln("failed!")
    }
}
