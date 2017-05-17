package main

import (
	"encoding/asn1"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Before marshalling: ", t.String())

	mdata, _ := asn1.Marshal(t)
	fmt.Println("Marshalled ok")

	var newTime = new(time.Time)
	//	_, _ = asn1.Unmarshal(mdata, newTime)
	asn1.Unmarshal(mdata, newTime)
	fmt.Println("After marshal/unmarshal: ", newTime.String())

	s := "hello \u00bc"
	fmt.Println("Before marshal/unmarshal: ", s)

	mdataString, _ := asn1.Marshal(s)
	fmt.Println("Marshalled ok")

	var newStr string
	//_, _ = asn1.Unmarshal(mdataString, &newStr)
	asn1.Unmarshal(mdataString, &newStr)

	fmt.Println("Afer marshal/unmarshal: ", newStr)
}
