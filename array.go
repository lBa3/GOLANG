package main 
import ("fmt")


func main(){
  
  var peliculas [3]string;
  peliculas[0] = "BATMAN";
  peliculas[1] = "SPIDERMAN";
  peliculas[2] = "CASA BLANCA";
  
  fmt.Println(peliculas);
  fmt.Println(peliculas[1]);

  //otra manera de definir array
  carros :=[3] string{"volvo", "ford", "KIA"}
  fmt.Println(carros);

  //array multi dimencionales
  var ropa  [3][2] string;
  ropa[0][0] = "breshka";
  ropa[0][1] = "woman";

  ropa[1][0] = "modatela";
  ropa[1][1] = "premire";

  ropa[2][0] = "zara";
  ropa[2][1] = "catwoman";


  fmt.Println(ropa);

  //slice
  perros := [] string{"bulldog", "terrier", "gran maltez"}
  perros = append(perros, "pastor aleman");
  perros = append(perros, "tovi");
  fmt.Println(perros);
  //tamaño del array
  fmt.Println(len (perros));




}


