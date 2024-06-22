package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
        //export DOCKER_API_VERSION='1.45'
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Containers found by Docker Golang client\n")
	for _, ctr := range containers {
		fmt.Printf("%s %s %s (status: %s)\n", ctr.ID, ctr.Image, ctr.State, ctr.Status)
	}

}
