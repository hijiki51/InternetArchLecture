package main

import (
	"os"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/goccy/go-yaml"

	"github.com/cdktf/cdktf-provider-google-go/google/v4/provider"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type User struct {
	UserID    string `yaml:"userID"`
	PublicKey string `yaml:"publicKey"`
}

var (
	admin        User
	participants []User
)

func loadUsers(file string) error {
	// yamlファイルを読み込む
	buf, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	// yamlファイルを構造体に変換する
	var users struct {
		Admin        User   `yaml:"admin"`
		Participants []User `yaml:"participants"`
	}
	err = yaml.Unmarshal(buf, &users)
	if err != nil {
		return err
	}
	admin = users.Admin
	participants = users.Participants

	return nil
}

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	provider.NewGoogleProvider(scope, jsii.String("InternetArchLecture"), &provider.GoogleProviderConfig{
		Project: jsii.String("internet-arch-lecture"),
		Region:  jsii.String("asia-northeast1"),
	})

	// ユーザー情報を読み込む
	err := loadUsers("../users.yaml")
	if err != nil {
		panic(err)
	}

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "gcp-tf-cdk")

	app.Synth()
}