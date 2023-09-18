package IRepository

type ITaskPort interface {
	Create() error
	GetAll() error
	GetById() error
}
