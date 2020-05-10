package repositories

import (
	"go-atomic/main/helpers"
	"go-atomic/main/models"
)

func GetAllProduct(p *[]models.Product) (err error) {
	if err = helpers.GetDB().Find(p).Error; err != nil {
		return err
	}
	return nil
}
