# azure-blob-storage

How to set Expiry Time for Azure Blob Storage URL.

## Introduction

In this article, we will learn how to set the expiry time for the Azure Blob Storage URL. We will use the Azure Storage SDK to generate the URL with the expiry time.

## Development

Create the new `.env` file and add the following configuration.

```env
ACCOUNT_NAME=xxxxxxx
ACCOUNT_KEY=xxxxxx
CONTAINER_NAME=testblob
BLOB_NAME=test.txt
DURATION=24h
```

Build the golang application.

```bash
make
```

Run the application.

```bash
./bin/azure-blob-storage
```
