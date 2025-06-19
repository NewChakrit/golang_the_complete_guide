package main

import "fmt"

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	prices = append(prices, 10.99)
	fmt.Println(prices)
}

//func main() {
//	var productNames [4]string = [4]string{"A Book", "", "Go"}
//	//productNames = [4]string{"A Book", "", "Go"}
//	prices := []float64{11.22, 44.55, 66.77, 99.55}
//	fmt.Println(prices)
//
//	productNames[3] = "A Mac"
//
//	fmt.Println(productNames)
//
//	fmt.Println(prices[2])
//
//	featuredPrice := prices[1:] // parameter1 = รวม index นั้น, parameter2 = ไม่รวม index นั้น (ใช้ index ถักไป)
//	featuredPrice[0] = 199.99
//	highlightedPrices := featuredPrice[:1]
//	fmt.Println(highlightedPrices) // 44.55
//	fmt.Println(prices)
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices)) // 1, 3
//
//	fmt.Println("----------")
//	highlightedPrices = highlightedPrices[:3]
//	fmt.Println(highlightedPrices) // 44.55
//	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
//
//	//deleteSlice(2, prices)
//
//	//sum := addText()
//	//fmt.Println(sum)
//}

//func addText() []string {
//	var val string
//	sum := []string{}
//
//	for {
//		fmt.Print("Add something: ")
//		fmt.Scanln(&val)
//		if val == "Exit" {
//			break
//		}
//
//		sum = append(sum, val)
//	}
//
//	return sum
//}

//func deleteSlice(i int, p []float64) {
//	a := p[:i]
//	b := p[i+1:]
//
//	for _, v := range b {
//		a = append(a, v)
//	}
//
//	fmt.Println(a)
//}
