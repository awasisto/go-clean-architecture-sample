package database

import (
	"go-clean-architecture-sample/application/common/errors"
	"go-clean-architecture-sample/domain/entities"
)

type InMemoryEmployeeDatabase struct {
	employees []entities.Employee
}

func NewInMemoryEmployeeDatabase() *InMemoryEmployeeDatabase {
	return &InMemoryEmployeeDatabase{}
}

func (d *InMemoryEmployeeDatabase) Add(employee entities.Employee) (int, error) {
	d.employees = append(d.employees, employee)
	return len(d.employees) - 1, nil
}

func (d *InMemoryEmployeeDatabase) GetAll() ([]entities.Employee, error) {
	employees := make([]entities.Employee, 0, len(d.employees))
	for _, employee := range d.employees {
		employees = append(employees, employee)
	}
	return employees, nil
}

func (d *InMemoryEmployeeDatabase) GetById(id int) (*entities.Employee, error) {
	if id < 0 || id >= len(d.employees) {
		return nil, errors.ErrNotFound
	}
	return &d.employees[id], nil
}
