package prices

import (
	"example.com/price_calculator/conversion"
	"example.com/price_calculator/iomanager"
	"fmt"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"` // todo การใช่้ "-" นั้นหมายความว่าเราขอให้ละเว้น json ค่านี้
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()
	if err != nil {
		return err
	}

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

	return job.IOManager.WriteResult(job) // if func err it return error, else return nil
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
