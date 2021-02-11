package usecases

import (
	"golang-clean-architecture-sample/pkg/core/data"
	"golang-clean-architecture-sample/pkg/core/entities"
)

type AddEmployeeUseCase struct {
	employee           entities.Employee
	employeeRepository data.EmployeeRepository
	avatarProvider     data.AvatarProvider
}

func NewAddEmployeeUseCase(
	employee entities.Employee,
	employeeRepository data.EmployeeRepository,
	avatarProvider data.AvatarProvider,
) *AddEmployeeUseCase {
	return &AddEmployeeUseCase{
		employee:           employee,
		employeeRepository: employeeRepository,
		avatarProvider:     avatarProvider,
	}
}

func (uc *AddEmployeeUseCase) Execute() (createdEmployee *entities.Employee, err error) {
	employeeId, err := uc.employeeRepository.Add(uc.employee)
	if err != nil {
		return nil, err
	}

	getEmployeeByIdUseCase := NewGetEmployeeByIdUseCase(employeeId, uc.employeeRepository, uc.avatarProvider)

	return getEmployeeByIdUseCase.Execute()
}
