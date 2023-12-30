package main

import (
	"database/sql"
	"log"
	"sync"
)


type DBPool struct {
	connString  string
	maxConn     int
	connections map[*sql.DB]bool
	sync.Mutex
}

func New(maxConn int, connString string) *DBPool {
	return &DBPool{
		connString:  connString,
		maxConn:     maxConn,
		connections: make(map[*sql.DB]bool),
	}
}

func (db *DBPool) CreateConnectionPool() {
	for i := 0; i < db.maxConn; i++ {
		conn, err := sql.Open("postgres", db.connString)
		if err != nil {
			log.Fatal("Cannot Create Connection Pool with DB: " + err.Error())
		}
		err = conn.Ping()
		if err != nil {
			log.Fatal("Cannot Ping DB: " + err.Error())
		}
		db.connections[conn] = false
	}
}

func (db *DBPool) ClosePool() {
	for conn, value := range db.connections {
		if value {
			err := conn.Close()
			db.connections[conn] = false
			if err != nil {
				log.Fatal("Cannot Close Connection in DBPool: " + err.Error())
			}
		}
	}
}

func (db *DBPool) getConn() *sql.DB {
	db.Lock()
	defer db.Unlock()
	for conn, value := range db.connections {
		if !value {
			db.connections[conn] = true
			return conn
		}
	}
	log.Fatal("Cannot Get connection! All Connections are in Use")
	return nil
}

func (db *DBPool) closeConn(conn *sql.DB) {
	db.Lock()
	defer db.Unlock()
	db.connections[conn] = false
}
