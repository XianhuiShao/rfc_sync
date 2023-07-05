package rfc_module

func RfcModule() (funcname string, params map[string]interface{}) {
	//fmt.Println("module")
	//	a = "str"
	//params = map[string]interface{}{}

	//funcname = "RFC_SYSTEM_INFO"

	funcname, params = a2()
	return funcname, params
}

func a1() (funcname string, params map[string]interface{}) {

	funcname = "RFC_SYSTEM_INFO"
	return funcname, params
}

func a2() (funcname string, params map[string]interface{}) {

	var url string
	var interface_id string

	interface_id = "SAP_0001"

	url = "https://10.4.134.124:443/api/capp-eetl/v1/receive/" + interface_id

	params = map[string]interface{}{
		"CAPP_INPUT": map[string]interface{}{
			"INTERFACE_ID": interface_id,
			"EXT":          "",
			"SPRAS":        "",
			"AS_URL":       url,
			"DEBUG":        "",
		},
	}

	funcname = "ZAS_CAPP_NOTICE"
	return funcname, params
}
