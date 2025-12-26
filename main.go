package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetRouter(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("Hello World"))

}

func main(){
	port:=os.Getenv("PORT")
	if port==""{
		port="8080"
	}

	http.HandleFunc("/",GetRouter)
	err:=http.ListenAndServe(fmt.Sprintf(":%s",port),nil)
	if err!=nil{
		log.Println("Error starting the server ",err)
		return
	}

}