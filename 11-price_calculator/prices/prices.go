package prices

import (
	"example.com/price_calculator/conversion"
	"example.com/price_calculator/filemanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := filemanager.Readlins("prices.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	//result := make(map[string]float64, len(job.TaxIncludedPrices))
	result := make(map[string]string, len(job.TaxIncludedPrices))

	for _, price := range job.InputPrices {

		calTaxIncludedPrice := price * (1 + job.TaxRate)
		//strconv.ParseFloat()
		//calTaxIncludedPriceFl, err := strconv.ParseFloat(fmt.Sprintf("%.2f", calTaxIncludedPrice), 64)
		//if err != nil {
		//	fmt.Println(errors.New("Converting price to float failed."))
		//	fmt.Println(err)
		//	return
		//}
		//result[fmt.Sprintf("%.2f", price)] = calTaxIncludedPriceFl
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", calTaxIncludedPrice)
	}
	//
	//for k, v := range result {
	//	fmt.Printf("Price: %s\n", k)
	//	fmt.Printf("Price Included Tax: %v\n", v)
	//}

	job.TaxIncludedPrices = result

	err := filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
	if err != nil {
		fmt.Println(err)
	}

}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
