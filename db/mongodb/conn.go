package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var Database *MongoDatabase

type MongoDatabase struct {
	DatabaseName string
	DatabaseURI  string
	Timeout      time.Duration
	conn         *mongo.Client
}

func ConnectDB(dbURI string, dbName string, timeout time.Duration) error {
	Database = &MongoDatabase{
		DatabaseName: dbName,
		DatabaseURI:  dbURI,
		Timeout:      timeout,
	}
	_, err := Database.GetConn()
	return err
}

func MustConnectDB(dbURI string, dbName string, timeout time.Duration) {
	if err := ConnectDB(dbURI, dbName, timeout); err != nil {
		panic(err)
	}
}

func (m *MongoDatabase) GetConn() (*mongo.Client, error) {
	var err error = nil
	if m.conn == nil {
		ctx, _ := context.WithTimeout(context.Background(), m.Timeout)
		m.conn, err = mongo.Connect(ctx, options.Client().ApplyURI(m.DatabaseURI))
		log.Println("connected to mongo db")
	}
	return m.conn, err
}

func (m *MongoDatabase) GetDB() *mongo.Database {
	return m.conn.Database(m.DatabaseName)
}
