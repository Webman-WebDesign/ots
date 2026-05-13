# Stage 1: Build
FROM golang:1.23-alpine3.21 AS builder

# GNU-tar und andere Tools installieren
RUN apk add --no-cache \
    curl \
    git \
    make \
    nodejs \
    npm \
    gcc \
    musl-dev \
    tar \
    gzip \
    bash

WORKDIR /app

# Quellcode kopieren
COPY . .

# 1. Abhängigkeiten für das Frontend (npm ci ist im Makefile unter generate-inner)
RUN npm ci --include=dev

# 2. Externe Libs (FontAwesome) manuell laden
# Wir nutzen hier GNU tar, das wir oben installiert haben
RUN VER_FONTAWESOME=6.4.0 && \
    mkdir -p frontend && \
    curl -sSfL https://github.com/FortAwesome/Font-Awesome/archive/${VER_FONTAWESOME}.tar.gz | \
    tar -vC frontend -xz --strip-components=1 --wildcards --exclude='*/js-packages' '*/css' '*/webfonts'

# 3. Assets generieren (Wir führen die Makefile-Schritte manuell aus, ohne 'docker run')
RUN node ./ci/build.mjs
RUN npx --yes @redocly/cli build-docs docs/openapi.yaml --disableGoogleFont true -o frontend/api.html

# 4. Go-Binary bauen (statisches Build für Alpine)
RUN export CGO_ENABLED=0 && \
    go build -o /ots -ldflags "-s -w -X main.version=$(git describe --tags --always || echo dev)" -mod=readonly -trimpath

# Stage 2: Runtime
FROM alpine:3.18

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Binary und Config kopieren
COPY --from=builder /ots /usr/local/bin/ots

COPY --from=builder /app/frontend /app/frontend

EXPOSE 3000

CMD ["ots"]