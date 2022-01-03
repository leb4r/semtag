package main

import (
	"github.com/leb4r/semtag/cmd"
	"github.com/leb4r/semtag/pkg/utils"
)

func main() {
	if err := cmd.Execute(); err != nil {
		utils.ThrowError(err)
	}
}
