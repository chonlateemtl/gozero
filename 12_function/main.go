package main

/* create custom type function */
type keep func(v int) bool

func filter(l []int, fn keep) []int {
	var r []int
	for _, v := range l {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}

/* no need custom type */
func filterV2(l []int, fn func(v int) bool) []int {
	var r []int
	for _, v := range l {
		if fn(v) {
			r = append(r, v)
		}
	}
	return r
}

type server struct {
	readTimeout  int
	writeTimeout int
	idleTimeout  int
	addr         string
}

/*optional pattern */
// type option func(s *server)

// func changeReadTimeout() option {
// 	return func(s *server) {
// 		s.readTimeout = 10
// 	}
// }

// func changeWriteTimeout(v int) option {
// 	return func(s *server) {
// 		s.writeTimeout = v
// 	}
// }

// func changeReadTimeout() func(s *server) {
// 	return func(s *server) {
// 		s.readTimeout = 10
// 	}
// }

// func changeWriteTimeout(v int) func(s *server) {
// 	return func(s *server) {
// 		s.writeTimeout = v
// 	}
// }

func main() {

	// even := func(val int) bool {
	// 	return val%2 == 0
	// }

	// l := filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, even)

	// l := filterV2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, even)

	// l := filterV2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(v int) bool {
	// 	return v%2 == 0
	// })

	// for _, v := range l {
	// 	fmt.Printf("v %v\n", v)
	// }

	// s := NewServer(changeReadTimeout(), changeWriteTimeout(20))
	// fmt.Printf("server %+v\n", s)

}

// func NewServer(options ...option) *server {
// 	svr := &server{}
// 	for _, o := range options {
// 		o(svr)
// 	}

// 	return svr
// }

// func NewServer(options ...func(*server)) *server {
// 	svr := &server{}
// 	for _, o := range options {
// 		o(svr)
// 	}

// 	return svr
// }
