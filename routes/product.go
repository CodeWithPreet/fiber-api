package routes

import (
	"github.com/CodeWithPreet/fiber-api/database"
	"github.com/CodeWithPreet/fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Product struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	SerialNo string `json:"serial_no"`
}


func CreateResponseProduct(productModel models.Product)  Product {
	return Product{
		ID: productModel.ID,
		Name: productModel.Name,
		SerialNo: productModel.SerialNo,
	}
	
}

func createProduct(c *fiber.Ctx) error {
	var product models.Product
	err:= c.BodyParser(&product)
	if err != nil {
		return c.Status(400).JSON(err.Error())
		
	}
	product.ID =uuid.New()
	database.DBI.Db.Create(&product)
	resProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(resProduct)
}
func listProducts(c*fiber.Ctx)error  {
	var products []models.Product
	err := database.DBI.Db.Find(&products).Error
	if err !=nil{
		return c.Status(400).JSON(err.Error())
	}



	resProducts := []Product{}
	for _, product := range products{
		resProducts = append(resProducts, CreateResponseProduct(product))
	}
	return c.Status(200).JSON(resProducts)
}


func getProduct(id uuid.UUID, product *models.Product) error {
	database.DBI.Db.Find(&product, "id = ?", id)
	if product.ID == uuid.Nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}
	return nil
}




func getProductById(c*fiber.Ctx)error{
	id := c.Params("id")
	var product models.Product

	uid ,err := uuid.Parse(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
	}
	if err := getProduct(uid, &product) ;err!=nil{
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	
	return c.Status(200).JSON(CreateResponseProduct(product))
}


func  deleteProduct(c* fiber.Ctx) error {
	var product models.Product
	id := c.Params("id")
	uid ,err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"}) 
	}
	if err := getProduct(uid ,&product); err!= nil{
		return c.Status(404).JSON(fiber.Map{"message": "Product not found","error":err.Error()})
	}
	
	if err := database.DBI.Db.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Product not found", "error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{
        "message": "successfully deleted",
        "code":    200,
        "product": CreateResponseProduct(product),
    })
}

func updateProduct(c *fiber.Ctx)  error{
	var product models.Product
	id := c.Params("id")
	uid ,err :=  uuid.Parse(id)
	if err != nil{
		return c.Status(400).JSON(fiber.Map{"error": "Invalid UUID"})
	}
	if err := getProduct(uid, &product);err!= nil{
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	type UpdateProduct struct{
		Name string `json:"name"`
		SerialNo string `json:"serial_no"`
	}
	var updateData UpdateProduct
	if err:= c.BodyParser(&updateData) ;err!=nil{
		return c.Status(500).JSON(err.Error())
	}
	product.Name = updateData.Name
	product.SerialNo=updateData.SerialNo
	database.DBI.Db.Save(&product)

	return c.Status(200).JSON(CreateResponseProduct(product))
}





func ProductControllers(app *fiber.App)  {
	productRoute := app.Group("api/products")
	productRoute.Post("/" ,createProduct)
	productRoute.Get("/",listProducts)
	productRoute.Get("/:id",getProductById)
	productRoute.Delete("/:id",deleteProduct)
	productRoute.Put("/:id",updateProduct)
	
}