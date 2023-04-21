/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"embed"
	"github.com/41443658/cgen/cmd"
)

//go:embed template/*
var template embed.FS

func main() {
	cmd.TFS = template
	cmd.Execute()
}
