package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /")
	io.WriteString(w, "Hello, world!\n")
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET /healthz")
	io.WriteString(w, "")
}

func main() {
	my_pid := os.Getpid()
	pidfile := "/tmp/httpserver.pid"
	if err := ioutil.WriteFile(pidfile, []byte(fmt.Sprintf("%d", my_pid)), 0644); err != nil {
		log.Fatalf("Cannot create pidfile %s\n", pidfile)
		os.Exit(-1)
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		s := <-c
		log.Println("Gracefully exit...", s)
		if err := os.Remove(pidfile); err != nil {
			panic(err)
		}
		os.Exit(0)
	}()

	host := "127.0.0.1"
	port := "8080"
	if _host := os.Getenv("HOST"); _host != "" {
		host = _host
	}
	if _port := os.Getenv("PORT"); _port != "" {
		port = _port
	}
	int_port, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		panic(err)
	}
	bindstr := fmt.Sprintf("%s:%d", host, int_port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/healthz", healthzHandler)
	log.Printf("Start HTTP server (%d) on %s ...\n", my_pid, bindstr)
	if err := http.ListenAndServe(bindstr, mux); err != nil {
		log.Fatalln("failed!")
		os.Remove(pidfile)
	}
}
