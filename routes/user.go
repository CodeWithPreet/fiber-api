package routes

import (
	"errors"
	"fmt"

	"github.com/CodeWithPreet/fiber-api/database"
	"github.com/CodeWithPreet/fiber-api/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	
}

func createResponseUser(userModel models.User) User {
	return User{
		ID: userModel.ID,
		Name: userModel.Name,
	}
	
}

func createUser(c *fiber.Ctx) error {
	var user models.User
	if err:= c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	} 
	user.ID = uuid.New()
	database.DBI.Db.Create(&user)
	responseUser := createResponseUser(user)
	return c.Status(201).JSON(responseUser)
}

func listUsers(c *fiber.Ctx) error{
	
	// var users []models.User
	var users []models.User
	err := database.DBI.Db.Find(&users).Error
	if err != nil {
		fmt.Println(err.Error())
		return c.Status(400).JSON(err.Error())
	}	
	resUsers:= []User{}
	 for _ ,user := range users { 
		resUsers = append(resUsers, createResponseUser(user))
	 }
	return c.Status(200).JSON(resUsers)
}

func getUserbyId(c *fiber.Ctx) error  {
	id := c.Params("id")
	var user models.User
	
	uid, err := uuid.Parse(id)
	if err != nil {
		return c.Status(400).JSON("invalid UUID format")
	}
	
	if err := findUser(uid, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	
	return c.Status(200).JSON(createResponseUser(user))

	 
	
}
func findUser(id uuid.UUID, user *models.User) error {
	database.DBI.Db.Find(&user, "id =?", id)
	if user.ID == uuid.Nil {
		return errors.New("user does not exist")
	}
	return nil
}

// updateuser create

func updateUser(c *fiber.Ctx) error {
	id:=  c.Params("id")
	var user models.User

	uid,err := uuid.Parse(id)

	if err!=nil{
		return c.Status(400).JSON("invalid UUID format")
	}

	if err := findUser(uid,&user); err!= nil{
		return c.Status(404).JSON(err.Error())
	}

	type UpdateUser struct{
		Name string `json:"name"`
	}
	var updateData UpdateUser

	if err:= c.BodyParser(&updateData) ;err!=nil{
		return c.Status(500).JSON(err.Error())
	}
	user.Name =updateData.Name
	database.DBI.Db.Save(&user)

	return c.Status(200).JSON(createResponseUser(user))
	
}

func deleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	uid,err:= uuid.Parse(id)
	if err!=nil{
		return c.Status(400).JSON("invalid UUID format")
	}
	if err:= findUser(uid,&user);err!=nil{
		return c.Status(404).JSON(err.Error())
	}
	if err:= database.DBI.Db.Delete(&user).Error;err!=nil{
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "successfully deleted",
		"code":    200,
		"user":    user,
	})

}


func UserControllers(app *fiber.App){
	usersRoute := app.Group("api/users")
	usersRoute.Post("/",createUser)
	usersRoute.Get("/" , listUsers)
	usersRoute.Get("/:id", getUserbyId)
	usersRoute.Put("/:id",updateUser)
	usersRoute.Delete("/:id",deleteUser)
	 
	// users
}