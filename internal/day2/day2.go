package day2

import (
    "os"
    "bufio"
    "strings"
    "strconv"

    "adventOfCode2024/pkg/utils"
)

func CheckSafety(data []string) (bool, int) {
    var goingUp = true // true => up / false => down
    first, err := strconv.Atoi(data[0])
    utils.Check(err)

    second, err := strconv.Atoi(data[1])
    utils.Check(err)

    if first > second {
        goingUp = false
    }

    for idx, num := range(data) {
        if idx == 0 {
            continue
        }

        num1, err := strconv.Atoi(data[idx - 1])
        utils.Check(err)

        num2, err := strconv.Atoi(num)
        utils.Check(err)

        diff := num1 - num2

        if diff > 3 || diff < -3 || diff == 0 {
            return false, idx
        }

        if (diff > 0 && goingUp) || (diff < 0 && !goingUp) {
            return false, idx
        }
    }
    return true, 0
}

func Part1 () int {
    file, err := os.Open("./assets/input2.txt")
    utils.Check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var sums int

    for scanner.Scan() {
        result := strings.Split(scanner.Text(), " ")
        safe, _ := CheckSafety(result)

        if safe {
            sums++
        }
    }
    return sums
}

func Part2() int {
    file, err := os.Open("./assets/input2.txt")
    utils.Check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var sums int

    for scanner.Scan() {
        result := strings.Split(scanner.Text(), " ")
        result1 := strings.Split(scanner.Text(), " ")
        result2 := strings.Split(scanner.Text(), " ")

        size := len(result)
        safe, idx := CheckSafety(result)

        if !safe && idx < size {
            newArr := utils.RemoveIndex(result, idx)
            safe, _ = CheckSafety(newArr)
        }

        if !safe && idx > 0 {
            newArr := utils.RemoveIndex(result1, idx - 1)
            safe, _ = CheckSafety(newArr)
        }

        if !safe && (idx > 1)  {
            newArr := utils.RemoveIndex(result2, idx - 2)
            safe, _ = CheckSafety(newArr)
        }

        if safe {
            sums++
        }
    }
    return sums
}
