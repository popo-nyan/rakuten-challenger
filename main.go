package main

import (
	"encoding/json"
	"fmt"
	"log"
	"rakuten-challenger/rc"
	"strings"
)

func main() {
	jsonData := strings.Replace(`{\"mask\":\"bebd\",\"key\":\"29\",\"seed\":4233684362}`, "\\", "", -1)

	var mdata rc.Mdata

	err := json.Unmarshal([]byte(jsonData), &mdata)
	if err != nil {
		log.Fatal(err)
	}
	hash := rc.SolvePow(mdata.Key, mdata.Seed, mdata.Mask)
	fmt.Println("hash:", hash)

}
