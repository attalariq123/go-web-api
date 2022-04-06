package database

import "web-api/pkg"

func MigrateDB() {
	DB.AutoMigrate(pkg.Book{}, pkg.User{})
}
