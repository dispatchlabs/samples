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
	"github.com/dispatchlabs/samples/common-util/config-helpers"

	golog "github.com/nic0lae/golog"
	gologC "github.com/nic0lae/golog/contracts"
	gologM "github.com/nic0lae/golog/modifiers"
	gologP "github.com/nic0lae/golog/persisters"
)

// SeedsCount - nr of SEED(s) to spawn
const SeedsCount = 1

// DelegatesCount - nr of DELEGATE(s) to spawn
const DelegatesCount = 5

// NamePrefix - VM name prefix
const NamePrefix = "perf-test-9-11"

// CodeBranch - Brnach of the code to deploy
const CodeBranch = "master"

// VMConfig - config for a VM in Google Cloud
type VMConfig struct {
	ImageProject        string
	ImageFamily         string
	MachineTemplate     string
	MachineType         string
	Tags                string
	NamePrefix          string
	ScriptPrefixURL     string
	ScriptNewNode       string
	ScriptUpdateNode    string
	ScriptNewScandis    string
	ScriptUpdateScandis string
	CodeBranch          string
}

// NodeConfigParams -
type NodeConfigParams struct {
	Config  *types.Config
	Node    *types.Node
	Account *types.Account
}

func main() {
	inmemoryLogger := configureLogging()
	fmt.Println("Running, please wait...")

	// 1. Create SEED VMs => NON configured seeds, but we have the VMs IPs
	// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
	seedVMConfig := getDefaultVMConfig()
	seedVMConfig.NamePrefix = NamePrefix + "-seed"
	createDisgoNodeVMs(SeedsCount, &seedVMConfig)

	var seedsAccounts = []*types.Account{}
	for i := 0; i < SeedsCount; i++ {
		var vmName = fmt.Sprintf("%s-%d", seedVMConfig.NamePrefix, i)
		seedsAccounts = append(seedsAccounts, configHelpers.CreateAccount(vmName))
	}

	var seedAddressToNodeConfigs = fetchSeedsVMsMappings(seedVMConfig.NamePrefix, SeedsCount, seedsAccounts)

	// 2. Create DELEGATE VMs => NON configured delegates, but we have the VMs IPs
	// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
	delegateVMConfig := getDefaultVMConfig()
	delegateVMConfig.NamePrefix = NamePrefix + "-delegate"
	createDisgoNodeVMs(DelegatesCount, &delegateVMConfig)

	var delegatesAccounts = []*types.Account{}
	for i := 0; i < DelegatesCount; i++ {
		var vmName = fmt.Sprintf("%s-%d", delegateVMConfig.NamePrefix, i)
		delegatesAccounts = append(delegatesAccounts, configHelpers.CreateAccount(vmName))
	}

	configDelegateVMs(delegateVMConfig.NamePrefix, DelegatesCount, delegatesAccounts, seedAddressToNodeConfigs)

	// 3. Update seeds configs
	// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
	configSeedVMs(seedAddressToNodeConfigs, delegatesAccounts)

	// 4. Stasrt SEEDs and Delegtes
	for i := 0; i < SeedsCount; i++ {
		var vmName = fmt.Sprintf("%s-%d", seedVMConfig.NamePrefix, i)
		restartDisgoNodeOnVM(vmName)
	}
	for i := 0; i < DelegatesCount; i++ {
		var vmName = fmt.Sprintf("%s-%d", delegateVMConfig.NamePrefix, i)
		restartDisgoNodeOnVM(vmName)
	}

	// 5. Create Scandis VM
	// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
	scandisVMConfig := getDefaultVMConfig()
	scandisVMConfig.Tags = "http-server,https-server"
	scandisVMConfig.NamePrefix = NamePrefix + "-scandis"
	scandisVMConfig.ScriptNewScandis = "vm-debian9-new-scandis.sh"

	createScandisVMs(&scandisVMConfig, seedVMConfig.NamePrefix+"-0")

	fmt.Println("Done.")

	// Dump log
	(inmemoryLogger.(*gologM.InmemoryLogger)).Flush()
}

