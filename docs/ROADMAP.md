# World Weaver — Development Roadmap

Milestones are learning checkpoints, not calendar deadlines. All are initially planned; verify existing repository work before marking any complete.

| Milestone | Deliverable | Learning focus |
|---|---|---|
| M0 | Project foundation | Go modules, Git, tests |
| M1 | Static world | Types, slices, methods |
| M2 | Seeded terrain | Random sources, reproducibility |
| M3 | Coherent geography | Noise, spatial algorithms |
| M4 | Settlements | Scoring, constraints, data modeling |
| M5 | Simulation clock | State transitions, testing |
| M6 | Resource economy and chronicle | Modeling, invariants, events |
| M7 | Trade and roads | Graphs, pathfinding |
| M8 | Population and factions | Relationships, decision rules |
| M9 | Autonomous inhabitants | Agent behavior, state machines |
| M10 | Observer and persistence | UI, serialization |
| M11 | Scale and polish | Profiling, concurrency, releases |

## M0 — Project foundation

- [ ] Review README, design and roadmap.
- [ ] Initialize Go module using the actual repository import path.
- [ ] Create a minimal executable and verify it runs.
- [ ] Verify `go test ./...` succeeds; add a basic test when meaningful.
- [ ] Establish feature branches and pull requests; optionally add CI.

Done when a fresh clone can build and run the minimal program with documented commands.

## M1 — Static world

- [ ] Decide how to represent terrain, coordinates and a rectangular grid.
- [ ] Create a small hand-authored map containing water, plains, forest and mountains.
- [ ] Render the map to the terminal using ASCII glyphs.
- [ ] Separate rendering from the world model.
- [ ] Test map dimensions, tile access and out-of-bounds behavior.

Example output (illustrative):

```text
~~~~~~~~~~
~~..^^..~~
~..TT^...~
~.TT.....~
~~......~~
~~~~~~~~~~
```

Done when the map renders consistently, each terrain type has a distinct representation, and map logic can be tested without terminal input/output. No procedural generation or external UI library required.

## M2 — Seeded terrain

- [ ] Accept dimensions and a seed; reject invalid dimensions.
- [ ] Generate simple randomized terrain using a local random source.
- [ ] Display the seed with the map.
- [ ] Test identical configuration and seed produce identical maps.

Done when worlds can be reproduced and valid dimensions and terrain are guaranteed. Natural-looking geography is not required yet.

## M3 — Coherent geography

- [ ] Generate an elevation field using a coherent noise technique.
- [ ] Classify land, water and mountains.
- [ ] Introduce moisture and temperature, then derive biomes.
- [ ] Explore rivers and drainage as a follow-up task.

Done when maps contain coherent geographic regions and generation remains reproducible. Rivers may be deferred until their model is understood.

## M4 — Settlements

- [ ] Model settlements with stable IDs, names, positions and population.
- [ ] Score candidate tiles using geography and resource suitability.
- [ ] Enforce land placement and minimum spacing; handle insufficient candidates.

Done when generated settlements obey placement constraints and can be inspected.

## M5 — Simulation clock

- [ ] Represent elapsed days and implement an explicit step operation.
- [ ] Advance a requested number of days without requiring rendering.
- [ ] Establish a structured event record.

Done when advancing N days advances exactly N ticks and repeat runs with the same inputs behave consistently.

## M6 — Resource economy and chronicle

- [ ] Track food production, consumption, storage and unmet demand.
- [ ] Record shortage onset and recovery.
- [ ] Test surplus, balance and deficit cases, including exhausted reserves.
- [ ] Display settlement status and chronological events.

Done when a seeded world can run for at least 30 days and explain resource changes through inspectable state and events. This is the first meaningful simulation prototype.

## Later milestones (provisional)

| Milestone | Completion target |
|---|---|
| M7 | Connected settlements exchange resources via valid routes. |
| M8 | Population changes and faction decisions respond to simulated conditions. |
| M9 | Named agents pursue goals and generate consequential events. |
| M10 | Users inspect maps and history, and save/load a world with its continuation state. |
| M11 | Benchmarks identify bottlenecks; safe optimizations and a documented release follow. |

Refine later milestones into issues only when prerequisites are understood.
