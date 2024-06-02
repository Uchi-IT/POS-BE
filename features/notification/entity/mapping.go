package entity

import "uchiiParfume/features/notification/model"

func NotifikasiModelToNotifikasiCore(notifikasi model.Notifikasi) NotifikasiCore{
	notifikasiCore := NotifikasiCore{
		Id: notifikasi.Id,
		TotalStock: notifikasi.TotalStock,
	}
	return notifikasiCore
}

func NotifikasiCoreToNotifikasiModel(notifikasi NotifikasiCore) model.Notifikasi{
	notifikasiModel := model.Notifikasi{
		Id: notifikasi.CabangId,
		TotalStock: notifikasi.TotalStock,
	}
	return notifikasiModel
}

func ListNotifikasiModelToNotifikasiCore(notifikasi []model.Notifikasi) []NotifikasiCore{
	coreNotifikasi := []NotifikasiCore{}
	for _, v := range notifikasi {
		notifications := NotifikasiModelToNotifikasiCore(v)
		coreNotifikasi = append(coreNotifikasi, notifications)
	}
	return coreNotifikasi
}