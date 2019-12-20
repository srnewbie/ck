package main

import (
	"github.com/srnewbie/ck/dispatcher"
	"github.com/srnewbie/ck/server"
)

func main() {
	d := dispatcher.New()
	d.Start()
	server.New(d).ListenAndServe()
}
