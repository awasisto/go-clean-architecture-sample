package usecases

import (
	"golang-clean-architecture-sample/pkg/core/data"
	"golang-clean-architecture-sample/pkg/core/entities"
)

type AddEmployeeUseCase struct {
	employee           entities.Employee
	employeeDataSource data.EmployeeDataSource
	avatarProvider     data.AvatarProvider
}

func NewAddEmployeeUseCase(
	employee entities.Employee,
	employeeDataSource data.EmployeeDataSource,
	avatarProvider data.AvatarProvider,
) *AddEmployeeUseCase {
	return &AddEmployeeUseCase{
		employee:           employee,
		employeeDataSource: employeeDataSource,
		avatarProvider:     avatarProvider,
	}
}

func (uc *AddEmployeeUseCase) Execute() (createdEmployee *entities.Employee, err error) {
	employeeId, err := uc.employeeDataSource.Add(uc.employee)
	if err != nil {
		return nil, err
	}

	getEmployeeByIdUseCase := NewGetEmployeeByIdUseCase(employeeId, uc.employeeDataSource, uc.avatarProvider)

	return getEmployeeByIdUseCase.Execute()
}
