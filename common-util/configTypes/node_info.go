package configTypes

import (
	"encoding/json"

	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
)

type NodeInfo struct {
	Name     string         `json:"name"`
	Host     string         `json:"host"`
	HttpPort int64          `json:"port"`
	GrpcPort int64          `json:"port"`
	Account  *types.Account `json:"nodeAccount"`
	Config   *types.Config  `json:"nodeConfig"`
}

func (this NodeInfo) ToPrettyJson() string {
	bytes, err := json.MarshalIndent(this, "", "  ")
	if err != nil {
		utils.Error("unable to marshal object", err)
		return ""
	}
	return string(bytes)
}
