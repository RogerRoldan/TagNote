package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/roger/workhub/controllers"
)

func Init() {

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500", "http://localhost:5500"}
	router.Use(cors.New(config))

	RoutesApiUser(router)
	RoutesApiTask(router)
	RoutesApiGroup(router)
	RoutesApiGroupUser(router)
	RouterApiAuth(router)
	RouterApiTaskUser(router)

	//Configurar el servidor para HTTPS
	router.RunTLS(":8085", "cert.pem", "key.pem")

	router.Run(":80")
	router.RunTLS(":443", "cert.pem", "key.pem")

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
	task.GET("/get-by-group/:id", controllers.GetTasksByGroupId)
}

func RoutesApiGroup(e *gin.Engine) {
	group := e.Group("/api/group")
	group.GET("/get-all", controllers.GetGroups)
	group.GET("/get/:id", controllers.GetGroupById)
	group.POST("/create", controllers.CreateGroup)
	group.PUT("/update", controllers.UpdateGroup)
	group.DELETE("/delete/:id", controllers.DeleteGroup)
	group.GET("/get-by-admin/:id", controllers.GetGroupsByAdminId)
}

func RoutesApiGroupUser(e *gin.Engine) {
	group := e.Group("/api/group-user")
	group.GET("/get-all", controllers.GetGroupUsers)
	group.GET("/get/:id", controllers.GetGroupUserById)
	group.POST("/create", controllers.CreateGroupUser)
	group.PUT("/update", controllers.UpdateGroupUser)
	group.DELETE("/delete", controllers.DeleteGroupUser)
	group.GET("/get-by-user/:id", controllers.GetGroupUsersByUserId)
	group.GET("/get-by-group/:id", controllers.GetGroupUsersByGroupId)
}

func RouterApiAuth(e *gin.Engine) {
	auth := e.Group("/api/auth")
	auth.POST("/login", controllers.Login)
}

func RouterApiTaskUser(e *gin.Engine) {
	taskUser := e.Group("/api/task-user")
	taskUser.GET("/get-all", controllers.GetTaskUsers)
	taskUser.GET("/get/:id", controllers.GetTaskUserById)
	taskUser.POST("/create", controllers.CreateTaskUser)
	taskUser.PUT("/update", controllers.UpdateTaskUser)
	taskUser.DELETE("/delete", controllers.DeleteTaskUser)
}
