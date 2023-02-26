package model

type Employee struct {
	Name         string `json:"name"`
	Id           string `json:"id"`
	PhoneNo      int    `json:"phoneNo"`
	DepartmentId int    `json:"departmentId"`
}
