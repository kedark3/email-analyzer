FROM registry.svc.ci.openshift.org/openshift/release:golang-1.10 as builder
RUN go get github.com/kedark3/email-analyzer
WORKDIR /go/src/github.com/kedark3/email-analyzer
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/analyzeEmails main.go

FROM scratch
LABEL author="Kedar Kulkarni"

WORKDIR /email-analyzer
COPY --from=builder /go/bin/analyzeEmails /email-analyzer/analyzeEmails
EXPOSE 8080

ENTRYPOINT ["./analyzeEmails"]