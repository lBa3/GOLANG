package main 

import ("fmt")
import ("time")

func main(){

  momento := time.Now();
  hoy := momento.Weekday();
  
  switch hoy {
      case 0: 
      	fmt.Println("Hoy es domingo");
      case 1: 
      	fmt.Println("Hoy es lunes");
      case 2: 
      	fmt.Println("Hoy es martes");
      case 3: 
      	fmt.Println("Hoy es miercoles");
      default:
      	fmt.Println("Es otro dia de la semana");

  }


}


