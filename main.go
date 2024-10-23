package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
	"github.com/appleboy/com/gh"
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

	accountName := getGlobalValue("account_name")
	accountKey := getGlobalValue("account_key")
	containerName := getGlobalValue("container_name")
	blobName := getGlobalValue("blob_name")
	duration := getGlobalValue("duration")
	github := getGlobalValue("github")

	if accountName == "" || accountKey == "" || containerName == "" || blobName == "" || duration == "" {
		fmt.Println("Please provide all the required parameters")
		return
	}

	txpireTime, err := time.ParseDuration(duration)
	if err != nil {
		panic(err)
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
		ExpiryTime: time.Now().UTC().Add(txpireTime),
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
	fmt.Println("Blob URL:", sasURL)
	fmt.Println("Expires at:", time.Now().UTC().Add(txpireTime))

	if github != "" {
		gh.SetOutput(map[string]string{
			"blob_url":       sasURL,
			"expire_at":      time.Now().UTC().Add(txpireTime).String(),
			"expire_at_unix": fmt.Sprintf("%d", time.Now().UTC().Add(txpireTime).Unix()),
		})
	}
}
