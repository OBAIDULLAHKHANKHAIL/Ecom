package handlers

// import (
// 	"github.com/go-openapi/runtime/middleware"

// 	runtime "obaid/Ecom"
// 	"obaid/gen/restapi/operations"
// 	"obaid/models"
// )

// // NewUpdateUser function is for edit User.
// func NewUpdateUser(rt *runtime.Runtime) operations.UpdateUserHandler {
// 	return &updateUser{rt: rt}
// }

// type updateUser struct {
// 	rt *runtime.Runtime
// }

// // Handle the put User request.
// func (d *updateUser) Handle(params operations.UpdateUserParams) middleware.Responder {
// 	_, err := d.rt.Service().GetUser(params.ID)
// 	if err != nil {
// 		return operations.NewGetUserNotFound()
// 	}
// 	user := models.User{
// 		Name:  params.User.Name,
// 		User:  params.User.User,
// 		Password: params.User.Password,
// 	}
// 	if err := d.rt.Service().UpdateUser(params.ID, &user); err != nil {
// 		return operations.NewUpdateUserInternalServerError()
// 	}

// 	return operations.NewUpdateUserOK()
// }