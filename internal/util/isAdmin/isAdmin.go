package isAdmin

import (
	db "compClub/internal/db"
	"net/http"
)

func Check(r *http.Request) bool{
	//Только зареганный админ может добавлять других админов. МастерАДМИН добавляется напрямую через бд
	tempAdmin := db.Admin{}
	clientPassword := r.Header.Get("Authorization")
	db.Db.Where("password = ?", clientPassword).Find(&tempAdmin)
	if tempAdmin == (db.Admin{}){
		return false
	} else {return true}
}