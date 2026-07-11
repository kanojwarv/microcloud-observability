ADR-0007: Artifact Generation Owns Persistence
Status

Accepted

Context

The repository introduced artifact generation through the generate package:

graph.json
graph.html

Initially, artifact persistence was implemented in the CLI layer:

internal/cli/generate.go

while other components, such as:

watch mode
server startup
future automation

invoked:

generate.All(...)

directly.

This separation caused inconsistent behavior, where generated artifacts were not always written to disk.

Decision

Artifact generation and artifact persistence will both be owned by the internal/generate package.

The CLI layer is responsible only for:

parsing commands;
invoking generators;
displaying status messages.

The generate package is responsible for:

producing artifacts;
creating output directories;
writing artifacts to disk.
Consequences

Benefits:

consistent behavior across commands;
simpler CLI implementation;
watch mode automatically regenerates artifacts;
easier future extensions.

Trade-offs:

generation package becomes responsible for filesystem writes;
artifact locations become part of the generation contract.