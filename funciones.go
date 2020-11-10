package main 
import ("fmt")


func main(){
  cebezera();

  var num1 float32 = 10;
  var num2 float32= 6;
  fmt.Println ("Los valores son : ", num1 , " Y ",  num2);
  calculadora(num1, num2);

  fmt.Println (":::::::::::::::::::::::::::::::::::::::");

  var num3 float32 = 44;
  var num4 float32 = 7;
  fmt.Println ("Los valores son : ", num3 , " Y ", num4);
  calculadora(num3, num4);




}

func cebezera(){
   fmt.Println (":::::::::::::::::CALCULADORA POR FUNCIONES::::::::::::::::::");
}

func operacion(num1 float32, num2 float32, operacion string) float32{
	var resultado float32;

	if (operacion == "+"){
		resultado = num1 + num2;
	}

	if (operacion == "-"){
		resultado = num1 - num2;
	}

	if (operacion == "*"){
		resultado = num1 * num2;
	}

	if (operacion == "/"){
		resultado = num1 / num2;
	}

	return resultado
}

func calculadora(num1 float32, num2 float32){

  fmt.Print("El resultado de la suma es: ");
  fmt.Println(operacion(num1, num2, "+"));

  fmt.Print("El resultado de la resta es: ");
  fmt.Println(operacion(num1, num2, "-"));

  fmt.Print("El resultado de la multiplicacion es: ");
  fmt.Println(operacion(num1, num2, "*"));

  fmt.Print("El resultado de la division es: ");
  fmt.Println(operacion(num1, num2, "/"));
}

