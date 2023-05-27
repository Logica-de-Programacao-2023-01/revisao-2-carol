package q3

import "errors"

//Você está desenvolvendo um sistema de gerenciamento de estoque para uma loja. Cada produto possui um código único, nome,
//preço unitário e quantidade em estoque. Você decidiu usar uma struct para representar as informações de cada produto.
//
//Agora, você precisa implementar uma função chamada "updateStock" que recebe como parâmetro um ponteiro para um objeto do
//tipo Product e um mapa que representa as vendas realizadas, onde as chaves são os códigos dos produtos vendidos e os
//valores são as quantidades vendidas. A função deve atualizar a quantidade em estoque do produto com base nas vendas
//realizadas.
//
//Sua tarefa é escrever a função "updateStock" e usá-la para atualizar a quantidade em estoque dos produtos vendidos.

type Product struct {
	Code     string
	Name     string
	Price    float64
	Quantity int
}

func UpdateStock(product *Product, sales map[string]int) error {
	package main
	if prod == nil {
		return errors.New("product pointer is nil")
	}

	for code, quantity := range sales {
		if quantity < 0 {
			return errors.New("negative sale quantity")
		}

		if code == prod.Code {
			newQuantity := prod.Quantity - quantity
			if newQuantity < 0 {
				return errors.New("insufficient stock")
			}
			prod.Quantity = newQuantity
		}
	}

	return nil
}

func main() {
	product := &Product{
		Code:     "P001",
		Name:     "Phone",
		Price:    1000.0,
		Quantity: 10,
	}

	sales := map[string]int{
		"P001": 3,
		"P002": 5,
	}

	err := updateStock(product, sales)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Product:", product.Name)
	fmt.Println("Quantity in Stock:", product.Quantity)
}

	return errors.New("Not implemented yet")
}
