package handlers

// import (
// 	"github.com/go-openapi/runtime/middleware"

// 	runtime "obaid/Ecom"
// 	domainErr "obaid/errors"
// 	"obaid/gen/restapi/operations"
// )

// // NewDeleteProduct function will delete the Product.
// func NewDeleteProduct(rt *runtime.Runtime) operations.DeleteProductHandler {
// 	return &deleteProduct{rt: rt}
// }

// type deleteProduct struct {
// 	rt *runtime.Runtime
// }

// // Handle the delete entry request.
// func (d *deleteProduct) Handle(params operations.DeleteProductParams) middleware.Responder {
// 	if err := d.rt.Service().DeleteProduct(params.ID); err != nil {
// 		switch apiErr := err.(*domainErr.APIError); {
// 		case apiErr.IsError(domainErr.NotFound):
// 			return operations.NewGetProductNotFound()
// 		default:
// 			return operations.NewDeleteProductInternalServerError()
// 		}
// 	}
// 	return operations.NewDeleteProductNoContent()
// }