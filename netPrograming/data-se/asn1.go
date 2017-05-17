package main

import (
	"encoding/asn1"
	"fmt"
	"os"
)

func main() {
	v := 998
	fmt.Println("Before marshal/unmarshal: ", v)

	mdata, err := asn1.Marshal(v)
	checkError(err)

	var i int
	what, err := asn1.Unmarshal(mdata, &i)
	checkError(err)

	fmt.Println("After marshal/unmarshal: ", i)
	fmt.Println("and ", what)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(2)
	}
}
