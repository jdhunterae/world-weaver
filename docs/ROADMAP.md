# World Weaver — Development Roadmap

Revised baseline. Milestones are learning checkpoints, not calendar deadlines. All are planned until verified against the repository. Each milestone should leave a runnable or independently testable result.

## Release gates

| Gate | Milestones | Demonstrable result |
|---|---|---|
| Foundation | M0–M2 | Seeded terminal map |
| Survival MVP | M3–M6 | Renewable ecosystem and competing food-seeking tribes |
| Governance and interaction | M7–M10 | Adaptive ranges, encounters, leadership and political change |
| Living collectives | M11–M12 | Fragmentation and autonomous predator packs |
| Observer and experiments | M13–M14 | Inspectable history, repeatable comparative runs |

## M0 — Project foundation

- [ ] Review and commit revised README, design and roadmap.
- [ ] Initialize Go module using actual repository import path; create minimal executable.
- [ ] Verify `go run ./cmd/world-weaver` and `go test ./...`; add CI when useful.
- [ ] Establish focused issues, branches and self-reviewed pull requests.

Done when a fresh clone builds and runs with documented commands. Learning: modules, Git, tests.

## M1 — Static world

- [ ] Define coordinates, terrain, tile and rectangular grid; validate dimensions and bounds.
- [ ] Build a small hand-authored map with five terrain types.
- [ ] Render ASCII independently of world logic; test dimensions, access and terrain glyphs.

Done when a static map renders reliably and tests run without terminal I/O. Learning: types, slices, methods. No life or procedural generation yet.

## M2 — Seeded geography

- [ ] Accept dimensions and seed; generate valid terrain with a local random source.
- [ ] Improve geographic coherence using a simple spatial technique; defer detailed climate and rivers.
- [ ] Display seed; test identical inputs produce identical maps.

Done when reproducible maps contain coherent regions. Learning: randomness, spatial algorithms.

## M3 — Clock, vegetation and events

- [ ] Introduce explicit daily step and structured event records.
- [ ] Model terrain-dependent vegetation capacity and regeneration.
- [ ] Advance N days without rendering; inspect biomass and test bounds.

Done when a depleted patch recovers according to rules and time advances exactly N ticks. Learning: state transitions, bounded growth.

## M4 — Prey ecology

- [ ] Add aggregate prey populations, feeding, reproduction and starvation.
- [ ] Define consumption and growth order; handle zero populations and depleted vegetation.
- [ ] Test vegetation/prey interactions and seeded replay.

Done when prey respond to renewable food availability without negative counts or duplicated resources. Learning: feedback loops, invariants.

## M5 — Tribes and basic survival

- [ ] Add stable IDs, home placement, aggregate population, food stockpile and hunger.
- [ ] Implement simple local resource acquisition from actual available prey or vegetation, consumption and population change.
- [ ] Test shortage, recovery, extinction and competing resource claims.

Done when multiple tribes can survive or decline over a bounded run. Local acquisition is temporary scaffolding, not an extra source retained alongside later expeditions. Learning: data modeling, resource accounting.

## M6 — Hunting expeditions and fixed influence (Survival MVP)

- [ ] Implement fixed reachable range using uniform movement costs and BFS.
- [ ] Assign hunting groups; track travel, harvesting, return, casualties and cargo.
- [ ] Replace direct local acquisition with explicit expedition acquisition where applicable.
- [ ] Render homes, influence and overlap; inspect food, population and events.
- [ ] Test member allocation, shared prey depletion, return accounting and deterministic replay.

Done when a seeded multi-tribe world runs for a configurable period (initial test target: 30 days) and visibly demonstrates resource competition, survival and decline. No political or predator-pack simulation required. Learning: graph traversal, state machines.

## M7 — Dynamic influence and destination choice

