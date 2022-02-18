package controllers

import (
	"strconv"
	"time"

	"github.com/AnggaDanarP/Online-Learning-Plathform/database"
	"github.com/AnggaDanarP/Online-Learning-Plathform/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser((&data))
	if err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	creatAt := time.Now()

	user := models.User{
		Name:     data["name"],
		Email: 	  data["email"],
		Password: password,
		CreatedAt: creatAt,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser((&data))
	if err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "incorect password",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Login Success",
	})
}

// set the middleware
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorized",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.Next()
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}

func GetAllCategory(c *fiber.Ctx) error {
	var RequestCategory []string

	var course []models.Course

	err := database.DB.Select("category").Find(&course).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	for i := 0; i < len(course); i++ {
		RequestCategory = append(RequestCategory, course[i].Category)
	}

	return c.JSON(RequestCategory)
}

func GetFavorite(c *fiber.Ctx) error {
	var Output []string

	var course []models.Course

	err := database.DB.Select("name", "favorites").Find(&course).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	for i := 0; i < len(course); i++ {
		Output = append(Output, course[i].Name+" Likes: "+course[i].Favorites+" Peoples")
	}

	return c.JSON(Output)
}

func GetSearchCourse(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser((&data))
	if err != nil {
		return err
	}

	var course models.Course

	database.DB.Where("name = ?", data["name"]).First(&course)

	if course.Name == "" {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "course not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Found",
		"data": course,
	})
}

func GetCourseFree(c *fiber.Ctx) error {
	var course []models.Course

	database.DB.Where("price = ?", "free").Find(&course)

	return c.JSON(fiber.Map{
		"message": "Success",
		"data": course,
	})
}

//missing:
// Get Popular Category Course
// Get Course
// Get Detail Course
// Sort Lowest and Highest Price