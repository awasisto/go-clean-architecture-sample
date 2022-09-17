package routers

import (
	"github.com/gorilla/mux"
	"go-clean-architecture-sample/api/controllers"
	"net/http"
)

type EmployeeRouter struct {
	MuxRouter  mux.Router
	controller controllers.EmployeeController
}

func NewEmployeeRouter(
	controller controllers.EmployeeController,
) *EmployeeRouter {
	employeeRouter := EmployeeRouter{
		MuxRouter:  *mux.NewRouter(),
		controller: controller,
	}

	employeeRouter.MuxRouter.
		Path("/api/employees").
		Methods(http.MethodPost).
		HandlerFunc(controller.AddEmployee)

	employeeRouter.MuxRouter.
		Path("/api/employees").
		Methods(http.MethodGet).
		HandlerFunc(controller.GetAllEmployees)

	employeeRouter.MuxRouter.
		Path("/api/employees/{employee_id}").
		Methods(http.MethodGet).
		HandlerFunc(controller.GetEmployeeById)

	return &employeeRouter
}
