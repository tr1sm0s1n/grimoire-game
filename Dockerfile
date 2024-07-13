ARG node_base_image=node:20-alpine
ARG go_base_image=golang:1.22.3-alpine

FROM ${node_base_image} AS node_build

WORKDIR /app
RUN npm i -g pnpm
COPY ./client/package.json ./package.json
RUN pnpm i
COPY ./client/ .
RUN pnpm run build

FROM ${go_base_image} AS go_build

RUN apk --no-cache add ca-certificates
WORKDIR /server
COPY ./server/go.mod .
COPY ./server/go.sum .
RUN go mod download
COPY ./server/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o main

FROM scratch

COPY --from=go_build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=go_build /server /
COPY --from=node_build /dist /

CMD ./main
