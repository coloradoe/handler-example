package main

import (
	"fmt"
	"net/http"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {

	tracer.Start(tracer.WithAnalytics(true))
	defer tracer.Stop()

	// Create a traced mux router.
	mux := muxtrace.NewRouter(muxtrace.WithServiceName("Howdy"))

	// Continue using the router as you normally would.
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	fmt.Println("Starting on port:8080")
	http.ListenAndServe(":8080", mux)
}
