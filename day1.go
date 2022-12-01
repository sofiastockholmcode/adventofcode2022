package main

import (
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    input, _ := os.ReadFile("day1input.txt")
    //input := []string{"1000","2000","3000","","4000","","5000","6000","","7000","8000","9000","","1000",}
    inputList := strings.Split(string(input), "\n")
    maxVal := 0
    currVal := 0
    elfesWithCals := make(map[int]int)

    for i, val := range inputList {
        if len(val) == 0 {
            fmt.Println("SEPARATOR")
            if currVal > maxVal {
                maxVal = currVal
            }
            elfesWithCals[i] = currVal
            currVal = 0
        } else {
            intVal, _ := strconv.Atoi(val)
            currVal += intVal
            fmt.Println(currVal)
        }
    }

    fmt.Println("")
    fmt.Println(elfesWithCals)
    keys := make([]int, 0, len(elfesWithCals))

    for key := range elfesWithCals {
        keys = append(keys, key)
    }

    sort.SliceStable(keys, func(i, j int) bool {
        return elfesWithCals[keys[i]] < elfesWithCals[keys[j]]
    })

    topElfesWithCalsVal := elfesWithCals[keys[len(keys)-1]] + elfesWithCals[keys[len(keys)-2]] + elfesWithCals[keys[len(keys)-3]]
    fmt.Println(topElfesWithCalsVal)

}
