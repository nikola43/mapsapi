package controllers

import (
	"gorm.io/gorm"
)

var GormDB *gorm.DB

func Migrate() {
	/*
		// DROP
		GormDB.Migrator().DropTable(&models.Client{})
		GormDB.Migrator().DropTable(&models.Employee{})
		GormDB.Migrator().DropTable(&models.Clinic{})
		GormDB.Migrator().DropTable(&models.Video{})
		GormDB.Migrator().DropTable(&models.Image{})
		GormDB.Migrator().DropTable(&models.Holographic{})
		GormDB.Migrator().DropTable(&models.Heartbeat{})
		GormDB.Migrator().DropTable(&streamingModels.Streaming{})
		GormDB.Migrator().DropTable(&recovery.UserRecovery{})
		GormDB.Migrator().DropTable(&models.PushNotificationData{})
		GormDB.Migrator().DropTable(&models.Promo{})
		GormDB.Migrator().DropTable(&models.CalculatorDetail{})
		GormDB.Migrator().DropTable(&kicks.Kick{})
		GormDB.Migrator().DropTable(&models.Company{})
		GormDB.Migrator().DropTable(&models.Payment{})
		GormDB.Migrator().DropTable(&models.Invitation{})
		GormDB.Migrator().DropTable(&models.ClinicClient{})
		GormDB.Migrator().DropTable(&models.ClinicPromo{})

		// CREATE
		GormDB.AutoMigrate(&models.Client{})
		GormDB.AutoMigrate(&models.Employee{})
		GormDB.AutoMigrate(&models.Clinic{})
		GormDB.AutoMigrate(&models.Video{})
		GormDB.AutoMigrate(&models.Image{})
		GormDB.AutoMigrate(&models.Holographic{})
		GormDB.AutoMigrate(&models.Heartbeat{})
		GormDB.AutoMigrate(&streamingModels.Streaming{})
		GormDB.AutoMigrate(&recovery.UserRecovery{})
		GormDB.AutoMigrate(&models.PushNotificationData{})
		GormDB.AutoMigrate(&models.Promo{})
		GormDB.AutoMigrate(&models.CalculatorDetail{})
		GormDB.AutoMigrate(&kicks.Kick{})
		GormDB.AutoMigrate(&models.Company{})
		GormDB.AutoMigrate(&models.Payment{})
		GormDB.AutoMigrate(&models.Invitation{})
		GormDB.AutoMigrate(&models.ClinicClient{})
		GormDB.AutoMigrate(&models.ClinicPromo{})

		//GormDB.SetupJoinTable(&models.Client{}, "Addresses", &models.Clinic{})
	*/
}
