package day6

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"adventOfCode2024/pkg/utils"
)

type Direction int

const (
    up Direction = iota
    down
    left
    right
)

type enemy struct {
    posX int
    posY int
    placesVisited int
    lookingDirection Direction
}

func (e enemy) String() string {
    return fmt.Sprintf("O inimigo está em (%v, %v) olhando para %v", e.posX , e.posY, e.lookingDirection)
}

func updateCell(gameMap *[]string, posX, posY int, val rune) {
    line := []rune((*gameMap)[posY])
    line[posX] = val
    (*gameMap)[posY] = string(line)
}

func nextCell(gameMap []string, posX, posY int) rune {
    return rune(gameMap[posY][posX])
}

func Part1 () int {
    file, err := os.Open("./assets/input6.txt")
    utils.Check(err)
    defer file.Close()

    var gameMap []string
    var guard enemy
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        gameMap = append(gameMap, line)

        idx := strings.Index(line, "^")
        if idx != -1 {
            guard.posX = idx
            guard.posY = len(gameMap) - 1
            guard.placesVisited = 1
            guard.lookingDirection = up
        }
    }

    gameLoop:
    for {
        switch guard.lookingDirection {
        case up:
            // Look for end of board
            nextY := guard.posY - 1
            if nextY < 0 {
                break gameLoop
            }

            next := nextCell(gameMap, guard.posX, nextY)

            // Look for obstacles
            if next == '#' {
                guard.lookingDirection = right
                continue
            }

            guard.posY--

            // Look cells already stepped
            if next != 'X' {
                guard.placesVisited++
                updateCell(&gameMap, guard.posX, guard.posY, 'X')
            }

        case right:
            // Look for end of board
            nextX := guard.posX + 1
            if nextX >= len(gameMap[guard.posY])  {
                break gameLoop
            }

            next := nextCell(gameMap, nextX, guard.posY)

            // Look for obstacles
            if next == '#' {
                guard.lookingDirection = down
                continue
            }

            guard.posX++

            // Look cells already stepped
            if next != 'X' {
                guard.placesVisited++
                updateCell(&gameMap, guard.posX, guard.posY, 'X')
            }

        case down:
            // Look for end of board
            nextY := guard.posY + 1
            if nextY >= len(gameMap)  {
                break gameLoop
            }

            next := nextCell(gameMap, guard.posX, nextY)

            // Look for obstacles
            if next == '#' {
                guard.lookingDirection = left
                continue
            }

            guard.posY++

            // Look cells already stepped
            if next != 'X' {
                guard.placesVisited++
                updateCell(&gameMap, guard.posX, guard.posY, 'X')
            }

        case left:
            // Look for end of board
            nextX := guard.posX - 1
            if nextX < 0  {
                break gameLoop
            }

            next := nextCell(gameMap, nextX, guard.posY)

            // Look for obstacles
            if next == '#' {
                guard.lookingDirection = up
                continue
            }

            guard.posX--

            // Look cells already stepped
            if next != 'X' {
                guard.placesVisited++
                updateCell(&gameMap, guard.posX, guard.posY, 'X')
            }
        }
    }

    return guard.placesVisited
}

func Part2() int {
    file, err := os.Open("./assets/input6.txt")
    utils.Check(err)
    defer file.Close()

    var gameMap []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
        gameMap = append(gameMap, line)

        idx := strings.Index(line, "^")
        fmt.Println(idx)
    }

    return 1
}
