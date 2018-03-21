package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

var seedsCount = 1
var seedImageProject = "debian-cloud"
var seedImageFamily = "debian-9"
var seedMachineType = "f1-micro"
var seedTags = "disgo-node"
var seedStartupScript = "vm-debian9-configure.sh"

// var delegatesCount = 21
// var nodesCount = 50
var vmPrefix = "test-net-1-1"

func main() {
	rawData, _ := ioutil.ReadFile(seedStartupScript)
	seedStartupScriptContent := string(rawData)

	for i := 0; i < seedsCount; i++ {
		exec.Command("sh", "-c", fmt.Sprintf(
			"gcloud compute instances create %s-seed-%d --image-project %s --image-family %s --machine-type %s --tags %s --metadata startup-script='%s'",
			vmPrefix,
			i,
			seedImageProject,
			seedImageFamily,
			seedMachineType,
			seedTags,
			seedStartupScriptContent,
		)).Run()
	}
}
