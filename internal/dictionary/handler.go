package dictionary

import "net/http"


type Dictionary struct{
	app *http.ServeMux
}

func NewDictionary(app *http.ServeMux)*Dictionary{
	dictionary := &Dictionary{
		app: app,
	}

	return  dictionary
}