package main

import (
  "fmt"
  "math"
)

func main(){
  var grau int
  var rad float64

  fmt.Print("Ângulo em º: ")
  fmt.Scanln(&grau)

  rad = float64(grau) * math.Pi / 180
  cosseno := math.Cos(float64(rad))
  seno := math.Sin(float64(rad))

  fmt.Printf("\nCos: %.2f\nSen: %.2f\n", cosseno, seno)
}
