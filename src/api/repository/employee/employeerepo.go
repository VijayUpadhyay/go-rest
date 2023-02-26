package employee

import "learn/go-rest/src/api/model"

type EmployeesDao interface {
	FetchEmployeeById(id string) model.Employee
	SaveEmployee(employee model.Employee) string
}
