# World Weaver

A procedurally generated fantasy ecosystem and society simulator written in Go.

## Overview

World Weaver explores how geography, renewable resources, prey populations, predator packs and human tribes interact to produce emergent history. Generate a seeded world, advance its clock and observe how its inhabitants survive, compete, cooperate, migrate and change over time.

The observer does not direct individual inhabitants. Stories arise from the consequences of simulation rules, not scripted events or predetermined winners.

## Project goals

- Generate reproducible worlds with coherent terrain and renewable vegetation.
- Model food chains, resource scarcity and population survival.
- Give autonomous collectives homes, dynamic ranges of influence and group activities.
- Model leadership, population dissatisfaction, confidence, cooperation and political change.
- Allow incidental conflict, alliances, migration and fragmentation to emerge naturally.
- Record meaningful events in an inspectable historical chronicle.
- Learn idiomatic Go, algorithms, testing and incremental software architecture.

## Design principles

- Influence is not ownership; overlap is not hostility.
- Leadership is not absolute control; intentions are not guaranteed outcomes.
- Dissatisfaction is not necessarily overthrow; expansion is not necessarily prosperity.
- Conflict does not require intent; history is an outcome, not a scripted event.
- Simulate consequences before details. Share mechanisms without forcing every species to behave alike.
- Keep simulation independent of presentation, deterministic when seeded, and independently testable.

## Status and scope

**Planning / early development.** These are design goals, not implemented features. The first MVP is a terminal-observable ecosystem with vegetation, prey and multiple food-seeking tribes. Governance, autonomous predator packs and fragmentation follow in staged releases. No individual NPC simulation, magic, detailed economy or graphical UI is required for the MVP.

## Documentation

- [Design document](docs/DESIGN.md)
- [Development roadmap](docs/ROADMAP.md)

## Technology and license

Go, initially using the standard library. Build and usage instructions will be added once a runnable prototype exists. A license has not yet been selected; public visibility alone does not grant reuse rights.
