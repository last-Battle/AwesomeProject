package function

import "fmt"

func Range() {
	//kvs := map[string]string{"a": "b", "b": "c"}
	//for k, v := range kvs {
	//	fmt.Printf("%s -> %s\n", k, v)
	//}

	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "capital", countryCapitalMap[country])
	}
}
