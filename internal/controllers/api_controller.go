package controllers

import (
	"compClub/internal/db"
	"compClub/internal/redis"
	"compClub/internal/rent"
	"compClub/internal/util/isAdmin"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

// 97c94ebe5d767a353b77f3c0ce2d429741f2e8c99473c3c150e2faa3d14c9da6
func (c *Controller) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	if !isAdmin.Check(r) {
		w.WriteHeader(http.StatusForbidden)
	} else {
		admin := db.Admin{}
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &admin)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		passwordHashed := fmt.Sprintf("%x", sha256.Sum256([]byte(admin.Password)))
		admin.Password = passwordHashed

		insert := db.Db.Create(&admin)

		if insert.Error != nil {
			log.Default()
		}
	}
}

func (c *Controller) AddNewUser(w http.ResponseWriter, r *http.Request) {
	if !isAdmin.Check(r) {
		w.WriteHeader(http.StatusForbidden)
	} else {
		user := db.User{}
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		passwordHashed := fmt.Sprintf("%x", sha256.Sum256([]byte(user.Password)))
		user.Password = passwordHashed

		insert := db.Db.Create(&user)

		if insert.Error != nil {
			log.Default()
		}
	}
}

func (c *Controller) AddNewPC(w http.ResponseWriter, r *http.Request) {
	if !isAdmin.Check(r) {
		w.WriteHeader(http.StatusForbidden)
	} else {
		pc := db.Computer{}
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &pc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		redis.AddPcInfo(pc)
		insert := db.Db.Create(&pc)
		if insert.Error != nil {
			log.Default()
		}
	}
}

/*
{
    "name": "Aorta tvoey babushki 2.0",
    "price": 301,
    "gpu": "RTX 2080",
    "cpu": "Intel Core I5 9700k",
    "ram": "16 gb"
}
*/

func (c *Controller) AddNewShift(w http.ResponseWriter, r *http.Request) {
	if !isAdmin.Check(r) {
		w.WriteHeader(http.StatusForbidden)
	} else {
		shift := db.Shift{}
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &shift)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		insert := db.Db.Create(&shift)
		if insert.Error != nil {
			log.Default()
		}
	}
}

func (c *Controller) GetAvailablePc(w http.ResponseWriter, r *http.Request) {
	body := make(map[string]db.Computer)
	keys, _ := redis.Rdb.Keys(redis.Ctx, "*").Result()
	for _, key := range keys {
		val, _ := redis.Rdb.Get(redis.Ctx, key).Bytes()
		var tempPc db.Computer
		_ = json.Unmarshal(val, &tempPc)
		body[key] = tempPc
	}
	bytes, _ := json.Marshal(body)
	w.Write([]byte(bytes))
}

func (c *Controller) NewRent(w http.ResponseWriter, r *http.Request) {
	if !isAdmin.Check(r) {
		w.WriteHeader(http.StatusForbidden)
	} else {
		newRent := db.Rent{}
		body, _ := io.ReadAll(r.Body)
		err := json.Unmarshal(body, &newRent)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		fmt.Println("pcinfo -")
		fmt.Println(redis.GetPcInfo(newRent.ComputerID))
		fmt.Println("rent started. ComputerID -" + fmt.Sprint(newRent.ComputerID))
		go rent.RentInit(newRent)
	}
}

//getavailablepc
