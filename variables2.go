package main

import "fmt"

var razaGoku = "Saiyajin"

func main(){
  var(	
    saludo = "Hola mundo!"
  )
  fmt.Println(saludo)

  //Scope

  imprimir()
}

func imprimir(){
  fmt.Println("La raza de Goku es: " + razaGoku)
}
