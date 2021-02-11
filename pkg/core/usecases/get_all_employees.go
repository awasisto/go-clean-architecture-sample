package usecases

import (
	"golang-clean-architecture-sample/pkg/core/data"
	"golang-clean-architecture-sample/pkg/core/entities"
)

type GetAllEmployeesUseCase struct {
	employeeRepository data.EmployeeRepository
	avatarProvider     data.AvatarProvider
}

func NewGetAllEmployeesUseCase(
	employeeRepository data.EmployeeRepository,
	avatarProvider data.AvatarProvider,
) *GetAllEmployeesUseCase {
	return &GetAllEmployeesUseCase{
		employeeRepository: employeeRepository,
		avatarProvider:     avatarProvider,
	}
}

func (uc *GetAllEmployeesUseCase) Execute() ([]entities.Employee, error) {
	employees, err := uc.employeeRepository.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range employees {
		avatarUrl, err := uc.avatarProvider.GetAvatarUrlByEmail(employees[i].Email)
		if err != nil {
			return nil, err
		}

		employees[i].AvatarUrl = avatarUrl
	}

	return employees, nil
}
