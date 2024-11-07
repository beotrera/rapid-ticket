package handlers

import "gorm.io/gorm"

var db *gorm.DB 

func SetDB(DB *gorm.DB) {
	db = DB
}
