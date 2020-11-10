package main 
import ("fmt")



func main(){
   /*Variables*/
   var nombre string = "Rogelio"
   var apellido string ="Batres"
   apellido = "warneros"
   var suma int = 8+8;
   
   /*declaracion de variables sin la declaracion de tipo*/
   pais := "Mexico"
   prueba := true
   flonte :=  1.233

   /*constantes*/
   const year = 2020

   /*impreciones*/
   fmt.Println("Hello Freelancer !!!"+ nombre + " " + apellido)
   fmt.Println("El resultado de la suma es: ", suma)
   fmt.Println(pais)
   fmt.Println(prueba)
   fmt.Println(flonte)
   fmt.Println(year)
}

