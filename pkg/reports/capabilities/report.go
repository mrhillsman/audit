package capabilities

import (
	"encoding/json"

	"github.com/operator-framework/audit/pkg"
)

type Report struct {
	Columns    []Column
	Flags      BindFlags
	GenerateAt string
}

func (r *Report) writeJSON() error {
	data, err := json.Marshal(r)
	if err != nil {
		return err
	}

	const reportType = "capabilities"
	return pkg.WriteJSON(data, r.Flags.IndexImage, r.Flags.OutputPath, reportType)
}
