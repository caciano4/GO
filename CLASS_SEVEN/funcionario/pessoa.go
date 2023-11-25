package funcionario

type Pessoa struct {
	nome  string
	idade int
}

func (p *Pessoa) setName(name string) {
	p.nome = name
}
