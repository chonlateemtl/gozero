package main

import "fmt"

func main() {

	/* initialize array */
	// var a [5]int
	// fmt.Println("emp:", a)

	/* assign value to array */
	// a[0] = 100
	// fmt.Println("emp:", a)

	/* initialize array with value */
	// var a = [...]int{1, 2, 3, 4, 5}
	// fmt.Println("emp:", a)

	/* initialize slice with value */
	// var a = []int{1, 2, 3, 4, 5}
	// fmt.Println("emp:", a)

	// a = append(a, 6)
	// fmt.Println("emp:", a)

	// slice topic
	// var s []string
	// fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s = make([]string, 3)
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s[0] = "a"
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	/* ---------- start len cap append topic -------------- */
	// s = make([]string, 3)
	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// s = make([]string, 0, 3)
	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	/* loop over slice */
	// s = make([]string, 0, 3)
	// s = append(s, "a")
	// s = append(s, "b")
	// s = append(s, "c")

	// for _, v := range s {
	// 	fmt.Printf("value: %s\n", v)
	// }

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("value: %s\n", s[i])
	// }

	/* ---------- start len cap append topic -------------- */

	/* --------------------- start copy slice -------------------*/
	// var s []string
	// s = append(s, "a")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// b := make([]string, 1)
	// copy(b, s)

	// fmt.Println("emp:", b, "len:", len(b), "cap:", cap(b))

	// var s []string
	// s = append(s, "a")
	// s = append(s, "b")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// b := make([]string, 1)
	// copy(b, s)

	// fmt.Println("emp:", b, "len:", len(b), "cap:", cap(b))

	// var s []string
	// s = append(s, "a")
	// s = append(s, "b")
	// fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// var b []string
	// copy(b, s)

	// fmt.Println("emp:", b, "len:", len(b), "cap:", cap(b))

	// var s []string
	// s = append(s, "a")
	// s = append(s, "b")
	// fmt.Println("slice s:", s, "len:", len(s), "cap:", cap(s))

	// b := make([]string, 1)
	// copy(b, s)

	// fmt.Println("slice b:", b, "len:", len(b), "cap:", cap(b))

	// b[0] = "c"

	// fmt.Println("slice s:", s, "len:", len(s), "cap:", cap(s))
	// fmt.Println("slice b:", b, "len:", len(b), "cap:", cap(b))

	/* --------------------- end copy slice -------------------*/

	/* ----------------- start modify slice topic ----------------*/
	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("slice s before call modify:", s, "len:", len(s), "cap:", cap(s))

	// modifySlice(s)

	/* change slice by return value from func */
	// fmt.Println("slice s after call modify:", s, "len:", len(s), "cap:", cap(s))

	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("slice s before call retSlice:", s, "len:", len(s), "cap:", cap(s))

	// s = retSlice(s)

	// fmt.Println("slice s after call retSlice:", s, "len:", len(s), "cap:", cap(s))

	/* change slice by poiter */

	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// fmt.Println("slice s before call ptrSlice:", s, "len:", len(s), "cap:", cap(s))

	// ptrSlice(&s) // yellow flag if slice growth

	// fmt.Println("slice s after call ptrSlice:", s, "len:", len(s), "cap:", cap(s))

	/* re-slice by pointer */
	// s := make([]string, 3)
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"

	// b := s

	// fmt.Println("slice s before call reSlice:", s, "len:", len(s), "cap:", cap(s))
	// fmt.Println("slice b before call reSlice:", b, "len:", len(b), "cap:", cap(b))

	// reSlice(&s) // yellow flag in here

	// fmt.Println("slice s after call reSlice:", s, "len:", len(s), "cap:", cap(s))
	// fmt.Println("slice b before call reSlice:", b, "len:", len(b), "cap:", cap(b))

	/* ----------------- end modify slice topic ----------------*/

	/*
		in Go, a slice itself is a reference type, meaning it contains a pointer to the underlying array.
		So when you pass a slice to a function, you're already passing a reference to the data, not a copy
	*/
}

func modifySlice(s []string) {
	s = append(s, "new value")
	fmt.Println("slice s in func:", s, "len:", len(s), "cap:", cap(s))
}

func retSlice(s []string) []string {
	s = append(s, "new value")
	fmt.Println("slice s in func:", s, "len:", len(s), "cap:", cap(s))
	return s
}

func ptrSlice(s *[]string) {
	*s = append(*s, "new value")
	fmt.Println("slice s in func:", s, "len:", len(*s), "cap:", cap(*s))
}

func reSlice(s *[]string) {
	*s = []string{"a"}
}
