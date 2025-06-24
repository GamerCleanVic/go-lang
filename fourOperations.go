package main

import "fmt"

func soma(num1 float64, num2 float64) float64 {
  return(num1 + num2)
}
func sub(num1 float64, num2 float64) float64 {
  return(num1 - num2)
}
func multi(num1 float64, num2 float64) float64 {
  return(num1 * num2)
}
func divi(num1 float64, num2 float64) float64 {
  return(num1 / num2)
}

func main(){
  somar := soma(11.1, 9.2)
  fmt.Println("Soma = ", somar)
  subt := sub(11.1, 9.2)
  fmt.Println("Subtração = ", subt)
  mult := multi(11.1, 9.2)
  fmt.Println("Multiplicação = ", mult)
  div := divi(11.1, 9.2)
  fmt.Println("Divisão = ", div)
}
