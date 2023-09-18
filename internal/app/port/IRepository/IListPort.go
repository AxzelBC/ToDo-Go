package IRepository

type IListPort interface {
	Create() error
	GetAll() error
	GetById() error
}
