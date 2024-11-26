package config

import (
	"carthage/services/gateway/handlers"
	"carthage/services/gateway/routes"
)

func NewConf() map[string]routes.HandlerFunc {
	bootcampSvc := handlers.BootcampHandler()
	userSvc := handlers.UserHandler()

	HandlersMap := map[string]routes.HandlerFunc{
		// Bootcamp Service
		"GetBootcamps": bootcampSvc.GetBootcamps(),
		// User Service
		"GetUsers": userSvc.GetUsers(),
		"GetUser":  userSvc.GetUser(),
	}

	return HandlersMap
}
