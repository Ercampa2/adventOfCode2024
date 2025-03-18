package day5

import (
    "os"
    "bufio"
    
    "adventOfCode2024/pkg/utils"
)

func Part1 () int {
    file, err := os.Open("./assets/input5.txt")
    utils.Check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return 1
}

func Part2() int {
    file, err := os.Open("./assets/input5.txt")
    utils.Check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    return 1
}
