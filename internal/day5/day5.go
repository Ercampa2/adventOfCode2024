package day5

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"adventOfCode2024/pkg/utils"
)

func Part1 () int {
    file, err := os.Open("./assets/input5.txt")
    utils.Check(err)
    defer file.Close()

    var total int
    rules := make(map[string][]string)

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        var line string = scanner.Text()
        // Find rulesets
        if strings.Contains(line, "|") {
            parts := strings.Split(line, "|")
            _, ok := rules[parts[0]]
            if !ok {
                rules[parts[0]] = []string{}
            }

            rules[parts[0]] = append(rules[parts[0]], parts[1])
        }

        // Find values
        if strings.Contains(line, ",") {
            var valid bool = true;
            shown := make(map[string]int)
            numbers := strings.Split(line, ",")

            for _, val := range(numbers) {
                for _, ruleVal := range(rules[val]) {
                    _, ok := shown[ruleVal]
                    if ok {
                        valid = false
                    }
                    shown[val] = 0
                }
            }
            // Update pages
            if (valid) {
                val, err := strconv.Atoi(numbers[(len(numbers)-1) / 2])
                utils.Check(err)
                total += val
            }
        }
    }
    return total
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

    return 4
}
