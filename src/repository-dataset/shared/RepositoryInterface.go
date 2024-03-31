package shared

type RepositoryInterface interface {
	Save(repo *RepositoryDto) error
	SaveMultiple(repos []RepositoryDto) error
	FindByID(id int64) (*RepositoryDto, error)
	FindByName(name string) (*RepositoryDto, error)
}
