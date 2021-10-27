package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	instance *db
	doOnce   sync.Once
	doLock             = &sync.Mutex{}
	out      io.Writer = os.Stdout
)

type db struct {
	Connection string
}

func setInstance(connection string) {
	instance = &db{Connection: connection}
	fmt.Fprintln(out, "Singleton instance has been created.")
}

func (d *db) GetConnection() string {
	return d.Connection
}

func GetInstanceByDoOnce(connection string) *db {
	doOnce.Do(func() {
		setInstance(connection)
	})
	return instance
}

func GetInstanceByDoLock(connection string) *db {
	if instance == nil {
		doLock.Lock()
		defer doLock.Unlock()
		setInstance(connection)
	}
	return instance
}
