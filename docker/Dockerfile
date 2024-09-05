FROM alpine:3.20

ARG TARGETOS
ARG TARGETARCH

LABEL maintainer="Bo-Yi Wu <appleboy.tw@gmail.com>" \
  org.label-schema.name="Azure Blob Storage" \
  org.label-schema.vendor="Bo-Yi Wu" \
  org.label-schema.schema-version="1.0"

LABEL org.opencontainers.image.source=https://github.com/appleboy/azure-blob-storage
LABEL org.opencontainers.image.description="Sync Gitea Secrets"
LABEL org.opencontainers.image.licenses=MIT

RUN apk add --no-cache ca-certificates && \
  rm -rf /var/cache/apk/*

COPY release/${TARGETOS}/${TARGETARCH}/azure-blob-storage /bin/

ENTRYPOINT ["/bin/azure-blob-storage"]
