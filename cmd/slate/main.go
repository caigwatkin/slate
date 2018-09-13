/*
Copyright 2018 Cai Gwatkin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
	"slate/internal/app/slate/api"
	"slate/internal/pkg/context"
	"slate/internal/pkg/log"
	"slate/internal/pkg/secrets"
)

var (
	debug        bool
	env          string
	gcpProjectID string
	port         int
)

func init() {
	flag.BoolVar(&debug, "debug", true, "Debug mode on/off")
	flag.StringVar(&env, "env", "dev", "Deployment environment")
	flag.StringVar(&gcpProjectID, "gcp_project_id", "slate-00", "GCP project ID")
	flag.IntVar(&port, "port", 8080, "Port")
	flag.Parse()
}

func main() {
	ctx := context.StartUp()

	log.Init(debug)
	log.Info(ctx, "Logger initialised",
		log.FmtBool(debug, "debug"),
		log.FmtString(env, "env"),
		log.FmtString(gcpProjectID, "gcpProjectID"),
		log.FmtInt(port, "port"),
		log.FmtStrings(os.Environ(), "os.Environ()"),
	)

	log.Info(ctx, "Creating secrets client")
	secretsClient, err := secrets.NewClient(ctx)
	if err != nil {
		log.Fatal(ctx, "Failed creating secrets client", log.FmtError(err))
	}
	log.Info(ctx, "Created secrets client")

	log.Info(ctx, "Creating API client")
	apiClient := api.NewClient(api.Config{
		Env:          env,
		GCPProjectID: gcpProjectID,
		Port:         fmt.Sprintf(":%d", port),
	}, secretsClient)
	log.Info(ctx, "Created API client")

	if err := apiClient.ListenAndServe(ctx); err != nil {
		log.Fatal(ctx, "Slate client unexpectedly returned with error from listening and serving; terminating", log.FmtError(err))
		return
	}
	log.Fatal(ctx, "Slate client unexpectedly returned from listening and serving; terminating")
}
