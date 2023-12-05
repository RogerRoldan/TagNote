package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/roger/workhub/controllers"
)

func Init() {

	router := gin.Default()

	//add cors
	config := cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}
	router.Use(cors.New(config))

	RoutesApiUser(router)
	RoutesApiTask(router)
	RoutesApiGroup(router)

	//Start server
	router.Run(":8085")
}

func RoutesApiUser(e *gin.Engine) {
	user := e.Group("/api/user")
	user.GET("/get-all", controllers.GetUsers)
	user.GET("/get/:id", controllers.GetUserById)
	user.POST("/create", controllers.CreateUser)
	user.PUT("/update", controllers.UpdateUser)
	user.DELETE("/delete", controllers.DeleteUser)
}

func RoutesApiTask(e *gin.Engine) {
	task := e.Group("/api/task")
	task.GET("/get-all", controllers.GetTasks)
	task.GET("/get/:id", controllers.GetTaskById)
	task.POST("/create", controllers.CreateTask)
	task.PUT("/update", controllers.UpdateTask)
	task.DELETE("/delete", controllers.DeleteTask)
}

func RoutesApiGroup(e *gin.Engine) {
	group := e.Group("/api/group")
	group.GET("/get-all", controllers.GetGroups)
	group.GET("/get/:id", controllers.GetGroupById)
	group.POST("/create", controllers.CreateGroup)
	group.PUT("/update", controllers.UpdateGroup)
	group.DELETE("/delete", controllers.DeleteGroup)
}
