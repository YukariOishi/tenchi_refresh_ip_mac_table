package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elliotchance/phpserialize"
	"github.com/mostlygeek/arp"
)

func main() {
	var (
		d = flag.String("d", "./", "laravel cache directory")
	)
	flag.Parse()
	if *d == "" {
		// larabel cache file path
		// how to use
		log.Println("ex. main.exe -d ./storage/framework/cache/data")
		return
	}

	table := arp.Table()
	for ip, mac := range table {
		nm := strings.ToUpper(strings.ReplaceAll(mac, ":", ""))
		key := fmt.Sprintf("key:%s", nm)
		log.Println("key->ip:", key, "->", ip) // debug log

		s := string(phpserialize.MarshalString(ip))
		// cache hold seconds

		fc := fmt.Sprintf("9999999999%s", s) // fc: file contents
		h := sha1.New()
		h.Write([]byte([]byte(key)))
		hk := fmt.Sprintf("%x", h.Sum(nil))

		p1 := hk[0:2]
		p2 := hk[2:4]
		hd := fmt.Sprintf("%s/%s/%s", *d, p1, p2)
		if err := os.MkdirAll(hd, 0755); err != nil {
			// create directry error
			log.Println("cannot create directory", err)
			continue
		}

		fp := fmt.Sprintf("%s/%s", hd, hk) // fc: file path

		if err := os.WriteFile(fp, []byte(fc), 0644); err != nil {
			// create cache file error
			log.Println("failed write cache file", err)
			continue
		}

		log.Println("fp=>fc", fp, "->", fc) //debug log
	}
}
