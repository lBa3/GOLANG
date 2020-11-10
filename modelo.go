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
   var gorra_negra  = Gorra {marca : "Nike",color : "negro",precio: 25.90,plana : false}
   var gorra_addidas  = Gorra {"addidas","AZUL", 29.90,true}

      fmt.Println (gorra_negra)
      //aceder a un elemento del objeto
      fmt.Println (gorra_negra.marca)
      fmt.Println (gorra_negra.precio)

      fmt.Println (gorra_addidas)

}