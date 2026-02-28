# Organization API Documentation

Endpoints para gerenciamento de organizações e seus membros. Todas as rotas requerem autenticação via JWT.

## Base URL
`/v1/organization`

## Autenticação

Todas as rotas exigem o header:
```
Authorization: Bearer <token>
```

O `user_id` do usuário autenticado é extraído automaticamente do token — não é necessário enviá-lo no body.

---

## Endpoints

### 1. Criar Organização

Cria uma nova organização. O usuário autenticado se torna automaticamente o **owner** e é adicionado como membro com role `OWNER`.

- **URL:** `POST /v1/organization/`
- **Auth:** Requerida

**Request Body:**

| Campo | Tipo   | Descrição               | Obrigatório |
|-------|--------|-------------------------|-------------|
| `name`| string | Nome da organização     | Sim         |

```json
{
  "name": "Minha Organização"
}
```

**Response (201 Created):**

```json
{
  "message": "organização criada com sucesso",
  "organization": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Minha Organização",
    "owner_id": "user-uuid",
    "banner_id": null,
    "validated_at": null,
    "deleted_at": null,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z",
    "members": [
      {
        "id": "member-uuid",
        "user_id": "user-uuid",
        "organization_id": "550e8400-e29b-41d4-a716-446655440000",
        "role": "OWNER",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      }
    ]
  }
}
```

**Erros:**

| Status | Descrição |
|--------|-----------|
| `401`  | Token inválido ou expirado |
| `500`  | Erro interno no servidor |

---

### 2. Listar Organizações do Usuário

Retorna todas as organizações onde o usuário autenticado é o **owner**.

- **URL:** `GET /v1/organization/`
- **Auth:** Requerida

**Response (200 OK):**

```json
{
  "organizations": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "Organização Alpha",
      "owner_id": "user-uuid",
      "banner_id": null,
      "validated_at": null,
      "deleted_at": null,
      "created_at": "2024-01-01T10:00:00Z",
      "updated_at": "2024-01-01T10:00:00Z",
      "members": null
    },
    {
      "id": "661f9511-f30c-52e5-b827-557766551111",
      "name": "Organização Beta",
      "owner_id": "user-uuid",
      "banner_id": null,
      "validated_at": "2024-01-03T15:30:00Z",
      "deleted_at": null,
      "created_at": "2024-01-02T10:00:00Z",
      "updated_at": "2024-01-03T15:30:00Z",
      "members": null
    }
  ]
}
```

---

### 3. Buscar Organização por ID

Retorna os dados de uma organização específica, incluindo a lista completa de membros.

- **URL:** `GET /v1/organization/:id`
- **Auth:** Requerida
- **URL Params:** `id` (string) — UUID da organização

**Response (200 OK):**

```json
{
  "organization": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "Organização Alpha",
    "owner_id": "user-uuid",
    "banner_id": null,
    "validated_at": null,
    "deleted_at": null,
    "created_at": "2024-01-01T10:00:00Z",
    "updated_at": "2024-01-01T10:00:00Z",
    "members": [
      {
        "id": "member-uuid-1",
        "user_id": "user-uuid",
        "organization_id": "550e8400-e29b-41d4-a716-446655440000",
        "role": "OWNER",
        "created_at": "2024-01-01T10:00:00Z",
        "updated_at": "2024-01-01T10:00:00Z"
      },
      {
        "id": "member-uuid-2",
        "user_id": "other-user-uuid",
        "organization_id": "550e8400-e29b-41d4-a716-446655440000",
        "role": "MEMBER",
        "created_at": "2024-01-03T10:00:00Z",
        "updated_at": "2024-01-03T10:00:00Z"
      }
    ]
  }
}
```

**Erros:**

| Status | Descrição |
|--------|-----------|
| `404`  | Organização não encontrada |
| `500`  | Erro interno no servidor |

---

### 4. Atualizar Organização

Atualiza o nome de uma organização. Somente o **owner** pode realizar esta ação.

- **URL:** `PATCH /v1/organization/:id`
- **Auth:** Requerida
- **URL Params:** `id` (string) — UUID da organização

**Request Body:**

| Campo | Tipo   | Descrição               | Obrigatório |
|-------|--------|-------------------------|-------------|
| `name`| string | Novo nome da organização| Sim         |

