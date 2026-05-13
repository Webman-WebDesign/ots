# Stage 1: Build
FROM alpine:3.18 AS builder

# Notwendige Build-Tools installieren
RUN apk add --no-cache \
    curl \
    git \
    go \
    make \
    nodejs \
    npm \
    gcc \
    musl-dev

WORKDIR /app

# 1. Quellcode kopieren
COPY . .

# 2. Frontend-Assets und Code-Generierung
# Falls make fehlschlägt, liegt es oft an fehlenden npm-Berechtigungen
RUN npm install && make download_libs
RUN make generate-inner
RUN make generate-apidocs

# 3. Go-Build vorbereiten
# Wir setzen CGO_ENABLED=0 für eine statische Binary (besser für Alpine)
# Wir nutzen 'go build' statt 'go install', um die Kontrolle über das Output-Verzeichnis zu haben
RUN export CGO_ENABLED=0 && \
    go build -o /ots -ldflags "-X main.version=$(git describe --tags --always || echo dev)" -mod=readonly

# Stage 2: Runtime
FROM alpine:3.18

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Die fertig gebaute Binary und die benötigten Runtime-Files kopieren
COPY --from=builder /ots /usr/local/bin/ots
COPY --from=builder /app/customize.yaml.example /etc/ots/customize.yaml

# Ports und Startbefehl
EXPOSE 3000

ENTRYPOINT ["ots"]
CMD ["-config", "/etc/ots/customize.yaml"]
