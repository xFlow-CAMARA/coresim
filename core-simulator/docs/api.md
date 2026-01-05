# API Usage Guide

This document provides detailed information about the APIs exposed by the 5G Core Simulator.

## Overview

The simulator exposes two categories of REST APIs:

- **3GPP SBI APIs** (Service-Based Interfaces): Used to simulate AMF, SMF, and PCF event exposure.  
- **OAM APIs** (Operations and Maintenance): Used to control simulations and retrieve metrics.

---

## 3GPP SBI APIs (default: :8080)

The simulator aligns with **3GPP TS 29-series (Release 17)**.

### Nsmf_EventExposure (TS 29.502 Rel-17)
Session management event exposure.

### Namf_Events (TS 29.518 Rel-17)
UE mobility and registration event exposure.

### Npcf_PolicyAuthorization (TS 29.514 Rel-17)
Policy control and authorization for UEs.

---

## OAM APIs (default: :8081)

### POST /core-simulator/v1/start
Start a simulation using the current configuration.

### POST /core-simulator/v1/stop
Stop the running simulation.

### GET /core-simulator/v1/status
Retrieve the current simulation status.

### POST /core-simulator/v1/configure
Send a configuration payload to update simulation parameters.