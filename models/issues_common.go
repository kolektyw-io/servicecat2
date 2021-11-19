package models

import "gorm.io/gorm"

type IssueType struct {
	gorm.Model
	Name string
	Icon string
}

type IssueState struct {
	gorm.Model
	Name    string
	AsDraft bool
	AsDone  bool
}

type IssuePriority struct {
	gorm.Model
	Name string
	Icon string
}
