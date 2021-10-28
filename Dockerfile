############################
# STEP 1 build executable binary
############################
FROM golang:1.12-alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

RUN adduser -D -g '' appuser
WORKDIR $GOPATH/src/github.com/evzpav/documents-crud-refactored/

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/documents-crud-refactored ./cmd/server/main.go

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /etc/passwd /etc/passwd

COPY --from=builder /go/bin/documents-crud-refactored /go/bin/documents-crud-refactored

USER appuser

ENTRYPOINT ["/go/bin/documents-crud-refactored"]