- [ ] Track recent hunting yields, reserves and travel costs; calculate desired range.
- [ ] Adjust range gradually; select feasible destinations by weighted utility.
- [ ] Add terrain costs and Dijkstra reachability; test barriers and budget bounds.

Done when scarcity expands activity, abundance can contract it, and geography shapes reachable regions. Learning: priority queues, scoring.

## M8 — Leadership and population response

- [ ] Add leader traits as decision biases, not rigid commands.
- [ ] Derive population pressures and implement pure dissatisfaction evaluation.
- [ ] Apply bounded compliance effects to actual group activity; test that intentions and outcomes differ.

Done when identical circumstances can yield different decisions under different leaders and evaluation never mutates state. Learning: pure functions, table-driven tests.

## M9 — Encounters and relationships

- [ ] Resolve co-located expeditions using explicit presence rules.
- [ ] Model avoidance, tolerance, incidental conflict and casualties.
- [ ] Track pairwise relationships; introduce basic cooperation, alliances and intentional conflict.

Done when overlapping ranges can coexist peacefully or generate conflict through actual interactions. Learning: interaction resolution, deterministic ordering.

## M10 — Confidence, instability and elections

- [ ] Add persistent confidence, gradual loss/recovery and separate instability evaluation.
- [ ] Resolve political action without automatic overthrow at zero confidence.
- [ ] Generate three candidates; use the same pure evaluation; elect lowest dissatisfaction.
- [ ] Restore moderate confidence only; preserve material state and record succession.

Done when a leader can recover from hardship or be replaced, and inherited crises do not force immediate succession by definition. Test repeated crises and candidate ties. Learning: state transitions, deterministic randomness.

## M11 — Cohesion, migration and fragmentation

- [ ] Separate group cohesion from leadership confidence.
- [ ] Add relocation and viable split outcomes; handle home contention or migration.
- [ ] Partition people, resources and outstanding expeditions without duplication; record ancestry.

Done when new independent tribes can emerge from existing ones and conservation tests pass. Joining/merging is optional follow-up. Learning: entity lifecycle, invariants.

## M12 — Autonomous predator packs

- [ ] Add packs with prey needs, den, movement, influence and hunting behavior.
- [ ] Reuse proven geographic/resource mechanisms through composition; avoid human political rules for animals.
- [ ] Resolve competition and encounters with tribes and other packs; add species-appropriate dispersal and fragmentation.
- [ ] Test prey accounting, pack survival and ecological feedback into tribal conditions.

Done when packs can independently change hunting ranges and affect tribal survival, and pack splits conserve members. Mobile prey herds are a later extension. Learning: refactoring, composition, integration tests.

## M13 — Observer and persistence

- [ ] Improve terminal inspection, overlays and event filtering; expose decision scores and political history.
- [ ] Add save/load with sufficient continuation state, including RNG strategy.
- [ ] Test replay equivalence across save/load.

Done when a user can inspect why major events occurred and resume a world. Learning: serialization, UI boundaries.

## M14 — Experiments and scale

- [ ] Run controlled scenarios across multiple seeds; report survival, population, resource and political metrics.
- [ ] Compare strategies with equivalent starting conditions and document limitations.
- [ ] Benchmark; parallelize independent runs only if beneficial, with isolated RNG/state.

Done when repeatable experiments produce comparable summaries and performance changes are measured. Learning: benchmarking, concurrency, statistics.

## Working agreement

The project author writes the implementation. AI assistance defaults to requirements, algorithm explanations, pseudocode, test suggestions, debugging and code review; complete implementations only when requested. Open focused issues as prerequisites become understood. Avoid implementing future abstractions before concrete use cases exist.

## Next issue 

**M0: Establish revised project foundation.** Review these documents, commit them on a documentation branch, and verify the module/executable status before starting M1. Suggested branch: `docs/revise-simulation-scope`; suggested commit: `docs: redefine World Weaver as an emergent ecosystem simulator`. Do not mark milestones complete based on this document alone.
