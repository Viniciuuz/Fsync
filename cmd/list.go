/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

// pathsCmd represents the paths command
var pathsCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all folders \"queued\" to sync",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		ListPaths()
	},
}

func init() {
	rootCmd.AddCommand(pathsCmd)
}

func ListPaths(){
	cfg, err := ini.Load("config.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }

	src := cfg.Section("paths").Key("src").String()
	dest := cfg.Section("paths").Key("dest").String()

	src_slice := strings.Split(src, ",")
	dest_slice := strings.Split(dest, ",")

	fmt.Printf("Total Paths: %v\n\n", len(src_slice) - 1)

	for index := range src_slice {
		if src_slice[index] == ""{
			continue
		}
		fmt.Printf("src: %v\ndest: %v\n\n", src_slice[index], dest_slice[index])
	}

}