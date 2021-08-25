package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	Password    string
	IsSuperUser bool
	IsStaff     bool
	Sessions    []Session
	IsActive	bool
}

type Group struct {
	gorm.Model
	Name        string
	Description string
}

type Session struct {
	gorm.Model
	Token  string
	UserID uint
}
