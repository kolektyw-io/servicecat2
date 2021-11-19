package models

import "gorm.io/gorm"

type DataType struct {
	gorm.Model
	Name string
	Type string
}

func (d *DataType) SetType(datatype string) {
	d.Type = datatype
}
