package main

import "fmt"

func main(){
  var altura float32
  var base float32
  var area float32

  fmt.Println("Base do triângulo: ")
  fmt.Scanln(&base)
  fmt.Println("Altura do triângulo: ")
  fmt.Scanln(&altura)

  area = base * altura / 2

  fmt.Printf("Área: %.1f\n", area)
}
