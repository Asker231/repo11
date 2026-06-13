package utils

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(pass []byte)(string, error){
	hashCode,err := bcrypt.GenerateFromPassword(pass,bcrypt.DefaultCost)
	if err != nil{
		return "",err
	}	
	return string(hashCode), nil 
}

func ComparePassword(hashed []byte,password []byte)error{
	return bcrypt.CompareHashAndPassword(hashed,password)
}