package main

import (
  "bufio"
  "fmt"
  "ensino/instituicao"
  "ensino/escola"
  "os"
  "strings"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  
  fmt.Print("\nDigite o nome da Faculdade: ")
  nomeFaculdade, _ := reader.ReadString('\n')
  nomeFaculdade = strings.TrimSpace(nomeFaculdade)

  fmt.Print("\nDigite o nome do Colégio: ")
  nomeColegio, _ := reader.ReadString('\n')
  nomeColegio = strings.TrimSpace(nomeColegio)

  faculdade := instituicao.NovaFaculdade(nomeFaculdade)
  colegio := escola.NovoColegio(nomeColegio)

  if nomeFaculdade == "" || nomeColegio == ""{
    fmt.Println("\nDigite o nome da instituição.\n")
  }else{
    faculdade.ImprimirDados()
    colegio.ImprimirDados()
  }
}
