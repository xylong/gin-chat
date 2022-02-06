package model

func migrate() {
	DB.AutoMigrate(&User{})
}
