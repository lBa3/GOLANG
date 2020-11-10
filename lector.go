package main 

import ("fmt")
import ("io/ioutil")
import ("os")

func main(){
  
  //dato por POST
  nuevo_texto := os.Args[1];

  archivo, err := os.OpenFile("archivo_externo.txt", os.O_APPEND|os.O_WRONLY, 0777)
  showError(err);


  escribir, err := archivo.WriteString(nuevo_texto);
  fmt.Println(escribir);
  showError(err);

  archivo.Close();


  // leer archivo
  fichero, errorFichero := ioutil.ReadFile("archivo_externo.txt");
  showError(errorFichero);

  fmt.Println(string(fichero))

}

func showError(e error){
   if (e != nil){
    panic (e);
   }
}
