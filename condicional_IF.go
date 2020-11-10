package main 

import ("fmt")
import ("os")
import ("strconv")


func main(){
  
  //var edad int =  99;

//valor por post 
//ejecutar go run condicionalID.go Rogelio 20
fmt.Println("Dame tu edad: "+ os.Args[1]+ " - Calculando...");
edad, err := strconv.Atoi(os.Args[2]);
fmt.Println(err)


  if (edad >= 18 && edad != 25 && edad != 99){
    fmt.Println("Es mayor de edad");
  }else if edad ==25{
    fmt.Println("Es mayor de edad y tines 25");
  }else if edad == 99 {
    fmt.Println("Es mayor de edad y tines 99");
  }else{
    fmt.Println("Es menor de edad");
  }


}


