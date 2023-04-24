package main

import "fmt"

type Test struct {
	int1 int
	int2 int
	int3 int
	int4 int
}

type Test1 struct {
}

func main() {
	test := Test{1, 2, 3, 4}
	fmt.Println(test)
	fmt.Printf("%p %p %p %p\n", &test.int1, &test.int2, &test.int3, &test.int4)
}

/*
{1 2 3 4}
0xc00012a000 0xc00012a008 0xc00012a010 0xc00012a018
*/
