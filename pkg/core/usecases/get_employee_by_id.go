package usecases

import (
	"golang-clean-architecture-sample/pkg/core/data"
	"golang-clean-architecture-sample/pkg/core/entities"
)

type GetEmployeeByIdUseCase struct {
	employeeId         int64
	employeeRepository data.EmployeeRepository
	avatarProvider     data.AvatarProvider
}

func NewGetEmployeeByIdUseCase(
	employeeId int64,
	employeeRepository data.EmployeeRepository,
	avatarProvider data.AvatarProvider,
) *GetEmployeeByIdUseCase {
	return &GetEmployeeByIdUseCase{
		employeeId:         employeeId,
		employeeRepository: employeeRepository,
		avatarProvider:     avatarProvider,
	}
}

func (uc *GetEmployeeByIdUseCase) Execute() (*entities.Employee, error) {
	employee, err := uc.employeeRepository.GetById(uc.employeeId)
	if err != nil {
		return nil, err
	}

	avatarUrl, err := uc.avatarProvider.GetAvatarUrlByEmail(employee.Email)
	if err != nil {
		return nil, err
	}

	return &entities.Employee{
		Id:        employee.Id,
		Name:      employee.Name,
		Email:     employee.Email,
		AvatarUrl: avatarUrl,
	}, nil
}
