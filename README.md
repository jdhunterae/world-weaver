# World Weaver

A procedurally generated fantasy world and autonomous civilization simulator written in Go.

## Overview

World Weaver explores how geography, resources, settlements, factions, and individuals can interact to create a living fantasy world with its own history. Inspired by tabletop role-playing games and emergent simulation games, the project aims to generate worlds whose stories arise from simulation rules rather than predetermined scripts.

The initial interface will be terminal-based. Users will be able to generate a world, advance time, and eventually inspect its inhabitants, settlements, and history.

## Project goals

- Generate reproducible worlds from seeds.
- Create coherent terrain, climates, biomes, and settlements.
- Simulate resources, population, trade, and political relationships.
- Allow autonomous agents to make decisions and affect the world.
- Record meaningful events and their consequences as world history.
- Learn idiomatic Go, algorithm design, testing, and software architecture through incremental development.

## Status

**Early development / planning.** The features described above are goals, not implemented functionality. Build and usage instructions will be added once a runnable prototype exists.

## Design principles

- Simulation rules are independent of the user interface.
- Prefer reproducible, testable behavior and simple models before complexity.
- Stories emerge from interactions between systems.
- Add abstractions and concurrency only when they solve a demonstrated problem.

## Documentation

- [Design document](docs/DESIGN.md)
- [Development roadmap](docs/ROADMAP.md)

## Technology

Go, initially using the standard library. Terminal UI libraries and persistence technologies will be evaluated as the project develops.

## License

A license has not yet been selected. Public repository visibility does not itself grant permission to reuse or redistribute the code.
