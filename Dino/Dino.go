package main

import (
	web "EndToEnd/Dino/Portal"
	"encoding/json"
	"fmt"
	"os"
)

type config struct{

	WebServer string `json:"webserver"`
}


func main() {

	file,err:=os.Open("Dino/config.json")
    if err!=nil{
    	fmt.Println("file reading in error")
        os.Exit(1)
	}

	conf:=new(config)
	err=json.NewDecoder(file).Decode(conf)

	err=web.RunWebPortal(conf.WebServer)
	if err!=nil{
		fmt.Println("Error occured")
	}
}
