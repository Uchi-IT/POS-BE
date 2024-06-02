package repository

import (
	"uchiiParfume/features/notification/entity"

	"gorm.io/gorm"
)

type notifikasiRepository struct {
	db *gorm.DB
}

func NewNotifikasiRepository(db *gorm.DB) entity.NotifikasiRepositoryInterface {
	return &notifikasiRepository{
		db: db,
	}
}

// DeleteNotifikasi implements entity.NotifikasiRepositoryInterface.
func (n *notifikasiRepository) DeleteNotifikasi(id string) error {
	panic("unimplemented")
}

// GetAllNotifikasi implements entity.NotifikasiRepositoryInterface.
func (n *notifikasiRepository) GetAllNotifikasi() ([]entity.NotifikasiCore, error) {
	panic("unimplemented")
}

// GetById implements entity.NotifikasiRepositoryInterface.
func (n *notifikasiRepository) GetById(id string) (entity.NotifikasiCore, error) {
	panic("unimplemented")
}

// PushNotifikasi implements entity.NotifikasiRepositoryInterface.
func (n *notifikasiRepository) PushNotifikasi(data entity.NotifikasiCore) (entity.NotifikasiCore, error) {
	panic("unimplemented")
}

// UpdateNotifikasi implements entity.NotifikasiRepositoryInterface.
func (n *notifikasiRepository) UpdateNotifikasi(id string, data entity.NotifikasiCore) error {
	panic("unimplemented")
}

