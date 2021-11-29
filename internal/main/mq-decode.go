package main

import (
	"flag"
	"fmt"
	"husd.com/trace/trace"
)

func main1() {

	file := flag.String("f", ".", "tcpdump的文件内容")
	flag.Parse()
	fmt.Println("------------- start -----------------file:", *file)
	*file = "/tmp/a1.txt"
	trace.DecodeTcpBytes(*file)
}
