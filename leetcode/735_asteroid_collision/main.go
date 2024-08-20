package main

import "math"

func main() {
	// TODO: implement here
}

func asteroidCollisionV1(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		if len(stack) == 0 {
			stack = append(stack, asteroid)
			continue
		}

		var existed bool
		for len(stack) > 0 {
			if l, r := stack[len(stack)-1], asteroid; isNotCollision(l, r) {
				existed = true
				break
			} else if leftABS, rightABS := math.Abs(float64(l)), math.Abs(float64(r)); leftABS == rightABS {
				stack = stack[:len(stack)-1]
				break
			} else if leftABS > rightABS {
				break
			} else {
				// leftABS < rightABS
				stack = stack[:len(stack)-1]
			}
			existed = len(stack) == 0
		}
		if existed {
			stack = append(stack, asteroid)
		}
	}
	return stack
}

func isNotCollision(l, r int) bool {
	return (l > 0 && r > 0) || (l < 0 && r < 0) || (l < 0 && r > 0)
}

// refer to: https://leetcode.com/problems/asteroid-collision/solutions/4025988/go-solution-easy-explanation-o-n-time-space-complexity/?envType=study-plan-v2&envId=leetcode-75
func asteroidCollisionV2(asteroids []int) []int {
	var stack []int
	for i := 0; i < len(asteroids); i++ {
		asteroid := asteroids[i]
		if len(stack) == 0 || stack[len(stack)-1] < 0 || asteroid > 0 {
			stack = append(stack, asteroid)
		} else if prev := stack[len(stack)-1]; prev == -asteroid {
			stack = stack[:len(stack)-1]
		} else if prev < -asteroid {
			stack = stack[:len(stack)-1]
			i--
		}
	}
	return stack
}
