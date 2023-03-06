package main

import (
	"fmt"
	"strconv"

	"github.com/xluohome/phonedata"
)

func main() {

	var phone_num = 1300000
	for {
		if phone_num > 1999999 {
			break
		}
		phone_num++
		pr, err := phonedata.Find(strconv.Itoa(phone_num))
		if err == nil {
			fmt.Printf("%s,%s,%s,%s,%s,%s \n",
				pr.AreaZone, pr.CardType, pr.City, pr.PhoneNum, pr.Province, pr.ZipCode)
		}
	}
}
