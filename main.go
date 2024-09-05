package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/sas"
)

type input struct {
	accountName   string
	accountKey    string
	containerName string
	blobName      string
}

func main() {
	params := &input{}
	flag.StringVar(&params.accountName, "accountName", "", "The storage account name")
	flag.StringVar(&params.accountKey, "accountKey", "", "The storage account key")
	flag.StringVar(&params.containerName, "containerName", "", "The container name")
	flag.StringVar(&params.blobName, "blobName", "", "The blob name")
	flag.Parse()

	if params.accountName == "" || params.accountKey == "" || params.containerName == "" || params.blobName == "" {
		fmt.Println("Please provide all the required parameters")
		return
	}

	credential, err := azblob.NewSharedKeyCredential(params.accountName, params.accountKey)
	if err != nil {
		panic(err)
	}

	accountURL := fmt.Sprintf("https://%s.blob.core.windows.net", params.accountName)

	sasQueryParams, err := sas.BlobSignatureValues{
		BlobName:   params.blobName,
		Protocol:   sas.ProtocolHTTPS,
		StartTime:  time.Now().UTC(),
		ExpiryTime: time.Now().UTC().Add(48 * time.Hour),
		Permissions: to.Ptr(sas.BlobPermissions{
			Read: true,
		}).String(),
		ContainerName: params.containerName,
	}.SignWithSharedKey(credential)
	if err != nil {
		panic(err)
	}

	sasURL := fmt.Sprintf("%s/%s/%s?%s",
		accountURL,
		params.containerName,
		params.blobName,
		sasQueryParams.Encode(),
	)
	fmt.Println(sasURL)
}
