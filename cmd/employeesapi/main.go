package main

import (
	"golang-clean-architecture-sample/pkg/adapters/db"
	"golang-clean-architecture-sample/pkg/adapters/githubavatarprovider"
	"golang-clean-architecture-sample/pkg/adapters/routers"
	"net/http"
)

func main() {
	employeeDataSource := db.NewInMemoryEmployeeDatabase()
	avatarProvider := githubavatarprovider.NewGithubAvatarProvider()
	employeeRouter := routers.NewEmployeeRouter(employeeDataSource, avatarProvider)
	http.ListenAndServe("0.0.0.0:8080", &employeeRouter.MuxRouter)
}
