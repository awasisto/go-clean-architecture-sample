package query

import (
	"go-clean-architecture-sample/application/common/interfaces"
	"go-clean-architecture-sample/domain/entities"
)

type GetEmployeeByIdQuery struct {
	EmployeeId int
}

type GetEmployeeByIdQueryHandler struct {
	employeeDataSource interfaces.EmployeeDataSource
	avatarProvider     interfaces.AvatarProvider
}

func NewGetEmployeeByIdQueryHandler(
	employeeDataSource interfaces.EmployeeDataSource,
	avatarProvider interfaces.AvatarProvider,
) *GetEmployeeByIdQueryHandler {
	return &GetEmployeeByIdQueryHandler{
		employeeDataSource: employeeDataSource,
		avatarProvider:     avatarProvider,
	}
}

func (h *GetEmployeeByIdQueryHandler) Handle(request GetEmployeeByIdQuery) (*entities.Employee, error) {
	employee, err := h.employeeDataSource.GetById(request.EmployeeId)
	if err != nil {
		return nil, err
	}

	employee.AvatarUrl, err = h.avatarProvider.GetAvatarUrlByEmail(employee.Email)
	if err != nil {
		return nil, err
	}

	return employee, nil
}
