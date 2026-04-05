/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	// "os"
	"path/filepath"

	"github.com/artur-spr/gotask/cmd"
	"github.com/artur-spr/gotask/conf"
)

func main() {
	cmd.Execute()
	log.Println(conf.Global)
}
