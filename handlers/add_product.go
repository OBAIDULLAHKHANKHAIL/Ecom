package handlers

// import (
// 	"github.com/go-openapi/runtime/middleware"

// 	runtime "obaid/Ecom"
// 	"obaid/gen/restapi/operations"
// 	"obaid/models"
// )

// // NewAddUser handles request for saving product.
// func NewAddProduct(rt *runtime.Runtime) operations.AddProductHandler {
// 	return &addProduct{rt: rt}
// }

// type addProduct struct {
// 	rt *runtime.Runtime
// }

// // Handle the add product request.
// func (d *add) Handle(params operations.AddProductParams) middleware.Responder {
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

// 	id, err := d.rt.Service().AddProduct(&product)
// 	if err != nil {
// 		log().Errorf("failed to add product: %s", err)
// 		return operations.NewAddProductConflict()
// 	}


// 	params.Student.ID = id
// 	return operations.NewAddProductCreated().WithPayload(params.User)
// }