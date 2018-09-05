package main

import (
	"flag"
	"os"
	"slate/internal/app/slate"
	"slate/internal/pkg/context"
	"slate/internal/pkg/log"
	"slate/internal/pkg/secret"
)

var (
	debug bool
)

func init() {
	flag.BoolVar(&debug, "debug", true, "Set debug mode")
	flag.Parse()
}

func main() {
	ctx := context.StartUp()

	log.Init(debug)
	log.Info(ctx, "Logger initialised", log.FmtStrings(os.Environ(), "os.Environ()"))

	secretClient, err := secret.NewClient(ctx)
	if err != nil {
		log.Fatal(ctx, "Failed creating secret client", log.FmtError(err))
	}
	log.Info(ctx, "Secret client created")

	slateClient := slate.NewClient(secretClient)
	log.Info(ctx, "Slate client created")

	slateClient.ListenAndServe(ctx, 8080)
	log.Fatal(ctx, "Slate client unexpectedly returned from listening and serving, terminating")
}
