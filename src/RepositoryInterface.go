package main

type RepositoryInterface interface {
	Save(repo *Repository) error
	SaveMultiple(repos []Repository) error
	FindByID(id int64) (*Repository, error)
	FindByName(name string) (*Repository, error)
}
