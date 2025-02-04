package main

import (
	"fmt"
	"os"

	"setup/generated/hashicorp/google/provider"
	"setup/generated/hashicorp/google_beta/googlecomputeinstancefrommachineimage"
	bprovider "setup/generated/hashicorp/google_beta/provider"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/goccy/go-yaml"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// asia-northeast1-a, asia-northeast2-a, asia-northeast3-a
// 各リージョン7個ずつ
const zone = "asia-northeast2-a"

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
	as := 60020

	stack := cdktf.NewTerraformStack(scope, &id)

	credfile, err := os.ReadFile("../gcp-credentials.json")
	if err != nil {
		panic(err)
	}

	provider.NewGoogleProvider(stack, jsii.String("InternetArchLectureProject"), &provider.GoogleProviderConfig{
		Project:     jsii.String("internetarchlecture-372008"),
		Zone:        jsii.String(zone),
		Credentials: jsii.String(string(credfile)),
	})

	bprov := bprovider.NewGoogleBetaProvider(stack, jsii.String("InternetArchLectureProjectBeta"), &bprovider.GoogleBetaProviderConfig{
		Project:     jsii.String("internetarchlecture-372008"),
		Zone:        jsii.String(zone),
		Credentials: jsii.String(string(credfile)),
	})

	// ユーザー情報を読み込む
	err = loadUsers("../users.yaml")
	if err != nil {
		panic(err)
	}

	// network := computenetwork.NewComputeNetwork(stack, jsii.String("InternetArchLectureNetwork"), &computenetwork.ComputeNetworkConfig{
	// 	Name: jsii.String("internet-arch-lecture-network"),
	// })

	// subnet := computesubnetwork.NewComputeSubnetwork(stack, jsii.String("InternetArchLecturePrivateSubnet"), &computesubnetwork.ComputeSubnetworkConfig{
	// 	Name:        jsii.String("internet-arch-lecture-private-subnet"),
	// 	Network:     network.Id(),
	// 	IpCidrRange: jsii.String(fmt.Sprintf("192.168..0/24", as)),
	// })

	for _, user := range participants {
		pubKey := map[string]*string{
			"ssh-keys": jsii.String(fmt.Sprintf("%s:%s", user.UserID, user.PublicKey)),
		}

		googlecomputeinstancefrommachineimage.NewGoogleComputeInstanceFromMachineImage(
			stack,
			jsii.String(fmt.Sprintf("InternetArchLectureInstance-%s-%d", user.UserID, as)),
			&googlecomputeinstancefrommachineimage.GoogleComputeInstanceFromMachineImageConfig{
				Provider:           bprov,
				Name:               jsii.String(fmt.Sprintf("%s-%d", user.UserID, as)),
				MachineType:        jsii.String("e2-micro"),
				SourceMachineImage: jsii.String("projects/internetarchlecture-372008/global/machineImages/internetarchlecture-participants"),
				Zone:               jsii.String(zone),
				Metadata:           &pubKey,
			},
		)

		// computeinstance.NewComputeInstance(stack, jsii.String(fmt.Sprintf("InternetArchLectureInstance-%s-%d", user.UserID, as)), &computeinstance.ComputeInstanceConfig{
		// 	Name:        jsii.String(fmt.Sprintf("%s-%d", user.UserID, as)),
		// 	MachineType: jsii.String("e2-micro"),
		// 	BootDisk: &computeinstance.ComputeInstanceBootDisk{
		// 		InitializeParams: &computeinstance.ComputeInstanceBootDiskInitializeParams{
		// 			Image: jsii.String("ubuntu-os-cloud/ubuntu-2004-lts"),
		// 			Size:  jsii.Number(30),
		// 			Type:  jsii.String("pd-standard"),
		// 		},
		// 	},
		// NetworkInterface: []computeinstance.ComputeInstanceNetworkInterface{
		// 	{
		// 		Network: jsii.String("default"),
		// 		AccessConfig: []computeinstance.ComputeInstanceNetworkInterfaceAccessConfig{
		// 			{
		// 				NatIp: address.Address(),
		// 			},
		// 		},
		// 	},
		// },
		// 	// MetadataStartupScript: jsii.String(fmt.Sprintf(`#!/bin/bash`),
		// 	Zone:     jsii.String("us-central1-a"),
		// 	Metadata: &pubKey,
		// })

		as++
	}

	return stack
}

func main() {
	app := cdktf.NewApp(nil)

	// gcp-tf-cdk
	// gcp-tf-cdk-2
	// gcp-tf-cdk-3
	NewMyStack(app, "gcp-tf-cdk-3")

	app.Synth()
}
