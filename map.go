package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//ages := map[string]int{
	//	"alice": 31,
	//	"charlie": 34,
	//}
	////ages = make(map[string]int)
	//ages["bob"]++
	//for name, age := range ages {
	//	fmt.Printf("%s\t%d\n", name, age)
	//}
	//names := make([]string, 0, len(ages))
	//fmt.Println(names)
	//for name := range ages {
	//	names = append(names, name)
	//}
	//sort.Strings(names)
	//for _, name := range names {
	//	fmt.Printf("%s\t%d\n", name, ages[name])
	//}

	//ages := make(map[string]int)
	//var ages map[string]int
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages == nil)    // "true"
	fmt.Println(len(ages) == 0) // "true"
	ages["carol"] = 21          // panic: assignment to entry in nil map
	if age, ok := ages["bob"]; !ok {
		fmt.Println(ok)
		fmt.Println(age)
	}
	//if age, ok := ages["bob"]; !ok { /* ... */ }
	fmt.Println(equal(ages, ages))

	//scanLine()

	testList := []string{"hhh", "xxx", "qqq"}
	fmt.Println(k(testList))
	Add(testList)
	fmt.Println(Count(testList))
}

// 判断两个map是否相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	// 不能用以下方法
	// 例如equal(map[string]int{"A": 0}, map[string]int{"B": 42})
	//for k, xv := range x {
	//	if xv != y[k] {
	//		return false
	//	}
	//}
	return true
}

func scanLine() {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

var m = make(map[string]int)

func k(list []string) string  { return fmt.Sprintf("%q", list) }
func Add(list []string)       { m[k(list)]++ }
func Count(list []string) int { return m[k(list)] }
