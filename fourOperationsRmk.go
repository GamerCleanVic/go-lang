package main

import "fmt"

func calc(num1 float64, num2 float64){
  soma := num1 + num2
  sub := num1 - num2
  multi := num1 * num2
  divi := num1 / num2

  fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, soma)
  fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, sub)
  fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, multi)
  fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, divi)
}
func main(){
  calc(11.1, 12.3)
}
