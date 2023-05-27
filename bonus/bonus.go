package bonus

import "errors"

// Você recebe uma lista de caminhos, onde `caminhos[i] = [cidadeAi, cidadeBi]` significa que existe um caminho direto que
//vai de cidadeAi para cidadeBi. Retorne a cidade de destino, ou seja, a cidade sem nenhum caminho que saia dela.

func Destino(caminhos [][2]string) (string, error) {
	cityOutgoing := make(map[string]bool)

	for _, path := range paths {
		cityOutgoing[path[0]] = true
	}

	for _, path := range paths {
		if !cityOutgoing[path[1]] {
			return path[1], nil
		}
	}
	return "", errors.New("Not implemented yet")
}
