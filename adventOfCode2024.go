package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Utils
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func absolute(val int) int {
    if val > 0 {
        return val
    }
    return val * -1
}

func checkSafety(data []string) (bool, int) {
    var goingUp = true // true => up / false => down
    first, err := strconv.Atoi(data[0])
    check(err)

    second, err := strconv.Atoi(data[1])
    check(err)

    if first > second {
        goingUp = false
    }

    for idx, num := range(data) {
        if idx == 0 {
            continue
        }

        num1, err := strconv.Atoi(data[idx - 1])
        check(err)

        num2, err := strconv.Atoi(num)
        check(err)

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

func removeIndex(s[]string, index int) []string {
    return append(s[:index], s[(index+1):]...)
}

// Days
func day1Part1() {
    var arr1 = make([]int, 1000)
    var arr2 = make([]int, 1000)
    var sum int

    file, err := os.Open("./input.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        result := strings.Split(scanner.Text(), "   ")

        num1, err := strconv.Atoi(result[0])
        check(err)

        num2, err := strconv.Atoi(result[1])
        check(err)

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
        sum += absolute(arr1[index] - arr2[index])
    }
    fmt.Println(sum)
}

func day1Part2() {
    var dict = make(map[int]int)

    var arr1 = make([]int, 1000)
    var arr2 = make([]int, 1000)
    var sum int

    file, err := os.Open("./input.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        result := strings.Split(scanner.Text(), "   ")

        num1, err := strconv.Atoi(result[0])
        check(err)

        num2, err := strconv.Atoi(result[1])
        check(err)

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
    fmt.Println(sum)
}

func day2Part1() {
    file, err := os.Open("./input2.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var sums int

    for scanner.Scan() {
        result := strings.Split(scanner.Text(), " ")
        safe, _ := checkSafety(result)

        if safe {
            sums++
        }
    }
    fmt.Println(sums)
}

func day2Part2() {
    file, err := os.Open("./input2.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var sums int

    for scanner.Scan() {
        result := strings.Split(scanner.Text(), " ")
        result1 := strings.Split(scanner.Text(), " ")
        result2 := strings.Split(scanner.Text(), " ")

        size := len(result)
        safe, idx := checkSafety(result)

        if !safe && idx < size {
            newArr := removeIndex(result, idx)
            safe, _ = checkSafety(newArr)
        }

        if !safe && idx > 0 {
            newArr := removeIndex(result1, idx - 1)
            safe, _ = checkSafety(newArr)
        }

        if !safe && (idx > 1)  {
            newArr := removeIndex(result2, idx - 2)
            safe, _ = checkSafety(newArr)
        }

        if safe {
            sums++
        }
    }
    fmt.Println(sums)
}

func day3Part1() {
    file, err := os.Open("./input3.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var sum int

    for scanner.Scan() {
        result := scanner.Text()
        re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
        matches := re.FindAllStringSubmatch(result, -1)

        for _, val := range matches {
            str := strings.Split(val[0], "(")
            nums := strings.Split(str[1], ",")
            nums[1] = strings.Trim(nums[1], ")")
            num1 ,err := strconv.Atoi(nums[0])
            check(err)
            num2 ,err := strconv.Atoi(nums[1])
            check(err)
            sum += num1*num2
        }
    }
    fmt.Println(sum)
}

func day3Part2() {
    file, err := os.Open("./input3.txt")
    check(err)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var sum int
    var enabled bool = true


    for scanner.Scan() {
        result := scanner.Text()
        re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
        matches := re.FindAllStringSubmatch(result, -1)

        for _, val := range matches {
            if val[0] == "do()" {
                enabled = true
                continue
            } else if val[0] == "don't()" {
                enabled = false
                continue
            }

            if !enabled {
                continue
            }

            str := strings.Split(val[0], "(")
            nums := strings.Split(str[1], ",")
            nums[1] = strings.Trim(nums[1], ")")

            num1 ,err := strconv.Atoi(nums[0])
            check(err)
            num2 ,err := strconv.Atoi(nums[1])
            check(err)
            sum += num1*num2
        }
    }
    fmt.Println(sum)
}

func main() {
    // day1Part1()
    // day1Part2()
    // day2Part1()
    // day2Part2()
    // day3Part1()
    // day3Part2()
}

