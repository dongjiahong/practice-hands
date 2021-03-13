package models

import (
	"github.com/jinzhu/gorm"
)

type HelloDao struct {
	BaseDao
}

func NewHelloDao(conn *gorm.DB) *HelloDao {
	return &HelloDao{BaseDao: NewBaseDao(conn)}
}

func (d HelloDao) GetHello() string {

}
