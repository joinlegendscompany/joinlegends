import Link from "next/link";
import Image from "next/image";

// TODO: Buscar dados da organização via API
// const organizacao = await fetchOrganizacao(params.id);

export default function OrganizacaoPage({
  params,
}: {
  params: { id: string };
}) {
  // Dados temporários - será substituído por dados da API
  const organizacao = {
    id: params.id,
    nome: "LOUD",
    tagline: "A maior organização de esports do Brasil",
    banner: "/assets/organization/loud.jpg",
    logo: "/assets/x0.jpeg",
    descricao: `<p class="mb-4">A LOUD é uma das principais organizações brasileiras de esportes eletrônicos, fundada em 2019. Com equipes competindo em diversos jogos, a LOUD rapidamente se destacou no cenário global, tornando-se a organização de esports com maior número de seguidores nas redes sociais no Brasil e a segunda maior do mundo.</p>

<p class="mb-4">A organização conquistou títulos importantes em diversas modalidades, incluindo o Free Fire World Series em 2021, o Valorant Champions 2022, e múltiplos títulos da Liga Brasileira de Free Fire. A LOUD foi eleita a melhor organização no Prêmio Esports Brasil 2023, refletindo seu impacto significativo no cenário de esports.</p>

<h3 class="text-[#00f5ff] font-bold mb-2 mt-6">História e Conquistas</h3>
<p class="mb-4">Com uma trajetória marcada por conquistas expressivas e uma base de fãs engajada, a LOUD continua a ser uma força dominante nos esportes eletrônicos, representando o Brasil em competições internacionais e ampliando sua presença no cenário global.</p>`,
    totalJogadores: 45,
    jogos: [
      { nome: "Free Fire", logo: "/assets/games/freefire.png" },
      { nome: "Valorant", logo: "/assets/games/bloodstrike.png" },
      { nome: "League of Legends", logo: "/assets/games/deltaforce.png" },
      { nome: "Fortnite", logo: "/assets/games/freefire.png" },
    ],
    estatisticas: {
      xtreinosCriados: 128,
      xtreinosJogados: 342,
      xtreinosGanhos: 215,
      xtreinosPerdidos: 127,
      winRate: 62.9,
    },
    redesSociais: {
      youtube: "https://youtube.com/@LOUD",
      twitch: "https://twitch.tv/loud",
      tiktok: "https://tiktok.com/@loud",
      instagram: "https://instagram.com/loudgg",
      twitter: "https://twitter.com/LOUDgg",
    },
    participantes: [
      {
        nome: "Robo",
        nomeCompleto: "Leonardo Souza",
        avatar: "/assets/x0.jpeg",
        jogo: "League of Legends",
        role: "Top Laner",
        rank: "Pro Player",
        descricao: "7x Campeão do CBLOL",
      },
      {
        nome: "Cauanzin",
        nomeCompleto: "Cauan Pereira",
        avatar: "/assets/x1.jpeg",
        jogo: "Free Fire",
        role: "IGL",
        rank: "Pro Player",
        descricao: "Campeão Mundial Free Fire",
      },
      {
        nome: "Tacolol",
        nomeCompleto: "Gabriel Lima",
        avatar: "/assets/x2.jpeg",
        jogo: "Valorant",
        role: "Duelist",
        rank: "Pro Player",
        descricao: "Campeão Valorant Champions 2022",
      },
      {
        nome: "Less",
        nomeCompleto: "Felipe Basso",
        avatar: "/assets/x3.jpg",
        jogo: "Valorant",
        role: "Initiator",
        rank: "Pro Player",
        descricao: "Campeão Valorant Champions 2022",
      },
      {
        nome: "Saadhak",
        nomeCompleto: "Matias Delipetro",
        avatar: "/assets/x0.jpeg",
        jogo: "Valorant",
        role: "IGL",
        rank: "Pro Player",
        descricao: "Líder e Campeão Mundial",
      },
      {
        nome: "Aspas",
        nomeCompleto: "Erick Santos",
        avatar: "/assets/x1.jpeg",
        jogo: "Valorant",
        role: "Duelist",
        rank: "Pro Player",
        descricao: "MVP Valorant Champions 2022",
      },
      {
        nome: "Tuyz",
        nomeCompleto: "Arthur Vieira",
        avatar: "/assets/x2.jpeg",
        jogo: "Valorant",
        role: "Controller",
        rank: "Pro Player",
        descricao: "Campeão Valorant Champions 2022",
      },
      {
        nome: "Croc",
        nomeCompleto: "Park Yong-ho",
        avatar: "/assets/x3.jpg",
        jogo: "League of Legends",
        role: "Jungler",
        rank: "Pro Player",
        descricao: "4x Campeão CBLOL",
      },
      {
        nome: "Tinowns",
        nomeCompleto: "Thiago Sartori",
        avatar: "/assets/x0.jpeg",
        jogo: "League of Legends",
        role: "Mid Laner",
        rank: "Pro Player",
        descricao: "4x Campeão CBLOL",
      },
      {
        nome: "Route",
        nomeCompleto: "Felipe Tadeu",
        avatar: "/assets/x1.jpeg",
        jogo: "Free Fire",
        role: "Support",
        rank: "Pro Player",
        descricao: "Campeão Mundial Free Fire",
      },
      {
        nome: "Luccas",
        nomeCompleto: "Luccas Lima",
        avatar: "/assets/x2.jpeg",
        jogo: "Free Fire",
        role: "Fragger",
        rank: "Pro Player",
        descricao: "Campeão Mundial Free Fire",
      },
      {
        nome: "Voltax",
        nomeCompleto: "Gabriel Voltan",
        avatar: "/assets/x3.jpg",
        jogo: "Free Fire",
        role: "Support",
        rank: "Pro Player",
        descricao: "Campeão Mundial Free Fire",
      },
    ],
  };

  const winRateColor =
    organizacao.estatisticas.winRate >= 60
      ? "text-[#00ff41]"
      : organizacao.estatisticas.winRate >= 50
      ? "text-[#00f5ff]"
      : "text-[#ff00ff]";

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

      {/* Banner Hero */}
      <section className="relative pt-16">
        <div className="relative w-full h-64 md:h-96 overflow-hidden">
          <Image
            src={organizacao.banner}
            alt={organizacao.nome}
            fill
            className="object-cover"
          />
          <div className="absolute inset-0 bg-gradient-to-t from-[#0a0a0a] via-[#0a0a0a]/70 to-transparent"></div>

          {/* Logo e Nome da Organização */}
          <div className="absolute bottom-8 left-8 right-8">
            <div className="flex items-end gap-6">
              <div className="w-24 h-24 md:w-32 md:h-32 rounded-lg overflow-hidden border-4 border-[#00f5ff] bg-[#0a0a0a] flex-shrink-0">
                <Image
                  src={organizacao.logo}
                  alt={organizacao.nome}
                  width={128}
                  height={128}
                  className="object-cover w-full h-full"
                />
              </div>
              <div className="flex-1">
                <h1 className="text-4xl md:text-5xl font-bold text-white mb-2">
                  {organizacao.nome}
                </h1>
                <p className="text-lg md:text-xl text-gray-300">
                  {organizacao.tagline}
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Estatísticas Principais */}
      <section className="py-12 px-4 sm:px-6 lg:px-8 bg-[#1a1a2e]/30">
        <div className="max-w-7xl mx-auto">
          <div className="grid grid-cols-2 md:grid-cols-4 gap-4 md:gap-6">
            {/* Total de Jogadores */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6 text-center hover:border-[#00f5ff] hover:shadow-[0_0_20px_rgba(0,245,255,0.3)] transition-all">
              <div className="text-3xl md:text-4xl font-bold text-[#00f5ff] mb-2">
                {organizacao.totalJogadores}
              </div>
              <div className="text-sm text-gray-400">Jogadores</div>
            </div>

            {/* XTreinos Criados */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6 text-center hover:border-[#00f5ff] hover:shadow-[0_0_20px_rgba(0,245,255,0.3)] transition-all">
              <div className="text-3xl md:text-4xl font-bold text-[#00f5ff] mb-2">
                {organizacao.estatisticas.xtreinosCriados}
              </div>
              <div className="text-sm text-gray-400">XTreinos Criados</div>
            </div>

            {/* XTreinos Jogados */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6 text-center hover:border-[#00f5ff] hover:shadow-[0_0_20px_rgba(0,245,255,0.3)] transition-all">
              <div className="text-3xl md:text-4xl font-bold text-[#00f5ff] mb-2">
                {organizacao.estatisticas.xtreinosJogados}
              </div>
              <div className="text-sm text-gray-400">XTreinos Jogados</div>
            </div>

            {/* Win Rate */}
            <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6 text-center hover:border-[#00f5ff] hover:shadow-[0_0_20px_rgba(0,245,255,0.3)] transition-all">
              <div
                className={`text-3xl md:text-4xl font-bold mb-2 ${winRateColor}`}
              >
                {organizacao.estatisticas.winRate}%
              </div>
              <div className="text-sm text-gray-400">Win Rate</div>
            </div>
          </div>
        </div>
      </section>

      {/* Conteúdo Principal */}
      <section className="py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-7xl mx-auto">
          <div className="grid lg:grid-cols-3 gap-8">
            {/* Conteúdo Principal */}
            <div className="lg:col-span-2 space-y-8">
              {/* Descrição */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
                <h2 className="text-2xl font-bold text-[#00f5ff] mb-4">
                  Sobre a Organização
                </h2>
                <div
                  className="wysiwyg-content"
                  dangerouslySetInnerHTML={{ __html: organizacao.descricao }}
                />
              </div>

              {/* Jogos */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
                <h2 className="text-2xl font-bold text-[#00f5ff] mb-4">
                  Jogos que Participamos
                </h2>
                <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
                  {organizacao.jogos.map((jogo, index) => (
                    <div
                      key={index}
                      className="flex flex-col items-center p-4 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20 hover:border-[#00f5ff] transition-all"
                    >
                      <div className="w-16 h-16 mb-3 relative">
                        <Image
                          src={jogo.logo}
                          alt={jogo.nome}
                          fill
                          className="object-contain"
                        />
                      </div>
                      <span className="text-sm text-gray-300 text-center">
                        {jogo.nome}
                      </span>
                    </div>
                  ))}
                </div>
              </div>

              {/* Estatísticas Detalhadas de XTreinos */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
                <h2 className="text-2xl font-bold text-[#00f5ff] mb-4">
                  Estatísticas de XTreinos
                </h2>
                <div className="grid md:grid-cols-2 gap-6">
                  <div className="p-4 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20">
                    <div className="flex items-center justify-between mb-2">
                      <span className="text-gray-400">XTreinos Ganhos</span>
                      <span className="text-2xl font-bold text-[#00ff41]">
                        {organizacao.estatisticas.xtreinosGanhos}
                      </span>
                    </div>
                    <div className="w-full bg-[#0a0a0a] rounded-full h-2">
                      <div
                        className="bg-[#00ff41] h-2 rounded-full"
                        style={{
                          width: `${
                            (organizacao.estatisticas.xtreinosGanhos /
                              organizacao.estatisticas.xtreinosJogados) *
                            100
                          }%`,
                        }}
                      ></div>
                    </div>
                  </div>

                  <div className="p-4 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20">
                    <div className="flex items-center justify-between mb-2">
                      <span className="text-gray-400">XTreinos Perdidos</span>
                      <span className="text-2xl font-bold text-[#ff00ff]">
                        {organizacao.estatisticas.xtreinosPerdidos}
                      </span>
                    </div>
                    <div className="w-full bg-[#0a0a0a] rounded-full h-2">
                      <div
                        className="bg-[#ff00ff] h-2 rounded-full"
                        style={{
                          width: `${
                            (organizacao.estatisticas.xtreinosPerdidos /
                              organizacao.estatisticas.xtreinosJogados) *
                            100
                          }%`,
                        }}
                      ></div>
                    </div>
                  </div>
                </div>
              </div>

              {/* Participantes */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
                <h2 className="text-2xl font-bold text-[#00f5ff] mb-4">
                  Membros da Organização ({organizacao.participantes.length})
                </h2>
                <div className="flex gap-4 overflow-x-auto pb-4 scrollbar-thin scrollbar-thumb-[#00f5ff] scrollbar-track-[#0a0a0a]">
                  {organizacao.participantes.map((participante, index) => (
                    <div
                      key={index}
                      className="group relative flex-shrink-0 w-48 p-4 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20 hover:border-[#00f5ff] hover:shadow-[0_0_20px_rgba(0,245,255,0.3)] transition-all"
                    >
                      <div className="flex flex-col items-center text-center">
                        <div className="relative w-20 h-20 mb-3 rounded-full overflow-hidden border-2 border-[#00f5ff]/30 group-hover:border-[#00f5ff] transition-all">
                          <Image
                            src={participante.avatar}
                            alt={participante.nome}
                            fill
                            className="object-cover group-hover:scale-110 transition-transform"
                          />
                        </div>
                        <h3 className="text-lg font-bold text-white mb-1">
                          {participante.nome}
                        </h3>
                        <p className="text-xs text-gray-400 mb-2">
                          {participante.nomeCompleto}
                        </p>
                        <div className="flex items-center gap-2 mb-2">
                          <span className="px-2 py-1 text-xs bg-[#00f5ff]/20 text-[#00f5ff] rounded border border-[#00f5ff]/30">
                            {participante.jogo}
                          </span>
                        </div>
                        <p className="text-xs text-[#00f5ff] font-semibold mb-1">
                          {participante.role}
                        </p>
                        <p className="text-xs text-gray-500 italic">
                          {participante.descricao}
                        </p>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            </div>

            {/* Sidebar */}
            <div className="lg:col-span-1 space-y-6">
              {/* Redes Sociais */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6 sticky top-24">
                <h3 className="text-xl font-bold text-[#00f5ff] mb-4">
                  Redes Sociais
                </h3>
                <div className="space-y-3">
                  {organizacao.redesSociais.youtube && (
                    <a
                      href={organizacao.redesSociais.youtube}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center gap-3 p-3 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20 hover:border-[#00f5ff] hover:bg-[#00f5ff]/10 transition-all group"
                    >
                      <svg
                        className="w-6 h-6 text-[#ff0000] group-hover:scale-110 transition-transform"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path d="M23.498 6.186a3.016 3.016 0 0 0-2.122-2.136C19.505 3.545 12 3.545 12 3.545s-7.505 0-9.377.505A3.017 3.017 0 0 0 .502 6.186C0 8.07 0 12 0 12s0 3.93.502 5.814a3.016 3.016 0 0 0 2.122 2.136c1.871.505 9.376.505 9.376.505s7.505 0 9.377-.505a3.015 3.015 0 0 0 2.122-2.136C24 15.93 24 12 24 12s0-3.93-.502-5.814zM9.545 15.568V8.432L15.818 12l-6.273 3.568z" />
                      </svg>
                      <span className="text-gray-300 group-hover:text-white">
                        YouTube
                      </span>
                    </a>
                  )}

                  {organizacao.redesSociais.twitch && (
                    <a
                      href={organizacao.redesSociais.twitch}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center gap-3 p-3 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20 hover:border-[#00f5ff] hover:bg-[#00f5ff]/10 transition-all group"
                    >
                      <svg
                        className="w-6 h-6 text-[#9146ff] group-hover:scale-110 transition-transform"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path d="M11.571 4.714h1.715v5.143H11.57zm4.715 0H18v5.143h-1.714zM6 0L1.714 4.286v15.428h5.143V24l4.286-4.286h3.428L22.286 12V0zm14.571 11.143l-3.428 3.428h-3.429l-3 3v-3H6.857V1.714h13.714Z" />
                      </svg>
                      <span className="text-gray-300 group-hover:text-white">
                        Twitch
                      </span>
                    </a>
                  )}

                  {organizacao.redesSociais.tiktok && (
                    <a
                      href={organizacao.redesSociais.tiktok}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center gap-3 p-3 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20 hover:border-[#00f5ff] hover:bg-[#00f5ff]/10 transition-all group"
                    >
                      <svg
                        className="w-6 h-6 text-white group-hover:scale-110 transition-transform"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path d="M19.59 6.69a4.83 4.83 0 0 1-3.77-4.25V2h-3.45v13.67a2.89 2.89 0 0 1-5.2 1.74 2.89 2.89 0 0 1 2.31-4.64 2.93 2.93 0 0 1 .88.13V9.4a6.84 6.84 0 0 0-1-.05A6.33 6.33 0 0 0 5 20.1a6.34 6.34 0 0 0 10.86-4.43v-7a8.16 8.16 0 0 0 4.77 1.52v-3.4a4.85 4.85 0 0 1-1-.1z" />
                      </svg>
                      <span className="text-gray-300 group-hover:text-white">
                        TikTok
                      </span>
                    </a>
                  )}

                  {organizacao.redesSociais.instagram && (
                    <a
                      href={organizacao.redesSociais.instagram}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center gap-3 p-3 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20 hover:border-[#00f5ff] hover:bg-[#00f5ff]/10 transition-all group"
                    >
                      <svg
                        className="w-6 h-6 text-[#E4405F] group-hover:scale-110 transition-transform"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path d="M12 2.163c3.204 0 3.584.012 4.85.07 3.252.148 4.771 1.691 4.919 4.919.058 1.265.069 1.645.069 4.849 0 3.205-.012 3.584-.069 4.849-.149 3.225-1.664 4.771-4.919 4.919-1.266.058-1.644.07-4.85.07-3.204 0-3.584-.012-4.849-.07-3.26-.149-4.771-1.699-4.919-4.92-.058-1.265-.07-1.644-.07-4.849 0-3.204.013-3.583.07-4.849.149-3.227 1.664-4.771 4.919-4.919 1.266-.057 1.645-.069 4.849-.069zm0-2.163c-3.259 0-3.667.014-4.947.072-4.358.2-6.78 2.618-6.98 6.98-.059 1.281-.073 1.689-.073 4.948 0 3.259.014 3.668.072 4.948.2 4.358 2.618 6.78 6.98 6.98 1.281.058 1.689.072 4.948.072 3.259 0 3.668-.014 4.948-.072 4.354-.2 6.782-2.618 6.979-6.98.059-1.28.073-1.689.073-4.948 0-3.259-.014-3.667-.072-4.947-.196-4.354-2.617-6.78-6.979-6.98-1.281-.059-1.69-.073-4.949-.073zm0 5.838c-3.403 0-6.162 2.759-6.162 6.162s2.759 6.163 6.162 6.163 6.162-2.759 6.162-6.163c0-3.403-2.759-6.162-6.162-6.162zm0 10.162c-2.209 0-4-1.79-4-4 0-2.209 1.791-4 4-4s4 1.791 4 4c0 2.21-1.791 4-4 4zm6.406-11.845c-.796 0-1.441.645-1.441 1.44s.645 1.44 1.441 1.44c.795 0 1.439-.645 1.439-1.44s-.644-1.44-1.439-1.44z" />
                      </svg>
                      <span className="text-gray-300 group-hover:text-white">
                        Instagram
                      </span>
                    </a>
                  )}

                  {organizacao.redesSociais.twitter && (
                    <a
                      href={organizacao.redesSociais.twitter}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="flex items-center gap-3 p-3 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/20 hover:border-[#00f5ff] hover:bg-[#00f5ff]/10 transition-all group"
                    >
                      <svg
                        className="w-6 h-6 text-[#1DA1F2] group-hover:scale-110 transition-transform"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z" />
                      </svg>
                      <span className="text-gray-300 group-hover:text-white">
                        Twitter
                      </span>
                    </a>
                  )}
                </div>
              </div>
            </div>
          </div>
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
