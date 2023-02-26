package service

import "learn/go-rest/src/api/model"

type EmployeeService interface {
	GetEmployee(id string) model.Employee
	CreateEmployee(employee model.Employee) string
}
