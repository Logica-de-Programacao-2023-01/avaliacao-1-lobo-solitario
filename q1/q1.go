package q1

import (
	"errors"
)

func DivideWatermelon(weight int) (bool, error) {
	if weight <= 0 {
		return false, errors.New("Peso inválido")
	}
	par := false
	if weight == 2{
		return false, nil
	}
	if weight%2 == 0{
		par = true
		return true, nil
	}
	if par{
		return false, nil
	}
	
	return false, nil
}
