# JoinLegends Backend

Backend desenvolvido em Go para a plataforma JoinLegends - uma plataforma onde jogadores podem entrar em comunidades, criar e participar de XTreinos (eventos de treinamento/competição).

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Tecnologias](#tecnologias)
- [Estrutura do Projeto](#estrutura-do-projeto)
- [Arquitetura](#arquitetura)
- [Configuração](#configuração)
- [Como Executar](#como-executar)
- [Documentação da API](#documentação-da-api)

## 🎯 Visão Geral

Este projeto implementa uma API RESTful usando Go com Fiber, seguindo uma arquitetura em camadas (Clean Architecture) com separação clara de responsabilidades. O sistema utiliza autenticação JWT com sessões para maior segurança e controle de acesso.

### Principais Funcionalidades

- ✅ Autenticação e autorização com JWT + Sessões
- ✅ Gerenciamento de usuários e sessões
- ✅ Recuperação de senha via email
- ✅ Upload e streaming de vídeos
- ✅ WebSocket para atualizações em tempo real
- ✅ Integração com Redis para filas e pub/sub
- ✅ Observabilidade com Grafana e Loki

## 🛠 Tecnologias

- **Go 1.24.3** - Linguagem principal
- **Fiber v2** - Framework web
- **PostgreSQL** - Banco de dados relacional
- **Redis** - Cache e filas
- **WebSocket** - Comunicação em tempo real
- **JWT** - Autenticação
- **Grafana + Loki** - Observabilidade e logs
- **Docker & Docker Compose** - Containerização

## 📁 Estrutura do Projeto

```
joinlegends-backend/
│
├── cmd/
│   └── server/
│       └── main.go                 # Ponto de entrada da aplicação
│
├── internal/                        # Código interno da aplicação
│   │
│   ├── domains/                    # Lógica de negócio por domínio
│   │   ├── auth/                   # Domínio de autenticação
│   │   │   ├── auth.controller.go # Controlador HTTP
│   │   │   ├── auth.dto.go         # Data Transfer Objects
│   │   │   ├── auth.service.go     # Lógica de negócio
│   │   │   └── tests/              # Testes unitários
│   │   │       ├── auth_mocks_test.go
│   │   │       ├── auth_service_recovery_test.go
│   │   │       ├── auth_service_session_test.go
│   │   │       ├── auth_service_signin_session_test.go
│   │   │       ├── auth_service_signin_test.go
│   │   │       ├── auth_service_signup_session_test.go
│   │   │       └── auth_service_signup_test.go
│   │   │
│   │   ├── organization/           # Domínio de organizações
│   │   │   └── organization.go
│   │   │
│   │   ├── stream/                 # Domínio de streaming
│   │   │   └── stream.go          # Handler de streaming de vídeos
│   │   │
│   │   └── upload/                 # Domínio de upload
│   │       └── upload.go          # Handler de upload de arquivos
│   │
│   ├── infrastructure/             # Infraestrutura e adaptadores
│   │   ├── api/
│   │   │   └── routes/
│   │   │       └── routes.go      # Configuração de rotas da API
│   │   │
│   │   ├── database/
│   │   │   ├── database.go        # Conexão e configuração do DB
│   │   │   └── schema.go          # Schema do banco de dados
│   │   │
│   │   ├── factory/
│   │   │   └── factory.go         # Factory para criação de dependências
│   │   │
│   │   ├── redisclient/
│   │   │   ├── messages.go        # Mensagens para Redis
│   │   │   ├── producer.go        # Produtor de mensagens
│   │   │   ├── queues.go          # Configuração de filas
│   │   │   └── redisclient.go     # Cliente Redis
│   │   │
│   │   └── repositories/          # Camada de acesso a dados
│   │       ├── recovery/
│   │       │   ├── recovery.repository-contract.go  # Interface
│   │       │   └── recovery.repository.go          # Implementação
│   │       ├── session/
│   │       │   ├── session.repository-contract.go  # Interface
│   │       │   └── session.repository.go          # Implementação
│   │       └── user/
│   │           ├── user.repository-contract.go     # Interface
│   │           └── user.repository.go             # Implementação
│   │
│   ├── models/                     # Modelos de dados
│   │   └── models.go              # Definições de structs (User, Session, Recovery, etc.)
│   │
│   └── utilities/                 # Utilitários e helpers
│       ├── code/
│       │   └── code.go           # Geração de códigos (ex: recuperação)
│       ├── config/
│       │   └── config.go         # Configurações e variáveis de ambiente
│       ├── date/
│       │   └── date.go           # Utilitários de data/hora
│       ├── jwt/
│       │   └── jwt.go            # Geração e validação de tokens JWT
│       ├── logger/
│       │   └── logger.go         # Sistema de logging
│       ├── mail/
│       │   └── mail.go           # Serviço de envio de emails
│       └── middlewares/
│           ├── middleware.go     # Middlewares HTTP
│           └── requestLogger.go # Logger de requisições
│
├── api-doc/                       # Documentação da API
│   ├── index.html
│   ├── styles.css
│   └── doc.png
│
├── docs/                          # Documentação adicional
│   └── api/
│       └── README.md
│
├── docker-compose.yml             # Configuração Docker Compose
├── go.mod                         # Dependências Go
├── go.sum                         # Checksums das dependências
├── local-config.yaml              # Configuração local (Loki)
└── README.md                      # Este arquivo
```

## 🏗 Arquitetura

O projeto segue uma arquitetura em camadas (Clean Architecture) com separação clara de responsabilidades:

### Camadas

1. **Presentation Layer** (`cmd/server/main.go`, `internal/infrastructure/api/routes/`)
   - Ponto de entrada da aplicação
   - Configuração de rotas HTTP
   - Middlewares de autenticação e logging

2. **Domain Layer** (`internal/domains/`)
   - Lógica de negócio pura
   - Controllers, Services e DTOs
   - Independente de frameworks e infraestrutura

3. **Infrastructure Layer** (`internal/infrastructure/`)
   - Implementações concretas (repositories, database, Redis)
   - Adaptadores para serviços externos
   - Factory para injeção de dependências

4. **Models Layer** (`internal/models/`)
   - Estruturas de dados compartilhadas
   - Entidades do domínio

5. **Utilities Layer** (`internal/utilities/`)
   - Funções auxiliares reutilizáveis
   - Configurações, logging, JWT, email, etc.

### Fluxo de Dados

```
HTTP Request
    ↓
Routes (routes.go)
    ↓
Controller (auth.controller.go)
    ↓
Service (auth.service.go)
    ↓
Repository (user.repository.go, session.repository.go, etc.)
    ↓
Database (PostgreSQL)
```

### Autenticação JWT com Sessões

O sistema implementa autenticação JWT combinada com sessões para maior segurança:

- **Token JWT** assinado com o ID da sessão (não do usuário)
- **Sessões** armazenadas no banco de dados
- **Controle** de múltiplas sessões por usuário
- **Invalidação** imediata ao alterar senha ou deletar sessão

## ⚙️ Configuração

### Variáveis de Ambiente

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

```env
# PostgreSQL
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="postgres"
DB_HOST="localhost"
POSTGRES_PORT=5432
POSTGRES_DB="jwt_sessions_db"

# JWT
JWT_SEC_KEY="sua-chave-secreta-super-segura-aqui"

# Email (recomendado: Mailtrap para testes)
MAIL_HOST="smtp.mailtrap.io"
MAIL_PORT=2525
MAIL_USER="seu-usuario"
MAIL_PASS="sua-senha"
ROOT_EMAIL="noreply@joinlegends.com"

# Grafana/Loki (opcional)
GRAFANA_INITIAL_USER=admin
GRAFANA_INITIAL_PASSWORD=admin
LOKI_CONNECTION="http://localhost:3100"
```

## 🚀 Como Executar

### Pré-requisitos

- [Go](https://go.dev/) 1.24.3 ou superior
- [Docker](https://www.docker.com/) e Docker Compose
- Git

### Passos

1. **Clone o repositório** (se ainda não tiver feito)
   ```bash
   git clone <repository-url>
   cd joinlegends-backend
   ```

2. **Instale as dependências**
   ```bash
   go mod tidy
   ```

3. **Configure as variáveis de ambiente**
   - Crie o arquivo `.env` na raiz do projeto
   - Preencha com os valores apropriados (veja seção [Configuração](#configuração))

4. **Suba os serviços de infraestrutura**
   ```bash
   docker compose up -d
   ```
   
   Isso iniciará:
   - PostgreSQL (porta 5432)
   - Redis (porta 6379)
   - Loki (porta 3100)
   - Grafana (porta 9090)

5. **Execute a aplicação**
   ```bash
   go run cmd/server/main.go
   ```

   Ou compile e execute:
   ```bash
   go build -o bin/server cmd/server/main.go
   ./bin/server
   ```

6. **Acesse a aplicação**
   - API: http://localhost:8080
   - Documentação: http://localhost:8080 (servida estaticamente)
   - Grafana: http://localhost:9090

## 📚 Documentação da API

A documentação da API está disponível em `http://localhost:8080` após iniciar o servidor.

### Principais Endpoints

#### Autenticação (`/v1/auth/session`)

- `POST /v1/auth/session/sign-up` - Cadastro de usuário
- `POST /v1/auth/session/sign-in` - Login
- `GET /v1/auth/session/` - Validar sessão (requer autenticação)
- `GET /v1/auth/session/all` - Listar todas as sessões do usuário
- `DELETE /v1/auth/session/:sessionId` - Deletar uma sessão específica

#### Recuperação de Senha (`/v1/auth/recovery`)

- `POST /v1/auth/recovery/request/:email` - Solicitar código de recuperação
- `POST /v1/auth/recovery/change-password` - Alterar senha com código

#### Upload (`/v1/upload`)

- `POST /v1/upload/video` - Upload de vídeo

#### Streaming (`/v1/stream`)

- `GET /v1/stream/videos/:filename` - Stream de vídeo

#### WebSocket

- `WS /ws` - Conexão WebSocket para atualizações em tempo real

## 🧪 Testes

Execute os testes com:

```bash
go test ./...
```

Para executar testes de um pacote específico:

```bash
go test ./internal/domains/auth/tests/...
```

## 📝 Modelos de Dados

### User
- `ID` (string) - UUID do usuário
- `Name` (string) - Nome completo
- `Email` (string) - Email único
- `Password` (string) - Senha hasheada
- `Role` (string) - ADMIN, USER, ROOT
- `CreatedAt`, `UpdatedAt` (time.Time)

### Session
- `ID` (string) - UUID da sessão
- `UserID` (string) - ID do usuário
- `Browser` (*string) - Navegador utilizado
- `IP` (*string) - Endereço IP
- `CreatedAt`, `ExpiresAt` (time.Time)

### Recovery
- `ID` (int) - ID único
- `UserID` (string) - ID do usuário
- `Email` (string) - Email para recuperação
- `Code` (string) - Código de recuperação
- `Attempts` (int) - Tentativas de uso
- `Expired` (bool) - Status de expiração
- `ExpiresAt`, `CreatedAt`, `UpdatedAt` (time.Time)

### Organization
- `ID` (string) - UUID da organização
- `Name` (string) - Nome da organização
- `OwnerID` (string) - ID do dono
- `BannerID` (*string) - ID da imagem banner
- `Members` ([]Member) - Membros da organização
- `CreatedAt`, `UpdatedAt`, `DeletedAt` (time.Time)

### Video
- `ID` (string) - UUID do vídeo
- `FileName`, `FilePath` (string) - Informações do arquivo
- `FileSize`, `FileDuration`, `FileBitrate` (int64) - Metadados
- `FileType`, `FileHash`, `FileResolution`, `FileCodec` (string) - Detalhes técnicos
- `OwnerID`, `OwnerName` (string) - Informações do proprietário
- `CreatedAt`, `UpdatedAt`, `DeletedAt` (time.Time)

## 🔒 Segurança

- Autenticação JWT com sessões para controle de acesso
- Senhas hasheadas (não armazenadas em texto plano)
- Middleware de autenticação em rotas protegidas
- Validação de entrada em todos os endpoints
- Códigos de recuperação com expiração e limite de tentativas
- Logging de requisições para auditoria

## 📦 Dependências Principais

- `github.com/gofiber/fiber/v2` - Framework web
- `github.com/gofiber/websocket/v2` - WebSocket support
- `github.com/go-goe/goe` - ORM para PostgreSQL
- `github.com/redis/go-redis/v9` - Cliente Redis
- `github.com/dgrijalva/jwt-go` - JWT
- `gopkg.in/mail.v2` - Envio de emails
- `github.com/joho/godotenv` - Gerenciamento de variáveis de ambiente

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob licença proprietária.

## 👥 Autores

- Equipe JoinLegends

---

**Nota**: Este README documenta a estrutura atual do projeto. Para mais detalhes sobre a implementação de JWT com sessões, consulte o README original do projeto.
