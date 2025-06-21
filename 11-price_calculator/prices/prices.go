package prices

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println(errors.New("Cannot open file prices.txt!"))
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file) // todo เป็นการเปิดเผยการใช้เนื้อหา

	var lines []string

	for scanner.Scan() { // todo อ่านค่าใน file แบบ line by line จึงใช้ for เพื่อ scan ทุกบรรทัด (response type: bool)
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err() // scanner.Scan() บอกแค่ true false จึงต้องใช้ scanner.Err() ในการเช็ค err
	if err != nil {
		fmt.Println(errors.New("Reading the file content failed."))
		fmt.Println(err)
		file.Close()
		return
	}
}

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64, len(job.TaxIncludedPrices))

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	for k, v := range result {
		fmt.Printf("Price: %s\n", k)
		fmt.Printf("Price Included Tax: %v\n", v)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
