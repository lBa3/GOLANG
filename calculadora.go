package main 
import ("fmt")


//modelo de datos
type Gorra struct {
 marca string
 color string
 precio float32
 plana bool
}

func main(){
   var num1 int = 10;
   var num2 int = 6;

   //suma
   fmt.Print("la suma es: ");
   fmt.Println(num1 + num2);

   //resta
   fmt.Print("la resta es: ");
   fmt.Println(num1 - num2);

   //multiplicacion o producto
   fmt.Print("la multiplicacion es: ");
   fmt.Println(num1 * num2);

   //divicion
   fmt.Print("la divicion es: ");
   fmt.Println(num1 / num2);

   //divicion resto
   fmt.Print("la resto de la divicion es: ");
   fmt.Println(num1 % num2);

}
