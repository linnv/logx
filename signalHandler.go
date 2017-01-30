package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	sigs := make(chan os.Signal, 2)
	// The only signal values guaranteed to be present on all systems
	// are Interrupt (send the process an interrupt) and Kill (force
	// the process to exit).
	signal.Notify(sigs, os.Interrupt, os.Kill)
	log.Printf("use c-c to exit\n")
	r := <-sigs
	log.Printf("got signal : %+v\n", r)
}
