package main

/*func main() {
	var out []*int
	for i := 0; i < 3; i++ {
		u := i
		out = append(out, &u)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}
*/

/*func main() {
	var values []*int
	for _, val := range values {
		go func(val interface{}) {
			fmt.Println(val)
		}(val)
	}
}
*/
/*
func DeduplicateStrings(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}*/

/*func main() {

	in := []int{3, 2, 1, 4, 3, 2, 1, 4, 1} // any item can be sorted
	sort.Ints(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		in[j] = in[i]
	}
	result := in[:1+j]
	fmt.Println(result)

}
*/
// moveToFront moves needle to the front of haystack, in place if possible.
func moveToFront(needle string, haystack []string) []string {
	if len(haystack) != 0 && haystack[0] == needle {
		return haystack
	}
	prev := needle
	for i, elem := range haystack {
		switch {
		case i == 0:
			haystack[0] = needle
			prev = elem
		case elem == needle:
			haystack[i] = prev
			return haystack
		default:
			haystack[i] = prev
			prev = elem
		}
	}
	return append(haystack, prev)
}

func main() {

	/*haystack := []string{"a", "b", "c", "d", "e"} // [a b c d e]
	haystack = moveToFront("c", haystack)         // [c a b d e]
	haystack = moveToFront("f", haystack)         // [f c a b d e]
	fmt.Println(haystack)*/

}
