package capabilities

import (
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/operator-framework/audit/pkg"
	"github.com/operator-framework/audit/pkg/models"
)

type Data struct {
	AuditBundle       []models.AuditBundle
	Flags             BindFlags
	IndexImageInspect pkg.DockerInspect
}

func (d *Data) BuildCapabilitiesQuery() (string, error) {
	query := sq.Select("o.name, o.bundlepath").From(
		"operatorbundle o")

	if d.Flags.HeadOnly {
		query = sq.Select("o.name, o.bundlepath").From(
			"operatorbundle o, channel c")
		query = query.Where("c.head_operatorbundle_name == o.name")
	}
	if d.Flags.Limit > 0 {
		query = query.Limit(uint64(d.Flags.Limit))
	}
	if len(d.Flags.Filter) > 0 {
		query = sq.Select("o.name, o.bundlepath").From(
			"operatorbundle o, channel_entry c")
		like := "'%" + d.Flags.Filter + "%'"
		query = query.Where(fmt.Sprintf("c.operatorbundle_name == o.name AND c.package_name like %s", like))
	}

	query.OrderBy("o.name")

	sql, _, err := query.ToSql()
	if err != nil {
		return "", fmt.Errorf("unable to create sql : %s", err)
	}
	return sql, nil
}
