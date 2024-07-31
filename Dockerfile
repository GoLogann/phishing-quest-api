FROM golang:1.21.3-alpine3.18 as build

WORKDIR /app
COPY . .
RUN go mod tidy && go build -tags musl -o main

# Etapa de runtime
FROM alpine:3.18

WORKDIR /app

# Adicionando dados de timezone
RUN apk add --no-cache tzdata
ENV TZ America/Sao_Paulo

# Adicionando certificados raiz
RUN apk update && apk --no-cache add ca-certificates && rm -rf /var/cache/apk/*

# Copiando o binário do app compilado
COPY --from=build /app/main .

# Comando padrão
CMD [ "/app/main" ]
