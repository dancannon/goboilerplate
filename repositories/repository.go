package repositories

import (
	. "github.com/dancannon/gonews/models"
)

type Repository interface {
	FindById(id int) (Model, error)
	FindAll() ([]Model, error)
	Store(model Model) (Model, error)
	Delete(id int) error
}

type PostRepository interface {
	FindById(id int) (Post, error)
	FindAll() ([]Post, error)
	Store(model Post) (Post, error)
	Delete(id int) error
}
