package q2

import "errors"

//Você é um desenvolvedor de software em uma empresa financeira e está trabalhando em um sistema de folha de pagamento.
//Cada funcionário possui um ID único, nome, cargo, salário base e um conjunto de bônus mensais. Você decidiu usar uma
//struct para representar as informações de cada funcionário.
//
//Agora, você precisa implementar uma função chamada "calculateTotalSalary" que recebe como parâmetro um ponteiro para um
//objeto do tipo Employee e calcula o salário total do funcionário, considerando o salário base e a soma dos bônus
//mensais. Caso a soma dos bônus mensais seja maior que 1500.0, a titulação do funcionário deve ser atualizada com o
//prefixo "Senior".

type Employee struct {
	ID         int
	Name       string
	Title      string
	BaseSalary float64
	Bonuses    []float64
}

func CalculateTotalSalary(employee *Employee) (float64, error) {
	package main
	}

	totalBonus := 0.0
	for _, bonus := range emp.MonthlyBonus {
		totalBonus += bonus
	}

	totalSalary := emp.BaseSalary + totalBonus

	if totalBonus > 1500.0 {
		emp.Position = "Senior " + emp.Position
	}

	return totalSalary, nil
}

func main() {
	employee := &Employee{
		ID:         1,
		Name:       "John Doe",
		Position:   "Developer",
		BaseSalary: 5000.0,
		MonthlyBonus: []float64{
			1000.0,
			800.0,
			750.0,
		},
	}

	totalSalary, err := calculateTotalSalary(employee)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Total Salary:", totalSalary)
	fmt.Println("Position:", employee.Position)
}

	return 0, errors.New("Not implemented yet")
}
