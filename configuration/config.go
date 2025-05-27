package configuration

import (
	"database/sql"
	"sync"

	"github.com/mrrahbarnia/Design-Patterns-In-Golang/models"
)

var instance *Application
var once sync.Once
var db *sql.DB

type Application struct {
	Models *models.Models
}

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{Models: models.New(db)}
	})
	return instance
}
