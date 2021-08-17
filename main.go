package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

func main() {
	out, err := os.Create("index.html")
	defer out.Close()
	check(err)

	data := map[string]interface{}{}

	buf, err := ioutil.ReadFile("data.yml")
	check(err)

	err = yaml.Unmarshal(buf, &data)
	check(err)

	t, err := template.ParseGlob("./templates/*")
	check(err)

	fmt.Println(t)

	err = t.Execute(out, data)
	check(err)
}
