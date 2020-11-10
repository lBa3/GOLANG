
package main

import ("fmt");
import("net/http");
import("log");
import ("github.com/gorilla/mux");

func main (){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index);
    router.HandleFunc("/contacto", Contacto);


   fmt.Println ("Servidor iniciado en http://localhost:9090");
   server := http.ListenAndServe(":9090", router)
   log.Fatal(server);
}

func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"Hola muno server GO");
}

func Contacto(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w,"Esto es contacto");
}
