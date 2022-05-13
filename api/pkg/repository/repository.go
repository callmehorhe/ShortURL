package repository

import (
	"github.com/callmehorhe/shorturl/api/pkg/models"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateURL(model models.UrlModel) error {
	return r.db.Table("url_model").Create(&model).Error
}

func (r *Repository) GetURL(url string) (string, error) {
	var u models.UrlModel
	err := r.db.Table("url_model").Where("new_url=?", url).First(&u).Error
	return u.OldURL, err
}

func (r *Repository) IsUsed(url string) bool {
	var u models.UrlModel
	if err := r.db.Table("url_model").Where("new_url=?", url).First(&u).Error; err != nil {
		return false
	}
	return true
}

func (r *Repository) IsCreated(url string) string {
	var u string
	r.db.Select("new_url").Table("url_model").Where("old_url=?", url).Scan(&u)
	return u
}
