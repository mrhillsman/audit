# Build the audit-tool binary
FROM --platform=$BUILDPLATFORM golang:1.17 as builder
ARG TARGETARCH

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . .

# Build
RUN GOOS=linux GOARCH=$TARGETARCH go build -a -o build/audit-tool cmd/main.go

# Final image.
FROM registry.access.redhat.com/ubi8/ubi

ENV HOME=/opt/audit-tool \
    USER_NAME=audit-tool \
    USER_UID=1001

RUN echo "${USER_NAME}:x:${USER_UID}:0:${USER_NAME} user:${HOME}:/sbin/nologin" >> /etc/passwd
RUN dnf install -y podman

WORKDIR ${HOME}

# Add operator-sdk binary
RUN curl -Lfo /usr/local/bin/operator-sdk https://github.com/operator-framework/operator-sdk/releases/download/v${OPERATOR_SDK_VERSION:-1.14.0}/operator-sdk_${OS:-linux}_${ARCH:-amd64} \
    && chmod +x /usr/local/bin/operator-sdk

RUN chown -R audit-tool: ${HOME}

#COPY --from=builder /workspace/bundlelist.json /opt/audit-tool/bundlelist.json
COPY --from=builder /workspace/build/audit-tool /usr/local/bin/audit-tool

ENTRYPOINT ["/usr/local/bin/audit-tool"]

USER ${USER_UID}
