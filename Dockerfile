FROM ghcr.io/foundry-rs/foundry AS contracts-builder

WORKDIR /contracts

COPY contracts/ ./

RUN forge install --no-commit --no-git foundry-rs/forge-std

RUN forge build --extra-output-files abi metadata --skip script --force

FROM golang:latest AS cli-builder

WORKDIR cli

COPY ./ ./

RUN go env -w GOPROXY=https://proxy.golang.org \
    && go mod download

RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-w -s" \
    -a -installsuffix cgo \
    -o /go/bin/godiamond ./cmd/godiamond

FROM alpine

ARG USER_ID=${UID}
ARG GROUP_ID=${GID}

WORKDIR /app

RUN adduser \
    --disabled-password \
    --gecos '' \
    --uid $USER_ID \
    diamonduser \
    && chown -R diamonduser:diamonduser /app

USER diamonduser

COPY --from=cli-builder /go/bin/godiamond /go/cli/config.yaml ./ 

CMD ["./diamond-cli", "loupe", "--rpc", "local", "--debug"]
