FROM golang:1.12 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -ldflags -s -o bin/mtlmeta cmd/*.go

FROM scratch as run
COPY --from=build /app/bin/mtlmeta /app/mtlmeta
ENTRYPOINT [ "/app/mtlmeta" ]
