package client_user

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Openid    string    `json:"openid"`
	Avatar    string    `json:"avatar"`
	Nickname  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Model struct {
	Db   *gorm.DB
	Name string
}

func NewModel(db *gorm.DB) *Model {
	return &Model{
		Db:   db,
		Name: "client_users",
	}
}

func (s *Model) GetByOpenid(openid string) (ret User, err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("openid = ?", openid).First(&ret).Error
	return
}

func (s *Model) GetByUid(uid int) (ret User, err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("id = ?", uid).First(&ret).Error
	return
}

func (s *Model) Create(in User) (ret User, err error) {
	db := s.Db.Table(s.Name)
	err = db.Create(&in).Error
	ret = in
	return
}

func (s *Model) UpdateById(in User) (err error) {
	db := s.Db.Table(s.Name)
	err = db.Where("id = ?", in.Id).Updates(map[string]interface{}{
		"avatar":   in.Avatar,
		"nickname": in.Nickname,
	}).Error
	return
}
