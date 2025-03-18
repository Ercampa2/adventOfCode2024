package main

import (
	"fmt"
	"os"
	"strconv"

	"adventOfCode2024/pkg/utils"
	"adventOfCode2024/internal/day1"
	"adventOfCode2024/internal/day2"
	"adventOfCode2024/internal/day3"
	"adventOfCode2024/internal/day4"
	"adventOfCode2024/internal/day5"
)

// Main
func main() {
    if len(os.Args) < 2 {
        fmt.Println("Escolha um dia e parte")
        return
    }

    dia, err := strconv.Atoi(os.Args[1])
    utils.Check(err)

    if dia == 0 {
        fmt.Println("Escolha um dia a partir de 1")
        return
    }

    parte, err := strconv.Atoi(os.Args[2])
    utils.Check(err)

    if parte != 1 && parte != 2 {
        fmt.Println("Escolha a parte 1 ou 2")
        return
    }

    var operacoes = [][]func() int {
        {nil, nil},
        {day1.Part1, day1.Part2},
        {day2.Part1, day2.Part2},
        {day3.Part1, day3.Part2},
        {day4.Part1, day4.Part2},
        {day5.Part1, day5.Part2},
    }

    if (dia > len(operacoes) - 1) {
        fmt.Println("Dia não cadastrado")
        return
    }

    fmt.Println(operacoes[int(dia)][int(parte)]())
    fmt.Println("OK")
}
