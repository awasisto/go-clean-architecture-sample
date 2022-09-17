package query

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-clean-architecture-sample/application/common/interfaces/mocks"
	"go-clean-architecture-sample/domain/entities"
	"testing"
)

func TestGetAllEmployees(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockEmployeeDataSource := mocks.NewMockEmployeeDataSource(ctrl)
	mockAvatarProvider := mocks.NewMockAvatarProvider(ctrl)

	getAllEmployeesQueryHandler := NewGetAllEmployeesQueryHandler(
		mockEmployeeDataSource,
		mockAvatarProvider,
	)

	mockEmployeeDataSource.EXPECT().
		GetAll().
		Return(
			[]entities.Employee{
				{
					Id:    42,
					Name:  "John Smith",
					Email: "john.smith@example.com",
				},
				{
					Id:    43,
					Name:  "Jane Smith",
					Email: "jane.smith@example.com",
				},
			},
			nil,
		)

	mockAvatarProvider.EXPECT().
		GetAvatarUrlByEmail("john.smith@example.com").
		Return("http://example.com/john_smith.jpg", nil)

	mockAvatarProvider.EXPECT().
		GetAvatarUrlByEmail("jane.smith@example.com").
		Return("http://example.com/jane_smith.jpg", nil)

	want := []entities.Employee{
		{
			Id:        42,
			Name:      "John Smith",
			Email:     "john.smith@example.com",
			AvatarUrl: "http://example.com/john_smith.jpg",
		},
		{
			Id:        43,
			Name:      "Jane Smith",
			Email:     "jane.smith@example.com",
			AvatarUrl: "http://example.com/jane_smith.jpg",
		},
	}

	got, _ := getAllEmployeesQueryHandler.Handle(GetAllEmployeesQuery{})

	assert.Equal(t, want, got)
}
