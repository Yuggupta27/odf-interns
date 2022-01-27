package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Demo application to reduce the size of image")
})

	log.Fatal(http.ListenAndServe(":8081", nil))
}