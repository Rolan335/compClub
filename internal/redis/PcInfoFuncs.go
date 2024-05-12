package redis

import (
	"compClub/internal/db"
	"encoding/json"
	"fmt"
)

func InitPcInfo() {
	var pc []db.Computer
	result := db.Db.Find(&pc)
	if result.Error != nil{
		fmt.Println(result.Error)
	}
	fmt.Println(pc)
	for _, v := range pc {
		bytes, _ := json.Marshal(v)
		err := Rdb.Set(Ctx, v.Name, bytes,0).Err()
		if err != nil{
			fmt.Println(err)
		}
	}
}

func AddPcInfo(pc db.Computer){
	bytes, _ := json.Marshal(pc)
	err := Rdb.Set(Ctx, pc.Name, bytes,0).Err()
		if err != nil{
			fmt.Println(err)
		}
}

func GetPcInfo(id int) db.Computer{
	var pc db.Computer
	val, _ := Rdb.Get(Ctx, fmt.Sprint(id)).Bytes()
	_ = json.Unmarshal(val,&pc)
	return pc
}
