package main

import (
	"fmt"
	"github.com/alexflint/go-arg"
	"log"
	"net/http"
	"tjweldon/infi-site/util"
	"tjweldon/infi-site/web"
)

var args struct {
	Port int `arg:"positional" default:"8080"`
}

var (
	green = util.RGB(50, 255, 50)
	red   = util.RGB(255, 50, 50)
)

func main() {
	ParseArgs()
	http.HandleFunc("/", web.EchoPath.With(web.LogRequests))

	portSpec := fmt.Sprintf(":%d", args.Port)
	log.Fatal(http.ListenAndServe(portSpec, nil))
}

func ParseArgs() {
	parser := arg.MustParse(&args)
	lower, upper := 0, 1<<16
	if args.Port < lower || upper < args.Port {
		failMsg := fmt.Sprintf(
			"Port must be in the range %d - %d, got %d",
			lower, upper, args.Port,
		)
		parser.Fail(red(failMsg))
	}

	fmt.Printf(green("Serving Infi-Site on port %d!\n"), args.Port)
}
