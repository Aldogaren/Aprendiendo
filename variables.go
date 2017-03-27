package main

import "fmt"

func main()  {
  var numero int
  numero = 25
  fmt.Println(numero)
  numero = 40
  fmt.Println(numero)
  nombre := "Alejandro"
  nombre2 := "Pedro"
  fmt.Println(nombre, nombre2)
  nombre, nombre2 = nombre2, nombre
  fmt.Println(nombre, nombre2) 

}
