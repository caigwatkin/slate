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
	"os/signal"
	"syscall"

	go_context "github.com/caigwatkin/go/context"
	go_http "github.com/caigwatkin/go/http"
	go_log "github.com/caigwatkin/go/log"
	go_secrets "github.com/caigwatkin/go/secrets"
	"github.com/caigwatkin/slate/internal/app"
	"github.com/caigwatkin/slate/internal/http"
	"github.com/caigwatkin/slate/internal/lib/firestore"
)

var (
	cloudkmsKey      string
	cloudkmsKeyRing  string
	debug            bool
	env              string
	gcpProjectID     string
	port             int
	secretsBucket    string
	secretsBucketDir string
	serviceName      string
)

func init() {
	flag.BoolVar(&debug, "debug", true, "Debug mode on/off")
	flag.StringVar(&env, "env", "dev", "Deployment environment")
	flag.StringVar(&gcpProjectID, "gcpProjectID", "slate-00", "GCP project ID")
	flag.StringVar(&cloudkmsKey, "cloudkmsKey", "slate", "GCP cloud KMS key")
	flag.StringVar(&cloudkmsKeyRing, "cloudkmsKeyRing", "slate", "GCP cloud KMS key ring")
	flag.IntVar(&port, "port", 8080, "Port")
	flag.StringVar(&secretsBucket, "secretsBucket", "slate-api-config", "GCP bucket storing secrets directory")
	flag.StringVar(&secretsBucketDir, "secretsBucketDir", "secrets", "GCP bucket secrets directory storing secrets")
	flag.StringVar(&serviceName, "serviceName", "Slate-Api", "Service name in canonical case for header")
	flag.Parse()
}

func main() {
	ctx := go_context.StartUp()

	logClient := go_log.NewClient(debug)
	logClient.Info(ctx, "Logger initialised",
		go_log.FmtBool(debug, "debug"),
		go_log.FmtString(env, "env"),
		go_log.FmtString(gcpProjectID, "gcpProjectID"),
		go_log.FmtString(cloudkmsKey, "cloudkmsKey"),
		go_log.FmtString(cloudkmsKeyRing, "cloudkmsKeyRing"),
		go_log.FmtInt(port, "port"),
		go_log.FmtString(secretsBucket, "secretsBucket"),
		go_log.FmtString(secretsBucketDir, "secretsBucketDir"),
		go_log.FmtString(serviceName, "serviceName"),
		go_log.FmtStrings(os.Environ(), "os.Environ()"),
	)

	logClient.Info(ctx, "Creating go HTTP client")
	goHTTPClient := go_http.NewClient(logClient, "Slate")
	logClient.Info(ctx, "Created go HTTP client")

	logClient.Info(ctx, "Creating secrets client")
	secretsClient, err := go_secrets.NewClient(ctx, env, gcpProjectID, cloudkmsKeyRing, cloudkmsKey)
	if err != nil {
		logClient.Fatal(ctx, "Failed creating secrets client", go_log.FmtError(err))
	}
	logClient.Info(ctx, "Created secrets client")

	requiredSecrets := firestore.RequiredSecrets()
	logClient.Info(ctx, "Dowloading required secrets", go_log.FmtAny(requiredSecrets, "requiredSecrets"))
	if err := secretsClient.DownloadAndDecryptAndCache(ctx, secretsBucket, secretsBucketDir, requiredSecrets); err != nil {
		logClient.Fatal(ctx, "Failed downloading and decrypting and caching required secrets", go_log.FmtError(err))
	}
	logClient.Info(ctx, "Dowloaded required secrets")

	logClient.Info(ctx, "Creating firestore client")
	firestoreClient, err := firestore.NewClient(ctx, logClient, secretsClient)
	if err != nil {
		logClient.Fatal(ctx, "Failed creating firestore client", go_log.FmtError(err))
	}
	defer firestoreClient.Close()
	logClient.Info(ctx, "Created firestore client")

	logClient.Info(ctx, "Creating app client")
	appClient := app.NewClient(firestoreClient, logClient)
	logClient.Info(ctx, "Created app client")

	logClient.Info(ctx, "Creating HTTP client")
	httpClient := http.NewClient(http.Config{
		Env:          env,
		GCPProjectID: gcpProjectID,
		Port:         fmt.Sprintf(":%d", port),
		ServiceName:  serviceName,
	}, appClient, goHTTPClient, logClient)
	logClient.Info(ctx, "Created HTTP client")

	logClient.Info(ctx, "Preparing clean up")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		ctx := go_context.ShutDown()
		logClient.Notice(ctx, "Cleaning up")
		firestoreClient.Close()
		logClient.Notice(ctx, "Cleaned up")
		os.Exit(2)
	}()
	logClient.Info(ctx, "Prepared clean up")

	logClient.Info(ctx, "Listening and serving", go_log.FmtInt(port, "port"))
	if err := httpClient.ListenAndServe(); err != nil {
		logClient.Fatal(ctx, "HTTP client unexpectedly returned with error from listening and serving; terminating", go_log.FmtError(err))
		return
	}
	logClient.Fatal(ctx, "HTTP client unexpectedly returned from listening and serving; terminating")
}
