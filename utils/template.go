package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func ParseTemplate(templateFileName string, data interface{}) string {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		fmt.Println("err", err)
		return ""
	}

	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		fmt.Println("err", err)
		return ""
	}

	return buffer.String()
}
