package main

import (
	"os"
	"testing"

	"github.com/mrrahbarnia/Design-Patterns-In-Golang/models"
)

var testApp application

func TestMain(m *testing.M) {
	testApp = application{
		Models: *models.New(nil),
	}

	os.Exit(m.Run())
}
