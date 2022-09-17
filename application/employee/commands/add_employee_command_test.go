package commands

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-clean-architecture-sample/application/common/interfaces/mocks"
	"go-clean-architecture-sample/domain/entities"
	"testing"
)

func TestAddEmployee(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockEmployeeDataSource := mocks.NewMockEmployeeDataSource(ctrl)
	mockAvatarProvider := mocks.NewMockAvatarProvider(ctrl)

	addEmployeeCommandHandler := NewAddEmployeeCommandHandler(
		mockEmployeeDataSource,
		mockAvatarProvider,
	)

	mockEmployeeDataSource.EXPECT().
		Add(entities.Employee{
			Name:  "John Smith",
			Email: "john.smith@example.com",
		}).
		Return(42, nil)

	mockAvatarProvider.EXPECT().
		GetAvatarUrlByEmail("john.smith@example.com").
		Return("http://example.com/john_smith.jpg", nil)

	want := entities.Employee{
		Id:        42,
		Name:      "John Smith",
		Email:     "john.smith@example.com",
		AvatarUrl: "http://example.com/john_smith.jpg",
	}

	got, _ := addEmployeeCommandHandler.Handle(AddEmployeeCommand{
		Name:  "John Smith",
		Email: "john.smith@example.com",
	})

	assert.Equal(t, want, *got)
}
