package database

import (
	"log"

	"github.com/charmbracelet/charm/kv"
)

//CreateClient returns instance of the DB
func CreateClient() *kv.KV {
	db, err := kv.OpenWithDefaults("todo_badger")
	if err != nil {
		log.Fatalf("Error with DB %s \n", err)
	}
	db.Sync()
	return db
}
