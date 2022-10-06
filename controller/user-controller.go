package controller

import (
	"golangproject/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController interface{
	FormUser(ctx *gin.Context)
	FormLogin(ctx *gin.Context)
	Dashboard(ctx *gin.Context)
	DaftarAbsen(ctx *gin.Context)
	DaftarLogbook(ctx *gin.Context)
	DaftarCuti(ctx *gin.Context)
	DaftarKalender(ctx *gin.Context)
	SaveToDB(ctx *gin.Context)
	ShowDaftarUser(ctx *gin.Context)
}

type controller struct {
	service service.UserService
}

func New(service service.UserService) UserController  {
	return &controller{
		service: service,
	}
}

// TAMPILAN
func (c *controller) FormUser(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "formuser.html", nil)
}

func (c *controller) FormLogin(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "formlogin.html", nil)
}

func (c *controller) Dashboard(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "dashboard.html", nil)
}
func (c *controller) DaftarAbsen(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "daftarabsen.html", nil)
}
func (c *controller) DaftarLogbook(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "daftarlogbook.html", nil)
}
func (c *controller) DaftarKalender(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "daftarkalender.html", nil)
}
func (c *controller) DaftarCuti(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK, "daftarcuti.html", nil)
}

// Data
func (c *controller) SaveToDB(ctx *gin.Context)  {
	c.service.SaveToDB(ctx)
}

func (c *controller)  ShowDaftarUser(ctx *gin.Context){
	users := c.service.ShowDaftarUser()
	data := gin.H{
		"title" : "User Page",
		"users" : users,
	}
	ctx.HTML(http.StatusOK, "daftaruser.html", data)
}