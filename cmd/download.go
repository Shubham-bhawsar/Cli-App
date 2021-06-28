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
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var url string
var downloadPath string

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `This is a Cli App that enable you to directly download any file with the help of URL,
	You are just needed to pass the URL as argument,And the path where you wanna save that particular file.
	 For example:
	 Cli-App download --url="https://golangcode.com/logo.svg" --path="/$HOME/Downloads/abc.svg"`,
	Run: func(cmd *cobra.Command, args []string) {

		err := DownloadFile(url)
		if err != nil {
			panic(err)
		}
		fmt.Println("Downloaded: " + url)
	},
}

func DownloadFile(url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(downloadPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func init() {

	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().StringVar(
		&url, "url", "https://golangcode.com/logo.svg", "URL for the specified file needed to download",
	)
	downloadCmd.Flags().StringVar(
		&downloadPath, "path", "d", "Path where you wanted to get your file saved",
	)
}
