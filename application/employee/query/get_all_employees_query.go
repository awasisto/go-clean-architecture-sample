package query

import (
	"go-clean-architecture-sample/application/common/interfaces"
	"go-clean-architecture-sample/domain/entities"
)

type GetAllEmployeesQuery struct{}

type GetAllEmployeesQueryHandler struct {
	employeeDataSource interfaces.EmployeeDataSource
	avatarProvider     interfaces.AvatarProvider
}

func NewGetAllEmployeesQueryHandler(
	employeeDataSource interfaces.EmployeeDataSource,
	avatarProvider interfaces.AvatarProvider,
) *GetAllEmployeesQueryHandler {
	return &GetAllEmployeesQueryHandler{
		employeeDataSource: employeeDataSource,
		avatarProvider:     avatarProvider,
	}
}

func (h *GetAllEmployeesQueryHandler) Handle(request GetAllEmployeesQuery) ([]entities.Employee, error) {
	employees, err := h.employeeDataSource.GetAll()
	if err != nil {
		return nil, err
	}

	for i := range employees {
		employees[i].AvatarUrl, err = h.avatarProvider.GetAvatarUrlByEmail(employees[i].Email)
		if err != nil {
			return nil, err
		}
	}

	return employees, nil
}
