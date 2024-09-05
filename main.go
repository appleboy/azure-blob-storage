package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/joho/godotenv"
)

var (
	Version     string
	Commit      string
	showVersion bool
)

func main() {
	var envfile string
	flag.StringVar(&envfile, "env-file", ".env", "Read in a file of environment variables")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.Parse()

	if showVersion {
		fmt.Printf("Version: %s Commit: %s\n", Version, Commit)
		return
	}

	_ = godotenv.Load(envfile)

	accountName := getGlobalValue("accountName")
	accountKey := getGlobalValue("accountKey")
	containerName := getGlobalValue("containerName")
	blobName := getGlobalValue("blobName")

	if accountName == "" || accountKey == "" || containerName == "" || blobName == "" {
		fmt.Println("Please provide all the required parameters")
		return
	}

	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		panic(err)
	}

	accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", accountName)

	sasQueryParams, err := sas.BlobSignatureValues{
		BlobName:   blobName,
		Protocol:   sas.ProtocolHTTPS,
		StartTime:  time.Now().UTC(),
		ExpiryTime: time.Now().UTC().Add(48 * time.Hour),
		Permissions: to.Ptr(sas.BlobPermissions{
			Read: true,
		}).String(),
		ContainerName: containerName,
	}.SignWithSharedKey(credential)
	if err != nil {
		panic(err)
	}

	sasURL := fmt.Sprintf("%s/%s/%s?%s",
		accountURL,
		containerName,
		blobName,
		sasQueryParams.Encode(),
	)
	fmt.Println(sasURL)
}
