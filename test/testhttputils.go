package main

import (
	"awesomeProject/utils"
	"fmt"
)

func main() {
	client1 := utils.NewBrowser()
	params := make(map[string]string)
	params["app_id"] = "378"
	params["channel_userid"] = ""
	params["juhe_token"] = "8e0616b10461c9698a701396c105f6c4"
	params["juhe_userid"] = "68147086"
	url := "https://juhesdk.3975ad.com/api/oauth/verify_token"
	fmt.Println(url)
	fmt.Println(params)
	dataBytes := client1.HttpPostJson(url, params)
	result := string(dataBytes)
	fmt.Println(result)
}
