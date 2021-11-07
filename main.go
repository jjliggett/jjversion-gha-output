package main

import (
	"fmt"

	jjvercore "github.com/jjliggett/jjversion/jjvercore"
)

func main() {
	v := jjvercore.CalculateVersion()

	println("::group::jjversion application JSON output")
	println(v.Json())
	println("::endgroup::")

	fmt.Printf("::set-output name=major::%d\n", v.Major)
	fmt.Printf("::set-output name=minor::%d\n", v.Minor)
	fmt.Printf("::set-output name=patch::%d\n", v.Patch)
	fmt.Printf("::set-output name=majorMinorPatch::%d.%d.%d\n", v.Major, v.Minor, v.Patch)
	fmt.Printf("::set-output name=sha::%s\n", v.Sha)
	fmt.Printf("::set-output name=shortSha::%s\n", v.Sha[:7])

}
