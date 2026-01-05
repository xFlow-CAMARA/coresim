# open-exposure Core Simulator — Quickstart & Deployment Guide

This document explains how to run the open-exposure Core Simulator included in this repository. It focuses on running the simulator as a standalone service using the provided Docker Compose manifest and includes pointers to the simulator CLI and API documentation.

Contents
- What you'll find in this folder
- Prerequisites
- Building images locally
- Deploying the Core Simulator (standalone)
- Using the simulator CLI
- Useful checks and smoke tests
- Troubleshooting and tips
- File map and purpose

## What you'll find here

Under `coresim/artifacts/` the repository provides a ready-to-run compose manifest and example configuration files to run the core simulator:

- `docker-compose/` — contains `docker-compose.yaml` and `config/coreSim.yaml` plus optional `grafana/` and `prometheus/` example configs

The simulator code, CLI and docs live under `coresim/core-simulator/`:

- `core-simulator/` — simulator implementation, CLI, Dockerfiles and docs
- `core-simulator/cli/cnsim-cli.py` — Python CLI script for interacting with the simulator
- `core-simulator/docs/cli.md` — CLI usage documentation
- `core-simulator/docs/api.md` — API documentation

## Prerequisites

- Docker Engine
- Docker Compose v2 (use `docker compose`) or `docker-compose` v1
- Python 3.x if you plan to run the CLI locally
- At least 4-8 GB free RAM for the simulator and optional observability services

## Building the images locally

Prebuilt images are available from the project's CI registry. To build the core-simulator image locally, run the docker build command from the `core-simulator` directory:

```bash
cd coresim/core-simulator
docker build -t openexposure/core-simulator:local -f docker/Dockerfile .
```

If you prefer building only the CLI image (for running the CLI in a container):

```bash
cd coresim/core-simulator/cli
docker build -t openexposure/core-simulator-cli:local -f Dockerfile .
```

Remember to update `coresim/artifacts/docker-compose/docker-compose.yaml` if you change the image tag and want the compose file to use your local image.

## Deploying the Core Simulator (standalone)

1. Change into the compose folder:

```bash
cd coresim/artifacts/docker-compose
```

2. Inspect `docker-compose.yaml` and `config/coreSim.yaml` to review service settings and ports. Adjust configuration values (ports, persistent volumes, environment variables) to suit your environment.

3. Start the simulator stack in the background:

```bash
docker compose up -d
```

4. Confirm containers are up and healthy:

```bash
docker compose ps
```

5. Tail logs for the core-simulator service:

```bash
docker compose logs -f core-simulator
```

6. When finished, stop and remove the stack:

```bash
docker compose down
```

## Using the simulator CLI

The Core Simulator provides an interactive CLI called `cnsim-cli`. The canonical documentation for the CLI commands is at `coresim/core-simulator/docs/cli.md` — follow that doc for the authoritative command list (examples include `loadprofile`, `init`, `start`, `status`, and `stop`).

Recommended way to run the CLI

The preferred and most reliable way to run the CLI in a deployed compose environment is to execute the built `cnsim-cli` binary inside the running `core-simulator` container. This ensures the CLI talks to the simulator instance with the correct network namespace and configuration.

1. Start the compose stack (if not already running):

```bash
cd coresim/artifacts/docker-compose
docker compose up -d
```

2. Open an interactive shell inside the running core-simulator container and run the CLI

```bash
docker compose exec core-simulator /bin/sh -c "cnsim-cli"
```

If your environment provides `bash` instead of `sh`, you can use `/bin/bash -c "cnsim-cli"` instead. As an alternative (when `docker compose exec` is not available) you can use the container ID:

```bash
docker exec -it $(docker compose ps -q core-simulator) /bin/sh -c "cnsim-cli"
```

3. Example interactive session (commands are taken from `docs/cli.md`):

```text
$ cnsim-cli
cnsim> loadprofile basic
Profile 'basic' loaded.
cnsim> init
Simulation configured with profile 'basic'.
cnsim> start
Simulation started.
cnsim> status
Simulation running with X UEs, Y gNBs.
cnsim> stop
Simulation stopped.
```

Notes

- The CLI commands and profiles are documented in `coresim/core-simulator/docs/cli.md` and `coresim/core-simulator/cli/cnsim-profile.yaml`.
- Running the CLI inside the container uses the same binaries and environment used by the simulator service and avoids issues due to mismatched package versions or network configuration.
- For development purposes you can still run the Python variant locally (see `coresim/core-simulator/cli`), but it is not the recommended workflow for interacting with a running compose deployment.

## Useful checks and a basic smoke test

- Container health and status: `docker compose ps` and the `STATUS` column
- Logs: `docker compose logs -f core-simulator`


## Troubleshooting and tips

- Port conflicts: Edit `docker-compose.yaml` to change published ports if another service is listening on the same port.
- Logs: Use `docker compose logs --tail 200 core-simulator` to see recent startup output.
- If the CLI fails due to missing Python packages, ensure you installed `requirements.txt` or use the CLI docker image.
- Resource limits: If containers are killed for memory, increase available RAM or disable optional observability containers (Prometheus/Grafana) in the compose file.

## File map and purpose

- `coresim/artifacts/docker-compose/docker-compose.yaml` — Compose manifest to run the core simulator (and optional observability services)
- `coresim/artifacts/docker-compose/config/coreSim.yaml` — Example simulator configuration
- `coresim/core-simulator/` — simulator source code, Dockerfiles and docs
- `coresim/core-simulator/cli/cnsim-cli.py` — Python CLI script
- `coresim/core-simulator/docs/cli.md` — CLI documentation and examples
- `coresim/core-simulator/docs/api.md` — API documentation for the simulator

## Next steps and suggestions

- Customize `config/coreSim.yaml` to match your test scenarios and networking environment
- If you plan to integrate the simulator with other open-exposure components, use the simulator compose network and check service discovery settings
