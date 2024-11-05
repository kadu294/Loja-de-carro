package models

import (
	"github.com/kadu294/Loja-de-carro/db"
)

type Carro struct {
	Id        int
	Modelo    string
	Ano       int
	Marca     string
	Descricao string
	Preco     float64
}

func BuscaTodosOsCarros() []Carro {
	db := db.ConectaComBancoDeDados()
	selectDeTodosOsCarros, err := db.Query("select * from carros order by id asc")
	if err != nil {
		panic(err.Error())
	}

	c := Carro{}
	carros := []Carro{}

	for selectDeTodosOsCarros.Next() {
		var id, ano int
		var modelo, descricao, marca string
		var preco float64

		err = selectDeTodosOsCarros.Scan(&id, &modelo, &ano, &marca, &descricao, &preco)
		if err != nil {
			panic(err.Error())
		}
		c.Id = id
		c.Modelo = modelo
		c.Ano = ano
		c.Marca = marca
		c.Descricao = descricao
		c.Preco = preco

		carros = append(carros, c)

	}
	defer db.Close()
	return carros
}

func CriaNovoCarro(modelo string, ano int, marca, descricao string, preco float64) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into carros (modelo, ano, marca, descricao, preco) values ($1, $2, $3, $4, $5)")
	if err != nil {
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(modelo, ano, marca, descricao, preco)
	defer db.Close()
}

func DeletaCarro(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOCarro, err := db.Prepare("delete from carros where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOCarro.Exec(id)
	defer db.Close()
}

func EditaCarro(id string) Carro {
	db := db.ConectaComBancoDeDados()

	carroDoBanco, err := db.Query("select * from carros where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	carroParaAtualizar := Carro{}

	for carroDoBanco.Next() {
		var id, ano int
		var modelo, descricao, marca string
		var preco float64

		err = carroDoBanco.Scan(&id, &modelo, &ano, &marca, &descricao, &preco)
		if err != nil {
			panic(err.Error())
		}
		carroParaAtualizar.Id = id
		carroParaAtualizar.Modelo = modelo
		carroParaAtualizar.Ano = ano
		carroParaAtualizar.Marca = marca
		carroParaAtualizar.Descricao = descricao
		carroParaAtualizar.Preco = preco
	}
	defer db.Close()
	return carroParaAtualizar
}

func AtualizaCarro(id int, modelo string, ano int, marca, descricao string, preco float64) {
	db := db.ConectaComBancoDeDados()

	AtualizaCarro, err := db.Prepare("update carros set modelo=$1, ano=$2, marca=$3, descricao=$4, preco=$5 where id=$6")
	if err != nil {
		panic(err.Error())
	}
	AtualizaCarro.Exec(modelo, ano, marca, descricao, preco, id)
	defer db.Close()
}
