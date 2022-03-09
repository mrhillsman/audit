package capabilities

import "github.com/operator-framework/audit/pkg/models"

type Column struct {
	PackageName     string   `json:"packageName"`
	BundleImagePath string   `json:"bundleImagePath,omitempty"`
	DefaultChannel  string   `json:"defaultChannel,omitempty"`
	Channels        []string `json:"bundleChannel,omitempty"`
	AuditErrors     []string `json:"errors,omitempty"`
	Capabilities    bool     `json:"capabilities"`
}

func NewColumn(v models.AuditCapabilities) *Column {
	col := Column{}
	col.PackageName = v.PackageName
	col.BundleImagePath = v.OperatorBundleImagePath
	col.DefaultChannel = v.DefaultChannel
	col.Channels = v.Channels
	col.AuditErrors = v.Errors
	col.Capabilities = v.Capabilities

	return &col
}
