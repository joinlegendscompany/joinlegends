# JoinLegends Backend – Contexto para o agente

Backend em **Go** (módulo `go-backend-stream`), Fiber v2, PostgreSQL (goe), Redis, JWT com sessões.

**Referência principal para LLMs**: leia `docs/CONTEXT_FOR_LLM.md` para arquitetura, estrutura de pastas, padrões (Controller → Service → Repository), convenções de repositórios, models, factory, rotas e como adicionar novos domínios.

Resumo rápido:
- Domínios em `internal/domains/` (auth, organization, stream, upload).
- Repositórios em `internal/infrastructure/repositories/<entity>/` (interface em `*-contract.go`, impl em `*.repository.go`).
- Models em `internal/models/*.go`. Schema e conexão em `internal/infrastructure/database/`.
- Rotas e DI em `internal/infrastructure/api/routes/routes.go` e `internal/infrastructure/factory/factory.go`.
- Usuário autenticado: `c.Locals("userId").(string)` após `NewJwtSessionMiddleware(db)`.

Ao implementar ou refatorar, siga os padrões descritos em `docs/CONTEXT_FOR_LLM.md`.
