package validation

import (
	"github.com/go-playground/validator"
	"strconv"
	"strings"
)

func ValidateCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	var eq = true
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			eq = false
		}
	}

	if eq {
		return false
	}

	var numbers [11]int

	for i := range cpf {
		numbers[i], _ = strconv.Atoi(string(cpf[i]))
	}

	var sum = 0
	for i := 0; i < 9; i++ {
		sum += numbers[i] * (10 - i)
	}

	var remainder = sum % 11
	if remainder < 2 {
		if numbers[9] != 0 {
			return false
		}
	} else if numbers[9] != 11-remainder {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += numbers[i] * (11 - i)
	}

	remainder = sum % 11
	if remainder < 2 {
		if numbers[10] != 0 {
			return false
		}
	} else if numbers[10] != 11-remainder {
		return false
	}

	return true
}

func ValidateCPFType(fl validator.FieldLevel) bool {
	return ValidateCPF(fl.Field().String())
}
