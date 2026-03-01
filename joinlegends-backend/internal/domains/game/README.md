# Game API Documentation

Endpoints para gerenciamento de jogos cadastrados na plataforma.

## Base URL
`/v1/game`

---

## Modelo

| Campo | Tipo | Descrição |
|---|---|---|
| `id` | string (UUID) | Identificador único do jogo |
| `name` | string | Nome do jogo |
| `banner_id` | string \| null | ID do banner (referência ao upload) |
| `category` | string | Categoria do jogo (ex: FPS, MOBA, RPG) |
| `description` | string | Descrição do jogo |
| `developer` | string | Desenvolvedora do jogo |
| `publisher` | string \| null | Publicadora do jogo |
| `release_year` | int | Ano de lançamento |
| `is_active` | bool | Se o jogo está ativo na plataforma |
| `created_at` | string (ISO 8601) | Data de criação |
| `updated_at` | string (ISO 8601) | Data da última atualização |
| `deleted_at` | string \| null | Data de remoção (soft delete) |

---

## Endpoints

### 1. Criar jogo
Cadastra um novo jogo na plataforma.

- **URL:** `POST /v1/game`
- **Auth:** `Authorization: Bearer <token>` (obrigatório)
- **Content-Type:** `application/json`

**Request Body:**

| Campo | Tipo | Obrigatório | Descrição |
|---|---|---|---|
| `name` | string | Sim | Nome do jogo |
| `category` | string | Sim | Categoria do jogo |
| `description` | string | Sim | Descrição do jogo |
| `developer` | string | Sim | Desenvolvedora |
| `release_year` | int | Sim | Ano de lançamento |
| `banner_id` | string | Não | ID do banner já enviado via upload |
| `publisher` | string | Não | Publicadora do jogo |

```json
{
  "name": "Counter-Strike 2",
  "category": "FPS",
  "description": "O FPS tático mais jogado do mundo.",
  "developer": "Valve",
  "publisher": "Valve",
  "release_year": 2023,
  "banner_id": "550e8400-e29b-41d4-a716-446655440000"
}
```

**Response `201 Created`:**
```json
{
  "message": "jogo cadastrado com sucesso",
  "game": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "Counter-Strike 2",
    "banner_id": "550e8400-e29b-41d4-a716-446655440000",
    "category": "FPS",
    "description": "O FPS tático mais jogado do mundo.",
    "developer": "Valve",
    "publisher": "Valve",
    "release_year": 2023,
    "is_active": true,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "deleted_at": null
  }
}
```

**Response `500 Internal Server Error`:**
```json
{
  "message": "internal server error when create a new game",
  "error": "..."
}
```

---

### 2. Listar jogos
Retorna todos os jogos ativos. Aceita filtro opcional por categoria via query param.

- **URL:** `GET /v1/game`
- **Auth:** Não obrigatório

**Query Params (opcionais):**

| Param | Tipo | Descrição |
|---|---|---|
| `category` | string | Filtra jogos por categoria |

**Exemplos:**
```
GET /v1/game
GET /v1/game?category=FPS
GET /v1/game?category=MOBA
```

**Response `200 OK`:**
```json
{
  "games": [
    {
      "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
      "name": "Counter-Strike 2",
      "banner_id": "550e8400-e29b-41d4-a716-446655440000",
      "category": "FPS",
      "description": "O FPS tático mais jogado do mundo.",
      "developer": "Valve",
      "publisher": "Valve",
      "release_year": 2023,
      "is_active": true,
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z",
      "deleted_at": null
    }
  ]
}
```

---

### 3. Buscar jogo por ID
Retorna os dados de um jogo específico.

- **URL:** `GET /v1/game/:id`
- **Auth:** Não obrigatório

**URL Params:**

| Param | Tipo | Descrição |
|---|---|---|
| `id` | string (UUID) | ID do jogo |

**Response `200 OK`:**
```json
{
  "game": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "Counter-Strike 2",
    "banner_id": "550e8400-e29b-41d4-a716-446655440000",
    "category": "FPS",
    "description": "O FPS tático mais jogado do mundo.",
    "developer": "Valve",
    "publisher": "Valve",
    "release_year": 2023,
    "is_active": true,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T10:30:00Z",
    "deleted_at": null
  }
}
```

**Response `404 Not Found`:**
```json
{
  "message": "game with id a1b2c3d4-e5f6-7890-abcd-ef1234567890 not found"
}
```

---

### 4. Atualizar jogo
Atualiza parcialmente os dados de um jogo. Todos os campos são opcionais — apenas os campos enviados serão alterados.

- **URL:** `PATCH /v1/game/:id`
- **Auth:** `Authorization: Bearer <token>` (obrigatório)
- **Content-Type:** `application/json`

**URL Params:**

| Param | Tipo | Descrição |
|---|---|---|
| `id` | string (UUID) | ID do jogo |

**Request Body (todos opcionais):**

| Campo | Tipo | Descrição |
|---|---|---|
| `name` | string | Novo nome do jogo |
| `banner_id` | string | Novo ID do banner |
| `category` | string | Nova categoria |
| `description` | string | Nova descrição |
| `developer` | string | Nova desenvolvedora |
| `publisher` | string | Nova publicadora |
| `release_year` | int | Novo ano de lançamento |
| `is_active` | bool | Ativar ou desativar o jogo |

```json
{
  "name": "CS2 - Updated",
  "is_active": false
}
```

**Response `200 OK`:**
```json
{
  "message": "jogo atualizado com sucesso",
  "game": {
    "id": "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
    "name": "CS2 - Updated",
    "banner_id": "550e8400-e29b-41d4-a716-446655440000",
    "category": "FPS",
    "description": "O FPS tático mais jogado do mundo.",
    "developer": "Valve",
    "publisher": "Valve",
    "release_year": 2023,
    "is_active": false,
    "created_at": "2024-01-15T10:30:00Z",
    "updated_at": "2024-01-15T11:00:00Z",
    "deleted_at": null
  }
}
```

**Response `404 Not Found`:**
```json
{
  "message": "game with id a1b2c3d4-e5f6-7890-abcd-ef1234567890 not found"
}
```

---

### 5. Remover jogo
Remove um jogo da plataforma via soft delete (o registro permanece no banco com `deleted_at` preenchido).

- **URL:** `DELETE /v1/game/:id`
- **Auth:** `Authorization: Bearer <token>` (obrigatório)

**URL Params:**

| Param | Tipo | Descrição |
|---|---|---|
| `id` | string (UUID) | ID do jogo |

**Response `200 OK`:**
```json
{
  "message": "jogo removido com sucesso"
}
```

**Response `404 Not Found`:**
```json
{
  "message": "game with id a1b2c3d4-e5f6-7890-abcd-ef1234567890 not found"
}
```

---

## Categorias sugeridas

| Valor | Descrição |
|---|---|
| `FPS` | First Person Shooter |
| `MOBA` | Multiplayer Online Battle Arena |
| `Battle Royale` | Batalha até o último sobrevivente |
| `RPG` | Role Playing Game |
| `Sports` | Jogos esportivos |
| `Strategy` | Jogos de estratégia |
| `Fighting` | Jogos de luta |
| `Sandbox` | Mundo aberto e criação livre |
| `Simulation` | Simuladores |
| `Horror` | Terror e suspense |
| `Adventure` | Aventura e exploração |
