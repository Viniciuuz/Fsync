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

	"github.com/Viniciuuz/Fsync/app"
	"github.com/spf13/cobra"
	"gopkg.in/ini.v1"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync folders",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		Sync()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
}

func Sync(){
	cfg, err := ini.Load("config.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }

	src := cfg.Section("paths").Key("src").String()
	dest := cfg.Section("paths").Key("dest").String()

	src_slice := strings.Split(src, ",")
	dest_slice := strings.Split(dest, ",")

	fmt.Println("[Starting backup...]")

	for index := range src_slice {
		if src_slice[index] == ""{
			continue
		}
		fmt.Println("\nBacking up: "+ src_slice[index])
		app.CopyDir(src_slice[index], dest_slice[index])
	}

	fmt.Println("\n[Backup complete]")
}