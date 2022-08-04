FROM golang:1.18.2-stretch

## UPDATE THE OS
RUN apt-get update && \
    apt-get install -y tzdata && \
    go install -v golang.org/x/tools/gopls@latest

WORKDIR /go/src

## SET ENVIRONMENT
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV TZ America/Sao_Paulo

## COPY NECESSARY FILES
COPY . .

# ## INSTALL MY STANDARD LIBRARIES 
# RUN go get github.com/spf13/viper && \
#     go get github.com/spf13/cobra && \
#     go get github.com/satori/go.uuid && \
#     go get github.com/gofiber/fiber/v2 && \
#     go get github.com/gofiber/jwt/v2 && \
#     go get github.com/gofiber/jwt/v2 && \
#     go get github.com/stretchr/testify

## TIDY THE PROJECT
RUN go mod download && \
    go mod tidy

## KEEP THE CONTAINER RUNNiNG
CMD ["tail", "-f", "/dev/null"]