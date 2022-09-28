package handlers

// import (
// 	"github.com/go-openapi/runtime/middleware"

// 	runtime "obaid/Ecom"
// 	domainErr "obaid/errors"
// 	"obaid/models"
// 	"obaid/gen/restapi/operations"
// )

// // NewGetProduct handles a request for retrieving Product.
// func NewGetProduct(rt *runtime.Runtime) operations.GetProductHandler {
// 	return &getProduct{rt: rt}
// }

// type getProduct struct {
// 	rt *runtime.Runtime
// }

// // Handle the get Product request.
// func (d *getProduct) Handle(params operations.GetProductParams) middleware.Responder {
// 	product, err := d.rt.Service().GetProduct(params.ID)
// 	if err != nil {
// 		switch apiErr := err.(*domainErr.APIError); {
// 		case apiErr.IsError(domainErr.NotFound):
// 			return operations.NewGetProductNotFound()
// 		default:
// 			return operations.NewGetProductInternalServerError()
// 		}
// 	}

// 	return operations.NewGetProductOK().WithPayload(&models.Product{
// 		Name:        product.Name,
// 		Price:       product.Price,
// 		Description: product.Description,
// 		Password:    product.Password,
// 		CreatedBy:   product.CreatedBy,
// 		CreatedAt:   product.CreatedAt,
// 		UpdatedBy:   product.UpdatedBy,
// 		UpdatedAt:   product.UpdatedAt,
// 	})
// }