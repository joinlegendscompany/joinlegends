# JoinLegends Frontend

Plataforma de comunidades para jogadores com foco em treinos, competições e networking. Interface dark com tema cyberpunk.

## 🎮 Sobre o Projeto

JoinLegends é uma plataforma onde jogadores podem:
- **Entrar em comunidades** de jogos específicos
- **Criar e participar de XTreinos** (sessões de treino personalizadas)
- **Conectar-se com outros jogadores** e formar equipes
- **Acompanhar progresso** e estatísticas de performance
- **Participar de competições** e eventos

## 🎨 Design System

### Tema: Dark Cyberpunk

A plataforma utiliza um design system inspirado no estilo cyberpunk com:
- **Paleta de cores**: Tons escuros (#0a0a0a, #1a1a2e) com acentos neon (ciano, magenta, verde)
- **Tipografia**: Fontes modernas e tecnológicas
- **Efeitos visuais**: Glows, gradientes, animações suaves
- **UI Elements**: Bordas brilhantes, sombras coloridas, elementos futuristas

### Cores Principais
```
Background: #0a0a0a (Preto profundo)
Surface: #1a1a2e (Azul escuro)
Primary: #00f5ff (Ciano neon)
Secondary: #ff00ff (Magenta neon)
Accent: #00ff41 (Verde neon)
Text: #ededed (Branco suave)
```

## 🛠️ Tecnologias

- **Next.js 16** - Framework React com App Router
- **React 19** - Biblioteca UI
- **TypeScript** - Tipagem estática
- **Tailwind CSS 4** - Estilização utilitária
- **ESLint** - Linting de código

## 📁 Estrutura do Projeto

```
joinlegends-frontend/
├── app/
│   ├── globals.css          # Estilos globais e tema
│   ├── layout.tsx           # Layout raiz
│   └── page.tsx             # Página inicial
├── public/                  # Assets estáticos
├── package.json
├── tsconfig.json
└── README.md
```

## 🚀 Como Rodar

### Pré-requisitos
- Node.js 18+ instalado
- npm ou yarn

### Instalação

```bash
# Instalar dependências
npm install

# Rodar em modo desenvolvimento
npm run dev

# Build para produção
npm run build

# Iniciar servidor de produção
npm start
```

A aplicação estará disponível em `http://localhost:3000`

## 🗺️ Roadmap

### Fase 1: Fundação e Design System (Atual) ⏳
- [x] Setup inicial do projeto Next.js
- [x] Configuração do Tailwind CSS
- [ ] Implementar design system cyberpunk completo
- [ ] Criar componentes base (Button, Card, Input)
- [ ] Configurar tema dark/cyberpunk
- [ ] Sistema de cores e variáveis CSS
- [ ] Tipografia e fontes

### Fase 2: Autenticação e Onboarding 🎯
- [ ] Página de Login
- [ ] Página de Registro
- [ ] Integração com backend de autenticação
- [ ] Fluxo de recuperação de senha
- [ ] Onboarding inicial para novos usuários
- [ ] Perfil básico do usuário

### Fase 3: Dashboard Principal 🏠
- [ ] Layout principal com sidebar/navbar
- [ ] Dashboard home com estatísticas
- [ ] Feed de atividades recentes
- [ ] Notificações em tempo real
- [ ] Busca global
- [ ] Menu de navegação

### Fase 4: Comunidades 👥
- [ ] Listagem de comunidades disponíveis
- [ ] Página de detalhes da comunidade
- [ ] Sistema de busca e filtros
- [ ] Criar nova comunidade
- [ ] Entrar/sair de comunidades
- [ ] Feed da comunidade
- [ ] Membros e roles
- [ ] Configurações da comunidade (admin)

### Fase 5: XTreinos 💪
- [ ] Criar XTreino
- [ ] Listar XTreinos disponíveis
- [ ] Detalhes do XTreino
- [ ] Participar de XTreino
- [ ] Agendar sessões
- [ ] Histórico de treinos
- [ ] Sistema de convites
- [ ] Calendário de treinos

### Fase 6: Perfil e Progresso 📊
- [ ] Página de perfil completo
- [ ] Estatísticas de performance
- [ ] Histórico de treinos
- [ ] Conquistas e badges
- [ ] Gráficos e visualizações
- [ ] Editar perfil
- [ ] Upload de avatar/banner

### Fase 7: Social e Networking 🤝
- [ ] Sistema de amigos/seguidores
- [ ] Chat em tempo real
- [ ] Mensagens diretas
- [ ] Feed social
- [ ] Compartilhar conquistas
- [ ] Recomendações de jogadores

### Fase 8: Competições e Eventos 🏆
- [ ] Listar competições
- [ ] Criar competição
- [ ] Participar de competições
- [ ] Leaderboards
- [ ] Resultados e premiações
- [ ] Calendário de eventos

### Fase 9: Notificações e Real-time 🔔
- [ ] Sistema de notificações push
- [ ] WebSocket para atualizações em tempo real
- [ ] Notificações de convites
- [ ] Alertas de treinos próximos
- [ ] Notificações de comunidade

### Fase 10: Otimização e Polimento ✨
- [ ] Performance optimization
- [ ] SEO
- [ ] Acessibilidade (a11y)
- [ ] Testes E2E
- [ ] Responsividade mobile
- [ ] PWA (Progressive Web App)
- [ ] Internacionalização (i18n)

## 🎯 Próximos Passos Imediatos

1. **Implementar Design System Cyberpunk**
   - Criar componentes base com tema dark
   - Configurar paleta de cores neon
   - Adicionar efeitos de glow e animações

2. **Páginas de Autenticação**
   - Login e Registro com estilo cyberpunk
   - Integração com API backend

3. **Layout Principal**
   - Sidebar com navegação
   - Header com busca e notificações
   - Container principal responsivo

## 📝 Notas de Desenvolvimento

- O projeto utiliza App Router do Next.js 16
- Tailwind CSS 4 para estilização
- TypeScript para type safety
- Design mobile-first e responsivo

## 🤝 Contribuindo

Este é um projeto em desenvolvimento ativo. Para contribuir:
1. Siga o design system cyberpunk estabelecido
2. Mantenha consistência com o código existente
3. Teste suas mudanças antes de commitar

## 📄 Licença

[Adicionar licença conforme necessário]

---

**Status**: 🚧 Em Desenvolvimento Ativo

**Última atualização**: 2024
