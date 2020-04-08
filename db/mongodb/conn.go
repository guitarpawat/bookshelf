package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var database *mongoDatabase

type mongoDatabase struct {
	DatabaseName string
	DatabaseURI  string
	Timeout      time.Duration
	conn         *mongo.Client
}

// ConnectDB connects program to the MongoDB with specified options.
func ConnectDB(dbURI string, dbName string, timeout time.Duration) error {
	database = &mongoDatabase{
		DatabaseName: dbName,
		DatabaseURI:  dbURI,
		Timeout:      timeout,
	}
	_, err := database.getConn()
	return err
}

// MustConnectDB calls ConnectDB but panic when has an error.
func MustConnectDB(dbURI string, dbName string, timeout time.Duration) {
	if err := ConnectDB(dbURI, dbName, timeout); err != nil {
		panic(err)
	}
}

func (m *mongoDatabase) getConn() (*mongo.Client, error) {
	var err error = nil
	if m.conn == nil {
		ctx, _ := context.WithTimeout(context.Background(), m.Timeout)
		m.conn, err = mongo.Connect(ctx, options.Client().ApplyURI(m.DatabaseURI))
		log.Println("connected to mongo db")
	}
	return m.conn, err
}

func (m *mongoDatabase) getDB() *mongo.Database {
	return m.conn.Database(m.DatabaseName)
}
