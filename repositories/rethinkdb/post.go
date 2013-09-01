package rethinkdb

import (
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
	. "github.com/dancannon/gonews/models"
)

func FindById(id int) (Post, error) {
	post := Post{}
	return post, fmt.Errorf("Not yet implemented")
}

func FindAll() ([]Post, error) {
	return nil, fmt.Errorf("Not yet implemented")
}

func Store(post Post) (Post, error) {
	session := CreateSession()
	rows := r.Table("posts").Insert(post).Run(session)

	err := rows.Err()
	if err != nil {
		return post, fmt.Errorf("error inserting into db: %s", err)
	}

	return post, nil
}

func Delete(id int) error {
	return fmt.Errorf("Not yet implemented")
}
