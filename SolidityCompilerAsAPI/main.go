package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/nic0lae/JerryMouse/Servers"
)

// EvmVersion
// 		homestead,
// 		tangerineWhistle,
// 		spuriousDragon,
// 		byzantium (default)
// 		constantinople
type SolcInput struct {
	CodeAsBase64 string `json:"codeAsBase64"`
	EvmVersion   string `json:"evmVersion,omitempty"`
}

type SolcOutput struct {
	CompilerMessages       string `json:"compilerMessages,omitempty"`
	CompilerMetadaAsBase64 string `json:"compilerMetadaAsBase64,omitempty"`
}

func codeCompilerHandler(data []byte) Servers.JsonResponse {
	// 1. Read code to compile
	var solcInput = &SolcInput{}

	err := json.Unmarshal(data, solcInput)
	if err != nil {
		return Servers.JsonResponse{Error: fmt.Sprintf("%v", err)}
	}

	solidityCode, _ := base64.StdEncoding.DecodeString(solcInput.CodeAsBase64)

	var tempFile, _ = ioutil.TempFile("", "")
	tempFile.WriteString(string(solidityCode))

	// 2. Run Solidity Compiler
	if len(strings.TrimSpace(solcInput.EvmVersion)) == 0 {
		solcInput.EvmVersion = "byzantium"
	}

	var execScript = fmt.Sprintf(
		"solc %s --evm-version %s --combined-json abi,bin",
		tempFile.Name(),
		solcInput.EvmVersion,
	)
	cmd := exec.Command("sh", "-c", execScript)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	outputData, _ := ioutil.ReadAll(stdout)
	errorData, _ := ioutil.ReadAll(stderr)

	// 3. Send back the Compiler output
	var response Servers.JsonResponse
	response.Data = &SolcOutput{
		CompilerMessages:       string(errorData),
		CompilerMetadaAsBase64: base64.StdEncoding.EncodeToString(outputData),
	}

	return response
}

func main() {
	apiServer := Servers.Api()

	apiServer.SetJsonHandlers([]Servers.JsonHandler{
		Servers.JsonHandler{
			Route:      "/Compile",
			Handler:    codeCompilerHandler,
			JsonObject: &SolcInput{},
		},
	})

	apiServer.Run(":9999")
}
