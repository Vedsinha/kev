/**
 * Copyright 2020 Appvia Ltd <info@appvia.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"os"

	"github.com/appvia/kev/pkg/kev"
	"github.com/appvia/kev/pkg/kev/log"
	"github.com/spf13/cobra"
)

var renderLongDesc = `(render) render Kubernetes manifests in selected format.

Examples:

  ### Render an app Kubernetes manifests (default) for all environments
  $ kev render

  ### Render an app Kubernetes manifests (default) for a specific environment(s)
  $ kev render -e staging [-e production ...]`

var renderCmd = &cobra.Command{
	Use:   "render",
	Short: "Generates application's deployment artefacts according to the specified output format for a given environment (ALL environments by default).",
	Long:  renderLongDesc,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		fns := []func(cmd *cobra.Command, args []string) error{
			runReconcileCmd,
			runDetectSecretsCmd,
		}
		for _, fn := range fns {
			if err := fn(cmd, args); err != nil {
				return err
			}
			os.Stdout.Write([]byte("\n"))
		}
		return nil
	},
	RunE: runRenderCmd,
}

func init() {
	flags := renderCmd.Flags()
	flags.SortFlags = false

	flags.StringP(
		"format",
		"f",
		"kubernetes", // default: native kubernetes manifests
		"Deployment files format. Default: Kubernetes manifests.",
	)

	flags.BoolP(
		"single",
		"s",
		false, // default: produce multiple files. If true then a single file will be produced.
		"Controls whether to produce individual manifests or a single file output. Default: false",
	)

	flags.StringP(
		"dir",
		"d",
		"", // default: will output kubernetes manifests in k8s/<env>...
		"Override default Kubernetes manifests output directory. Default: k8s/<env>",
	)

	flags.StringSliceP(
		"environment",
		"e",
		[]string{},
		"Target environment for which deployment files should be rendered",
	)

	rootCmd.AddCommand(renderCmd)
}

func runRenderCmd(cmd *cobra.Command, _ []string) error {
	cmdName := "Render"

	format, err := cmd.Flags().GetString("format")
	singleFile, err := cmd.Flags().GetBool("single")
	dir, err := cmd.Flags().GetString("dir")
	envs, err := cmd.Flags().GetStringSlice("environment")
	verbose, _ := cmd.Root().Flags().GetBool("verbose")

	if err != nil {
		return err
	}

	setReporting(verbose)
	displayCmdStarted(cmdName)

	workingDir, err := os.Getwd()
	if err != nil {
		return displayError(err)
	}

	log.DebugTitlef("Output format: %s", format)
	if err := kev.Render(workingDir, format, singleFile, dir, envs); err != nil {
		return err
	}

	os.Stdout.Write([]byte("\n"))

	return nil
}
