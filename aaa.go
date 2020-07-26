package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("ծրագիրը ստուգում է մուտքագրված թվի պարզ լինելը")
	fmt.Println(" մուտքագրե՛ք թիվը")
	n, err := fmt.Scan(&num)
	if n < 1 {
		fmt.Println("մուտքային տվյալների սխալ", err)
		return
	}
	if !(num > 0) {
		fmt.Println("մուտքային տվյալների սխալ,պահանջվում է բնական թիվ")
		return
	}
	if num == 1 {
		fmt.Println("1-ը յուրահաատուկ թիվ է")
		return
	}

	k := 2
	l := 2
	for ; k < num; k++ {

		if num%l == 0 {
			fmt.Println("թիվը բաղադրյալ է")
			return
		}

		l++
	}
	fmt.Println("թիվը պարզ է")
	return
}
