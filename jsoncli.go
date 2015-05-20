package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bitly/go-simplejson"
)

func main() {

	r := bufio.NewReader(os.Stdin)

	d, err := r.ReadString('\n')
	j, err := simplejson.NewJson(d)
	if err != nil {
		fmt.Println("simplejson.NewJson error: %s", err)
	}

	v, err := j.EncodePretty()
	if err != nil {
		fmt.Println("simplejson.EncodePretty error", err)
	}
	fmt.Println(string(v))
}
