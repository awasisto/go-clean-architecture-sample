package usecases

import (
	"github.com/golang/mock/gomock"
	"golang-clean-architecture-sample/pkg/core/data/mocks"
	"golang-clean-architecture-sample/pkg/core/entities"
	"testing"
)

func TestGetEmployeeByIdUseCase_Execute(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	mockEmployeeDataSource := mocks.NewMockEmployeeDataSource(ctrl)
	mockAvatarProvider := mocks.NewMockAvatarProvider(ctrl)

	getEmployeeByIdUseCase := NewGetEmployeeByIdUseCase(
		42,
		mockEmployeeDataSource,
		mockAvatarProvider,
	)

	mockEmployeeDataSource.EXPECT().
		GetById(int64(42)).
		Return(
			&entities.Employee{
				Id:    42,
				Name:  "John Smith",
				Email: "john.smith@example.com",
			},
			nil,
		)

	mockAvatarProvider.EXPECT().
		GetAvatarUrlByEmail("john.smith@example.com").
		Return("http://example.com/john_smith.jpg", nil)

	want := entities.Employee{
		Id:        42,
		Name:      "John Smith",
		Email:     "john.smith@example.com",
		AvatarUrl: "http://example.com/john_smith.jpg",
	}

	got, _ := getEmployeeByIdUseCase.Execute()

	if *got != want {
		t.Errorf("Execute() = %v; want %v", got, want)
	}
}
