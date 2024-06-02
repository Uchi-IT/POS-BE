package service

import "uchiiParfume/features/notification/entity"

type notifikasiService struct {
	NotifikasiRepository entity.NotifikasiRepositoryInterface
}

func NewNotifikasiService(notifikasi entity.NotifikasiRepositoryInterface) entity.NotifikasiServiceInterface {
	return &notifikasiService{
		NotifikasiRepository: notifikasi,
	}
}

// DeleteNotifikasi implements entity.NotifikasiServiceInterface.
func (n *notifikasiService) DeleteNotifikasi(id string) error {
	panic("unimplemented")
}

// GetAllNotifikasi implements entity.NotifikasiServiceInterface.
func (n *notifikasiService) GetAllNotifikasi() ([]entity.NotifikasiCore, error) {
	panic("unimplemented")
}

// GetById implements entity.NotifikasiServiceInterface.
func (n *notifikasiService) GetById(id string) (entity.NotifikasiCore, error) {
	panic("unimplemented")
}

// PushNotifikasi implements entity.NotifikasiServiceInterface.
func (n *notifikasiService) PushNotifikasi(data entity.NotifikasiCore) (entity.NotifikasiCore, error) {
	panic("unimplemented")
}

// UpdateNotifikasi implements entity.NotifikasiServiceInterface.
func (n *notifikasiService) UpdateNotifikasi(id string, data entity.NotifikasiCore) error {
	panic("unimplemented")
}
