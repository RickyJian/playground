package main

import (
	"fmt"
)

func main() {
	fmt.Println(dailyTemperaturesBrute([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperaturesBrute([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperaturesBrute([]int{50, 40, 20, 30}))

	fmt.Println(dailyTemperaturesStack([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Println(dailyTemperaturesStack([]int{30, 40, 50, 60}))
	fmt.Println(dailyTemperaturesStack([]int{50, 40, 20, 30}))
}

func dailyTemperaturesBrute(temperatures []int) []int {
	temperatureLen := len(temperatures)
	if temperatureLen == 1 {
		return []int{0}
	}

	results := make([]int, 0, len(temperatures))
	for i := 0; i < temperatureLen; i++ {
		for j := i + 1; j < temperatureLen; j++ {
			if temperatures[j] > temperatures[i] {
				results = append(results, j-i)
				break
			} else if j == len(temperatures)-1 {
				results = append(results, 0)
			}
		}
	}
	results = append(results, 0)
	return results
}

func dailyTemperaturesStack(temperatures []int) []int {
	temperatureLen := len(temperatures)
	// 存 prevDay <= currentDay 的位置
	stack := make([]int, 0, temperatureLen)
	results := make([]int, temperatureLen)
	for i, t := range temperatures {
		// 若當下氣溫 > 過去氣溫，會將 stack 中所有符合的條件 pop 出來，
		// 並計算 prevDay ~ currentDay 的天數
		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			prevDayIdx := stack[len(stack)-1]
			results[prevDayIdx] = i - prevDayIdx
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return results
}
