// Conversor de binarios para decimais
// programa feito para fins educativos

package main

import (
	"errors"
	"fmt"
	"strconv"
)

// PowInt retorna base elevado a expoente
// a funcao padrao da biblioteca math retorna float

func PowInt(base, exp int) (int, error) {

	if exp == 0 {
		return 1, nil
	}
	// No contexto  de numeros binarios,numeros negativos nao sao validos

	if exp < 0 {
		return -1, errors.New("numeros negativos nao aceitos")
	}
	pow := base

	for i := 1; i < exp; i++ {
		pow *= base

	}
	return pow, nil
}

//Reverte a string para que o binario
//seja lido na ordem correta

func ReverseStr(bin string) []byte {
	var rev []byte

	// eh preciso pegar os caracteres da string como bytes
	for i := len(bin) - 1; i >= 0; i-- {
		rev = append(rev, bin[i])

	}

	return rev

}

// Converte uma string binaria em decimal
func BinToDec(bin string) (int, error) {
	bin = string(ReverseStr(bin))
	const base int = 2
	dec := 0
	for exp, digit := range bin {

		if digit != '1' && digit != '0' {
			return 0, errors.New("Digito invalido")
		}
		cur, _ := strconv.Atoi(string(digit))
		pow, err := PowInt(base, exp)
		if err != nil {
			return -1, err
		}
		cur *= pow
		dec += cur

	}
	return dec, nil

}

func main() {
	// testando
	a, err := BinToDec("")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}

}
