package main

import (
	"log"

	"ServiceStructure/server"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func init() {
	//set logger or other initial setup
}

func main() {
	// Start the tracer
	tracer.Start()

	// When the tracer is stopped, it will flush everything it has to the Datadog Agent before quitting.
	// Make sure this line stays in your main function.
	defer tracer.Stop()

	// Create new server instance
	srv := server.New()
	// Log the server host and port
	log.Printf("Server running on %s", srv.Address)
	// Start HTTP Server
	srv.ServeHTTP()
}
