package q1

//Você está trabalhando em um projeto de gerenciamento de uma escola. O sistema precisa armazenar informações sobre os alunos, incluindo seu nome, idade e as matérias em que estão matriculados, juntamente com suas respectivas notas. Você decidiu usar structs e mapas para representar essas informações.
//
//No entanto, você descobriu que existem dois conjuntos de dados diferentes sobre os alunos. Cada conjunto de dados é representado por um mapa. O mapa "studentData1" contém informações sobre as notas dos alunos para a primeira metade do semestre, enquanto o mapa "studentData2" contém informações sobre as notas para a segunda metade do semestre.
//
//Sua tarefa é criar uma função chamada "mergeStudentData" que recebe os mapas "studentData1" e "studentData2" como parâmetros e retorna um novo mapa que contém as informações combinadas dos dois conjuntos de dados.
//
//O objetivo é combinar as informações de cada aluno, preservando o nome e a idade, e atualizando as matérias e notas de acordo com o mapa mais recente. Lembre-se de que um aluno pode estar matriculado em diferentes matérias em cada metade do semestre.

type Student struct {
	Name     string
	Age      int
	Subjects map[string]float64
}

func MergeStudentData(studentData1 map[string]Student, studentData2 map[string]Student) map[string]Student {
	package main
	// Itera sobre os alunos do mapa studentData1
	for key, student := range studentData1 {
		mergedData[key] = student
	}

	// Itera sobre os alunos do mapa studentData2
	for key, student := range studentData2 {
		if _, exists := mergedData[key]; exists {
			// Se o aluno já existir no mapa mergedData, atualiza as matérias e notas
			for subject, grade := range student.Subjects {
				mergedData[key].Subjects[subject] = grade
			}
		} else {
			// Se o aluno não existir no mapa mergedData, adiciona-o
			mergedData[key] = student
		}
	}

	return mergedData
}

func main() {
	// Dados dos alunos para a primeira metade do semestre
	studentData1 := map[string]Student{
		"João": {
			Name: "João",
			Age:  18,
			Subjects: map[string]float64{
				"Matemática": 8.5,
				"Física":     7.8,
			},
		},
		"Maria": {
			Name: "Maria",
			Age:  17,
			Subjects: map[string]float64{
				"Inglês":  9.0,
				"História": 8.2,
			},
		},
	}

	// Dados dos alunos para a segunda metade do semestre
	studentData2 := map[string]Student{
		"João": {
			Name: "João",
			Age:  18,
			Subjects: map[string]float64{
				"Matemática": 9.0,
				"Física":     8.2,
				"Química":    7.5,
			},
		},
		"Ana": {
			Name: "Ana",
			Age:  16,
			Subjects: map[string]float64{
				"Biologia": 9.5,
				"Inglês":   8.7,
			},
		},
	}

	// Combinar os dados dos alunos
	mergedData := mergeStudentData(studentData1, studentData2)

	// Imprimir os dados combinados dos alunos
	for _, student := range mergedData {
		fmt.Printf("Nome: %s\n", student.Name)
		fmt.Printf("Idade: %d\n", student.Age)
		fmt.Println("Matérias e Notas:")
		for subject, grade := range student.Subjects {
			fmt.Printf("- %s: %.2f\n", subject, grade)
		}
		fmt.Println()
	}

	return nil
}
