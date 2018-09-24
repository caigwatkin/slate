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

	go_context "github.com/caigwatkin/go/context"
	go_http "github.com/caigwatkin/go/http"
	go_log "github.com/caigwatkin/go/log"
	go_secrets "github.com/caigwatkin/go/secrets"
	"github.com/caigwatkin/slate/app/slate/api"
)

var (
	debug        bool
	env          string
	gcpProjectID string
	port         int
	serviceName  string
)

func init() {
	flag.BoolVar(&debug, "debug", true, "Debug mode on/off")
	flag.StringVar(&env, "env", "dev", "Deployment environment")
	flag.StringVar(&gcpProjectID, "gcpProjectID", "slate-00", "GCP project ID")
	flag.IntVar(&port, "port", 8080, "Port")
	flag.StringVar(&serviceName, "serviceName", "Slate", "Service name in canonical case for header")
	flag.Parse()
}

func main() {
	ctx := go_context.StartUp()

	logClient := go_log.NewClient(debug)
	logClient.Info(ctx, "Logger initialised",
		go_log.FmtBool(debug, "debug"),
		go_log.FmtString(env, "env"),
		go_log.FmtString(gcpProjectID, "gcpProjectID"),
		go_log.FmtInt(port, "port"),
		go_log.FmtString(serviceName, "serviceName"),
		go_log.FmtStrings(os.Environ(), "os.Environ()"),
	)

	logClient.Info(ctx, "Creating secrets client")
	secretsClient, err := go_secrets.NewClient(ctx)
	if err != nil {
		logClient.Fatal(ctx, "Failed creating secrets client", go_log.FmtError(err))
	}
	logClient.Info(ctx, "Created secrets client")

	logClient.Info(ctx, "Creating http client")
	httpClient := go_http.NewClient(logClient, "Slate")
	logClient.Info(ctx, "Created http client")

	logClient.Info(ctx, "Creating API client")
	apiClient := api.NewClient(api.Config{
		Env:          env,
		GCPProjectID: gcpProjectID,
		Port:         fmt.Sprintf(":%d", port),
	}, httpClient, logClient, secretsClient)
	logClient.Info(ctx, "Created API client")

	if err := apiClient.ListenAndServe(ctx); err != nil {
		logClient.Fatal(ctx, "Slate client unexpectedly returned with error from listening and serving; terminating", go_log.FmtError(err))
		return
	}
	logClient.Fatal(ctx, "Slate client unexpectedly returned from listening and serving; terminating")
}
