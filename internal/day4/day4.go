package day4

import (
    "os"
    "bufio"

    "adventOfCode2024/pkg/utils"
)

func SearchForWord(data []string, row, col int, word string) int {
    var found = 0
    for x := -1; x <= 1; x++ {
        for y := -1; y <= 1; y++ {
            if (SearchInDirection(data, row, col, x, y, 1, word)) {
                found++
            }
        }
    }
    return found
}

func SearchInDirection(data []string, row, col, dirX, dirY, position int, word string) bool{
    if (len(word) == position) {
        return true
    }

    nextRow := row + (dirX * position)
    nextCol := col + (dirY * position)

    if (nextRow < 0 || nextRow > (len(data) - 1) || nextCol < 0 || nextCol > (len(data[nextRow]) - 1)) {
        return false
    }

    if (data[nextRow][nextCol] == word[position]) {
        return SearchInDirection(data, row, col, dirX, dirY, position + 1, word)
    }
    return false
}

func searchDiagonals(data []string, row, col int) bool {
    encontrados := 0;
    if row == 0 || col == 0 || row >= len(data) - 1 || col >= len(data[row]) - 1 {
        return false
    }

    if string(data[row-1][col-1]) == "M" && string(data[row+1][col+1]) == "S" {
        encontrados++
    }

    if string(data[row+1][col+1]) == "M" && string(data[row-1][col-1]) == "S" {
        encontrados++
    }

    if string(data[row-1][col+1]) == "M" && string(data[row+1][col-1]) == "S" {
        encontrados++
    }

    if string(data[row+1][col-1]) == "M" && string(data[row-1][col+1]) == "S" {
        encontrados++
    }
    return encontrados == 2
}

func Part1 () int {
    file, err := os.Open("./assets/input4.txt")
    utils.Check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    var word = "XMAS"
    var total = 0

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    for row, rowText := range lines {
        for col, char := range rowText {
            if (string(char) == "X") {
                total += SearchForWord(lines, row, col, word)
            }
        }
    }

    return total
}

func Part2() int {
    file, err := os.Open("./assets/input4.txt")
    utils.Check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    var total = 0
    var center string = "A"

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    for row, rowText := range lines {
        for col, char := range rowText {
            if string(char) == center {
                if searchDiagonals(lines, row, col) {
                    total++
                }
            }
        }
    }

    return total
}
