package repositories

type Repository interface {
}

type repository struct {
	CollectionName string
}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) List() {

}

func (repo *repository) Retrieve() {

}

func (repo *repository) Create() {

}

func (repo *repository) Update() {

}

func (repo *repository) Delete() {

}
