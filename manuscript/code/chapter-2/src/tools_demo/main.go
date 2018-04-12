package main

import (
	"fmt"
	"urlutil"
)

func main() {
	urls := []string{
		"http://www.github.com",
		"https://www.no-such-url-really.org",
		"blah://ftp.xyz.org",
	}

	for _, url := range urls {
		verdict := "not a valid"
		if urlutil.IsValidUrl(url) {
			verdict = "a valid"
		}

		fmt.Printf("%s is %s url\n", url, verdict)

		verdict = ""
		if !urlutil.IsReachable(url) {
			verdict = "not "
		}
		fmt.Printf("%s is %sreachable\n", url, verdict)
	}
}
