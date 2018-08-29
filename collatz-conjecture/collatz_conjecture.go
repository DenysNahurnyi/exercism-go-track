package collatzconjecture

import "errors"

func CollatzConjecture(num int) (steps int, err error) {
	if num <= 0 {
		return 0, errors.New("input number does not suitable for conditions")
	}
	for ; num != 1; steps++ {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
	}
	return steps, nil
}
