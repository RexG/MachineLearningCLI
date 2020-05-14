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

type Deployment struct {
	Stg            bool
	Prod           bool
	modelFramework string
	modelUrl       string
	apiName        string
	cpu            int
	memory         int
	replicaSet     int
}

var deployment = Deployment{}

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your model in  as a model-api (Restful API)",
	Long:  `Deploy your model in  as a model-api (Restful API)`,
	Run: func(cmd *cobra.Command, args []string) {

		var env = " STG"
		if deployment.Prod {
			env = " PROD"
		}

		fmt.Println("deploying your model ", deployment.apiName, "in", env, ", from:", deployment.modelUrl, ", model framework: ", deployment.modelFramework, "",
			", cpu:", deployment.cpu, ", memory:", deployment.memory, "Gi, how many instances:", deployment.replicaSet,
		)
		fmt.Println("....")
		fmt.Println("https://rexsoft.com/", deployment.apiName)

	},
}

func init() {
	modelCmd.AddCommand(deployCmd)

	deployCmd.Flags().BoolVar(&deployment.Stg, "-stg", false, "deploy your model in  STG")
	deployCmd.Flags().BoolVar(&deployment.Prod, "-prod", false, "deploy your model in  PROD")

	deployCmd.Flags().StringVar(&deployment.modelFramework, "model-framework", "MLflow", "which ML framework you used fro training this model, e.g. MLflow/Tensorflow/XGBoost/SKLearn")
	deployCmd.Flags().StringVar(&deployment.modelUrl, "model-url", "", "where your model is, a URL")
	deployCmd.Flags().StringVar(&deployment.apiName, "api-name", "", "your model-api name")
	deployCmd.Flags().IntVar(&deployment.cpu, "cpu", 1, "how many CPUs")
	deployCmd.Flags().IntVar(&deployment.memory, "memory", 1, "how many Gi memory for 1 model-api")
	deployCmd.Flags().IntVar(&deployment.replicaSet, "replica-set", 1, "how many model-api instances for 1 model-api")

}
