"use client";

import Link from "next/link";
import Image from "next/image";
import { useState } from "react";

interface Participante {
  id: string;
  nome: string;
  email: string;
  avatar: string;
  rank?: string;
}

interface Equipe {
  id: string;
  nome: string;
  organizacao: {
    id: string;
    nome: string;
    logo: string;
  };
  participantes: Participante[];
}

export default function CriarXTreinoPage() {
  const [formData, setFormData] = useState({
    titulo: "",
    jogo: "",
    descricao: "",
    descricaoCompleta: "",
    dataInicio: "",
    dataEncerramento: "",
    horario: "",
    duracao: "",
    jogadoresPorEquipe: "",
    reservasPorEquipe: "",
    nivel: "Intermediário",
    valorInscricao: "",
    valorPremiacao: "",
    tags: "",
    imagem: "",
  });

  const [equipes, setEquipes] = useState<Equipe[]>([
    // Equipe mockada pré-cadastrada
    {
      id: "1",
      nome: "LOUD Valorant",
      organizacao: {
        id: "loud",
        nome: "LOUD",
        logo: "/assets/x0.jpeg",
      },
      participantes: [
        {
          id: "1",
          nome: "Robo",
          email: "robo@loud.com",
          avatar: "/assets/x0.jpeg",
          rank: "Pro Player",
        },
        {
          id: "2",
          nome: "Cauanzin",
          email: "cauanzin@loud.com",
          avatar: "/assets/x1.jpeg",
          rank: "Pro Player",
        },
        {
          id: "3",
          nome: "Tacolol",
          email: "tacolol@loud.com",
          avatar: "/assets/x2.jpeg",
          rank: "Pro Player",
        },
      ],
    },
  ]);
  const [novaEquipe, setNovaEquipe] = useState({
    nome: "",
    organizacaoId: "",
  });
  const [equipeSelecionada, setEquipeSelecionada] = useState<string>("");
  const [emailBusca, setEmailBusca] = useState("");
  const [resultadoBusca, setResultadoBusca] = useState<Participante | null>(
    null
  );
  const [buscando, setBuscando] = useState(false);

  // Organizações disponíveis (mockadas)
  const organizacoesDisponiveis = [
    {
      id: "loud",
      nome: "LOUD",
      logo: "/assets/x0.jpeg",
    },
    {
      id: "furia",
      nome: "FURIA",
      logo: "/assets/x1.jpeg",
    },
    {
      id: "pain",
      nome: "paiN Gaming",
      logo: "/assets/x2.jpeg",
    },
  ];

  const jogosDisponiveis = [
    { nome: "Free Fire", logo: "/assets/games/freefire.png" },
    { nome: "Bloodstrike", logo: "/assets/games/bloodstrike.png" },
    { nome: "Delta Force", logo: "/assets/games/deltaforce.png" },
  ];

  const handleChange = (
    e: React.ChangeEvent<
      HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement
    >
  ) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  // TODO: Buscar usuário por email via API
  const buscarPorEmail = async (email: string) => {
    if (!email) return;

    setBuscando(true);
    // Simulação de busca - será substituído por chamada à API
    setTimeout(() => {
      // Dados mockados - em produção viria da API
      const usuariosMockados: Record<string, Participante> = {
        "player1@example.com": {
          id: "10",
          nome: "Player1",
          email: "player1@example.com",
          avatar: "/assets/x1.jpeg",
          rank: "Diamante",
        },
        "player2@example.com": {
          id: "11",
          nome: "Player2",
          email: "player2@example.com",
          avatar: "/assets/x2.jpeg",
          rank: "Ouro",
        },
        "player3@example.com": {
          id: "12",
          nome: "Player3",
          email: "player3@example.com",
          avatar: "/assets/x3.jpg",
          rank: "Platina",
        },
        "less@loud.com": {
          id: "13",
          nome: "Less",
          email: "less@loud.com",
          avatar: "/assets/x3.jpg",
          rank: "Pro Player",
        },
        "aspas@loud.com": {
          id: "14",
          nome: "Aspas",
          email: "aspas@loud.com",
          avatar: "/assets/x0.jpeg",
          rank: "Pro Player",
        },
      };

      const usuario = usuariosMockados[email.toLowerCase()];
      if (usuario) {
        setResultadoBusca(usuario);
      } else {
        setResultadoBusca(null);
        alert("Usuário não encontrado com este email");
      }
      setBuscando(false);
    }, 500);
  };

  const criarEquipe = () => {
    if (!novaEquipe.nome || !novaEquipe.organizacaoId) {
      alert("Preencha o nome da equipe e selecione uma organização");
      return;
    }

    const organizacao = organizacoesDisponiveis.find(
      (org) => org.id === novaEquipe.organizacaoId
    );

    if (!organizacao) return;

    const novaEquipeObj: Equipe = {
      id: Date.now().toString(),
      nome: novaEquipe.nome,
      organizacao: organizacao,
      participantes: [],
    };

    setEquipes([...equipes, novaEquipeObj]);
    setNovaEquipe({ nome: "", organizacaoId: "" });
  };

  const removerEquipe = (id: string) => {
    setEquipes(equipes.filter((e) => e.id !== id));
    if (equipeSelecionada === id) {
      setEquipeSelecionada("");
    }
  };

  const adicionarParticipanteEquipe = (
    equipeId: string,
    participante: Participante
  ) => {
    setEquipes(
      equipes.map((equipe) => {
        if (equipe.id === equipeId) {
          // Verificar se já está na equipe
          if (equipe.participantes.some((p) => p.id === participante.id)) {
            alert("Este participante já está nesta equipe");
            return equipe;
          }
          return {
            ...equipe,
            participantes: [...equipe.participantes, participante],
          };
        }
        return equipe;
      })
    );
    setEmailBusca("");
    setResultadoBusca(null);
  };

  const removerParticipanteEquipe = (
    equipeId: string,
    participanteId: string
  ) => {
    setEquipes(
      equipes.map((equipe) => {
        if (equipe.id === equipeId) {
          return {
            ...equipe,
            participantes: equipe.participantes.filter(
              (p) => p.id !== participanteId
            ),
          };
        }
        return equipe;
      })
    );
  };


  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // TODO: Implementar envio para API
    console.log("Dados do formulário:", formData);
    console.log("Equipes:", equipes);
    alert("XTreino criado com sucesso! (Funcionalidade será implementada)");
  };

  return (
    <div className="min-h-screen bg-[#0a0a0a] text-[#ededed] overflow-x-hidden">
      {/* Navigation */}
      <nav className="fixed top-0 w-full z-50 bg-[#0a0a0a]/80 backdrop-blur-md border-b border-[#00f5ff]/20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center">
              <Link href="/" className="text-2xl font-bold text-[#00f5ff]">
                JoinLegends
              </Link>
            </div>
            <div className="hidden md:flex items-center space-x-8">
              <Link
                href="/#features"
                className="hover:text-[#00f5ff] transition-colors"
              >
                Features
              </Link>
              <Link
                href="/#comunidades"
                className="hover:text-[#00f5ff] transition-colors"
              >
                Comunidades
              </Link>
              <Link
                href="/xtreinos"
                className="hover:text-[#00f5ff] transition-colors"
              >
                XTreinos
              </Link>
              <button className="px-4 py-2 border border-[#00f5ff] text-[#00f5ff] hover:bg-[#00f5ff] hover:text-[#0a0a0a] transition-all">
                Entrar
              </button>
              <button className="px-4 py-2 bg-[#00f5ff] text-[#0a0a0a] font-semibold hover:shadow-[0_0_20px_rgba(0,245,255,0.5)] transition-all">
                Cadastrar
              </button>
            </div>
          </div>
        </div>
      </nav>

      {/* Hero Section */}
      <section className="relative pt-32 pb-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-4xl mx-auto">
          <Link
            href="/xtreinos"
            className="inline-flex items-center text-[#00f5ff] hover:text-[#00f5ff]/80 mb-6 transition-colors"
          >
            <svg
              className="w-5 h-5 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M15 19l-7-7 7-7"
              />
            </svg>
            Voltar para XTreinos
          </Link>

          <h1 className="text-4xl md:text-5xl font-bold mb-4">
            <span className="text-[#00f5ff]">Criar Novo XTreino</span>
          </h1>
          <p className="text-xl text-gray-400 mb-8">
            Preencha as informações abaixo para criar um novo XTreino
          </p>
        </div>
      </section>

      {/* Form Section */}
      <section className="px-4 sm:px-6 lg:px-8 pb-20">
        <div className="max-w-4xl mx-auto">
          <form onSubmit={handleSubmit} className="space-y-8">
            {/* Banner */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
              <h2 className="text-2xl font-bold text-[#00f5ff] mb-6">
                Banner do XTreino
              </h2>

              <div>
                <label
                  htmlFor="imagem"
                  className="block text-sm font-semibold text-gray-300 mb-2"
                >
                  Upload da Imagem *
                </label>
                <input
                  type="file"
                  id="imagem"
                  name="imagem"
                  accept="image/*"
                  onChange={(e) => {
                    const file = e.target.files?.[0];
                    if (file) {
                      // TODO: Implementar upload real para servidor
                      const reader = new FileReader();
                      reader.onloadend = () => {
                        setFormData((prev) => ({
                          ...prev,
                          imagem: reader.result as string,
                        }));
                      };
                      reader.readAsDataURL(file);
                    }
                  }}
                  className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-semibold file:bg-[#00f5ff] file:text-[#0a0a0a] file:cursor-pointer hover:file:bg-[#00f5ff]/90 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                />
                <p className="text-xs text-gray-500 mt-2">
                  Dimensões recomendadas: 1920x1080px (16:9). Formatos aceitos:
                  JPG, PNG, WebP. Tamanho máximo: 5MB.
                </p>
                {formData.imagem && (
                  <div className="mt-4 relative w-full h-48 rounded-lg overflow-hidden border border-[#00f5ff]/30">
                    <Image
                      src={formData.imagem}
                      alt="Preview do banner"
                      fill
                      className="object-cover"
                    />
                  </div>
                )}
              </div>
            </div>

            {/* Informações Básicas */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
              <h2 className="text-2xl font-bold text-[#00f5ff] mb-6">
                Informações Básicas
              </h2>

              <div className="space-y-6">
                {/* Título */}
                <div>
                  <label
                    htmlFor="titulo"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Título do XTreino *
                  </label>
                  <input
                    type="text"
                    id="titulo"
                    name="titulo"
                    value={formData.titulo}
                    onChange={handleChange}
                    required
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                    placeholder="Ex: Treino de Aim - Bloodstrike"
                  />
                </div>

                {/* Jogo */}
                <div>
                  <label
                    htmlFor="jogo"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Jogo *
                  </label>
                  <select
                    id="jogo"
                    name="jogo"
                    value={formData.jogo}
                    onChange={handleChange}
                    required
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                  >
                    <option value="">Selecione um jogo</option>
                    {jogosDisponiveis.map((jogo, index) => (
                      <option key={index} value={jogo.nome}>
                        {jogo.nome}
                      </option>
                    ))}
                  </select>
                </div>

                {/* Descrição Breve */}
                <div>
                  <label
                    htmlFor="descricao"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Descrição Breve *
                  </label>
                  <textarea
                    id="descricao"
                    name="descricao"
                    value={formData.descricao}
                    onChange={handleChange}
                    required
                    rows={3}
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all resize-none"
                    placeholder="Uma descrição curta que aparecerá nos cards de listagem..."
                  />
                </div>

                {/* Descrição Completa */}
                <div>
                  <label
                    htmlFor="descricaoCompleta"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Descrição Completa (HTML) *
                  </label>
                  <textarea
                    id="descricaoCompleta"
                    name="descricaoCompleta"
                    value={formData.descricaoCompleta}
                    onChange={handleChange}
                    required
                    rows={10}
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all resize-none font-mono text-sm"
                    placeholder="Use HTML para formatar a descrição completa. Ex: &lt;p&gt;Texto aqui&lt;/p&gt;"
                  />
                  <p className="text-xs text-gray-500 mt-2">
                    Você pode usar tags HTML como &lt;p&gt;, &lt;h3&gt;,
                    &lt;ul&gt;, &lt;li&gt;, &lt;strong&gt;, etc.
                  </p>
                </div>

                {/* Tags */}
                <div>
                  <label
                    htmlFor="tags"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Tags (separadas por vírgula)
                  </label>
                  <input
                    type="text"
                    id="tags"
                    name="tags"
                    value={formData.tags}
                    onChange={handleChange}
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                    placeholder="Ex: Aim, Precisão, Reflexos, Competitivo"
                  />
                </div>
              </div>
            </div>

            {/* Data e Horário */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
              <h2 className="text-2xl font-bold text-[#00f5ff] mb-6">
                Data e Horário
              </h2>

              <div className="grid md:grid-cols-2 gap-6">
                {/* Data de Início */}
                <div>
                  <label
                    htmlFor="dataInicio"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Data de Início *
                  </label>
                  <input
                    type="date"
                    id="dataInicio"
                    name="dataInicio"
                    value={formData.dataInicio}
                    onChange={handleChange}
                    required
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                  />
                </div>

                {/* Data de Encerramento */}
                <div>
                  <label
                    htmlFor="dataEncerramento"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Data de Encerramento *
                  </label>
                  <input
                    type="date"
                    id="dataEncerramento"
                    name="dataEncerramento"
                    value={formData.dataEncerramento}
                    onChange={handleChange}
                    required
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                  />
                </div>

                {/* Horário */}
                <div>
                  <label
                    htmlFor="horario"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Horário *
                  </label>
                  <input
                    type="time"
                    id="horario"
                    name="horario"
                    value={formData.horario}
                    onChange={handleChange}
                    required
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                  />
                </div>

                {/* Duração */}
                <div>
                  <label
                    htmlFor="duracao"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Duração *
                  </label>
                  <input
                    type="text"
                    id="duracao"
                    name="duracao"
                    value={formData.duracao}
                    onChange={handleChange}
                    required
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                    placeholder="Ex: 2 horas"
                  />
                </div>
              </div>
            </div>

            {/* Configuração de Equipes e Vagas */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
              <h2 className="text-2xl font-bold text-[#00f5ff] mb-6">
                Configuração de Equipes e Vagas
              </h2>

              <div className="grid md:grid-cols-3 gap-6 mb-6">
                {/* Jogadores por Equipe */}
                <div>
                  <label
                    htmlFor="jogadoresPorEquipe"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Jogadores por Equipe *
                  </label>
                  <input
                    type="number"
                    id="jogadoresPorEquipe"
                    name="jogadoresPorEquipe"
                    value={formData.jogadoresPorEquipe}
                    onChange={handleChange}
                    required
                    min="1"
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                    placeholder="Ex: 5"
                  />
                </div>

                {/* Reservas por Equipe */}
                <div>
                  <label
                    htmlFor="reservasPorEquipe"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Reservas por Equipe *
                  </label>
                  <input
                    type="number"
                    id="reservasPorEquipe"
                    name="reservasPorEquipe"
                    value={formData.reservasPorEquipe}
                    onChange={handleChange}
                    required
                    min="0"
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                    placeholder="Ex: 2"
                  />
                </div>

                {/* Nível */}
                <div>
                  <label
                    htmlFor="nivel"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Nível *
                  </label>
                  <select
                    id="nivel"
                    name="nivel"
                    value={formData.nivel}
                    onChange={handleChange}
                    required
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                  >
                    <option value="Iniciante">Iniciante</option>
                    <option value="Intermediário">Intermediário</option>
                    <option value="Avançado">Avançado</option>
                  </select>
                </div>
              </div>

              {/* Cálculo de Vagas */}
              {formData.jogadoresPorEquipe && formData.reservasPorEquipe && (
                <div className="p-4 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20">
                  <div className="flex items-center justify-between">
                    <div>
                      <p className="text-sm text-gray-400 mb-1">
                        Total de Equipes Cadastradas
                      </p>
                      <p className="text-2xl font-bold text-[#00f5ff]">
                        {equipes.length}
                      </p>
                    </div>
                    <div className="text-center">
                      <p className="text-sm text-gray-400 mb-1">
                        Vagas por Equipe
                      </p>
                      <p className="text-2xl font-bold text-[#00f5ff]">
                        {Number(formData.jogadoresPorEquipe) +
                          Number(formData.reservasPorEquipe)}
                      </p>
                      <p className="text-xs text-gray-500 mt-1">
                        ({formData.jogadoresPorEquipe} titulares +{" "}
                        {formData.reservasPorEquipe} reservas)
                      </p>
                    </div>
                    <div className="text-right">
                      <p className="text-sm text-gray-400 mb-1">
                        Total de Vagas
                      </p>
                      <p className="text-2xl font-bold text-[#00ff41]">
                        {equipes.length *
                          (Number(formData.jogadoresPorEquipe) +
                            Number(formData.reservasPorEquipe))}
                      </p>
                      <p className="text-xs text-gray-500 mt-1">
                        {equipes.length} equipe(s) ×{" "}
                        {Number(formData.jogadoresPorEquipe) +
                          Number(formData.reservasPorEquipe)}{" "}
                        vagas
                      </p>
                    </div>
                  </div>
                </div>
              )}
            </div>

            {/* Valores */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
              <h2 className="text-2xl font-bold text-[#00f5ff] mb-6">
                Valores
              </h2>

              <div className="grid md:grid-cols-2 gap-6">
                {/* Valor da Inscrição */}
                <div>
                  <label
                    htmlFor="valorInscricao"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Valor da Inscrição (R$) *
                  </label>
                  <input
                    type="number"
                    id="valorInscricao"
                    name="valorInscricao"
                    value={formData.valorInscricao}
                    onChange={handleChange}
                    required
                    min="0"
                    step="0.01"
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                    placeholder="0.00"
                  />
                </div>

                {/* Valor da Premiação */}
                <div>
                  <label
                    htmlFor="valorPremiacao"
                    className="block text-sm font-semibold text-gray-300 mb-2"
                  >
                    Valor da Premiação Total (R$) *
                  </label>
                  <input
                    type="number"
                    id="valorPremiacao"
                    name="valorPremiacao"
                    value={formData.valorPremiacao}
                    onChange={handleChange}
                    required
                    min="0"
                    step="0.01"
                    className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                    placeholder="0.00"
                  />
                </div>
              </div>
            </div>

            {/* Participantes e Equipes */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
              <h2 className="text-2xl font-bold text-[#00f5ff] mb-6">
                Participantes
              </h2>
              
              {/* Seção de Equipes */}
              <div className="mb-8">
                <h3 className="text-xl font-semibold text-[#00f5ff] mb-4">
                  Equipes
                </h3>

              {/* Criar Nova Equipe */}
              <div className="mb-6 p-4 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20">
                <h3 className="text-lg font-semibold text-gray-300 mb-4">
                  Criar Nova Equipe
                </h3>
                <div className="grid md:grid-cols-2 gap-4 mb-4">
                  <div>
                    <label
                      htmlFor="nomeEquipe"
                      className="block text-sm font-semibold text-gray-300 mb-2"
                    >
                      Nome da Equipe *
                    </label>
                    <input
                      type="text"
                      id="nomeEquipe"
                      value={novaEquipe.nome}
                      onChange={(e) =>
                        setNovaEquipe({ ...novaEquipe, nome: e.target.value })
                      }
                      className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                      placeholder="Ex: LOUD Valorant"
                    />
                  </div>
                  <div>
                    <label
                      htmlFor="organizacaoEquipe"
                      className="block text-sm font-semibold text-gray-300 mb-2"
                    >
                      Organização *
                    </label>
                    <select
                      id="organizacaoEquipe"
                      value={novaEquipe.organizacaoId}
                      onChange={(e) =>
                        setNovaEquipe({
                          ...novaEquipe,
                          organizacaoId: e.target.value,
                        })
                      }
                      className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                    >
                      <option value="">Selecione uma organização</option>
                      {organizacoesDisponiveis.map((org) => (
                        <option key={org.id} value={org.id}>
                          {org.nome}
                        </option>
                      ))}
                    </select>
                  </div>
                </div>
                <button
                  type="button"
                  onClick={criarEquipe}
                  className="px-4 py-2 bg-[#00f5ff] text-[#0a0a0a] font-semibold rounded-lg hover:shadow-[0_0_20px_rgba(0,245,255,0.5)] transition-all"
                >
                  Criar Equipe
                </button>
              </div>

              {/* Lista de Equipes */}
              {equipes.length > 0 && (
                <div className="space-y-4 mb-6">
                  {equipes.map((equipe) => (
                    <div
                      key={equipe.id}
                      className="p-4 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20"
                    >
                      <div className="flex items-start justify-between mb-4">
                        <div className="flex items-center gap-3">
                          <div className="w-12 h-12 rounded-lg overflow-hidden border-2 border-[#00f5ff]/30">
                            <Image
                              src={equipe.organizacao.logo}
                              alt={equipe.organizacao.nome}
                              width={48}
                              height={48}
                              className="object-cover w-full h-full"
                            />
                          </div>
                          <div>
                            <h3 className="text-lg font-semibold text-white">
                              {equipe.nome}
                            </h3>
                            <p className="text-sm text-gray-400">
                              {equipe.organizacao.nome}
                            </p>
                            <p className="text-xs text-[#00f5ff] mt-1">
                              {equipe.participantes.length} participante(s)
                            </p>
                          </div>
                        </div>
                        <button
                          type="button"
                          onClick={() => removerEquipe(equipe.id)}
                          className="p-2 text-[#ff00ff] hover:bg-[#ff00ff]/10 rounded-lg transition-all"
                        >
                          <svg
                            className="w-5 h-5"
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                          >
                            <path
                              strokeLinecap="round"
                              strokeLinejoin="round"
                              strokeWidth={2}
                              d="M6 18L18 6M6 6l12 12"
                            />
                          </svg>
                        </button>
                      </div>

                      {/* Adicionar Participante à Equipe */}
                      <div className="mb-4">
                        <label className="block text-sm font-semibold text-gray-300 mb-2">
                          Adicionar Participante à Equipe
                        </label>
                        <div className="flex gap-2">
                          <input
                            type="email"
                            value={
                              equipeSelecionada === equipe.id ? emailBusca : ""
                            }
                            onChange={(e) => {
                              setEmailBusca(e.target.value);
                              setEquipeSelecionada(equipe.id);
                            }}
                            onKeyPress={(e) => {
                              if (
                                e.key === "Enter" &&
                                equipeSelecionada === equipe.id
                              ) {
                                e.preventDefault();
                                buscarPorEmail(emailBusca);
                              }
                            }}
                            className="flex-1 px-4 py-2 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all text-sm"
                            placeholder="Buscar participante por email..."
                            onClick={() => setEquipeSelecionada(equipe.id)}
                          />
                          <button
                            type="button"
                            onClick={() => {
                              setEquipeSelecionada(equipe.id);
                              buscarPorEmail(emailBusca);
                            }}
                            disabled={buscando || !emailBusca}
                            className="px-4 py-2 bg-[#00f5ff] text-[#0a0a0a] font-semibold rounded-lg hover:shadow-[0_0_20px_rgba(0,245,255,0.5)] transition-all disabled:opacity-50 disabled:cursor-not-allowed text-sm"
                          >
                            {buscando && equipeSelecionada === equipe.id
                              ? "Buscando..."
                              : "Buscar"}
                          </button>
                        </div>

                        {/* Resultado da Busca para esta equipe */}
                        {resultadoBusca && equipeSelecionada === equipe.id && (
                          <div className="mt-3 p-3 bg-[#0a0a0a]/70 rounded-lg border border-[#00f5ff]/30">
                            <div className="flex items-center gap-3">
                              <div className="w-10 h-10 rounded-full overflow-hidden border-2 border-[#00f5ff]/30">
                                <Image
                                  src={resultadoBusca.avatar}
                                  alt={resultadoBusca.nome}
                                  width={40}
                                  height={40}
                                  className="object-cover w-full h-full"
                                />
                              </div>
                              <div className="flex-1">
                                <h4 className="text-sm font-semibold text-white">
                                  {resultadoBusca.nome}
                                </h4>
                                <p className="text-xs text-gray-400">
                                  {resultadoBusca.email}
                                </p>
                                {resultadoBusca.rank && (
                                  <p className="text-xs text-[#00f5ff] mt-1">
                                    {resultadoBusca.rank}
                                  </p>
                                )}
                              </div>
                              <button
                                type="button"
                                onClick={() => {
                                  adicionarParticipanteEquipe(
                                    equipe.id,
                                    resultadoBusca
                                  );
                                }}
                                className="px-3 py-1.5 bg-[#00ff41] text-[#0a0a0a] font-semibold rounded-lg hover:shadow-[0_0_20px_rgba(0,255,65,0.5)] transition-all text-sm"
                              >
                                Adicionar
                              </button>
                            </div>
                          </div>
                        )}
                      </div>

                      {/* Participantes da Equipe */}
                      {equipe.participantes.length > 0 && (
                        <div>
                          <h4 className="text-sm font-semibold text-gray-300 mb-2">
                            Participantes da Equipe
                          </h4>
                          <div className="flex gap-3 overflow-x-auto pb-2">
                            {equipe.participantes.map((participante) => (
                              <div
                                key={participante.id}
                                className="group relative flex-shrink-0 w-32 p-3 bg-[#0a0a0a]/70 rounded-lg border border-[#00f5ff]/20 hover:border-[#00f5ff] transition-all"
                              >
                                <button
                                  type="button"
                                  onClick={() =>
                                    removerParticipanteEquipe(
                                      equipe.id,
                                      participante.id
                                    )
                                  }
                                  className="absolute top-1 right-1 p-1 text-[#ff00ff] hover:bg-[#ff00ff]/20 rounded-full transition-all opacity-0 group-hover:opacity-100"
                                >
                                  <svg
                                    className="w-3 h-3"
                                    fill="none"
                                    stroke="currentColor"
                                    viewBox="0 0 24 24"
                                  >
                                    <path
                                      strokeLinecap="round"
                                      strokeLinejoin="round"
                                      strokeWidth={2}
                                      d="M6 18L18 6M6 6l12 12"
                                    />
                                  </svg>
                                </button>
                                <div className="flex flex-col items-center text-center">
                                  <div className="w-12 h-12 mb-2 rounded-full overflow-hidden border-2 border-[#00f5ff]/30">
                                    <Image
                                      src={participante.avatar}
                                      alt={participante.nome}
                                      width={48}
                                      height={48}
                                      className="object-cover w-full h-full"
                                    />
                                  </div>
                                  <p className="text-xs font-semibold text-white truncate w-full">
                                    {participante.nome}
                                  </p>
                                  {participante.rank && (
                                    <p className="text-xs text-[#00f5ff]">
                                      {participante.rank}
                                    </p>
                                  )}
                                </div>
                              </div>
                            ))}
                          </div>
                        </div>
                      )}
                    </div>
                  ))}
                </div>
              )}

              {equipes.length === 0 && (
                <p className="text-sm text-gray-500 text-center py-4">
                  Nenhuma equipe cadastrada. Crie uma equipe acima para começar.
                </p>
              )}
              </div>
            </div>

            {/* Botões de Ação */}
            <div className="flex flex-col sm:flex-row gap-4 justify-end">
              <Link
                href="/xtreinos"
                className="px-6 py-3 border-2 border-[#00f5ff] text-[#00f5ff] font-semibold rounded-lg hover:bg-[#00f5ff]/10 transition-all text-center"
              >
                Cancelar
              </Link>
              <button
                type="submit"
                className="px-6 py-3 bg-[#00f5ff] text-[#0a0a0a] font-bold rounded-lg hover:shadow-[0_0_30px_rgba(0,245,255,0.6)] transition-all transform hover:scale-105"
              >
                Criar XTreino
              </button>
            </div>
          </form>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-[#00f5ff]/20 bg-[#1a1a2e]/30 py-12 px-4 sm:px-6 lg:px-8 mt-20">
        <div className="max-w-7xl mx-auto">
          <div className="grid md:grid-cols-4 gap-8">
            <div>
              <h3 className="text-2xl font-bold text-[#00f5ff] mb-4">
                JoinLegends
              </h3>
              <p className="text-gray-400 text-sm">
                A plataforma definitiva para jogadores que buscam excelência.
              </p>
            </div>
            <div>
              <h4 className="text-[#00f5ff] font-semibold mb-4">Produto</h4>
              <ul className="space-y-2 text-gray-400 text-sm">
                <li>
                  <Link
                    href="/#features"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    Features
                  </Link>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    Preços
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    Roadmap
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="text-[#00f5ff] font-semibold mb-4">Comunidade</h4>
              <ul className="space-y-2 text-gray-400 text-sm">
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    Discord
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    Twitter
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    GitHub
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="text-[#00f5ff] font-semibold mb-4">Suporte</h4>
              <ul className="space-y-2 text-gray-400 text-sm">
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    Documentação
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    Contato
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    FAQ
                  </a>
                </li>
              </ul>
            </div>
          </div>
          <div className="mt-8 pt-8 border-t border-[#00f5ff]/20 text-center text-gray-400 text-sm">
            <p>© 2024 JoinLegends. Todos os direitos reservados.</p>
          </div>
        </div>
      </footer>
    </div>
  );
}
