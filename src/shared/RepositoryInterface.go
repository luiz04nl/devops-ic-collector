package shared

type RepositoryInterface interface {
	Save(repo *RepositoryDto) error
	SaveMultiple(repos []RepositoryDto) error
	FindByID(id int) (*RepositoryDto, error)
	GetAll() ([]RepositoryDto, error)
	GetCloned() ([]RepositoryDto, error)
	FindByName(name string) (*RepositoryDto, error)
	UpdateById(id int, newRepository RepositoryDto) error
}
