package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	jjvercore "github.com/jjliggett/jjversion/jjvercore"
)

func main() {
	v := jjvercore.CalculateVersion()

	println("::group::jjversion application JSON output")
	println(v.Json())
	println("::endgroup::")

	println("::group::$GITHUB_OUTPUT environment value")
	output := os.Getenv("GITHUB_OUTPUT")
	println(output)
	println("::endgroup::")

	file, err := os.OpenFile(output, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}

	content := fmt.Sprintf("major=%d\n"+
		"minor=%d\n"+
		"patch=%d\n"+
		"majorMinorPatch=%d.%d.%d\n"+
		"sha=%s\n"+
		"shortSha=%s\n",
		v.Major, v.Minor, v.Patch, v.Major, v.Minor, v.Patch, v.Sha, v.Sha[:7])

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	result, err := ioutil.ReadFile(output)
	if err != nil {
		log.Fatal(err)
	}

	println("::group::$GITHUB_OUTPUT value")
	text := string(result)
	println(text)
	println("::endgroup::")

}
