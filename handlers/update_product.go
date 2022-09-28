package handlers

// import (
// 	"github.com/go-openapi/runtime/middleware"

// 	runtime "obaid/Ecom"
// 	"obaid/gen/restapi/operations"
// 	"obaid/models"
// )

// // NewUpdateProduct function is for edit Product.
// func NewUpdateProduct(rt *runtime.Runtime) operations.UpdateProductHandler {
// 	return &updateProduct{rt: rt}
// }

// type updateProduct struct {
// 	rt *runtime.Runtime
// }

// // Handle the put Product request.
// func (d *updateProduct) Handle(params operations.UpdateProductParams) middleware.Responder {
// 	_, err := d.rt.Service().GetProduct(params.ID)
// 	if err != nil {
// 		return operations.NewGetProductNotFound()
// 	}
// 	product := models.Product{
// 		Name:        params.Product.Name,
// 		Price:       params.Product.Price,
// 		Description: params.Product.Description,
// 		Password:    params.Product.Password,
// 		CreatedBy:   params.Product.CreatedBy,
// 		CreatedAt:   params.Product.CreatedAt,
// 		UpdatedBy:   params.Product.UpdatedBy,
// 		UpdatedAt:   params.Product.UpdatedAt,
// 	}
// 	if err := d.rt.Service().UpdateProduct(params.ID, &product); err != nil {
// 		return operations.NewUpdateProductInternalServerError()
// 	}

// 	return operations.NewUpdateProductOK()
// }