package db

import (
	"golang-clean-architecture-sample/pkg/core"
	"golang-clean-architecture-sample/pkg/core/entities"
)

type InMemoryEmployeeDb struct {
	employees []entities.Employee
}

func NewInMemoryEmployeeDatabase() *InMemoryEmployeeDb {
	return &InMemoryEmployeeDb{
		employees: make([]entities.Employee, 0),
	}
}

func (i *InMemoryEmployeeDb) Add(employee entities.Employee) (int64, error) {
	i.employees = append(i.employees, employee)
	return int64(len(i.employees) - 1), nil
}

func (i *InMemoryEmployeeDb) GetAll() ([]entities.Employee, error) {
	employees := make([]entities.Employee, 0, len(i.employees))
	for _, employee := range i.employees {
		employees = append(employees, employee)
	}
	return employees, nil
}

func (i *InMemoryEmployeeDb) GetById(id int64) (*entities.Employee, error) {
	if id < 0 || id >= int64(len(i.employees)) {
		return nil, core.ErrNotFound
	}
	return &i.employees[id], nil
}
