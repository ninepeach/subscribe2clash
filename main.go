package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/magicst0ne/subscribe2clash/internal/subscribe"
	"github.com/magicst0ne/subscribe2clash/internal/xbase64"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: subscribe2clash -c config.yaml [url]\n")
    flag.PrintDefaults()
    os.Exit(2)
}

var (
    conf  string
)

func init() {
    flag.StringVar(&conf, "c", "config.yaml", "config file")
    flag.Usage = usage
}

func main() {

	flag.Usage = usage
    flag.Parse()

    args := flag.Args()
    if len(args) < 1 {
        usage()
    }

	content, err := subscribe.GetSubContent(args[0])
	if err != nil {
		fmt.Println(err)
	}

	proxies := subscribe.ParseProxy(content)
	config, err := subscribe.GenerateClashConfig(proxies, conf)
	if err != nil {
		fmt.Println(err)
	}

	config = config
	fmt.Println(xbase64.UnicodeEmojiDecode(string(config)))
}
