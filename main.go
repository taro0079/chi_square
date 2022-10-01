package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const TEST_VAL = 3.841

func main() {
	fmt.Println("Welcome to Chi Checker !!")
	fmt.Println("有意水準 a=0.5で計算します")
	fmt.Println("**検定の自由度は1です！！**")
	op, _ := strconv.ParseFloat(originalParamter(), 64)
	on, _ := strconv.ParseFloat(originalNext(), 64)
	tp, _ := strconv.ParseFloat(testParameter(), 64)
	tn, _ := strconv.ParseFloat(testNext(), 64)
	onn := originalNonConversion(op, on)
	tnn := testNonConversion(tp, tn)
	a := makevar(on, onn, tn, tnn)
	chi := sumvar(a)

	fmt.Println("chi2 = " + strconv.FormatFloat(chi, 'f', -1, 64))
	fmt.Println("chi2(free) = " + strconv.FormatFloat(TEST_VAL, 'f', -1, 64))
	judgement(chi)
}

func judgement(chi float64) bool {
	if chi < TEST_VAL {
		fmt.Println("有意な差があるとは言えません")
		return false
	} else {
		fmt.Println("有意な差があります")
		return true
	}

}
func originalNonConversion(op float64, on float64) float64 {
	return op - on
}
func testNonConversion(tp float64, tn float64) float64 {
	return tp - tn
}

func sumvar(v [][]float64) float64 {
	sum := 0.0
	for _, i := range v {
		for _, num := range i {
			sum += num
		}
	}
	return sum
}
func calcvar(r float64, i float64) float64 {
	return ((r - i) * (r - i)) / i
}

func makevar(on float64, onn float64, tn float64, tnn float64) [][]float64 {
	r := makearray(on, onn, tn, tnn)
	i := makeideal(on, onn, tn, tnn)
	return [][]float64{{calcvar(r[0][0], i[0][0]), calcvar(r[0][1], i[0][1])}, {calcvar(r[1][0], i[1][0]), calcvar(r[1][1], i[1][1])}}

}

func makearray(on float64, onn float64, tn float64, tnn float64) [][]float64 {
	return [][]float64{{on, onn, on + onn}, {tn, tnn, tn + tnn}, {on + tn, onn + tnn, on + onn + tn + tnn}}
}

func makeideal(on float64, onn float64, tn float64, tnn float64) [][]float64 {
	arr := makearray(on, onn, tn, tnn)
	return [][]float64{{arr[2][0] * arr[0][2] / arr[2][2], arr[2][1] * arr[0][2] / arr[2][2]}, {arr[2][0] * arr[1][2] / arr[2][2], arr[2][1] * arr[1][2] / arr[2][2]}}
}

func originalParamter() string {
	fmt.Print("オリジナルのサンプル数は？ > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	on := scanner.Text()
	return on

}

func originalNext() string {
	fmt.Print("オリジナルのコンバージョン数は？ > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	on := scanner.Text()
	return on
}

func testParameter() string {
	fmt.Print("テストのサンプル数は？ > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	on := scanner.Text()
	return on

}

func testNext() string {
	fmt.Print("テストのコンバージョン数は？ > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	on := scanner.Text()
	return on
}
