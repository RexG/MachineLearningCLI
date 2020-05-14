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
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/spf13/cobra"
)

var (
	local     bool
	localPort int
	localName string
	Stg, Prod bool
)

// trainCmd represents the train command
var trainCmd = &cobra.Command{
	Use:   "train",
	Short: "Train your ML model",
	Long:  `Train your ML model`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			panic(err)
		}

		//imageName := "jupyter/datascience-notebook:latest"
		imageName := "jupyter/datascience-notebook:r-3.6.2"

		if local {
			//pullJupyterImage(imageName, ctx, cli)
			createJupyterInstance(imageName, localName, localPort, ctx, cli)
			openbrowser("http://localhost:" + strconv.Itoa(localPort))
		} else if Stg {
			// create Jupyter notebook in  stg
		} else {
			// create Jupyter notebook in  prod
		}
	},
}

func init() {
	modelCmd.AddCommand(trainCmd)
	// for local
	trainCmd.Flags().BoolVarP(&local, "local", "l", false, "Train ML models in your local Jupyter notebook")
	trainCmd.Flags().IntVar(&localPort, "port", 8888, "Specify your local Jupyter instance port")
	trainCmd.Flags().StringVar(&localName, "name", "local-", "Name your local Jupyter instance")
	// for stg
	trainCmd.Flags().BoolVarP(&Stg, "-stg", "s", false, "Train ML models in  STG Jupyter notebooks")
	// for prod
	trainCmd.Flags().BoolVarP(&Prod, "-prod", "p", false, "Train ML models in  PROD Jupyter notebooks")
}

// pull  Jupyter image to local
func pullJupyterImage(imageName string, ctx context.Context, cli *client.Client) {
	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)
}

func createJupyterInstance(imageName string, localName string, localPort int, ctx context.Context, cli *client.Client) {

	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: strconv.Itoa(localPort),
	}
	containerPort, _ := nat.NewPort("tcp", "8888")
	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Cmd:   []string{"start-notebook.sh", "--NotebookApp.token=''", "--NotebookApp.password=''"},
	}, &container.HostConfig{
		PortBindings: portBinding,
	}, nil, localName)

	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}
