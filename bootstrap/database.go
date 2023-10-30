package bootstrap

import (
	env "go-base/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// DatabaseConnect establishes a connection to the database.
//
// It does not take any parameters.
// It returns a pointer to a gorm.DB object.
func DatabaseConnect() *gorm.DB {
	dsn := env.Load("POSTGRES_DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	println("Connected to the database")
	// Migrate the schema
	db.AutoMigrate()
	return db
}
func Repository(value interface{}) *gorm.DB {
	return db.Model(value)
}
