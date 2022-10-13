package handlers

// import (
// 	"github.com/go-openapi/runtime/middleware"

// 	runtime "obaid/Ecom"
// 	domainErr "obaid/errors"
// 	"obaid/models"
// 	"obaid/gen/restapi/operations"
// )

// // NewGetUser handles a request for retrieving user.
// func NewGetUser(rt *runtime.Runtime) operations.GetUserHandler {
// 	return &getUser{rt: rt}
// }

// type getUser struct {
// 	rt *runtime.Runtime
// }

// // Handle the get user request.
// func (d *getUser) Handle(params operations.GetUserParams) middleware.Responder {
// 	user, err := d.rt.Service().GetUser(params.ID)
// 	if err != nil {
// 		switch apiErr := err.(*domainErr.APIError); {
// 		case apiErr.IsError(domainErr.NotFound):
// 			return operations.NewGetUserNotFound()
// 		default:
// 			return operations.NewGetUserInternalServerError()
// 		}
// 	}

// 	return operations.NewGetUserOK().WithPayload(&models.User{
// 		Name:  user.Name,
// 		User:  user.User,
// 		Password: user.Password,
// 	})
// }