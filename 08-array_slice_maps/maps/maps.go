package maps

import "fmt"

type Product struct {
	id    uint
	title string
	price float64
}

func Maps() {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites["Google"])

	// Add map key value
	websites["Youtube"] = "https://youtube.com"
	fmt.Println(websites)

	// Change value
	websites["Google"] = "https://kookle.co.th"
	fmt.Println(websites["Google"])

	// Delete value
	delete(websites, "Google")
	fmt.Println(websites)

	//for k, v := range websites {
	//	fmt.Printf("Company: %s, Website: %s\n", k, v)
	//}
}
