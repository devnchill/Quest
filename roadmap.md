# Learning Roadmap

## Phase 0: Chess Engine

**Language:** C++ or Rust

### Learn

- [ ] Bitboards, magic bitboards, attack tables
- [ ] Minimax, negamax, alpha-beta, iterative deepening, quiescence
- [ ] Eval: material, PST, tapered
- [ ] Zobrist hashing, transposition tables
- [ ] UCI protocol, PolyGlot opening book
- [ ] Time management
- [ ] Profiling: perf, flamegraphs, criterion
- [ ] Property-based testing (fast-check) for move invariants

### Project

- [ ] Monorepo: `engine/` (C++/Rust, own flake.nix), `packages/chess-lib/` (TS), `web/` (Next.js)
- [ ] Engine CLI speaking UCI — bitboard representation, magic move gen, alpha-beta search, eval tables all live here
- [ ] Nix flake: standalone `nix build` for engine, top-level flake references subflake for unified build
- [ ] Engine binary baked into web Docker image via multi-stage build (also functional for headless child process in prod)
- [ ] Test engine in Docker container — ensure UCI I/O works in headless environment

---

## Phase 1: TS Chess Library + Frontend

### Learn

- [ ] TypeScript (generics, conditional, branded types)
- [ ] TS chess lib (board state, FEN/PGN, legal move gen, checkmate)
- [ ] Next.js App Router
- [ ] Zustand
- [ ] Tailwind CSS + shadcn/ui
- [ ] SVG/Canvas board rendering
- [ ] Zod
- [ ] WCAG accessibility
- [ ] Visual regression testing (Playwright snapshot)
- [ ] Responsive design (mobile-first)
- [ ] Storybook (component dev in isolation)

### Project

- [ ] **TS chess lib** — board state model, FEN/PGN parsing, legal move generation (~40 candidates, no search), check/checkmate/stalemate detection. Used client-side for instant UI feedback.
- [ ] **Chess UI** — Next.js App Router pages, Zustand for game state, Tailwind + shadcn/ui for board/pieces, SVG/Canvas rendering, drag-and-drop + click-to-move, legal move highlights (TS lib generates client-side), move history, undo/redo, PGN export, simple random-move bot
- [ ] WCAG: keyboard nav, screen reader announces moves
- [ ] Zod validates FEN/PGN inputs client-side
- [ ] Same codebase continues to Phase 2
- [ ] Error boundaries, loading skeletons, empty states for every view

---

## Phase 2: Production Backend + Engine Integration

### Learn

- [ ] GraphQL (Apollo Server)
- [ ] Vitest, Supertest, Playwright
- [ ] Socket.io + Redis adapter
- [ ] Redis (cache, sessions, rate limiting, pub/sub)
- [ ] Postgres + Prisma
- [ ] Bull queue
- [ ] SSE
- [ ] Child process (UCI engine)
- [ ] Sentry — free tier (5k events/mo) enough for learning
- [ ] OpenTelemetry + distributed tracing (Jaeger)
- [ ] SAST scanning (Semgrep or CodeQL) — Semgrep free tier, CodeQL free for public repos
- [ ] Mutation testing (Stryker)
- [ ] Conventional commits + semantic versioning + changelog
- [ ] Pre-commit hooks (husky, lint-staged)
- [ ] Database migration strategy (Prisma migrations in CI)
- [ ] PgBouncer
- [ ] Nginx (reverse proxy, rate limiting, SSL)
- [ ] Docker Compose
- [ ] GitHub Actions — free for public repos
- [ ] Feature flags
- [ ] Pino structured logging
- [ ] Health checks, graceful shutdown
- [ ] Secrets management (Doppler)

### Project

- [ ] **GraphQL schema** — users, games, moves, analysis
- [ ] **Move validation** — client sends UCI string, server re-validates via TS chess lib (security boundary), applies to server state, broadcasts to opponent
- [ ] **Game state sync** — server is source of truth, Socket.io with Redis adapter broadcasts moves, clients apply locally, no desync
- [ ] **Invite system** — private games via shareable link + friend usernames, WebSocket room membership gated by invite
- [ ] **Spectating** — read-only WebSocket connections to game room, privacy toggle (streamer mode hides usernames)
- [ ] **Game controls** — draw offer, resign, abort via WebSocket events + GraphQL mutations
- [ ] **Time controls** — per-player clock with increment, per-move timing stored in Postgres
- [ ] **Engine hook** — Bull queue receives analysis job on game end, worker spawns engine binary via child_process, feeds UCI commands, collects eval lines
- [ ] **SSE stream** — analysis lines pushed to client as they come in
- [ ] **Elo/Glicko rating** — computed after each rated game, stored in Postgres
- [ ] **Storage** — Postgres + Prisma for users, games, PGN, move history, per-move timing; DB indexes on user_id + date, game_id + move_number
- [ ] **Infra** — Nginx reverse proxy, rate limiting, static cache, SSL; PgBouncer for connection pooling; Sentry for errors
- [ ] **Docker Compose** — app + postgres + redis + nginx + worker + engine
- [ ] **CI/CD** — GitHub Actions: lint → test → build → push
- [ ] **Operability** — health check endpoints, graceful shutdown, structured logging (pino), feature flags, secrets management (Doppler)

---

## Phase 3: Cloud

### Learn

