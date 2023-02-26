package service

import "learn/go-rest/src/api/model"

type EmpSrvcImpl struct{}

var _ EmployeeService = EmpSrvcImpl{}

func (e EmpSrvcImpl) GetEmployee(id string) model.Employee {
	panic("")
}

func (e EmpSrvcImpl) CreateEmployee(employee model.Employee) string {
	panic("")
}
