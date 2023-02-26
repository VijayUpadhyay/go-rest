package employee

import (
	"github.com/sirupsen/logrus"
	"learn/go-rest/src/api/model"
	"sync"
)

type EmpRepoImpl struct{}

var EmpDao EmployeesDao = EmpRepoImpl{}

var empMap sync.Map

func (e EmpRepoImpl) FetchEmployeeById(id string) (emp model.Employee) {
	if val, ok := empMap.Load(id); !ok {
		logrus.Errorf("Employee with given ID: %s doesn't exist", id)
		return
	} else {
		emp = val.(model.Employee)
	}
	return
}

func (e EmpRepoImpl) SaveEmployee(employee model.Employee) string {
	empMap.Store(employee.Id, employee)
	return employee.Id
}
