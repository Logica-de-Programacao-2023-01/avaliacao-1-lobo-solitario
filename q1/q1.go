package q1

import (
	"errors"
	"fmt"
)
func DivideWatermelon(weight int) (bool, error) {
	if weight <= 0 {
		return false, errors.New("Peso inválido")
	}
	if weight%2 == 0 {
		metade := weight / 2
		if metade%2 == 0 {
			fmt.Print("É possível: ", true)
		}
	}
	return false, nil
}
