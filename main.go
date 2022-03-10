package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"pluck/pkg/config"
	"text/template"
)

func main() {
	fmt.Printf("%s\n", os.Args)
	valuesFile := flag.String("v", "", "-v <valuesFile.yaml>")
	templateFile := flag.String("t", "", "-t <templateFile.gotext>")
	flag.Parse()
	context, err := config.LoadGlobalConfig(*valuesFile)
	handleError(err)
	templateContext, err := ioutil.ReadFile(*templateFile)
	handleError(err)
	parse, err := template.New("compose").Parse(string(templateContext))
	handleError(err)
	err = parse.Execute(os.Stdout, context)
	handleError(err)

}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error while processing %s", err)
		os.Exit(1)
	}
}
