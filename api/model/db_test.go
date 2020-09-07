package model

import "testing"

func TestInitDB(t *testing.T) {
	InitDB()
	DB.AutoMigrate(&Task{}, &Rule{})
}
