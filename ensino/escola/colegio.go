package escola

import "fmt"

type Colegio struct {
  Nome string
}

func NovoColegio (nome string) *Colegio {
  return &Colegio{Nome: nome}
}

func (c *Colegio) ImprimirDados(){
  fmt.Printf("Nome do Colégio: %s\n\n", c.Nome)
}
