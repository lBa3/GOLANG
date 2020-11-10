package main 

import ("fmt")

func main(){
  for i := 0;  i < 5; i++{
    fmt.Println("HOLA",i);

  }

fmt.Println(":::::::::::::::::::::::::::::");
  
  //array con for
  peliculas := []string{"batman", "superman", "joker", "escuadron suicidad", "flash"}
  
   for _, listaPeliculas := range peliculas{
      fmt.Println(listaPeliculas)
   }



}


