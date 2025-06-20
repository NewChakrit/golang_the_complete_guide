package maps

import "fmt"

func Maps() {
	websites := map[string]string{
		"Google": "https://google.com",
		"AWS":    "https://aws.com",
	}

	fmt.Println(websites["Google"])

	//fmt.Println(websites)

	//for k, v := range websites {
	//	fmt.Printf("Company: %s, Website: %s\n", k, v)
	//}
	//

	// Add map key value
	websites["Youtube"] = "https://youtube.com"
	fmt.Println(websites)

	// Change value
	websites["Google"] = "https://kookle.co.th"
	fmt.Println(websites["Google"])

	// Delete value
	delete(websites, "Google")
	fmt.Println(websites)
}
