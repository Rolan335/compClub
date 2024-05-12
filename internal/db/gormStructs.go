package db

import (
	"time"

	"gorm.io/gorm"
)
type Computer struct {
    gorm.Model
    ID    uint
	Name  string  `json:"name"`
    Price int    `json:"price"` 
    GPU   string `json:"gpu"`
    CPU   string `json:"cpu"`
    RAM   string `json:"ram"`
}

type User struct {
    gorm.Model
    ID       uint
    Login    string `json:"login" gorm:"unique"`
    Password string `json:"password"`
    Phone    string `json:"phone"`
    Balance  int    `json:"balance"`
}

type Admin struct {
    gorm.Model
    ID       uint
    Surname  string `json:"surname"`
    Name     string `json:"name"`
    Phone    string `json:"phone"`
    Passport string `json:"passport" gorm:"type:varchar(10)"`
    ITN      string `json:"itn" gorm:"type:varchar(12)"`
    Password string `json:"password"`
}

/*
{
    "surname":"Test",
    "name":"Admin",
    "phone":"+79998883322",
    "passport":"9080742322",
    "itn":"111222333444",
    "password":"12345678"
}
*/

type Shift struct {
    gorm.Model
    ID       uint   
    Start    time.Time `json:"start"`
    End      time.Time `json:"end"`
    AdminID  int	   `json:"admin_id"`
    Admin    Admin
    Profit   int       `json:"profit"`
}

/*
{
    "start": "2023-10-31T10:00:00Z",
    "end":"2023-10-31T18:00:00Z",
    "admin_id": 8,
    "profit": 1000
}
*/

type Rent struct {
    gorm.Model
    ID         uint   
    Start      time.Time `json:"start"`
    End        time.Time `json:"end"`
    AdminID    int       `json:"admin_id"`
    Admin      Admin
    ComputerID int       `json:"computer_id"`
    Computer   Computer
}