```json
{
  "name": "Novo Nome da Organização"
}
```

**Response (200 OK):**

```json
{
  "message": "organização atualizada com sucesso"
}
```

**Erros:**

| Status | Descrição |
|--------|-----------|
| `403`  | Usuário não é o owner da organização |
| `404`  | Organização não encontrada |
| `500`  | Erro interno no servidor |

---

### 5. Deletar Organização

Realiza um **soft delete** na organização (campo `deleted_at` é preenchido). Somente o **owner** pode realizar esta ação.

- **URL:** `DELETE /v1/organization/:id`
- **Auth:** Requerida
- **URL Params:** `id` (string) — UUID da organização

**Response (200 OK):**

```json
{
  "message": "organização deletada com sucesso"
}
```

**Erros:**

| Status | Descrição |
|--------|-----------|
| `403`  | Usuário não é o owner da organização |
| `404`  | Organização não encontrada |
| `500`  | Erro interno no servidor |

---

### 6. Adicionar Membro

Adiciona um usuário como membro da organização. Somente o **owner** pode realizar esta ação.

- **URL:** `POST /v1/organization/:id/members`
- **Auth:** Requerida
- **URL Params:** `id` (string) — UUID da organização

**Request Body:**

| Campo     | Tipo   | Descrição                        | Obrigatório |
|-----------|--------|----------------------------------|-------------|
| `user_id` | string | UUID do usuário a ser adicionado | Sim         |
| `role`    | string | Role do membro: `ADMIN`, `MEMBER`| Sim         |

```json
{
  "user_id": "other-user-uuid",
  "role": "MEMBER"
}
```

**Response (201 Created):**

```json
{
  "message": "membro adicionado com sucesso",
  "member": {
    "id": "member-uuid",
    "user_id": "other-user-uuid",
    "organization_id": "550e8400-e29b-41d4-a716-446655440000",
    "role": "MEMBER",
    "created_at": "2024-01-03T10:00:00Z",
    "updated_at": "2024-01-03T10:00:00Z"
  }
}
```

**Erros:**

| Status | Descrição |
|--------|-----------|
| `403`  | Usuário não é o owner da organização |
| `404`  | Organização ou usuário não encontrado |
| `500`  | Erro interno no servidor |

---

### 7. Remover Membro

Remove um membro da organização. Somente o **owner** pode realizar esta ação. O próprio owner não pode ser removido.

- **URL:** `DELETE /v1/organization/:id/members/:memberId`
- **Auth:** Requerida
- **URL Params:**
  - `id` (string) — UUID da organização
  - `memberId` (string) — UUID do membro a ser removido

**Response (200 OK):**

```json
{
  "message": "membro removido com sucesso"
}
```

**Erros:**

| Status | Descrição |
|--------|-----------|
| `403`  | Usuário não é o owner, ou tentativa de remover o próprio owner |
| `404`  | Organização ou membro não encontrado |
| `500`  | Erro interno no servidor |

---

## Roles de Membros

| Role     | Descrição |
|----------|-----------|
| `OWNER`  | Criador da organização. Atribuído automaticamente na criação. Não pode ser removido. |
| `ADMIN`  | Membro com permissões administrativas (a serem definidas por regra de negócio). |
| `MEMBER` | Membro padrão da organização. |

---

## Modelos

### Organization

```json
{
  "id": "string (UUID)",
  "name": "string",
  "owner_id": "string (UUID)",
  "banner_id": "string | null",
  "validated_at": "string (ISO 8601) | null",
  "deleted_at": "string (ISO 8601) | null",
  "created_at": "string (ISO 8601)",
  "updated_at": "string (ISO 8601)",
  "members": "Member[]"
}
```

> `validated_at` é `null` enquanto a organização estiver pendente de validação. Quando preenchida, indica o momento em que a organização foi tornada pública/validada.
```

### Member

```json
{
  "id": "string (UUID)",
  "user_id": "string (UUID)",
  "organization_id": "string (UUID)",
  "role": "OWNER | ADMIN | MEMBER",
  "created_at": "string (ISO 8601)",
  "updated_at": "string (ISO 8601)"
}
```
