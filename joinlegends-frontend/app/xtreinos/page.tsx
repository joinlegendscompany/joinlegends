import Link from "next/link";
import Image from "next/image";

export default function XTreinosPage() {
  // Dados mockados de XTreinos
  const xtreinos = [
    {
      id: 1,
      titulo: "Treino de Aim - CS2",
      jogo: "Counter-Strike 2",
      criador: "ProPlayer123",
      data: "15/01/2024",
      horario: "20:00",
      vagas: "8/10",
      nivel: "Intermediário",
      descricao: "Sessão focada em melhorar precisão e reflexos. Trabalharemos técnicas de flick e tracking.",
      imagem: "/assets/x0.jpeg",
    },
    {
      id: 2,
      titulo: "Estratégias Táticas - Valorant",
      jogo: "Valorant",
      criador: "TacticalMaster",
      data: "16/01/2024",
      horario: "19:30",
      vagas: "5/8",
      nivel: "Avançado",
      descricao: "Análise de mapas e execução de estratégias em equipe. Ideal para jogadores que querem subir de rank.",
      imagem: "/assets/x1.jpeg",
    },
    {
      id: 3,
      titulo: "Fundamentos - League of Legends",
      jogo: "League of Legends",
      criador: "CoachLOL",
      data: "17/01/2024",
      horario: "18:00",
      vagas: "12/15",
      nivel: "Iniciante",
      descricao: "Aprenda os fundamentos do jogo: farm, posicionamento, visão e objetivos. Perfeito para novos jogadores.",
      imagem: "/assets/x2.jpeg",
    },
    {
      id: 4,
      titulo: "Scrims Competitivos - Fortnite",
      jogo: "Fortnite",
      criador: "ScrimOrganizer",
      data: "18/01/2024",
      horario: "21:00",
      vagas: "16/20",
      nivel: "Avançado",
      descricao: "Scrims organizados com análise pós-jogo. Ambiente competitivo para melhorar em equipe.",
      imagem: "/assets/x3.jpg",
    },
    {
      id: 5,
      titulo: "Mechanics Training - Rocket League",
      jogo: "Rocket League",
      criador: "RLPro",
      data: "19/01/2024",
      horario: "20:30",
      vagas: "6/8",
      nivel: "Intermediário",
      descricao: "Treino de mecânicas avançadas: arials, dribbles e passes. Foco em consistência e precisão.",
      imagem: "/assets/x0.jpeg",
    },
    {
      id: 6,
      titulo: "VOD Review - Apex Legends",
      jogo: "Apex Legends",
      criador: "ApexCoach",
      data: "20/01/2024",
      horario: "19:00",
      vagas: "10/12",
      nivel: "Todos os níveis",
      descricao: "Análise de gameplay gravado. Identifique erros e aprenda com situações reais de jogo.",
      imagem: "/assets/x1.jpeg",
    },
  ];

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
                className="text-[#00f5ff] transition-colors"
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
        <div className="max-w-7xl mx-auto">
          <div className="text-center">
            <h1 className="text-5xl md:text-6xl font-bold mb-4">
              <span className="text-[#00f5ff]">XTreinos</span>
            </h1>
            <p className="text-xl text-gray-400 mb-8 max-w-2xl mx-auto">
              Encontre sessões de treino personalizadas e eleve seu nível de jogo
            </p>
            
            {/* Filtros */}
            <div className="flex flex-col sm:flex-row gap-4 justify-center items-center mb-8">
              <input
                type="text"
                placeholder="Buscar XTreino..."
                className="px-6 py-3 bg-[#1a1a2e]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
              />
              <select className="px-6 py-3 bg-[#1a1a2e]/50 border border-[#00f5ff]/30 rounded-lg text-white focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all">
                <option value="">Todos os jogos</option>
                <option value="cs2">Counter-Strike 2</option>
                <option value="valorant">Valorant</option>
                <option value="lol">League of Legends</option>
                <option value="fortnite">Fortnite</option>
                <option value="rocket-league">Rocket League</option>
                <option value="apex">Apex Legends</option>
              </select>
              <select className="px-6 py-3 bg-[#1a1a2e]/50 border border-[#00f5ff]/30 rounded-lg text-white focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all">
                <option value="">Todos os níveis</option>
                <option value="iniciante">Iniciante</option>
                <option value="intermediario">Intermediário</option>
                <option value="avancado">Avançado</option>
              </select>
            </div>

            <button className="px-6 py-3 bg-[#00f5ff] text-[#0a0a0a] font-bold rounded-lg hover:shadow-[0_0_30px_rgba(0,245,255,0.6)] transition-all transform hover:scale-105">
              + Criar Novo XTreino
            </button>
          </div>
        </div>
      </section>

      {/* Lista de XTreinos */}
      <section className="py-8 px-4 sm:px-6 lg:px-8">
        <div className="max-w-7xl mx-auto">
          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-6">
            {xtreinos.map((xtreino) => (
              <div
                key={xtreino.id}
                className="group relative p-0 border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg hover:border-[#00f5ff] hover:shadow-[0_0_20px_rgba(0,245,255,0.3)] transition-all overflow-hidden"
              >
                <div className="absolute inset-0 bg-gradient-to-r from-[#00f5ff]/0 to-[#00f5ff]/0 group-hover:from-[#00f5ff]/5 group-hover:to-transparent rounded-lg transition-all"></div>
                
                {/* Imagem */}
                <div className="relative w-full h-48 overflow-hidden">
                  <Image
                    src={xtreino.imagem}
                    alt={xtreino.titulo}
                    fill
                    className="object-cover group-hover:scale-105 transition-transform duration-300"
                  />
                  <div className="absolute inset-0 bg-gradient-to-t from-[#1a1a2e] via-[#1a1a2e]/50 to-transparent"></div>
                  <span className="absolute top-4 right-4 px-3 py-1 text-xs font-semibold bg-[#00f5ff]/20 text-[#00f5ff] rounded-full border border-[#00f5ff]/30 backdrop-blur-sm">
                    {xtreino.nivel}
                  </span>
                </div>
                
                <div className="relative p-6">
                  {/* Header do Card */}
                  <div className="mb-4">
                    <h3 className="text-xl font-bold text-[#00f5ff] mb-1">
                      {xtreino.titulo}
                    </h3>
                    <p className="text-sm text-gray-400">{xtreino.jogo}</p>
                  </div>

                  {/* Descrição */}
                  <p className="text-gray-400 text-sm mb-4 line-clamp-3">
                    {xtreino.descricao}
                  </p>

                  {/* Informações */}
                  <div className="space-y-2 mb-4">
                    <div className="flex items-center text-sm text-gray-400">
                      <svg className="w-4 h-4 mr-2 text-[#00f5ff]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                      </svg>
                      Criado por: <span className="text-[#00f5ff] ml-1">{xtreino.criador}</span>
                    </div>
                    <div className="flex items-center text-sm text-gray-400">
                      <svg className="w-4 h-4 mr-2 text-[#00f5ff]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                      </svg>
                      {xtreino.data} às {xtreino.horario}
                    </div>
                    <div className="flex items-center text-sm text-gray-400">
                      <svg className="w-4 h-4 mr-2 text-[#00f5ff]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                      </svg>
                      Vagas: <span className="text-[#00f5ff] ml-1">{xtreino.vagas}</span>
                    </div>
                  </div>

                  {/* Botão de Inscrição */}
                  <button className="w-full px-4 py-2 bg-[#00f5ff] text-[#0a0a0a] font-semibold rounded-lg hover:shadow-[0_0_20px_rgba(0,245,255,0.5)] transition-all">
                    Inscrever-se
                  </button>
                </div>
              </div>
            ))}
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
                  <Link href="/#features" className="hover:text-[#00f5ff] transition-colors">
                    Features
                  </Link>
                </li>
                <li>
                  <a href="#" className="hover:text-[#00f5ff] transition-colors">
                    Preços
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-[#00f5ff] transition-colors">
                    Roadmap
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="text-[#00f5ff] font-semibold mb-4">Comunidade</h4>
              <ul className="space-y-2 text-gray-400 text-sm">
                <li>
                  <a href="#" className="hover:text-[#00f5ff] transition-colors">
                    Discord
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-[#00f5ff] transition-colors">
                    Twitter
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-[#00f5ff] transition-colors">
                    GitHub
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="text-[#00f5ff] font-semibold mb-4">Suporte</h4>
              <ul className="space-y-2 text-gray-400 text-sm">
                <li>
                  <a href="#" className="hover:text-[#00f5ff] transition-colors">
                    Documentação
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-[#00f5ff] transition-colors">
                    Contato
                  </a>
                </li>
                <li>
                  <a href="#" className="hover:text-[#00f5ff] transition-colors">
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
