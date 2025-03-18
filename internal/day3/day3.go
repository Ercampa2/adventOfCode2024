package day3

import (
    "os"
    "bufio"
    "strings"
    "regexp"
    "strconv"

    "adventOfCode2024/pkg/utils"
)

func Part1 () int {
    file, err := os.Open("./assets/input3.txt")
    utils.Check(err)
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
            utils.Check(err)
            num2 ,err := strconv.Atoi(nums[1])
            utils.Check(err)
            sum += num1*num2
        }
    }
    return sum
}

func Part2() int {
    file, err := os.Open("./assets/input3.txt")
    utils.Check(err)
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
            utils.Check(err)
            num2 ,err := strconv.Atoi(nums[1])
            utils.Check(err)
            sum += num1*num2
        }
    }
    return sum
}
