package main

import "fmt"

type floatMap map[string]float64

func (f floatMap) output() {
	fmt.Println(f)
}

func main() {
	//list.List()
	//maps.Maps()

	userNames := make([]string, 2, 5)
	// TODO ^ third paramet is capacity
	fmt.Println(len(userNames)) // 2 <= because it default value
	fmt.Println(cap(userNames)) // 5

	userNames[0] = "Meen"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel", "Man", "Mop", "Mo")

	fmt.Println(len(userNames)) // 7 add on default 2
	fmt.Println(cap(userNames)) // 10 // TODO add topup default capacity, EX if default cap 5 and then add value more than 5, cap is 10.
	fmt.Println(userNames)      // [ Meen * Max Manuel]

	// if use
	//name := []string{}
	//name[0] = "John" // TODO <= ERROR!

	//coursesRatings := map[string]float64{}  // TODO ถ้าทำท่านี้แล้วเรา append Go จะจัดสรรหน่วยความจำใหม่ทุกครั้งที่เราเพิ่มรายการใหม่ลงใน map
	//coursesRatings := make(map[string]float64, 3) // TODO แก้โดยกรใช้ make ครอบ ตือเราจำกัดหน่วยความจำไว้ล่วงหน้าสำหรับ maps
	coursesRatings := make(floatMap, 3) // ถ้ายาวไป แล้วใช้ซ้ำก็ทำ type ไว้
	coursesRatings["go"] = 4.7
	coursesRatings["react"] = 4.5
	coursesRatings["angular"] = 4.2

	//fmt.Println(coursesRatings) // map[go:4.7 react:4.8]
	coursesRatings.output()

	//
	for k, v := range coursesRatings {
		fmt.Printf("Language: %s, Rating: %v\n", k, v)
	}

	for i, v := range userNames {
		fmt.Println("Index:", i)
		fmt.Println("Valure:", v)

	}
}
