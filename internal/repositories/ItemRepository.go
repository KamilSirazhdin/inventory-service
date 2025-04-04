package repositories

import (
	"github.com/jinzhu/gorm"
	"inventory-service/internal/models"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (repo *ItemRepository) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := repo.db.Find(&items).Error
	return items, err
}

func (repo *ItemRepository) CreateItem(item *models.Item) error {
	return repo.db.Create(item).Error
}

func (repo *ItemRepository) UpdateItem(item *models.Item) error {
	return repo.db.Save(item).Error
}

func (repo *ItemRepository) DeleteItem(id uint) error {
	return repo.db.Delete(&models.Item{}, id).Error
}
