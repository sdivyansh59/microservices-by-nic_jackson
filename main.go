package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Oops"))
			
			// line 19 is === 16 & 17
			http.Error(w,"Oops", http.StatusBadRequest)
			return 
		}

		log.Printf("Data %s ", d) // convert []bytes in string and print
		log.Println(d) // print []byte (which is d)

		// write back to user using responseWriter
		fmt.Fprintf(w, "hello %s ","Divyansh")

		
	})

	http.HandleFunc("/goodbye", func (w http.ResponseWriter, r *http.Request){
		log.Println("Goodbye World")	
	})

	http.ListenAndServe(":9090", nil)
}