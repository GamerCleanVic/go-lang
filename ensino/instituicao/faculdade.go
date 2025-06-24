package instituicao

import "fmt"

type Faculdade struct{
  Nome string
}

func NovaFaculdade(nome string) *Faculdade{
  return &Faculdade{Nome: nome}
}

func (f *Faculdade) ImprimirDados(){
  fmt.Printf("\nNome da Faculdade: %s\n", f.Nome)
}
