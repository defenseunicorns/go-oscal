package main

import "github.com/defenseunicorns/go-oscal/src/cmd"

//go:generate ./hack/gen-types.sh
func main() {
	cmd.Execute()
}
