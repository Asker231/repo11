package utils

import (
	"encoding/json"
	"fmt"
	"io"
)

//Обобщенная функция 
func  ReadFile[T any](r io.Reader)T{
	var payload T
	if err := json.NewDecoder(r).Decode(&payload); err !=nil{
		fmt.Println(err.Error())
	}
	return  payload
}