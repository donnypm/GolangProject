package entity

type User struct{
	Id int
	Username string			`json:"username"`
	Password string			`json:"password"`
	NIK string 				`json:"nik"`
	Nama string				`json:"nama"`
	TempatLahir string		`json:"tempatlahir"`
	TanggalLahir string		`json:"tanggallahir"`
	Alamat string			`json:"alamat"`
	Role string				`json:"role"`
}