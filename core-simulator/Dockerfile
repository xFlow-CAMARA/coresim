# ---- Stage 1: build Go binary ----
FROM golang:1.23 AS gobuilder

WORKDIR /app

# Copy all sources
COPY . .

# Download Go dependencies
RUN go mod download

# Build core-simulator
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X main.gitBranch=$(git rev-parse --abbrev-ref HEAD) -X main.gitCommit=$(git rev-parse HEAD)" \
    -o /core-simulator cmd/core-simulator/main.go


# ---- Stage 2: build Python CLI ----
FROM python:3.11-slim AS pybuilder

# Install system deps for PyInstaller
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential \
    python3-dev \
    libffi-dev \
    gcc \
    && rm -rf /var/lib/apt/lists/*

# Install PyInstaller
RUN pip install --no-cache-dir pyinstaller

WORKDIR /cli

# Copy CLI sources
COPY cli/ .

# Install Python dependencies
RUN pip install --no-cache-dir -r requirements.txt

# Build standalone executable
RUN pyinstaller --onefile cnsim-cli.py -n cnsim-cli


# ---- Stage 3: final runtime ----
FROM debian:latest

WORKDIR /

# Copy Go binary
COPY --from=gobuilder /core-simulator /core-simulator

# Copy Python CLI binary
COPY --from=pybuilder /cli/dist/cnsim-cli /usr/local/bin/cnsim-cli
COPY --from=pybuilder /cli/cnsim-profile.yaml /

RUN apt update && apt install curl -y

EXPOSE 8080
EXPOSE 8081

# Default command stays core-simulator
CMD ["/core-simulator"]
