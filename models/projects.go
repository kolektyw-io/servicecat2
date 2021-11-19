package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string
	Mnemonic    string
	Icon        string
	Description string
}

type IssueSchema struct {
	gorm.Model
}

type Issue struct {
	gorm.Model
}

type IssueComment struct {
	IssueID uint
}
