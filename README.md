# 🏗️ terraform-provider-mock

A high-fidelity Terraform provider plugin written in Go for local infrastructure simulation, dry-run testing, schema policy enforcement, and cost-free CI/CD pipelines.

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Terraform](https://img.shields.io/badge/Terraform-1.5%2B-844FBA.svg)](https://terraform.io)
[![Go](https://img.shields.io/badge/Go-1.22%2B-00ADD8.svg)](https://golang.org)

## Why terraform-provider-mock?

- ⚡ **Zero Cloud Cost** — Test complex module topologies and HCL expressions without provisioning real AWS/GCP/Azure resources.
- 🧪 **Deterministic CI/CD** — Mock network latency, simulated 4xx/5xx API errors, and concurrent quota throttles in pipelines.
- 📐 **State Verification** — Full CRUD lifecycle simulation with state drift detection and rollback mechanics.
- 🔒 **Policy Compliance** — Validate Sentinel / OPA Rego rules against simulated infrastructure graphs.

## Quick Start

### Provider Configuration

```hcl
terraform {
  required_providers {
    mock = {
      source  = "ardacelik-sudo/mock"
      version = "~> 0.3.0"
    }
  }
}

provider "mock" {
  simulate_delay_ms = 50
  failure_rate      = 0.0
}
```

### Resource Usage Example

```hcl
# Simulate a VPC Network
resource "mock_network" "main" {
  name       = "production-vpc"
  cidr_block = "10.0.0.0/16"
  tags = {
    Environment = "production"
    ManagedBy   = "Terraform"
  }
}

# Simulate a Compute Instance
resource "mock_instance" "app_server" {
  network_id    = mock_network.main.id
  instance_type = "c6i.2xlarge"
  name          = "api-worker-01"
}

output "instance_private_ip" {
  value = mock_instance.app_server.ip_address
}
```

## Building From Source

```bash
git clone https://github.com/ardacelik-sudo/terraform-provider-mock.git
cd terraform-provider-mock
go build -o terraform-provider-mock
```

## License

MIT License - Copyright (c) 2024 Arda Çelik.