- [ ] Cloudflare Workers — free tier (100k req/day)
- [ ] AWS (IAM, S3, RDS, ElastiCache, EC2, ALB, CloudFront, Lambda) — $ most AWS services have 12mo free tier then cost
- [ ] Terraform — free (OpenTofu is free drop-in alternative)
- [ ] Prometheus + Grafana + Loki — self-host for free, or Grafana Cloud free tier (3 dashboards, 14d retention)
- [ ] Read replicas
- [ ] Uptime monitoring
- [ ] Cost optimization
- [ ] Cost monitoring + budgets (AWS Budgets, Grafana dashboards)
- [ ] Performance budgets (Lighthouse CI, bundle size regressions)

### Project

- [ ] Terraform defines all infra: VPC, RDS, ElastiCache, EC2/ECS, ALB, CloudFront, S3 — $ for production; learn concepts locally or swap for free alternatives:
  - **RDS →** Supabase/Neon free tier (500MB Postgres)
  - **ElastiCache →** Redis Cloud free tier (30MB)
  - **EC2/ECS →** Fly.io free tier (3 shared VMs) or local Docker
  - **ALB →** Cloudflare Tunnel (free)
  - **CloudFront + S3 →** Cloudflare Pages + R2 (free tier)
- [ ] Cloudflare Worker acts as auth middleware at the edge (free `*.workers.dev`)
- [ ] Prometheus + Grafana: metrics dashboards — self-host free or Grafana Cloud free tier
- [ ] Loki: log aggregation — self-host free or Grafana Cloud free tier
- [ ] Uptime monitoring + alerts — free: Better Uptime or Uptime Robot free tier
- [ ] Multi-layer caching: swap CloudFront for Cloudflare (free)
- [ ] Read replicas for query offloading — swap RDS read replica for Supabase read replica
- [ ] Cost optimization: right-sizing, S3 lifecycle policies (use R2 instead)
- [ ] GitHub Actions: terraform plan → apply → deploy — free for public repos

---

## Phase 4: Kubernetes

### Learn

- [ ] ECS Fargate — $ learn with kind/k3s/minikube locally instead (free)
- [ ] K8s (pods, deployments, services, ingress, ConfigMap, secrets, HPA) — kind runs K8s in Docker for free
- [ ] Helm
- [ ] Redis Cluster
- [ ] K6 load testing
- [ ] Database sharding (Citus)
- [ ] cert-manager, ExternalDNS — both free (OSS)
- [ ] K9s

### Project

- [ ] Multi-service: web, api, engine-worker, analysis-worker as separate deployments
- [ ] Helm charts per service with ConfigMap, secrets, resource limits
- [ ] HPA for engine analysis workers (CPU/memory + custom metric: Bull queue depth)
- [ ] ALB Ingress Controller with cert-manager for TLS — $ ALB costs; learn with kind + ingress-nginx locally (free)
- [ ] Redis Cluster for sharded caching (hot games)
- [ ] PgBouncer sidecar
- [ ] Read replicas for production query offloading
- [ ] Database sharding: partition games by time (Citus)
- [ ] K6 in CI/CD — simulate 1000 concurrent games
- [ ] Prometheus + Grafana + Loki — self-host in cluster for free
- [ ] Rollback strategy: Helm rollback, fix-forward vs revert
- [ ] Canary / blue-green deployment strategy — Argo Rollouts (free) for K8s

---

## Phase 5: Distributed Systems

### Learn

- [ ] DDIA (book)
- [ ] Kafka + Schema Registry — free (OSS), Redpanda is free drop-in alternative with no ZooKeeper dependency
- [ ] gRPC + Protobuf
- [ ] CQRS + event sourcing
- [ ] Debezium (CDC)
- [ ] Service mesh (Istio/Linkerd) — both OSS, free
- [ ] Saga, circuit breaker, API gateway
- [ ] ADRs, runbooks
- [ ] Post-mortem culture (blameless incident reviews)

### Project

- [ ] **Event sourcing** — game state as event stream (GameStarted, MoveMade, GameEnded, AnalysisComplete), replay any game from events
- [ ] **Kafka** — topics per event type, partitions by game_id, consumer groups for analysis + stats + puzzle generation
- [ ] **Schema Registry** — Avro/Protobuf schemas for all events
- [ ] **gRPC engine workers** — replace child_process UCI parsing with protobuf, streaming eval lines
- [ ] **CQRS** — write model (events), materialized views for reads (game summaries, leaderboards, opening stats)
- [ ] **Debezium** — CDC from Postgres to Kafka for legacy v1 tables
- [ ] **Strangler fig migration** — migrate service by service alongside v1
- [ ] **ADRs** for every architecture decision
- [ ] **Runbooks** for incident response

---

## Phase 6: AI/ML

### Learn

- [ ] numpy, pandas, scikit-learn
- [ ] NNUE
- [ ] pgvector
- [ ] RAG — $ requires LLM API (OpenAI/etc.) for production; learn locally with Ollama + local embedding models (free)

### Project

- [ ] Opening recommendation engine — pgvector similarity search over historical games
- [ ] Puzzle generation — mine tactics from your own game history
- [ ] NNUE evaluator for engine (if going deep)

---

## Cross-cutting

- [ ] Security: CSP, helmet, CORS, rate limiting, dependabot, SAST — all free tools
- [ ] Observability: Sentry, OpenTelemetry tracing, Prometheus/Grafana, Loki
- [ ] Linting: ESLint, Prettier
- [ ] CI/CD: GitHub Actions
- [ ] Release: conventional commits, semantic versioning, changelog
- [ ] DX: pre-commit hooks, dev container, one-command setup
- [ ] Docs: ADRs, runbooks
