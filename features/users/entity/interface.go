package entity

import "uchiiParfume/utils/pagination"

type UsersRepositoryInterface interface {
	Login(email, password string) (UsersCore, string, error)
	CreateUser(data UsersCore) (UsersCore, error)
	GetById(id string) (UsersCore, error)
	GetAllUser(page, limit int, search, sort, filter string) ([]UsersCore, pagination.PageInfo, int, error)
	UpdateUser(id string, data UsersCore) error
	DeleteUser(id string) error
}

type UsersServiceInterface interface {
	Login(email, password string) (UsersCore, string, error)
	CreateUser(data UsersCore) (UsersCore, error)
	GetById(id string) (UsersCore, error)
	GetAllUser(page, limit int, search, sort, filter string) ([]UsersCore, pagination.PageInfo, int, error)
	UpdateUser(id string, data UsersCore) error
	DeleteUser(id string) error
}
