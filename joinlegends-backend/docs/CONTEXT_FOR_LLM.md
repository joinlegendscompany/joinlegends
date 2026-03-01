# Contexto do projeto para LLMs

Este documento descreve a estrutura, convenções e padrões do **JoinLegends Backend** (módulo Go `go-backend-stream`) para servir de referência a assistentes e LLMs.

---

## 1. Visão geral

- **Projeto**: Backend da plataforma JoinLegends (comunidades, XTreinos, streaming).
- **Linguagem**: Go 1.24+.
- **Módulo**: `go-backend-stream` (imports usam esse prefixo).
- **Framework HTTP**: Fiber v2.
- **Banco**: PostgreSQL via ORM **go-goe/goe** (structs em `internal/models` são as tabelas).
- **Cache/filas**: Redis (go-redis).
- **Auth**: JWT com **sessões** (token contém session ID; sessões em DB).

Fluxo de uma requisição autenticada:

`Request → Middleware (JWT/sessão) → Controller → Service → Repository → DB`

---

## 2. Estrutura de pastas

```
joinlegends-backend/
├── cmd/server/main.go              # Entrada: Fiber, DB, Redis, WebSocket, routes
├── internal/
│   ├── domains/                    # Lógica de negócio por domínio
│   │   ├── auth/                   # Auth: controller, service, DTOs, testes
│   │   ├── organization/           # Organizações (controller, service)
│   │   ├── stream/                 # Streaming de vídeo (handlers)
│   │   └── upload/                 # Upload de vídeo (handlers)
│   ├── infrastructure/
│   │   ├── api/routes/routes.go    # Registro de rotas e grupos (v1, auth, etc.)
│   │   ├── database/               # Conexão (database.go), schema (StreamDB)
│   │   ├── factory/factory.go      # DI: repositórios → services → controllers
│   │   ├── redisclient/            # Redis, filas, pub/sub
│   │   └── repositories/          # Um pacote por entidade
│   │       └── <entity>/
│   │           ├── <entity>.repository-contract.go   # Interface do repositório
│   │           └── <entity>.repository.go            # Implementação
│   ├── models/                     # Structs de domínio/DB (um arquivo por área)
│   │   ├── user.go
│   │   ├── auth.go                 # Recovery, Session
│   │   ├── video.go
│   │   ├── organization.go         # Organization, Member
│   │   ├── events.go
│   │   └── doc.go
│   └── utilities/
│       ├── config/                 # LoadEnv(), variáveis globais (.env)
│       ├── jwt/                    # Geração e parse de JWT
│       ├── logger/                 # Info, Error, Debug
│       ├── middlewares/            # JWT/sessão, request logger
│       ├── mail/                   # Envio de email
│       ├── code/                   # Códigos (ex: recuperação)
│       └── date/                   # Helpers de data
├── docs/                           # Documentação (este arquivo, api, etc.)
└── api-doc/                        # Doc estática servida na raiz
```

---

## 3. Camadas e responsabilidades

| Camada        | Onde                         | Responsabilidade |
|---------------|------------------------------|------------------|
| **Routes**    | `infrastructure/api/routes`  | Agrupar rotas (v1, auth, upload…), aplicar middlewares, chamar controllers. |
| **Controller**| `domains/<domain>/*.controller.go` | Parse de body/params, chamada ao service, montagem da resposta JSON e status HTTP. |
| **Service**   | `domains/<domain>/*.service.go`     | Regras de negócio; usa repositórios e outros serviços (ex: mail). Interfaces para testes. |
| **Repository**| `infrastructure/repositories/<entity>/` | Acesso a dados (goe); interface em `*-contract.go`, implementação em `*.repository.go`. |
| **Models**    | `internal/models/*.go`       | Structs com tags `json` e `db`; usadas por repositórios e services. |
| **Factory**   | `infrastructure/factory/factory.go` | Instanciar repositórios, services e controllers e expor em `AppControllers`. |

Controllers não acessam DB nem Redis diretamente; tudo passa por services e repositórios.

---

## 4. Padrões importantes

### 4.1 Repositórios

- **Interface** no arquivo `*-contract.go` (ex: `UserRepository` em `user.repository-contract.go`).
- **Implementação** no arquivo `*.repository.go`, recebendo `*database.StreamDB`.
- Métodos recebem/retornam `*models.<Entity>` quando há “out” (ex: `GetByEmail(email string, user *models.User) error`).

### 4.2 Services

- **Interface** pública (ex: `AuthService`) para permitir mocks em testes.
- **Implementação** privada (ex: `authService`); construtor retorna a interface (ex: `NewAuthService(...) AuthService`).
- Services recebem repositórios e outros serviços por parâmetro (injeção de dependência).

