package rethinkdb

import (
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
)

func CreateSession() *r.Session {
	session, err := r.Connect("localhost:28015", "news")
	if err != nil {
		fmt.Println("error connecting to db:", err)
		return nil
	}

	return session
}
