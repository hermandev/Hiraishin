package database

import (
	"log"

	"github.com/pocketbase/pocketbase"
)

var App *pocketbase.PocketBase

func InitDatabase() {
	App = pocketbase.New()

	go func() {
		if err := App.Start(); err != nil {
			log.Fatal(err)
		}
	}()
}
