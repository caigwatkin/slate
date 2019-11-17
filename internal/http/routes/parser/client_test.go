package parser

import (
	go_log_mock "github.com/caigwatkin/go/log/mock"
)

var (
	clientSuccess = client{
		logClient: go_log_mock.MockClientSuccess,
	}
)
