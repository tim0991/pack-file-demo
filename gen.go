//+build ignore

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

var tplName = "packer.go"

var resources = make(map[string][]byte)

func main() {
	t, err := template.New("").Funcs(map[string]interface{}{"conv": FormatByteSlice}).Parse(`//DO NOT EDIT
package main

import "packer/box" 

func init(){
	{{- range $filename, $byte := . }}
    	box.Add("{{ $filename }}", []byte{ {{ conv $byte }} })
	{{- end }}
}
`)
	if err != nil {
		panic(err)
	}

	fi, err := os.Create(tplName)
	if err != nil {
		panic(err)
	}

	resources[".env"], _ = ioutil.ReadFile(".env")

	err = t.Execute(fi, resources)
	if err != nil {
		panic(err)
	}

}

func FormatByteSlice(sl []byte) string {
	builder := strings.Builder{}
	for _, v := range sl {
		builder.WriteString(fmt.Sprintf("%d,", int(v)))
	}
	return builder.String()
}
