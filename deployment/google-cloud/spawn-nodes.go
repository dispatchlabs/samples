package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/dispatchlabs/disgo/commons/types"
)

// SeedsCount - nr of SEED(s) to spawn
const SeedsCount = 1

// DelegatesCount - nr of DELEGATE(s) to spawn
const DelegatesCount = 3

// NodesCount - nr of NODE(s) to spawn
const NodesCount = 0

// NamePrefix - VM name prefix
const NamePrefix = "test-nicolae"

// VMConfig - config for a VM in Google Cloud
type VMConfig struct {
	ImageProject     string
	ImageFamily      string
	MachineType      string
	Tags             string
	NamePrefix       string
	ScriptConfigURL  string
	ScriptConfigFile string
	CodeBranch       string
}

func main() {
	var defaultVMConfig = VMConfig{
		ImageProject:     "debian-cloud",
		ImageFamily:      "debian-9",
		MachineType:      "n1-standard-2",
		Tags:             "disgo-node",
		NamePrefix:       NamePrefix + "-seed",
		ScriptConfigURL:  "https://raw.githubusercontent.com/dispatchlabs/samples/dev/google-cloud-spawn-vms",
		ScriptConfigFile: "vm-debian9-configure.sh",
		CodeBranch:       "dev",
	}

	var defaultNodeConfig = types.Config{
		HttpEndpoint: &types.Endpoint{
			Host: "0.0.0.0",
			Port: 1975,
		},
		GrpcEndpoint: &types.Endpoint{
			Host: "0.0.0.0",
			Port: 1973,
		},
		GrpcTimeout:        5,
		UseQuantumEntropy:  false,
		SeedEndpoints:      []*types.Endpoint{},
		GenesisTransaction: `{"hash":"a48ff2bd1fb99d9170e2bae2f4ed94ed79dbc8c1002986f8054a369655e29276","type":0,"from":"e6098cc0d5c20c6c31c4d69f0201a02975264e94","to":"3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c","value":10000000,"data":"","time":0,"signature":"03c1fdb91cd10aa441e0025dd21def5ebe045762c1eeea0f6a3f7e63b27deb9c40e08b656a744f6c69c55f7cb41751eebd49c1eedfbd10b861834f0352c510b200","hertz":0,"fromName":"","toName":""}`,
	}

	// Create SEED VMs
	seedVMConfig := defaultVMConfig
	seedVMConfig.NamePrefix = NamePrefix + "-seed"

	createVMs(SeedsCount, &seedVMConfig, &defaultNodeConfig)

	var seedEndpoints = getEndpoints(SeedsCount, NamePrefix+"-seed")

	// Create DELEGATE VMs
	delegateVMConfig := defaultVMConfig
	delegateVMConfig.NamePrefix = NamePrefix + "-delegate"

	delegateConfig := defaultNodeConfig
	delegateConfig.SeedEndpoints = seedEndpoints

	createVMs(DelegatesCount, &delegateVMConfig, &delegateConfig)

	// Create NODE VMs
	nodeVMConfig := defaultVMConfig
	nodeVMConfig.NamePrefix = NamePrefix + "-node"

	delegateNodeConfig := defaultNodeConfig
	delegateNodeConfig.SeedEndpoints = seedEndpoints

	createVMs(NodesCount, &nodeVMConfig, &delegateNodeConfig)
}

// Get the underlying OS command shell
func getOSC() string {

	osc := "sh"
	if runtime.GOOS == "windows" {
		osc = "cmd"
	}

	return osc
}

// Get the shell/command startup option to execute commands
func getOSE() string {

	ose := "-c"
	if runtime.GOOS == "windows" {
		ose = "/c"
	}
	return ose
}

func createVMs(count int, vmConfig *VMConfig, nodeConfig *types.Config) {
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		var vmName = fmt.Sprintf("%s-%d", vmConfig.NamePrefix, i)

		// Command to CREATE new VM Instance
		var createVM = fmt.Sprintf(
			"gcloud compute instances create %s --image-project %s --image-family %s --machine-type %s --tags %s",
			vmName,
			vmConfig.ImageProject,
			vmConfig.ImageFamily,
			vmConfig.MachineType,
			vmConfig.Tags,
		)

		// Command to DOWNLOAD BASH scripts to the newly created VM
		var downloadScriptFiles = fmt.Sprintf(
			"gcloud compute ssh %s --command 'curl %s/%s -o %s'",
			vmName,
			vmConfig.ScriptConfigURL,
			vmConfig.ScriptConfigFile,
			vmConfig.ScriptConfigFile,
		)

		// Commands to RUN scripts
		var execScript = fmt.Sprintf(
			"gcloud compute ssh %s --command 'bash %s %s'",
			vmName,
			vmConfig.ScriptConfigFile,
			vmConfig.CodeBranch,
		)

		// RUN VM creation in PARALLEL
		// Run COMMANDS inside the VM in SEQUENTIAL order
		wg.Add(1)
		go func(vmName string, disgoConfig *types.Config, cmds ...string) {

			osc := getOSC()
			ose := getOSE()

			for _, cmd := range cmds {
				fmt.Println(fmt.Sprintf("RUN -> %s", cmd))
				exec.Command(osc, ose, cmd).Run()
			}

			disgoConfig.GrpcEndpoint.Host = getVMIP(vmName)

			// Save JSON config to a temp file then upload that file to the VM
			var configFileName = randString(20) + ".json"
			file, error := os.Create(configFileName)
			if error == nil {
				bytes, error := json.Marshal(&disgoConfig)
				if error == nil {
					fmt.Fprintf(file, string(bytes))
					file.Close()

					var fullFileName, _ = filepath.Abs(configFileName)

					exec.Command(osc, ose, fmt.Sprintf("gcloud compute scp %s %s:~/config.json", fullFileName, vmName)).Run()
					exec.Command(osc, ose, fmt.Sprintf("gcloud compute ssh %s --command 'sudo mv ~/config.json /go-binaries/config/ && sudo chown -R dispatch-services:dispatch-services /go-binaries'", vmName)).Run()
					exec.Command(osc, ose, fmt.Sprintf("gcloud compute ssh %s --command 'sudo sudo systemctl restart dispatch-disgo-node'", vmName)).Run()

					os.Remove(configFileName)
				}
			}

			wg.Done()
		}(vmName, nodeConfig, createVM, downloadScriptFiles, execScript)
	}

	wg.Wait()
}

func getVMIP(vmName string) string {
	out, err := exec.Command(getOSC(), getOSE(), fmt.Sprintf(
		"gcloud compute instances describe %s | grep natIP:",
		vmName,
	)).Output()

	if err == nil {
		var outputAsString = string(out)
		outputAsString = strings.Replace(outputAsString, "natIP: ", "", -1)
		return strings.TrimSpace(outputAsString)
	}

	return ""
}

func getEndpoints(count int, namePrefix string) []*types.Endpoint {
	var endpoints = []*types.Endpoint{}

	for i := 0; i < count; i++ {
		var vmName = fmt.Sprintf("%s-%d", namePrefix, i)

		var ip = getVMIP(vmName)
		if ip != "" {
			endpoints = append(endpoints, &types.Endpoint{Host: ip, Port: 1973})
		}
	}

	return endpoints
}

func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
