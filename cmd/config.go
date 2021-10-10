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

	"Fsync/app"

	"github.com/spf13/cobra"
	"github.com/sqweek/dialog"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure a new path",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		config()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

func config(){
	os.OpenFile("config.ini", os.O_RDONLY|os.O_CREATE, 0666)

	src, dest := app.GetPath()

	ok := dialog.Message("src: %v\ndest: %v", src, dest).Title("Save these paths?").YesNo()

	if ok{
		src_path := fmt.Sprintf("%v,", src)
		dest_path := fmt.Sprintf("%v,", dest)

		app.WritePathINI(src_path, dest_path)
	}else{
		fmt.Println("canceled by user")
	}
}