func configureLogging() gologC.Logger {
	var inmemoryLogger = gologM.NewInmemoryLogger(
		gologM.NewSimpleFormatterLogger(
			gologM.NewMultiLogger(
				[]gologC.Logger{
					gologP.NewConsoleLogger(),
					gologP.NewFileLogger("new-nodes.log"),
				},
			),
		),
	)

	golog.StoreSingleton(
		golog.NewLogger(
			inmemoryLogger,
		),
	)

	return inmemoryLogger
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

func getVmCreateString(vmConfig *VMConfig, name string) string {
	var createVM string

	if len(vmConfig.MachineTemplate) > 0 {
		createVM = fmt.Sprintf(
			"gcloud compute instances create %s --source-instance-template %s --image-project %s --image-family %s --tags %s",
			name,
			vmConfig.MachineTemplate,
			vmConfig.ImageProject,
			vmConfig.ImageFamily,
			vmConfig.Tags,
		)
	} else {

		// Command to CREATE new VM Instance
		createVM = fmt.Sprintf(
			"gcloud compute instances create %s --image-project %s --image-family %s --machine-type %s --tags %s",
			name,
			vmConfig.ImageProject,
			vmConfig.ImageFamily,
			vmConfig.MachineType,
			vmConfig.Tags,
		)
	}

	return createVM
}
func createDisgoNodeVMs(count int, vmConfig *VMConfig) {
	var wg sync.WaitGroup

	for i := 0; i < count; i++ {
		var vmName = fmt.Sprintf("%s-%d", vmConfig.NamePrefix, i)

		// Command to CREATE new VM Instance
		var createVM = getVmCreateString(vmConfig, vmName)

		fmt.Println(createVM)

		// Command to DOWNLOAD BASH scripts to the newly created VM
		var downloadScriptFiles = fmt.Sprintf(
			"gcloud compute ssh %s --command 'curl %s/%s -o %s'",
			vmName,
			vmConfig.ScriptPrefixURL,
			vmConfig.ScriptNewNode,
			vmConfig.ScriptNewNode,
		)

		// Commands to RUN scripts
		var execScript = fmt.Sprintf(
			"gcloud compute ssh %s --command 'bash %s %s'",
			vmName,
			vmConfig.ScriptNewNode,
			vmConfig.CodeBranch,
		)

		// RUN VM creation in PARALLEL
		// Run COMMANDS inside the VM in SEQUENTIAL order
		wg.Add(1)
		go func(vritualMachineName string, cmds ...string) {
			golog.Instance().LogInfo(vritualMachineName, 0, "~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
			golog.Instance().LogInfo(vritualMachineName, 0, "DEPLOY START")

			for _, cmd := range cmds {
				golog.Instance().LogInfo(vritualMachineName, 4, cmd)
				exec.Command(getOSC(), getOSE(), cmd).Run()
			}

			wg.Done()
		}(vmName, createVM, downloadScriptFiles, execScript)
	}

	wg.Wait()
}

func createScandisVMs(vmConfig *VMConfig, seedVmName string) {
	var wg sync.WaitGroup
	var createVM string

	createVM = getVmCreateString(vmConfig, seedVmName)

	// Command to DOWNLOAD BASH scripts to the newly created VM
	var downloadScriptFiles = fmt.Sprintf(
		"gcloud compute ssh %s --command 'curl %s/%s -o %s'",
		vmConfig.NamePrefix,
		vmConfig.ScriptPrefixURL,
		vmConfig.ScriptNewScandis,
		vmConfig.ScriptNewScandis,
	)

	// Commands to RUN scripts
	var seedVMIP = getVMIP(seedVmName)

	var execScript = fmt.Sprintf(
		"gcloud compute ssh %s --command 'bash %s %s %s:1975'",
		vmConfig.NamePrefix,
		vmConfig.ScriptNewScandis,
		vmConfig.CodeBranch,
		seedVMIP,
	)

	// RUN VM creation in PARALLEL
	// Run COMMANDS inside the VM in SEQUENTIAL order
	wg.Add(1)
	go func(vritualMachineName string, cmds ...string) {
		golog.Instance().LogInfo(vritualMachineName, 0, "~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~")
		golog.Instance().LogInfo(vritualMachineName, 0, "DEPLOY SCANDIS START")

		for _, cmd := range cmds {
			golog.Instance().LogInfo(vritualMachineName, 4, cmd)
			exec.Command(getOSC(), getOSE(), cmd).Run()
		}

		wg.Done()
	}(vmConfig.NamePrefix, createVM, downloadScriptFiles, execScript)

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

func getVMsIPs(count int, namePrefix string) []string {
	var vmsIPs = []string{}

	for i := 0; i < count; i++ {
		var vmName = fmt.Sprintf("%s-%d", namePrefix, i)

		var ip = getVMIP(vmName)
		if ip != "" {
			vmsIPs = append(vmsIPs, ip)
		}
	}

	return vmsIPs
}

func randString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func replaceConfigOnVM(vmName string, nodeConfigParams *NodeConfigParams) {
	var configFileName = randString(20) + ".json"
	var accountFileName = randString(20) + ".json"

	file1, err1 := os.Create(configFileName)
	file2, err2 := os.Create(accountFileName)

	if err1 == nil && err2 == nil {
		bytes1, err1 := json.Marshal(nodeConfigParams.Config)
		bytes2, err2 := json.Marshal(nodeConfigParams.Account)

		if err1 == nil && err2 == nil {

			// Save JSON config to a temp file
			fmt.Fprintf(file1, string(bytes1))
			fmt.Fprintf(file2, string(bytes2))

			file1.Close()
			file2.Close()

			var fullFileName1, _ = filepath.Abs(configFileName)
			var fullFileName2, _ = filepath.Abs(accountFileName)

			// Upload temp file to the VM
			var cmd1 = fmt.Sprintf("gcloud compute scp %s %s:~/config.json", fullFileName1, vmName)
			var cmd2 = fmt.Sprintf("gcloud compute scp %s %s:~/account.json", fullFileName2, vmName)

			var cmd3 = fmt.Sprintf("gcloud compute ssh %s --command 'sudo mv ~/config.json /go-binaries/config/ && sudo chown -R dispatch-services:dispatch-services /go-binaries'", vmName)
			var cmd4 = fmt.Sprintf("gcloud compute ssh %s --command 'sudo mv ~/account.json /go-binaries/config/ && sudo chown -R dispatch-services:dispatch-services /go-binaries'", vmName)

			golog.Instance().LogInfo(vmName, 4, cmd1)
			golog.Instance().LogInfo(vmName, 4, cmd2)
			golog.Instance().LogInfo(vmName, 4, cmd3)
			golog.Instance().LogInfo(vmName, 4, cmd4)

			exec.Command(getOSC(), getOSE(), cmd1).Run()
			exec.Command(getOSC(), getOSE(), cmd2).Run()
			exec.Command(getOSC(), getOSE(), cmd3).Run()
			exec.Command(getOSC(), getOSE(), cmd4).Run()

			// Remove temp file
			os.Remove(fullFileName1)
			os.Remove(fullFileName2)
		}
	}
}

func restartDisgoNodeOnVM(vmName string) {
	var cmd1 = fmt.Sprintf("gcloud compute ssh %s --command 'sudo sudo systemctl restart dispatch-disgo-node'", vmName)
	golog.Instance().LogInfo(vmName, 4, cmd1)
	exec.Command(getOSC(), getOSE(), cmd1).Run()
}

func fetchSeedsVMsMappings(seedNamePrefix string, seedCount int, seedsAccounts []*types.Account) map[string]*NodeConfigParams {
	var defaultNodeConfig = types.GetDefaultConfig()

	var seedAddressToNodeConfigs = map[string]*NodeConfigParams{}

	for i := 0; i < seedCount; i++ {

		var vmName = fmt.Sprintf("%s-%d", seedNamePrefix, i)
		var seedIP = getVMIP(vmName)

		var seedAddress = seedsAccounts[i].Address

		var config = configHelpers.GetSeedConfig(
			seedIP,
			defaultNodeConfig.HttpEndpoint.Port,
			defaultNodeConfig.GrpcEndpoint.Port,
			seedsAccounts,
		)

		var node = &types.Node{
			Address:      seedAddress,
			GrpcEndpoint: config.GrpcEndpoint,
			HttpEndpoint: config.HttpEndpoint,
			Type:         types.TypeSeed,
		}

		seedAddressToNodeConfigs[vmName] = &NodeConfigParams{
			Config:  config,
			Node:    node,
			Account: seedsAccounts[i],
		}
	}

	return seedAddressToNodeConfigs
}

func configSeedVMs(
	seedVMToNodeConfigs map[string]*NodeConfigParams,
	delegatesAccounts []*types.Account,
) {
	var delegateAddresses = []string{}
	for _, adress := range delegatesAccounts {
		delegateAddresses = append(delegateAddresses, adress.Address)
	}

	for vmName, nodeConfigs := range seedVMToNodeConfigs {
		nodeConfigs.Config.DelegateAddresses = delegateAddresses
		replaceConfigOnVM(vmName, nodeConfigs)
	}
}

func configDelegateVMs(
	delegateNamePrefix string,
	delegateCount int,
	delegatesAccounts []*types.Account,
	seedVMToNodeConfigs map[string]*NodeConfigParams,
) {
	var defaultNodeConfig = types.GetDefaultConfig()

	listOfSeedNodes := []*types.Node{}
	for _, v := range seedVMToNodeConfigs {
		listOfSeedNodes = append(listOfSeedNodes, v.Node)
	}

	for i := 0; i < delegateCount; i++ {
		var vmName = fmt.Sprintf("%s-%d", delegateNamePrefix, i)
		vmIP := getVMIP(vmName)

		account := delegatesAccounts[i]

		config := configHelpers.GetDelegateConfig(
			vmIP,
			defaultNodeConfig.HttpEndpoint.Port,
			defaultNodeConfig.GrpcEndpoint.Port,
			listOfSeedNodes,
		)

		replaceConfigOnVM(vmName, &NodeConfigParams{
			Config:  config,
			Node:    nil,
			Account: account,
		})
	}
}

func getDefaultVMConfig() VMConfig {

	// NOTE:  MachineTemplate overrides MachineType

	return VMConfig{
		ImageProject:    "debian-cloud",
		ImageFamily:     "debian-9",
		MachineTemplate: "perf-test-1",
		MachineType:     "n1-standard-2",
		Tags:            "disgo-node",
		NamePrefix:      NamePrefix + "-abracadabra",
		ScriptPrefixURL: "https://raw.githubusercontent.com/dispatchlabs/samples/dev/deployment",
		ScriptNewNode:   "vm-debian9-new-node.sh",
		CodeBranch:      CodeBranch,
	}
}
