package entity

import (
	"errors"
	"strconv"
)

type CPF string

func NewCPF(c string) (CPF, error) {
	isValid := checkCPFIsValid(c)
	if !isValid {
		return CPF(c), errors.New("cpf must contain 11 digits and be composed only with numbers")
	}
	cpf := CPF(c)
	return cpf, nil
}

func checkCPFIsValid(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}
	if _, err := strconv.Atoi(cpf); err != nil {
		return false
	}
	return true
}
