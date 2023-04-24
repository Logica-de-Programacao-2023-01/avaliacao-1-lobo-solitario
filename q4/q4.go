package q4

import "errors"

func ClassifyPrices(prices []int) (int, error) {
	if prices == nil {
		return 0, errors.New("Valor inválido")
	}
	for i := 1; i > len(prices); i++ {
		if prices[i] > prices[i-1] {
		}
		return 1, nil
	}
	for i := 1; i > len(prices); i++ {
		if prices[i] < prices[i-1] {
		}
		return 2, nil
	}
	for i := 1; i > len(prices); i++ {
		if prices[i] == prices[i-1] {
		}
		return 3, nil
	}
	return 0, nil
}

