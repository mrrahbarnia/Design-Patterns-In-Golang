package main

import (
	"os"
	"testing"

	"github.com/mrrahbarnia/Design-Patterns-In-Golang/configuration"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		App: configuration.New(nil),
	}

	os.Exit(m.Run())
}
