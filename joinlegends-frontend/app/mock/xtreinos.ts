export type TreinosData = {
  id: number;
  titulo: string;
  jogo: string;
  criador: string;
  avatarCriador: string;
  organizacao?: {
    id: string;
    nome: string;
    logo: string;
  };
  data: string; // DD/MM/YYYY
  horario: string; // HH:mm
  duracao: string; // ex: "2 horas"
  vagas: number;
  vagasTotal: number;
  nivel: "Iniciante" | "Intermediário" | "Avançado";
  descricao: string;
  descricaoCompleta: string;
  imagem: string;
  logo: string;
  valorInscricao: number; // Valor em reais
  valorPremiacao: number; // Valor total da premiação em reais
  participantes: {
    nome: string;
    avatar: string;
    rank: string;
  }[];
  tags: string[];
};

// Dados mockados de XTreinos
export const xtreinos = [
  {
    id: 1,
    titulo: "Treino de Aim - Bloodstrike",
    jogo: "Bloodstrike",
    criador: "ProPlayer123",
    organizacao: {
      id: "loud",
      nome: "LOUD",
      logo: "/assets/x0.jpeg",
    },
    data: "15/01/2024",
    horario: "20:00",
    vagas: "8/10",
    nivel: "Intermediário",
    descricao:
      "Sessão focada em melhorar precisão e reflexos. Trabalharemos técnicas de flick e tracking.",
    imagem: "/assets/x0.jpeg",
    logo: "/assets/games/bloodstrike.png",
    valorInscricao: 25.00,
    valorPremiacao: 500.00,
  },
  {
    id: 2,
    titulo: "Estratégias Táticas - Delta Force",
    jogo: "Delta Force",
    criador: "TacticalMaster",
    data: "16/01/2024",
    horario: "19:30",
    vagas: "5/8",
    nivel: "Avançado",
    descricao:
      "Análise de mapas e execução de estratégias em equipe. Ideal para jogadores que querem subir de rank.",
    imagem: "/assets/x1.jpeg",
    logo: "/assets/games/deltaforce.png",
    valorInscricao: 35.00,
    valorPremiacao: 800.00,
  },
  {
    id: 3,
    titulo: "Fundamentos - Free Fire",
    jogo: "Free Fire",
    criador: "CoachFF",
    data: "17/01/2024",
    horario: "18:00",
    vagas: "12/15",
    nivel: "Iniciante",
    descricao:
      "Aprenda os fundamentos do jogo: posicionamento, estratégia e sobrevivência. Perfeito para novos jogadores.",
    imagem: "/assets/x2.jpeg",
    logo: "/assets/games/freefire.png",
    valorInscricao: 15.00,
    valorPremiacao: 300.00,
  },
  {
    id: 4,
    titulo: "Scrims Competitivos - Free Fire",
    jogo: "Free Fire",
    criador: "ScrimOrganizer",
    data: "18/01/2024",
    horario: "21:00",
    vagas: "16/20",
    nivel: "Avançado",
    descricao:
      "Scrims organizados com análise pós-jogo. Ambiente competitivo para melhorar em equipe.",
    imagem: "/assets/x3.jpg",
    logo: "/assets/games/freefire.png",
    valorInscricao: 50.00,
    valorPremiacao: 1200.00,
  },
  {
    id: 5,
    titulo: "Mechanics Training - Bloodstrike",
    jogo: "Bloodstrike",
    criador: "RLPro",
    data: "19/01/2024",
    horario: "20:30",
    vagas: "6/8",
    nivel: "Intermediário",
    descricao:
      "Treino de mecânicas avançadas: movimentação, mira e controle de spray. Foco em consistência e precisão.",
    imagem: "/assets/x0.jpeg",
    logo: "/assets/games/bloodstrike.png",
    valorInscricao: 30.00,
    valorPremiacao: 600.00,
  },
  {
    id: 6,
    titulo: "VOD Review - Delta Force",
    jogo: "Delta Force",
    criador: "DeltaCoach",
    data: "20/01/2024",
    horario: "19:00",
    vagas: "10/12",
    nivel: "Todos os níveis",
    descricao:
      "Análise de gameplay gravado. Identifique erros e aprenda com situações reais de jogo.",
    imagem: "/assets/x1.jpeg",
    logo: "/assets/games/deltaforce.png",
    valorInscricao: 20.00,
    valorPremiacao: 400.00,
  },
];

// Dados mockados - em produção viria de uma API
export const xtreinosData: TreinosData = {
  id: 1,
  titulo: "Treino de Aim - Bloodstrike",
  jogo: "Bloodstrike",
  criador: "ProPlayer123",
  avatarCriador: "/assets/x0.jpeg",
  organizacao: {
    id: "loud",
    nome: "LOUD",
    logo: "/assets/x0.jpeg",
  },
  data: "15/01/2024",
  horario: "20:00",
  duracao: "2 horas",
  vagas: 8,
  vagasTotal: 10,
  nivel: "Intermediário",
  descricao:
    "Sessão focada em melhorar precisão e reflexos. Trabalharemos técnicas de flick e tracking.",
  descricaoCompleta: `<p class="mb-4">Nesta sessão de treino, vamos focar em desenvolver habilidades fundamentais de mira e precisão no Bloodstrike.</p>

<p class="mb-4">O treino será dividido em três partes principais:</p>

<ol class="mb-4 space-y-2">
  <li>
    <strong>Aquecimento e Fundamentos (30min)</strong>
    <ul class="ml-6 mt-1">
      <li>Exercícios básicos de mira estática</li>
      <li>Prática de controle de spray</li>
      <li>Ajuste de sensibilidade individual</li>
    </ul>
  </li>
  <li>
    <strong>Técnicas Avançadas (60min)</strong>
    <ul class="ml-6 mt-1">
      <li>Flick shots e tracking de alvos em movimento</li>
      <li>Pre-aim e crosshair placement</li>
      <li>Peek timing e angle holding</li>
    </ul>
  </li>
  <li>
    <strong>Scrims Práticos (30min)</strong>
    <ul class="ml-6 mt-1">
      <li>Partidas curtas para aplicar o aprendizado</li>
      <li>Análise de erros em tempo real</li>
      <li>Dicas personalizadas para cada participante</li>
    </ul>
  </li>
</ol>

<h3 class="text-[#00f5ff] font-bold mb-2 mt-6">Requisitos:</h3>
<ul class="mb-4 space-y-1">
  <li>Nível intermediário ou superior</li>
  <li>Microfone para comunicação</li>
  <li>Conexão estável</li>
</ul>

<h3 class="text-[#00f5ff] font-bold mb-2 mt-6">O que você vai aprender:</h3>
<ul class="space-y-1">
  <li>Melhorar sua precisão em até 40%</li>
  <li>Técnicas de mira profissional</li>
  <li>Como treinar eficientemente sozinho</li>
</ul>`,
  imagem: "/assets/x0.jpeg",
  logo: "/assets/games/bloodstrike.png",
  valorInscricao: 25.00,
  valorPremiacao: 500.00,
  participantes: [
    { nome: "Player1", avatar: "/assets/x1.jpeg", rank: "Diamante" },
    { nome: "Player2", avatar: "/assets/x2.jpeg", rank: "Ouro" },
    { nome: "Player3", avatar: "/assets/x3.jpg", rank: "Platina" },
  ],
  tags: ["Aim", "Precisão", "Reflexos", "Competitivo"],
};
