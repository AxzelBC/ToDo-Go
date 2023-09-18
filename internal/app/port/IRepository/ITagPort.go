package IRepository

type ITagPort interface {
	Create() error
	GetAll() error
	GetById() error
}
