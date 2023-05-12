package main

import "fmt"

func pow(a float64, n int) float64 {
	res := float64(1.0)
	for i := 0; i < n; i++ {
		res = res * a
	}
	return res
}

func findNRoot(a float64, b float64, t float32, n int) float64 {
	l := (a + b) / 2
	for {
		pow_l_n := pow(l, n)
		comperer := pow_l_n / 10000
		if (pow_l_n-comperer < float64(t)) && (pow_l_n+comperer > float64(t)) {
			break
		}
		if float32(pow_l_n) > t {
			b = l
		} else {
			a = l
		}
		l = (a + b) / 2
		fmt.Println(pow_l_n)
	}
	return l
}

func main() {
	fmt.Println(findNRoot(-484.484, 0, -484.484, 3))
}
