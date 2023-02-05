package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"runtime"
)

// PrettyPrint to print struct in a readable way
func PrettyPrintStruct(i interface{}) string {
	s, _ := json.MarshalIndent(i, " ", "\t")
	return string(s)
}

// https://gosamples.dev/pretty-print-json/
func PrettyPrintJSONString(responseBody string) (string, error) {
	return PrettyPrintJSONResponse([]byte(responseBody))
}

func PrettyPrintJSONResponse(responseBody []byte) (string, error) {
	var buffer bytes.Buffer
	error := json.Indent(&buffer, responseBody, "", "  ")
	if error != nil {
		return "", error
	}
	return buffer.String(), nil
}

func CheckErr(err error) {
	if err != nil {
		pc, filename, line, _ := runtime.Caller(1)
		log.Fatalf("[error] %v (%s in %s:%d)", err, runtime.FuncForPC(pc).Name(), filename, line)
	}
}
