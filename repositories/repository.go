package repositories

type Repository interface {
	List() (interface{}, error)
	Retrieve() (interface{}, error)
	Create() (interface{}, error)
	Update() (interface{}, error)
	Delete() (interface{}, error)
}

type repository struct {
	CollectionName string
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) List() (interface{}, error) {
	return nil, nil
}

func (repo *repository) Retrieve() (interface{}, error) {
	return nil, nil
}

func (repo *repository) Create() (interface{}, error) {
	return nil, nil
}

func (repo *repository) Update() (interface{}, error) {
	return nil, nil
}

func (repo *repository) Delete() (interface{}, error) {
	return nil, nil
}
