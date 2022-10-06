package service

import (
	"fmt"
	"golangproject/config"
	"golangproject/entity"

	"github.com/gin-gonic/gin"
	_ "github.com/denisenkom/go-mssqldb"
)

type UserService interface {
	SaveToDB(ctx *gin.Context)
	ShowDaftarUser()[]entity.User
}

type userService struct{
	users []entity.User
}

func New() UserService{
	return &userService{}
}

func (service *userService) SaveToDB(ctx *gin.Context) {
	db, err := config.GetDB()

	if err != nil{
		fmt.Print(err.Error())
	}

	username := ctx.Request.PostFormValue("username")
	password := ctx.Request.PostFormValue("password")
	nik := ctx.Request.PostFormValue("nik")
	nama := ctx.Request.PostFormValue("nama")
	tempatlahir := ctx.Request.PostFormValue("tempatlahir")
	tanggallahir := ctx.Request.PostFormValue("tanggallahir")
	alamat := ctx.Request.PostFormValue("alamat")
	role := ctx.Request.PostFormValue("role")

	stmt, err := db.Prepare("INSERT INTO users (username, password, nik, nama, tempatlahir, tanggallahir, alamat, role) VALUES (?,?,?,?,?,?,?,?)")

	if err != nil{
		fmt.Print(err.Error())
	}

	_, err = stmt.Exec(username, password, nik, nama, tempatlahir, tanggallahir, alamat, role)

	if err != nil{
		fmt.Print(err.Error())
	}
}

func (service *userService) ShowDaftarUser() []entity.User {
	db, err := config.GetDB()

	if err != nil{
		fmt.Print(err.Error())
	}

	rows, errorQuery := db.Query("SELECT id, nik, nama, role FROM users")

	if errorQuery != nil{
		fmt.Print(err.Error())
	}

	var(data entity.User)

	// clear dulu
	service.users = nil

	for rows.Next(){
		// scan berguna untuk copy isi row ke variabel data
		err = rows.Scan(&data.Id, &data.NIK, &data.Nama, &data.Role)
		service.users = append(service.users, data)

		if err != nil{
		fmt.Print(err.Error())
		}
	}

	return service.users
}

