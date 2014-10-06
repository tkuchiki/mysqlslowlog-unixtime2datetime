package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func usage() {
	fmt.Println("Usage: mysqlslowlog-unixtime2datetime <mysqlslowlog> [-help]")
	fmt.Println("")
	fmt.Println("Options:")
	fmt.Println("  -h  show this help message and exit")
	os.Exit(0)
}

func main() {
	var help bool
	var fp *os.File
	var err error

	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.BoolVar(&help, "help", false, "help")
	f.Parse(os.Args[1:])

	if help {
		usage()
	}

	if len(os.Args) >= 2 {
		fp, err = os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer fp.Close()
	} else {
		usage()
	}

	assined := regexp.MustCompile("SET timestamp=(.*);")

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		bs := []byte(scanner.Text())
		group := assined.FindSubmatch(bs)
		if group != nil {
			i64, _ := strconv.ParseInt(string(group[1]), 10, 64)
			datetime := time.Unix(i64, 0)
			fmt.Printf("%s ", scanner.Text())
			fmt.Println(fmt.Sprintf("(%s)", datetime))
		} else {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
