
package main

import ("fmt");
import("net/http");
import("log");

func main (){

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "HOLA MUNDO DESDE WEBSERVER CON GO");
	})

   fmt.Println ("Servidor iniciado en http://localhost:9090");
   server := http.ListenAndServe(":9090", nil)
   log.Fatal(server);

}
