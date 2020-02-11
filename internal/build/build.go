package build

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/nori-io/makisu/internal/build/parser"
)

const (
	pluginFileName = "plugin.go"
)

func Build(dir string) error {
	meta := parser.ExtractMeta(path.Join(dir, pluginFileName))
	_, err := meta.GetCore().GetConstraint()
	if err != nil {
		return err
	}


	// todo: check pkg versions in go.mod - must fit for provided core version in plugin meta
	// todo: build plugin name from meta.ID as id_version: ID("nori/Auth@0.1.0") -> nori_auth_0.1.0
	// todo: select needed docker image and tag for plugin core constraint

	return nil
}


func docker(dir, img string) error {
	sh := exec.Command("docker", "run",
		"--rm",
		"-e", "build_mode=single",
		"-v", fmt.Sprintf("%s:/plugins", dir),
		img)
	sh.Stdout = os.Stdout
	sh.Stderr = os.Stderr
	err := sh.Run()
	if err != nil {
		fmt.Printf("%s\n",err.Error())
	}
	return nil
}