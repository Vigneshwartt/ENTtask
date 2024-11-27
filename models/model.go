package models

type Employee struct {
	EmployeeId int `gorm:"primaryKey"`
	Name       string
	Age        int
	Address    []Address `gorm:"foreignKey:employee_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Address struct {
	AddressId  int `gorm:"primaryKey"`
	EmployeeId int
	Employee   Employee `json:"-"`
	Street     string
	City       string
	Code       int
}
