package main

import (
	"ck/dispatcher"
	"ck/server"
)

func main() {
	d := dispatcher.New()
	d.Start()
	server.New(d).ListenAndServe()
}
