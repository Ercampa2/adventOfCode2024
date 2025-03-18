package day1

import (
    "os"
    "bufio"
    "strings"
	"strconv"
    "sort"

    "adventOfCode2024/pkg/utils"
)

func Part1 () int {
    var arr1 = make([]int, 1000)
    var arr2 = make([]int, 1000)
    var sum int

    file, err := os.Open("./assets/input.txt")
    utils.Check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        result := strings.Split(scanner.Text(), "   ")

        num1, err := strconv.Atoi(result[0])
        utils.Check(err)

        num2, err := strconv.Atoi(result[1])
        utils.Check(err)

        arr1 = append(arr1, num1)
        arr2 = append(arr2, num2)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    sort.Slice(arr1, func(i,j int) bool {
        return arr1[i] < arr1[j]
    })

    sort.Slice(arr2, func(i,j int) bool {
        return arr2[i] < arr2[j]
    })

    for index := range arr1 {
        sum += utils.Absolute(arr1[index] - arr2[index])
    }
    return sum
}

func Part2() int {
    var dict = make(map[int]int)

    var arr1 = make([]int, 1000)
    var arr2 = make([]int, 1000)
    var sum int

    file, err := os.Open("./assets/input.txt")
    utils.Check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        result := strings.Split(scanner.Text(), "   ")

        num1, err := strconv.Atoi(result[0])
        utils.Check(err)

        num2, err := strconv.Atoi(result[1])
        utils.Check(err)

        arr1 = append(arr1, num1)
        arr2 = append(arr2, num2)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    for _, num := range(arr2) {
        dict[num]++
    }

    for _, num := range(arr1) {
        sum += num * dict[num]
    }
    return sum
}
