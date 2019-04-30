package parser

import (
	go_log "github.com/caigwatkin/go/log"
)

var (
	clientSuccess = client{
		logClient: go_log.MockClientSuccess,
	}
)
