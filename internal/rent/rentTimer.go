package rent

import (
	"compClub/internal/db"
	"time"
)

func Timer(hours int, rent db.Rent) {
	//поменяй секунды на часы
	timer := time.NewTimer(time.Duration(hours) * time.Second)
	<-timer.C
}
