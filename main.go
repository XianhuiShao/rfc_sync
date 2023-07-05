package main

import (
	"fmt"
	"rfc_sync/rfc_module"

	"github.com/sap/gorfc/gorfc"
)

func abapSystem() gorfc.ConnectionParameters {
	return gorfc.ConnectionParameters{
		"user":   "user01",
		"passwd": "Sap@202301",
		"ashost": "10.4.112.97",
		"sysnr":  "00",
		"client": "800",
		"lang":   "ZH",
	}
}

func main() {
	c, err := gorfc.ConnectionFromParams(abapSystem())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Connected:", c.Alive())

	//function module name
	//var a string
	funcname, params := rfc_module.RfcModule()
	//fmt.Println(a)
	attrs, _ := c.GetConnectionAttributes()
	fmt.Println("Connection attributes", attrs)

	//params := map[string]interface{}{}
	r, e := c.Call(funcname, params)

	if e != nil {
		fmt.Println(e)
		return
	}
	// 输出结果
	fmt.Printf("Response: %#v \n", r)

	c.Close()
}
