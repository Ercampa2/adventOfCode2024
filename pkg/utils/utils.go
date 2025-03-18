package utils

func Check(e error) {
    if e != nil {
        panic(e)
    }
}

func Absolute(val int) int {
    if val > 0 {
        return val
    }
    return val * -1
}

func RemoveIndex(s[]string, index int) []string {
    return append(s[:index], s[(index+1):]...)
}
