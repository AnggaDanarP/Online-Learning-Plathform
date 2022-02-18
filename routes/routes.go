package routes

import (
	"github.com/AnggaDanarP/Online-Learning-Plathform/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user/getAllCourseByCategory", controllers.User, controllers.GetAllCategory)
	app.Get("/api/user/getAllFavorites", controllers.User, controllers.GetFavorite)
	app.Get("/api/user/searchCourse", controllers.User, controllers.GetSearchCourse)
	app.Get("/api/user/getCourseFree", controllers.User, controllers.GetCourseFree)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/admin/course/create", controllers.CreateCourse)
	app.Get("/api/admin/course/getAllCourse", controllers.GetAllCourse)
	app.Put("/api/admin/course/update/:id", controllers.UpdateCourse)
	app.Delete("/api/admin/course/delete/:id", controllers.DeleteCourse)
	app.Get("/api/admin/user/getAllUser", controllers.GetAllUser)
	app.Delete("/api/admin/user/deleteUser/:id", controllers.DeleteUser)
	app.Get("/api/admin/statistic", controllers.Statistic)

}
