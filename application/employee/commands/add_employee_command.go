package commands

import (
	"go-clean-architecture-sample/application/common/interfaces"
	"go-clean-architecture-sample/domain/entities"
)

type AddEmployeeCommand struct {
	Name  string
	Email string
}

type AddEmployeeCommandHandler struct {
	employeeDataSource interfaces.EmployeeDataSource
	avatarProvider     interfaces.AvatarProvider
}

func NewAddEmployeeCommandHandler(
	employeeDataSource interfaces.EmployeeDataSource,
	avatarProvider interfaces.AvatarProvider,
) *AddEmployeeCommandHandler {
	return &AddEmployeeCommandHandler{
		employeeDataSource: employeeDataSource,
		avatarProvider:     avatarProvider,
	}
}

func (h *AddEmployeeCommandHandler) Handle(request AddEmployeeCommand) (createdEmployee *entities.Employee, err error) {
	entity := entities.Employee{
		Name:  request.Name,
		Email: request.Email,
	}

	entity.Id, err = h.employeeDataSource.Add(entity)
	if err != nil {
		return nil, err
	}

	entity.AvatarUrl, err = h.avatarProvider.GetAvatarUrlByEmail(entity.Email)
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
