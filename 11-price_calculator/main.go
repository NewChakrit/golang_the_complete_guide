package main

import (
	"example.com/price_calculator/filemanager"
	"example.com/price_calculator/prices"
	"fmt"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		//err := priceJob.Process()
		//if err != nil {
		//	fmt.Println("Could not procress job!")
		//	fmt.Println(err)
		//}

		go priceJob.Process(doneChans[i], errorChans[i])
		//if err != nil {
		//	fmt.Println("Could not procress job!")
		//	fmt.Println(err)
		//}

	}

	// todo Control structure
	// todo เป็นโครงสร้างเพื่อควบคุม channel
	// todo select คือการรอ channel เดียวเพื่อส่งค่า และเมื่อการส่งค่าเกิดขึ้น มันจะดำเนินการต่อ โดยไม่สนใจกรณีอื่นๆ
	// todo case ไหนให้ค่ากับเราเป็น case แรกจะชนะ! และ case อื่นๆ จะถูกละทิ้ง และ channel อื่นๆจะถูกละเว้นในภายหลัง
	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done!")
		}
	}

	//for _, doneChan := range doneChans {
	//	<-doneChan
	//}
}
