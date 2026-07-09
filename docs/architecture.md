# MicroCloud Observability Architecture

## Vision
Build an observability platform for LXD, MicroCeph, GreptimeDB, Vector and Perses.

## Data Flow
LXD + MicroCeph -> Vector -> GreptimeDB -> Recording Rules -> Perses

## Principles
- Recording rules are the API.
- Dashboards consume recording rules.
- Alerts consume recording rules.
- Dashboards answer operational questions.
