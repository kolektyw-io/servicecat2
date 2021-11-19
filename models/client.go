package models

import "gorm.io/gorm"

type PortalGroup struct {
	gorm.Model
	Name  string
	Icon  string
	Order uint8
}

type PortalForm struct {
	gorm.Model
	Name   string
	Icon   string
	Fields []Field
}

type Field struct {
	gorm.Model
	Name string
	Type string
}
