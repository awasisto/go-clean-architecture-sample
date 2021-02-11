package routers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang-clean-architecture-sample/pkg/core"
	"golang-clean-architecture-sample/pkg/core/data"
	"golang-clean-architecture-sample/pkg/core/entities"
	"golang-clean-architecture-sample/pkg/core/usecases"
	"io/ioutil"
	"net/http"
	"strconv"
)

type EmployeeRouter struct {
	MuxRouter          mux.Router
	employeeDataSource data.EmployeeDataSource
	avatarProvider     data.AvatarProvider
}

type jsonableErrorResponse struct {
	Message string `json:"message"`
}

type jsonableEmployee struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
}

func NewEmployeeRouter(
	employeeDataSource data.EmployeeDataSource,
	avatarProvider data.AvatarProvider,
) EmployeeRouter {
	employeeRouter := EmployeeRouter{
		MuxRouter:          *mux.NewRouter(),
		employeeDataSource: employeeDataSource,
		avatarProvider:     avatarProvider,
	}

	employeeRouter.MuxRouter.
		Path("/employees").
		Methods(http.MethodPost).
		HandlerFunc(employeeRouter.addEmployeeHandler)

	employeeRouter.MuxRouter.
		Path("/employees").
		Methods(http.MethodGet).
		HandlerFunc(employeeRouter.getAllEmployeesHandler)

	employeeRouter.MuxRouter.
		Path("/employees/{employee_id}").
		Methods(http.MethodGet).
		HandlerFunc(employeeRouter.getEmployeeByIdHandler)

	return employeeRouter
}

func (e *EmployeeRouter) addEmployeeHandler(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	body, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusInternalServerError, jsonableErrorResponse{err.Error()})
		return
	}

	var requestObj jsonableEmployee
	err = json.Unmarshal(body, &requestObj)
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusInternalServerError, jsonableErrorResponse{err.Error()})
		return
	}

	addEmployeeUseCase := usecases.NewAddEmployeeUseCase(
		entities.Employee{
			Name:  requestObj.Name,
			Email: requestObj.Email,
		},
		e.employeeDataSource,
		e.avatarProvider,
	)

	createdEmployee, err := addEmployeeUseCase.Execute()
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusInternalServerError, jsonableErrorResponse{err.Error()})
		return
	}

	responseObj := toJsonableEmployee(*createdEmployee)

	writeJsonHttpResponse(httpResponseWriter, http.StatusCreated, responseObj)
}

func (e *EmployeeRouter) getAllEmployeesHandler(httpResponseWriter http.ResponseWriter, _ *http.Request) {
	getAllEmployeesUseCase := usecases.NewGetAllEmployeesUseCase(e.employeeDataSource, e.avatarProvider)

	employees, err := getAllEmployeesUseCase.Execute()
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusInternalServerError, jsonableErrorResponse{err.Error()})
		return
	}

	responseObj := make([]jsonableEmployee, 0)
	for _, employee := range employees {
		responseObj = append(responseObj, toJsonableEmployee(employee))
	}

	writeJsonHttpResponse(httpResponseWriter, http.StatusOK, responseObj)
}

func (e *EmployeeRouter) getEmployeeByIdHandler(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	strEmployeeId, employeeIdSpecified := mux.Vars(httpRequest)["employee_id"]
	if !employeeIdSpecified {
		writeJsonHttpResponse(httpResponseWriter, http.StatusBadRequest, jsonableErrorResponse{"employee_id not specified"})
		return
	}

	employeeId, err := strconv.ParseInt(strEmployeeId, 10, 64)
	if err != nil {
		writeJsonHttpResponse(httpResponseWriter, http.StatusBadRequest, jsonableErrorResponse{"invalid employee_id format"})
		return
	}

	getEmployeeByIdUseCase := usecases.NewGetEmployeeByIdUseCase(employeeId, e.employeeDataSource, e.avatarProvider)

	employee, err := getEmployeeByIdUseCase.Execute()
	if err != nil {
		var statusCode int
		if err == core.ErrNotFound {
			statusCode = http.StatusNotFound
		} else {
			statusCode = http.StatusInternalServerError
		}
		writeJsonHttpResponse(httpResponseWriter, statusCode, jsonableErrorResponse{err.Error()})
		return
	}

	responseObj := jsonableEmployee{
		Id:        employee.Id,
		Name:      employee.Name,
		Email:     employee.Email,
		AvatarUrl: employee.AvatarUrl,
	}

	writeJsonHttpResponse(httpResponseWriter, http.StatusOK, responseObj)
}

func toJsonableEmployee(employee entities.Employee) jsonableEmployee {
	return jsonableEmployee{
		Id:        employee.Id,
		Name:      employee.Name,
		Email:     employee.Email,
		AvatarUrl: employee.AvatarUrl,
	}
}

func writeJsonHttpResponse(httpResponseWriter http.ResponseWriter, statusCode int, responseObj interface{}) {
	responseJson, _ := json.Marshal(responseObj)
	httpResponseWriter.Header().Set("Content-Type", "application/json")
	httpResponseWriter.WriteHeader(statusCode)
	httpResponseWriter.Write(responseJson)
}
