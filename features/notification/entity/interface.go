package entity

type NotifikasiRepositoryInterface interface {
	PushNotifikasi(data NotifikasiCore) (NotifikasiCore, error)
	GetById(id string) (NotifikasiCore, error)
	GetAllNotifikasi() ([]NotifikasiCore, error)
	UpdateNotifikasi(id string, data NotifikasiCore) error
	DeleteNotifikasi(id string) error
}

type NotifikasiServiceInterface interface {
	PushNotifikasi(data NotifikasiCore) (NotifikasiCore, error)
	GetById(id string) (NotifikasiCore, error)
	GetAllNotifikasi() ([]NotifikasiCore, error)
	UpdateNotifikasi(id string, data NotifikasiCore) error
	DeleteNotifikasi(id string) error
}
