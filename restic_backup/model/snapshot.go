package model

import (
	"time"
)

type Snapshot struct {
	Id             string    `json:"id"               comment:"快照ID"`
	ShortId        string    `json:"short_id"         comment:"快照短ID"`
	Time           time.Time `json:"time"             comment:"快照时间"`
	Hostname       string    `json:"hostname"         comment:"主机名"`
	Tags           []string  `json:"tags"             comment:"标签"`
	Paths          []string  `json:"paths"            comment:"路径"`
	Tree           string    `json:"tree"             comment:"树"`
	Username       string    `json:"username"         comment:"用户名"`
	ProgramVersion string    `json:"program_version"  comment:"restic版本"`
}
