// Copyright 2019 Authors of protobuf-gadgets
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"fmt"
	"os"

	"github.com/dunbit/protobuf-gadgets/internal/protoc-gen-docmaster/docmaster"
	"github.com/dunbit/protobuf-gadgets/pkg/protoc"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "protoc-gen-docmaster",
	Short: "Docmaster Plugin for Protoc",
	Long:  `Docmaster Plugin for Protoc to generate documentation`,
	Run:   run,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func run(cmd *cobra.Command, args []string) {
	protoc.Run(docmaster.NewPlugin())
}
