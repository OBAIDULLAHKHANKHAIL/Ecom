package handlers

// import (
// 	"github.com/go-openapi/runtime/middleware"

// 	runtime "obaid/Ecom"
// 	"obaid/gen/restapi/operations"
// 	"obaid/models"
// )

// // NewAddUser handles request for saving user.
// func NewAddUser(rt *runtime.Runtime) operations.AddUserHandler {
// 	return &addUser{rt: rt}
// }

// type addUser struct {
// 	rt *runtime.Runtime
// }

// // Handle the add student request.
// func (d *add) Handle(params operations.AddUserParams) middleware.Responder {
// 	user := models.User{
// 		Name:  params.User.Name,
// 		User:  params.User.User,
// 		Password: params.User.Password,
// 	}

// 	id, err := d.rt.Service().AddUser(&user)
// 	if err != nil {
// 		log().Errorf("failed to add user: %s", err)
// 		return operations.NewAddUserConflict()
// 	}


// 	params.Student.ID = id
// 	return operations.NewAddUserCreated().WithPayload(params.User)
// }