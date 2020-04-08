package main

import (
	"github.com/guitarpawat/bookshelf/db"
	"github.com/guitarpawat/bookshelf/db/mongodb"
	"github.com/guitarpawat/bookshelf/router"
	"time"
)

func main() {
	mongodb.MustConnectDB("mongodb://localhost:27017", "bookshelf", 10*time.Second)
	router.Listen(8080, db.GetRepoInstance(db.MongoDB))
}
