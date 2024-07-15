package model

import (
	"fmt"
	"runtime"
)

// 存储库
type Repository struct {
	Uri         string `json:"uri"          yaml:"uri"          comment:"存储库uri"`
	Password    string `json:"password"     yaml:"password"     comment:"存储库密码"`
	AwsUser     string `json:"aws_user"     yaml:"aws_user"     comment:"AmazonS3秘钥ID"`
	AwsPassword string `json:"aws_password" yaml:"aws_password" comment:"AmazonS3秘钥"`
}

// 设置环境变量
func (r *Repository) SetEnv() string {
	switch runtime.GOOS {
	case "linux":
		return fmt.Sprintf(
			`export RESTIC_REPOSITORY="%s";export RESTIC_PASSWORD="%s";export AWS_ACCESS_KEY_ID="%s";export AWS_SECRET_ACCESS_KEY="%s";`,
			r.Uri, r.Password, r.AwsUser, r.AwsPassword)
	case "windows":
		return fmt.Sprintf(
			`set RESTIC_REPOSITORY=%s&&set RESTIC_PASSWORD=%s&&set AWS_ACCESS_KEY_ID=%s&&set AWS_SECRET_ACCESS_KEY=%s&&`,
			r.Uri, r.Password, r.AwsUser, r.AwsPassword)
	default:
		return ""
	}
}
