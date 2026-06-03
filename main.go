package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)


const ACCESS = 0644

func main() {
	app := http.NewServeMux()
	server := http.Server{
		Addr: ":8081",
		Handler: app,
	}
	_  = NewService(app)
	server.ListenAndServe()
}

type Service struct{
	app *http.ServeMux
}

func NewService(app *http.ServeMux)*Service{
	u := &Service{app: app}

	u.app.Handle("POST /msg",u.RemoveWord())

	return u
  	
}

type Msg struct{
	Msg string `json:"msg"`
}

func(s *Service)RemoveWord() http.HandlerFunc{
	var msg Msg
	return  func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil{
			fmt.Println(err.Error())
		}
		res := strings.Map(func(r rune) rune {
			if r == 'a'{
				return  -1
			}
			return  r
		},msg.Msg)
		fmt.Println(res)
	}
}


func SlidingWindow(){
   k:= 3	
   arr :=[]int{1,2,3,4,5,6,7}
   window := 0
   
   for i:=0; i < k; i++{
	window += arr[i] 
   } 

   max := window

   for i := k; i < len(arr);i++{
	window = window - arr[i-k]
	window = window + arr[i]
	if window > max{
		max =window
	}	 
   }
   fmt.Println(max)
}

func filter[T comparable](arr[]T,predicat func(T)bool)[]T{
	result := []T{}
	for _,v := range arr{
		res := predicat(v)
		if res == true{
			result = append(result, v)
		}
	}
	return  result
}

func isPolindrom()(bool){
	str := "lo1l"
	rstr := []rune(str)
	
	start := 0
	end := len(rstr)-1

	for start <= end{
		if rstr[start] != rstr[end]{
			return  false
		}
		start+=1
		end-=1
	}
	return  true
}


func formatRate(rate string)(string){
	var builder strings.Builder
	numberFormat ,err:= strconv.ParseFloat(rate,64)
	if err != nil{
		fmt.Println("nan")
	}
	stingFormat := strconv.FormatFloat(numberFormat,'f',5, 64)
    parts := strings.Split(stingFormat, ".")
	first := parts[0]
	second := parts[1]
	
	sec:= strings.TrimRight(second,"0")

	builder.WriteString(first)
	builder.WriteString(".")
	builder.WriteString(sec)

	result := builder.String()

	return result
}