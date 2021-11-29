package main

import (
	"flag"
	"fmt"
	"husd.com/trace/trace"
)

/**
 *
 * @author hushengdong
 */
func main() {

	port := flag.Int("port", 10910, "监听的端口，默认是9876")
	flag.Parse()
	fmt.Println("------------- start -----------------port:", *port)
	trace.StartTrace(*port)
}
