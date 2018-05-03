package main

import "fmt"

func main() {

	m := make(map[string]int) // 初始化map m key是string 值是int

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"] // _空标识符 任何语法需要变量名但是程序逻辑不需要这个变量

	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1 ,"bar": 2}
    fmt.Println("map:", n)
}

/*
map map[k2:13 k1:7]
v1: 7
len: 2
map: map[k1:7]
prs: false
map: map[foo:1 bar:2]
*/