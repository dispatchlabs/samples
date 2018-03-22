package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dispatchlabs/commons/config"
)

// VMsConfig - config for a VM in Google Cloud
type VMsConfig struct {
	ImageProject     string
	ImageFamily      string
	MachineType      string
	Tags             string
	NamePrefix       string
	ScriptConfigURL  string
	ScriptConfigFile string
}

func main() {
	var seedsCount = 3
	var delegatesCount = 6
	var nodesCount = 10
	var namePrefix = "temp-test-net-1-1"

	// Create SEEDs
	createVMs(
		seedsCount,
		VMsConfig{
			ImageProject:     "debian-cloud",
			ImageFamily:      "debian-9",
			MachineType:      "f1-micro",
			Tags:             "disgo-node",
			NamePrefix:       namePrefix + "-seed",
			ScriptConfigURL:  "https://raw.githubusercontent.com/dispatchlabs/samples/master/google-cloud-spawn-vms",
			ScriptConfigFile: "vm-debian9-configure.sh",
		},
		config.DisgoProperties{
			HttpPort:          1975,
			HttpHostIp:        "0.0.0.0",
			GrpcPort:          1973,
			GrpcTimeout:       5,
			UseQuantumEntropy: false,
			IsSeed:            true,
			IsDelegate:        false,
			SeedList:          []string{},
			DaposDelegates:    []string{},
			NodeId:            "",
			ThisIp:            "",
		},
	)

	var seedIPList = getSeedIPs(seedsCount, namePrefix+"-seed")

	// Create DELEGATEs
	createVMs(
		delegatesCount,
		VMsConfig{
			ImageProject:     "debian-cloud",
			ImageFamily:      "debian-9",
			MachineType:      "f1-micro",
			Tags:             "disgo-node",
			NamePrefix:       namePrefix + "-delegate",
			ScriptConfigURL:  "https://raw.githubusercontent.com/dispatchlabs/samples/master/google-cloud-spawn-vms",
			ScriptConfigFile: "vm-debian9-configure.sh",
		},
		config.DisgoProperties{
			HttpPort:          1975,
			HttpHostIp:        "0.0.0.0",
			GrpcPort:          1973,
			GrpcTimeout:       5,
			UseQuantumEntropy: false,
			IsSeed:            false,
			IsDelegate:        true,
			SeedList:          seedIPList,
			DaposDelegates:    []string{},
			NodeId:            "",
			ThisIp:            "",
		},
	)

	// Create NODEs
	createVMs(
		nodesCount,
		VMsConfig{
			ImageProject:     "debian-cloud",
			ImageFamily:      "debian-9",
			MachineType:      "f1-micro",
			Tags:             "disgo-node",
			NamePrefix:       namePrefix + "-node",
			ScriptConfigURL:  "https://raw.githubusercontent.com/dispatchlabs/samples/master/google-cloud-spawn-vms",
			ScriptConfigFile: "vm-debian9-configure.sh",
		},
		config.DisgoProperties{
			HttpPort:          1975,
			HttpHostIp:        "0.0.0.0",
			GrpcPort:          1973,
			GrpcTimeout:       5,
			UseQuantumEntropy: false,
			IsSeed:            false,
			IsDelegate:        false,
			SeedList:          seedIPList,
			DaposDelegates:    []string{},
			NodeId:            "",
			ThisIp:            "",
		},
	)
}

func createVMs(count int, vmsConfig VMsConfig, disgoConfig config.DisgoProperties) {
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {

		var vmName = fmt.Sprintf("%s-%d", vmsConfig.NamePrefix, i)

		// Command to CREATE new VM Instance
		var createVM = fmt.Sprintf(
			"gcloud compute instances create %s --image-project %s --image-family %s --machine-type %s --tags %s",
			vmName,
			vmsConfig.ImageProject,
			vmsConfig.ImageFamily,
			vmsConfig.MachineType,
			vmsConfig.Tags,
		)

		// Command to DOWNLOAD BASH scripts to the newly created VM
		var downloadScriptFiles = fmt.Sprintf(
			"gcloud compute ssh %s --command 'curl %s/%s -o %s'",
			vmName,
			vmsConfig.ScriptConfigURL,
			vmsConfig.ScriptConfigFile,
			vmsConfig.ScriptConfigFile,
		)

		// Commands to RUN scripts
		var execScript = fmt.Sprintf(
			"gcloud compute ssh %s --command 'bash %s'",
			vmName,
			vmsConfig.ScriptConfigFile,
		)

		// RUN VM creation in PARALLEL
		// Run COMMANDS inside the VM in SEQUENTIAL order
		wg.Add(1)
		go func(vmName string, disgoConfig config.DisgoProperties, cmds ...string) {
			for _, cmd := range cmds {
				exec.Command("sh", "-c", cmd).Run()
			}

			disgoConfig.NodeId = vmName
			disgoConfig.ThisIp = getVMIP(vmName)

			// Save JSON config to a temp file then upload that file to the VM
			var configFileName = randString(20) + ".json"
			file, error := os.Create(configFileName)
			if error == nil {
				bytes, error := json.Marshal(&config.Properties)
				if error == nil {
					fmt.Fprintf(file, string(bytes))
					file.Close()

					var fullFileName, _ = filepath.Abs(configFileName)

					exec.Command("sh", "-c", fmt.Sprintf("gcloud compute scp %s %s:~/config.json", fullFileName, vmName)).Run()
					exec.Command("sh", "-c", fmt.Sprintf("gcloud compute ssh %s --command 'sudo mv ~/config.json /go-binaries/config/ && sudo chown -R dispatch-services:dispatch-services /go-binaries'", vmName)).Run()

					os.Remove(configFileName)
				}
			}

			wg.Done()
		}(vmName, disgoConfig, createVM, downloadScriptFiles, execScript)
	}

	wg.Wait()
}

func getVMIP(vmName string) string {
	out, err := exec.Command("sh", "-c", fmt.Sprintf(
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

func getSeedIPs(seedsCount int, namePrefix string) []string {
	var seedIPList = []string{}

	for i := 0; i < seedsCount; i++ {
		var vmName = fmt.Sprintf("%s-%d", namePrefix, i)

		var ip = getVMIP(vmName)
		if ip != "" {
			seedIPList = append(seedIPList, ip)
		}
	}

	return seedIPList
}

func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
