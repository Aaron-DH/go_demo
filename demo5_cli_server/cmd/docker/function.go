package docker

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"os/exec"
	"strings"
)

func listimages(ctx *cli.Context) error {
	var cmd *exec.Cmd

	cmd = exec.Command("docker", "images")
	images, err := cmd.Output()
	if err != nil {
		fmt.Println("List images error", err)
		return err
	}

	fmt.Println(string(images))
	return nil
}

func pullimages(ctx *cli.Context) error {
	if len(ctx.Args()) < 1 {
		return fmt.Errorf("misssing image's repo and tag")
	}

	for _, imageName := range ctx.Args() {
		if !strings.Contains(imageName, ":") {
			imageName = imageName + ":latest"
		}
		cmd := exec.Command("docker", "pull", imageName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("failed to pull image %s: %v",
				imageName, err)
		}
	}
	return nil
}
