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
	"git.code.oa.com/geeker/awesome-work/pingcap/block"
	"github.com/spf13/cobra"
	"k8s.io/klog"
)

var size int
var path string
// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate data for test",
	Long: `generate data for test.`,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Infof("generate %d block,path:%s",size,path)
		block.GenerateBlock(size,path)
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntVar(&size,"size", 100,  "default size is 100")
	generateCmd.Flags().StringVar(&path,"path", "test.log",  "default out file :test.log")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
