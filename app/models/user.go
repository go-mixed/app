package models

import (
	"gopkg.in/go-mixed/framework.v1/database/orm"
)

type User struct {
	orm.Model
	Name   string
	Avatar string
	orm.SoftDeletes
}
