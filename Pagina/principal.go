package main

import(
	"log"
	"net/http"
)

func main(){
	fs:=http.FileServer(http.Dir("./static"))
	http.Handle("/",fs)

	log.Println("Escuchando el puerto 4000")

	err:=http.ListenAndServe(":4000",nil)

	if err!=nil{
		log.Fatal(err);
	}
}