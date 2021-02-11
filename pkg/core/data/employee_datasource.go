package data

import "golang-clean-architecture-sample/pkg/core/entities"

type EmployeeDataSource interface {
	Add(employee entities.Employee) (employeeId int64, err error)
	GetAll() (employees []entities.Employee, err error)
	GetById(id int64) (employee *entities.Employee, err error)
}
