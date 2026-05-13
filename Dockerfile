# Stage 1: Build
FROM alpine:3.18 AS builder

# Installiere Build-Abhängigkeiten für Go und Node.js
RUN apk add --no-cache \
    curl \
    git \
    go \
    make \
    nodejs \
    npm \
    tar \
    unzip \
    gcc \
    musl-dev

WORKDIR /app

# Kopiere die Source-Dateien
COPY . .

# Führe die Build-Schritte aus dem Original-Repo aus
# Wir nutzen --unsafe-perm, falls npm Probleme mit Root-Rechten hat
RUN make download_libs
RUN make generate-inner
RUN make generate-apidocs

# Go-Binary bauen
RUN go install -ldflags "-X main.version=$(git describe --tags --always || echo dev)" -mod=readonly

# Stage 2: Runtime
FROM alpine:3.18

# Installiere nur die Laufzeit-Abhängigkeiten
RUN apk add --no-cache \
    ca-certificates \
    tzdata

WORKDIR /app

# Kopiere die gebaute Binary vom Builder
COPY --from=builder /root/go/bin/ots /usr/local/bin/ots
COPY --from=builder /app/customize.yaml.example /etc/ots/customize.yaml

# Ports und Startbefehl
EXPOSE 3000

ENTRYPOINT ["ots"]
CMD ["-config", "/etc/ots/customize.yaml"]
