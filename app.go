package main

import "fmt"

type obj interface {
	SetVal(int)
}

type objX struct {
	val int
}

func (p *objX) SetVal(val int) {
	p.val = val
}

func main() {
	lomo := make(map[obj]string)

	a := objX{5}
	// b := objX{15}
	// c := objX{20}
	// d := objX{5}
	fmt.Println(a)

	lomo[obj(&a)] = "a"
	// lomo[obj(&b)] = "b"
	// lomo[obj(&c)] = "c"
	// lomo[obj(&d)] = "d"

	a0 := interface{}(&a)
	fmt.Println(a0)
	fmt.Println(a0.(obj))

}
