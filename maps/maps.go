package maps

import "fmt"

func RunMapsRelated() {
	println("direct creation>>>")
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	println("Initialicate, add, update, get, delete")
	m := make(map[string]float64)

	m["pi"] = 3.14   // Add a new key-value pair
	m["pi"] = 3.1416 // Update value
	fmt.Println(m)   // Print map: "map[pi:3.1416]"

	v := m["pi"] // Get value: v == 3.1416
	v = m["pie"] // Not found: v == 0 (zero value)
	println(v)

	_, found := m["pi"] // found == true
	_, found = m["pie"] // found == false
	println(found)

	if x, found := m["pi"]; found {
		fmt.Println(x)
	} // Prints "3.1416"

	delete(m, "pi") // Delete a key-value pair
	fmt.Println(m)  // Print map: "map[]"

	// looping
	m_ := map[string]float64{
		"pi": 3.1416,
		"e":  2.71828,
	}
	fmt.Println(m_) // "map[e:2.71828 pi:3.1416]"

	for key, value := range m_ { // Order not specified
		fmt.Println(key, value)
	}
}
