package main

import (
	"fmt"
	"os"

	"setup/generated/hashicorp/google/computeaddress"
	"setup/generated/hashicorp/google/computeinstance"
	"setup/generated/hashicorp/google/provider"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/goccy/go-yaml"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type User struct {
	UserID    string `yaml:"userID"`
	PublicKey string `yaml:"publicKey"`
}

// nolint
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
	as := 60000

	stack := cdktf.NewTerraformStack(scope, &id)

	credfile, err := os.ReadFile("../gcp-credentials.json")
	if err != nil {
		panic(err)
	}

	provider.NewGoogleProvider(stack, jsii.String("InternetArchLectureProject"), &provider.GoogleProviderConfig{
		Project:     jsii.String("internetarchlecture-372008"),
		Zone:        jsii.String("us-central1-a"),
		Credentials: jsii.String(string(credfile)),
	})

	// ユーザー情報を読み込む
	err = loadUsers("../users.yaml")
	if err != nil {
		panic(err)
	}

	for _, user := range participants {
		pubKey := map[string]*string{
			"ssh-keys": jsii.String(fmt.Sprintf("%s:%s", user.UserID, user.PublicKey)),
		}

		address := computeaddress.NewComputeAddress(stack, jsii.String("InternetArchLectureAddress"), &computeaddress.ComputeAddressConfig{
			Name:        jsii.String(fmt.Sprintf("address-%s-%d", user.UserID, as)),
			AddressType: jsii.String("EXTERNAL"),
		})

		computeinstance.NewComputeInstance(stack, jsii.String(fmt.Sprintf("InternetArchLectureInstance-%s-%d", user.UserID, as)), &computeinstance.ComputeInstanceConfig{
			Name:        jsii.String(fmt.Sprintf("%s-%d", user.UserID, as)),
			MachineType: jsii.String("e2-micro"),
			BootDisk: &computeinstance.ComputeInstanceBootDisk{
				InitializeParams: &computeinstance.ComputeInstanceBootDiskInitializeParams{
					Image: jsii.String("ubuntu-os-cloud/ubuntu-2004-lts"),
					Size:  jsii.Number(10),
					Type:  jsii.String("pd-standard"),
				},
			},
			NetworkInterface: []computeinstance.ComputeInstanceNetworkInterface{
				{
					Network: jsii.String("default"),
					AccessConfig: []computeinstance.ComputeInstanceNetworkInterfaceAccessConfig{
						{
							NatIp: address.Address(),
						},
					},
				},
			},
			Zone:     jsii.String("us-central1-a"),
			Metadata: &pubKey,
		})

		as++
	}

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "gcp-tf-cdk")

	app.Synth()
}