### 4.3 Controllers

- Struct com campo `service <Domain>Service` e construtor `New<Domain>Controller(service) *<Domain>Controller`.
- Handlers são métodos do controller (ex: `SignInController(ctx *fiber.Ctx) error`).
- Em erro: `ctx.Status(...).JSON(fiber.Map{"message": ..., "error": ...})`.
- Em sucesso: `ctx.JSON(fiber.Map{...})` ou `ctx.Status(...).JSON(...)`.

### 4.4 DTOs

- DTOs de request ficam no domínio (ex: `auth.dto.go`: `SignUpDto`, `SignInDto`, `ChangePasswordRequestRecoveryDto`).
- Usados apenas no controller para parse do body; o service recebe DTOs ou tipos simples.

### 4.5 Autenticação (rotas protegidas)

- Middleware: `middlewares.NewJwtSessionMiddleware(db)`.
- Após o middleware, o ID do usuário está em `c.Locals("userId").(string)`.
- O JWT contém o ID da **sessão**; o middleware valida sessão no DB e expiração.

### 4.6 Banco de dados (goe)

- Schema: `internal/infrastructure/database/schema.go` — struct `StreamDB` com campos `*models.User`, `*models.Recovery`, etc. e `*goe.DB`.
- Migrations: `db.AutoMigrate()` no `main.go`.
- Repositórios usam o `StreamDB` para acessar as “tabelas” (structs do goe).

---

## 5. Onde encontrar coisas

| Objetivo                    | Onde procurar |
|----------------------------|----------------|
| Adicionar nova rota        | `internal/infrastructure/api/routes/routes.go` |
| Novo domínio (ex: “orders”) | `internal/domains/<domain>/` (controller, service, DTOs) + registrar no `factory.go` e em `routes.go` |
| Novo repositório           | `internal/infrastructure/repositories/<entity>/` (contract + impl) + adicionar ao `StreamDB` em `schema.go` se for nova tabela |
| Novos models / tabelas     | `internal/models/*.go` + campo em `database.StreamDB` em `schema.go` |
| Config / env               | `internal/utilities/config/config.go` e `.env` |
| Middlewares                | `internal/utilities/middlewares/` |
| Logs                       | `internal/utilities/logger` (Info, Error, Debug) |

---

## 6. Modelos de dados (resumo)

- **User**: id, name, email, password, role (ADMIN, USER, ROOT), created_at, updated_at.
- **Session**: id, user_id, browser, ip, created_at, expires_at.
- **Recovery**: id, user_id, email, code, attempts, expired, expires_at, created_at, updated_at.
- **Organization**: id, name, owner_id, banner_id, deleted_at, created_at, updated_at; relação com Members.
- **Member**: id, user_id, organization_id, role (ADMIN, MEMBER, OWNER), created_at, updated_at.
- **Video**: id, file_*, owner_id, owner_name, created_at, updated_at, deleted_at.

Models estão separados em arquivos por domínio em `internal/models/`.

---

## 7. API (prefixo e grupos)

- Base: `v1` → rotas em `v1.Group("v1")`.
- Auth: `v1/auth/session` (sign-in, sign-up, listar/deletar sessões), `v1/auth/recovery` (request, change-password).
- Upload: `v1/upload/video` (POST).
- Stream: `v1/stream/videos/:filename` (GET).
- WebSocket: `GET /ws` (canal `videos_updates`).

Rotas protegidas usam `middlewares.NewJwtSessionMiddleware(db)` antes do handler.

---

## 8. Dependências principais

- `github.com/gofiber/fiber/v2` — HTTP.
- `github.com/go-goe/goe` + `github.com/go-goe/postgres` — ORM e migrations.
- `github.com/redis/go-redis/v9` — Redis.
- `github.com/dgrijalva/jwt-go` — JWT.
- `github.com/joho/godotenv` — .env.
- `github.com/google/uuid` — UUIDs.

---

## 9. Como adicionar um novo domínio (checklist)

1. Criar em `internal/domains/<domain>/`: controller, service (interface + impl), DTOs se necessário.
2. Se houver nova entidade persistida: model em `internal/models/`, campo em `StreamDB` em `schema.go`, repositório em `infrastructure/repositories/<entity>/` (contract + impl).
3. Em `factory/factory.go`: instanciar repositórios, service e controller; adicionar ao struct `AppControllers`.
4. Em `routes/routes.go`: criar grupo (ex: `v1.Group("organization")`), registrar rotas e middlewares.
5. Rodar `go build ./...` e testes relevantes.

Use este documento como referência ao ler ou modificar o código.
