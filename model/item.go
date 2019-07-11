package model

import "github.com/jinzhu/gorm"

type Item struct {
	ID    int
	Title string
	Done  bool
}

func (Item) TableName() string {
	return "items"
}

type ItemRepo interface {
	Create(item *Item) error
	ListAll() ([]Item, error)
}

type SQLItemRepo struct {
	DB *gorm.DB
}

func (s SQLItemRepo) Create(item *Item) error {
	if err := s.DB.Create(&item).Error; err != nil {
		return err
	}
	return nil
}

func (s SQLItemRepo) ListAll() ([]Item, error) {
	items := []Item{}
	if err := s.DB.Model(&Item{}).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}
