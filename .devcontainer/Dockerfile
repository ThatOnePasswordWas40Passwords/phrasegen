FROM ubuntu:noble

ARG GO_VERS=1.24.3
ARG GOLINT_VERS=2.1.6
ARG JUST_VERS=1.40.0

ENV DEBIAN_FRONTEND=NONINTERACTIVE
ENV GOROOT=/usr/local/bin/go
ENV GOPATH=/root/go
ENV PATH=${GOPATH}/bin:${GOROOT}/bin:$PATH

RUN apt-get update && \
    apt-get upgrade -yq && \
    # install base deps
    apt-get install -yq wget curl jq git tar && \
    # install golang
    wget https://go.dev/dl/go${GO_VERS}.linux-amd64.tar.gz && \
        tar xzvf go*.linux-amd64.tar.gz && \
        chmod +x ./go && \
        mv go /usr/local/bin/ && \
        rm -rf go* && \
        curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b $(go env GOPATH)/bin v${GOLINT_VERS} && \
    # install just
    curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | bash -s -- --to /usr/local/bin && \
    # cleanup temp deps
    apt-get autoremove --purge -yq curl wget && \
    apt-get clean

RUN go install -v golang.org/x/tools/gopls@latest