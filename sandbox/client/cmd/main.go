package main

import (
	"github.com/gophergala/GopherKombat/sandbox/client"
	"log"
)

func main() {
	// Test client
	code := `package main; import "time"; func main() { time.Sleep(100*time.Millisecond) }`
	res, err := client.CompareExecution(code, code)
	if err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("%#v\n", res)
}
