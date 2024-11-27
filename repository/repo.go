package repository

import (
	"fmt"
	"routersmux/models"

	"gorm.io/gorm"
)

type DataConnection struct {
	DB *gorm.DB
}

func (databaseConnect DataConnection) NewCreateUserDetails(user *models.Employee) *gorm.DB {
	data := databaseConnect.DB.Create(&user)

	return data
}

func (databaseConnect DataConnection) NewGetUser(user *[]models.Employee) *gorm.DB {
	value := databaseConnect.DB.Preload("Address").Find(&user)

	return value
}

func (databaseConnect DataConnection) NewSingleuser(id int, user *models.Employee) *gorm.DB {
	value := databaseConnect.DB.Preload("Address").Find(&user, id)

	return value
}

func (databaseConnect DataConnection) NewUpdateUsers(id int, user *models.Employee) *gorm.DB {
	value := databaseConnect.DB.Model(&models.Employee{}).Where("employee_id=?", id).Updates(map[string]interface{}{"name": user.Name, "age": user.Age})

	for _, datavalue := range user.Address {
		addressValue := databaseConnect.DB.Model(&models.Address{}).Where("address_id=?", datavalue.AddressId).Updates(map[string]interface{}{"street": datavalue.Street, "city": datavalue.City, "code": datavalue.Code})
		if addressValue.Error != nil {
			fmt.Println("Error occured", addressValue.Error)
			return nil
		}
	}
	return value
}

func (databaseConnect DataConnection) NewDeleteUser(id int, user *models.Employee) *gorm.DB {
	value := databaseConnect.DB.Delete(&user, id)

	return value
}
