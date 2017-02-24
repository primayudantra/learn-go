package main

import (
	"encoding/json"
	"fmt"
)

type DataNew struct{
	_id 			int `json:"_id"`
	Username 	string `json:"username"`
	Email 		string `json:"email"`
}

func main(){
	datax := &DataNew{Username:"primayudantra",Email:"prima.yudantra@gmail.com"}
	result, err := json.Marshal(datax)
	if err != nil {
        fmt.Println(err)
        return
    }
  fmt.Println(string(result))
}