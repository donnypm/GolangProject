package main

import (
	"net/http"
	"golangproject/service"
	"golangproject/controller"

	"github.com/gin-gonic/gin"
)

var (
	userService service.UserService =  service.New()
	userController controller.UserController = controller.New(userService)
)

func main()  {
	r := gin.Default()

	r.LoadHTMLGlob("template/*.html")

	r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })

  r.GET("/formuser", userController.FormUser)
  r.GET("/formlogin", userController.FormLogin)
  r.GET("/dashboard", userController.Dashboard)
  r.GET("/daftarabsen", userController.DaftarAbsen)
  r.GET("/daftarlogbook", userController.DaftarLogbook)
  r.GET("/daftarkalender", userController.DaftarKalender)
  r.GET("/daftarcuti", userController.DaftarCuti)
  r.GET("/daftaruser", userController.ShowDaftarUser)

  // create
	r.POST("/create", func(ctx *gin.Context) {
		userController.SaveToDB(ctx)
		ctx.Redirect(http.StatusFound, "/formuser")
	})


// jalanin server
  r.Run()
}