package handlers

// import (
// 	"github.com/go-openapi/runtime/middleware"
// 	"github.com/go-openapi/swag"

// 	runtime "obaid/Ecom"
// 	domainErr "obaid/errors"
// 	"obaid/gen/models"
// 	"obaid/gen/restapi/operations"
// )

// // NewListProduct handles a request for retrieving Products.
// func NewListProduct(rt *runtime.Runtime) operations.ListProductsHandler {
// 	return &listProducts{rt: rt}
// }

// type listProducts struct {
// 	rt *runtime.Runtime
// }

// // Handle the list employees request.
// func (e *listProducts) Handle(params operations.ListProductsParams) middleware.Responder {
// 	filter := make(map[string]interface{})
// 	if params.Name!=nil {
// 		filter["name"] = swag.StringValue(params.Name)
// 	}
// 	if params.Age!=nil {
// 		filter["age"] = swag.Int64Value(params.Age)
// 	}
// 	if params.Level!=nil {
// 		filter["level"] = swag.StringValue(params.Level)
// 	}
// 	if params.Phone!=nil {
// 		filter["phone"] = swag.StringValue(params.Phone)
// 	}

// 	products, err := e.rt.Service().ListProducts(filter, swag.Int64Value(params.Limit), swag.Int64Value(params.Offset))
// 	if err != nil {
// 		switch apiErr := err.(*domainErr.APIError); {
// 		case apiErr.IsError(domainErr.NotFound):
// 			return operations.NewGetProductNotFound()
// 		default:
// 			return operations.NewUpdateProductInternalServerError()
// 		}
// 	}

	// stdList := make([]*models.Product, 0)
	// for _,product := range products {
	// 	stdList = append(stdList, toProductGen(product))
	// }

// 	return operations.NewListProductsOK().WithPayload(stdList)
// }