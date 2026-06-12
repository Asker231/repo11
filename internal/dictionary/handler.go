package dictionary

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)


type Dictionary struct{
	app *http.ServeMux
}

func NewDictionary(app *http.ServeMux)*Dictionary{
	dictionary := &Dictionary{
		app: app,
	}

	app.Handle("GET /dict",dictionary.getDictionary())

	return  dictionary
}


func(d *Dictionary)getDictionary() http.HandlerFunc{
	var dictionaryes Config
	file,err := os.Open("internal/dictionary/dictionary.json")
	if err != nil{
		fmt.Println(err.Error())
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&dictionaryes)
	if err != nil{
		fmt.Println(err.Error())
	}
	return  func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(dictionaryes)
		if err != nil{
			fmt.Println(err.Error())
		}
	}
}