package main

import (
	"context"
	"log"
	"main/microservices/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// registers func to path on DefaultServeMux
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Hello World")
	// 	d, _ := ioutil.ReadAll(r.Body)

	// 	// log.Printf("Data %s\n", d)
	// 	fmt.Fprintf(w, "Hello %s", d)
	// })

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	newHandler := handlers.NewHello(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", newHandler)

	// Creating a webserver (port, handler)
	// http.ListenAndServe(":9090", nil) //Default Serve mux
	// http.ListenAndServe(":9090", serveMux)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill) //Broadcast msg on channel when kill command is given

	sig := <-signalChannel // <- recieve operator represents the idea of passing a value from a channel to a reference
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
