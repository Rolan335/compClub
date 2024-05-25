package rent

import (
	"compClub/internal/db"
	"compClub/internal/redis"
	"fmt"

	nats "github.com/nats-io/nats.go"
)

func RentInit(rent db.Rent) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}
	defer nc.Close()

	nc.Subscribe("rent_expiry",func(m *nats.Msg){
		fmt.Println("nats returned")
		fmt.Println(m)
	})

	//new row in table rent
	db.Db.Create(&rent)

	//get rented pc
	var pc db.Computer
	db.Db.Find(&pc, rent.ComputerID)

	//delete rented pc from available pc
	errRedis := redis.Rdb.Del(redis.Ctx, pc.Name)
	if errRedis != nil {
		fmt.Println(err)
	}

	//balance minus profit stonks
	rentHours := int(rent.End.Sub(rent.Start).Hours())
	profit := rentHours * pc.Price
	fmt.Println(profit)

	var shift db.Shift
	db.Db.Find(&shift, rent.ShiftID)
	shift.Profit += profit
	db.Db.Save(&shift)

	var user db.User
	db.Db.Find(&user, rent.UserID)
	user.Balance -= profit
	db.Db.Save(&user)

	//rent timer init
	Timer(rentHours, rent)
	nc.Publish("rent_expiry", []byte(fmt.Sprint(rent.ID)))
	fmt.Println("rent expired -" + fmt.Sprint(rent.ComputerID))
	//return pc to avialable for rent
	redis.AddPcInfo(pc)
}
