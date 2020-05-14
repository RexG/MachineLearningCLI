/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

var uploadModel struct {
	modelName string
	modelPath string
}

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload local model to model repository",
	Long:  `Upload local model to model repository`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("model upload successfully")
	},
}

func init() {
	modelCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVar(&uploadModel.modelName, "model-name", "", "your model name")
	uploadCmd.Flags().StringVar(&uploadModel.modelPath, "model-path", "", "your model local path")
}
