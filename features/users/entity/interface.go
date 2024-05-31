package entity

type UsersRepositoryInterface interface {
	CreateUser(data UsersCore) (UsersCore, error)
	GetById(id string) (UsersCore, error)
	GetAllUser() ([]UsersCore, error)
	UpdateUser(id string, data UsersCore) error
	DeleteUser(id string) error
}

type UsersServiceInterface interface {
	Login(email, password string) (UsersCore, string, error)
	CreateUser(data UsersCore) (UsersCore, error)
	GetById(id string) (UsersCore, error)
	GetAllUser() ([]UsersCore, error)
	UpdateUser(id string, data UsersCore) error
	DeleteUser(id string) error
}
