package controllers

import (
	"strconv"

	"github.com/AnggaDanarP/Online-Learning-Plathform/database"
	"github.com/AnggaDanarP/Online-Learning-Plathform/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCourse(c *fiber.Ctx) error {
	var data map[string]interface{}
	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	course := models.Course{
		Name: data["name"].(string),
		Describe: data["describe"].(string),
		Category: data["category"].(string),
		Price: data["price"].(string),
		Favorites: data["favorites"].(string),
		
	}

	errCreate := database.DB.Create(&course).Error
	if errCreate != nil {
		return c.JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(course)
}

func GetAllCourse(c *fiber.Ctx) error {
	var courses []models.Course

	err := database.DB.Find(&courses).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(courses)
}

func UpdateCourse(c *fiber.Ctx) error {
	RequestUpdate := new(models.RequestUpdateCourse)
	if err := c.BodyParser(RequestUpdate); err != nil {
		return c.JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	var course models.Course

	database.DB.Where("id = ?", id).First(&course)

	if course.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "course not found",
		})
	}

	var data map[string]interface{}
	err = c.BodyParser(&data)
	if err != nil {
		return err
	}

	if RequestUpdate.Name != "" {
		course.Name = RequestUpdate.Name
	}
	if RequestUpdate.Describe != "" {
		course.Describe = RequestUpdate.Describe
	}
	if RequestUpdate.Category != "" {
		course.Category = RequestUpdate.Category
	}
	if RequestUpdate.Price!= "" {
		course.Price = RequestUpdate.Price
	}

	errUpdate := database.DB.Save(&course).Error
	if errUpdate != nil {
		return c.JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "course updated",
	})
}

func DeleteCourse(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	var course models.Course

	database.DB.Where("id = ?", id).First(&course)

	if course.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "course not found",
		})
	}

	errDelete := database.DB.Delete(&course).Error
	if errDelete != nil {
		return c.JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "course deleted",
	})
}

func GetAllUser(c *fiber.Ctx) error {
	// get all user that deleted_at is null
	// if null will not show in the response
	var users []models.User

	database.DB.Where("deleted_at IS NULL").Find(&users)

	return c.JSON(users)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return c.JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "user deleted",
	})
}

func Statistic(c *fiber.Ctx) error {
	var users []models.User
	var courses []models.Course
	var free []models.Course

	database.DB.Find(&users)
	database.DB.Find(&courses)
	database.DB.Where("price = ?", "free").Find(&free)

	return c.JSON(fiber.Map{
		"total_user": len(users),
		"total_course": len(courses),
		"total_free_course": len(free),
	})
}
