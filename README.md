# MicroCloud Observability

MicroCloud Observability is a repository-driven observability platform for
MicroCloud, LXD, Kubernetes, and infrastructure workloads.

The project treats observability as code.

Metrics, recording rules, variables, dashboards, dependencies, and APIs
are described declaratively and compiled into portable artifacts.

Current capabilities:

- Repository validation
- Dependency graph generation
- Impact analysis
- Repository health checks
- JSON and HTML generation
- Local web server
- HTTP APIs

Architecture:-


                    +----------------+
                    |    metrics/    |
                    +----------------+
                             |
                             v
              +-----------------------------+
              |      recording-rules/       |
              +-----------------------------+
                             |
                             v
                    +----------------+
                    |   variables/   |
                    +----------------+
                             |
                             v
                    +----------------+
                    |  dashboards/   |
                    +----------------+
                             |
                             v
              +-----------------------------+
              |      Graph / Doctor         |
              +-----------------------------+
                       |             |
                       v             v
              +----------------+  +----------------+
              |    CLI/API     |  |   Generators   |
              +----------------+  +----------------+
                       |             |
                       +------+------+ 
                              |
                              v
                    +----------------+
                    | graph.html     |
                    | graph.json     |
                    +----------------+


Repository Layout:-

.
├── artifacts/
├── dashboards/
├── docs/
├── metrics/
├── recording-rules/
├── variables/
├── cmd/
│   └── mco/
├── internal/
│   ├── cli/
│   ├── doctor/
│   ├── generate/
│   ├── graph/
│   ├── loader/
│   ├── model/
│   ├── server/
│   └── validate/
└── Makefile

Quick Start:-

## Quick Start

Build:

```bash
make build
```

Run all checks:

```bash
make dev
```

Start the platform:

```bash
./bin/mco serve
```

Open:

```text
http://localhost:8080
```

CLI Commands:-

## CLI

```bash
./bin/mco validate

./bin/mco graph

./bin/mco impact vm_cpu_usage_seconds_rate

./bin/mco doctor

./bin/mco generate graph

./bin/mco generate graph-html

./bin/mco serve
```

HTTP API's:-

## HTTP APIs

```text
GET /api/graph

GET /api/doctor
```

Generated Artifacts:-

## Generated Artifacts

- artifacts/graph.json
- artifacts/graph.html

RoadMap:-

## Roadmap

### Completed

- [x] Validation engine
- [x] Dependency graph
- [x] Impact analysis
- [x] Repository doctor
- [x] JSON generation
- [x] HTML generation
- [x] Local runtime
- [x] HTTP APIs

### Planned

- [ ] Live reload
- [ ] Interactive graph UI
- [ ] Grafana generation
- [ ] Perses generation
- [ ] API server expansion


---

Built one commit at a time.
