package main

import "fmt"

/**
 *
 * @author hushengdong
 */
func main() {

	//port := flag.Int("port", 10910, "监听的端口，默认是9876")
	//flag.Parse()
	//fmt.Println("------------- start -----------------port:", *port)
	//trace.StartTrace(*port)

	a := 1
	defer fmt.Println(a)
	a = 2
}
