
package main

import ("fmt");
import ("net/http");
import ("log");


func main (){

   router := NewRouter();

   fmt.Println ("Servidor iniciado en http://localhost:9090");
   server := http.ListenAndServe(":9090", router)
   log.Fatal(server);
}


