package routes

import (
	"log"

	"github.com/CodeWithPreet/fiber-api/database"
	"github.com/CodeWithPreet/fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Order struct {
	ID uuid.UUID `json:"id"`
	Product Product `json:"product"`
	User User `json:"user"`
}

func CreateResponseOrder(order models.Order,product Product,user User) Order {
	return Order{ID: order.ID,Product: product,User: user}
}

func createOrder(c* fiber.Ctx)error  {
	var order models.Order

	if err:= c.BodyParser(&order);err!= nil{
		return c.Status(400).JSON(err.Error())
	
	}

	var product models.Product
	var user models.User

	if err:= getProduct(order.ProductID ,&product); err !=nil{
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}

	if err:= findUser(order.UserID,&user);err != nil{
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	// order = models.Order{
	// 	ID:        uuid.New(), // Function to generate a unique ID
	// 	ProductID: order.ProductID,
	// 	Product: product,
	// 	UserID:    order.UserID,
	// 	User: user,
	// }
	order.ID=uuid.New()
	database.DBI.Db.Create(&order)
	resProduct := CreateResponseProduct(product)
	resUser := CreateResponseUser(user)
	log.Printf("Product: %+v\n", resProduct)
	log.Printf("User: %+v\n", resUser)

	return c.Status(200).JSON(CreateResponseOrder(order, resProduct, resUser))
}

func listOrders(c *fiber.Ctx) error {
	var orders []models.Order
	
	err:=database.DBI.Db.Find(&orders).Error
	if err !=nil{
		return c.Status(400).JSON(err.Error())
	}
	resOrders := []Order{}

	for _ ,order := range orders{
		var product models.Product
		var user models.User

		if err := getProduct(order.ProductID, &product); err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
		}

		if err := findUser(order.UserID, &user); err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}

		resOrders = append(resOrders, CreateResponseOrder(order, CreateResponseProduct(product), CreateResponseUser(user)))
	}
	return c.Status(200).JSON(resOrders)
}


func getOrderById(c *fiber.Ctx) error{
	var order models.Order
	id := c.Params("id")
	uId ,err :=uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON("invalid UUID format")
	}
	
	if err := findOrder(uId, &order ); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	var user models.User

		if err := getProduct(order.ProductID, &product); err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
		}

		if err := findUser(order.UserID, &user); err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}

		resOrder := CreateResponseOrder(order, CreateResponseProduct(product), CreateResponseUser(user))
	

	 
	return c.Status(200).JSON(resOrder)
}



func findOrder(id uuid.UUID, order *models.Order) error {
	database.DBI.Db.First(order, "id = ?", id)
	if order.ID == uuid.Nil{
		return fiber.NewError(fiber.StatusNotFound, "Order not found")
	}
	return nil
}

func deleteOrder(c* fiber.Ctx)error  {
	id := c.Params("id")
	var order models.Order
	uid,err:= uuid.Parse(id)
	if err!=nil{
		return c.Status(400).JSON("invalid UUID format")
	}
	if err:= findOrder(uid,&order);err!=nil{
		return c.Status(404).JSON(err.Error())
	}
	if err:= database.DBI.Db.Delete(&order).Error;err!=nil{
		return c.Status(404).JSON(err.Error())
	}
	var product models.Product
	var user models.User

		if err := getProduct(order.ProductID, &product); err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
		}

		if err := findUser(order.UserID, &user); err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "User not found"})
		}

		resOrder := CreateResponseOrder(order, CreateResponseProduct(product), CreateResponseUser(user))
	


	return c.Status(200).JSON(fiber.Map{
		"message": "successfully deleted",
		"code":    200,
		"order":    resOrder,
	})
}


func OrderControllers(app *fiber.App)  {
	orderRoute := app.Group("api/orders")
	orderRoute.Post("/",createOrder)
	orderRoute.Get("/",listOrders)
	orderRoute.Get("/:id",getOrderById)
	orderRoute.Delete("/:id",deleteOrder)
}