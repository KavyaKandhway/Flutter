package db

import (
	"crypto/tls"
	"gopkg.in/mgo.v2"
	"net"
	"os"
)

type DBConnection struct {
	session *mgo.Session
}

func NewConnection() (conn *DBConnection) {
	dialInfo := mgo.DialInfo{
		Addrs:    []string{
			"fdb0-shard-00-00.wpqdt.mongodb.net:27017",
			"fdb0-shard-00-01.wpqdt.mongodb.net:27017",
			"fdb0-shard-00-02.wpqdt.mongodb.net:27017",
		},
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PWD"),
	}
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig) // add TLS config
		return conn, err
	}
	session, err := mgo.DialWithInfo(&dialInfo)

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnection{session}
	return conn
}

// Use handles connect to a certain collection
func (conn *DBConnection) Use(dbName, tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB(dbName).C(tableName)
}

// Close handles closing a database connection
func (conn *DBConnection) Close() {
	// This closes the connection
	conn.session.Close()
	return
}
