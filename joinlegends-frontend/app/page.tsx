export default function Home() {
  return (
    <div className="min-h-screen bg-[#0a0a0a] text-[#ededed] overflow-x-hidden">
      {/* Navigation */}
      <nav className="fixed top-0 w-full z-50 bg-[#0a0a0a]/80 backdrop-blur-md border-b border-[#00f5ff]/20">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center">
              <h1 className="text-2xl font-bold text-[#00f5ff]">JoinLegends</h1>
            </div>
            <div className="hidden md:flex items-center space-x-8">
              <a
                href="#features"
                className="hover:text-[#00f5ff] transition-colors"
              >
                Features
              </a>
              <a
                href="#comunidades"
                className="hover:text-[#00f5ff] transition-colors"
              >
                Comunidades
              </a>
              <a
                href="#xtreinos"
                className="hover:text-[#00f5ff] transition-colors"
              >
                XTreinos
              </a>
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
      <section className="relative pt-32 pb-20 px-4 sm:px-6 lg:px-8">
        <div className="max-w-7xl mx-auto">
          <div className="text-center">
            <div className="inline-block mb-6 px-4 py-2 border border-[#00f5ff]/50 bg-[#00f5ff]/10 rounded-full">
              <span className="text-[#00f5ff] text-sm font-mono">
                BETA VERSION 0.1.0
              </span>
            </div>
            <h1 className="text-5xl md:text-7xl font-bold mb-6 leading-tight">
              <span className="text-[#00f5ff]">A Plataforma Definitiva</span>
              <br />
              <span className="text-white">para Jogadores</span>
            </h1>
            <p className="text-xl md:text-2xl text-gray-400 mb-8 max-w-3xl mx-auto">
              Conecte-se com jogadores, participe de{" "}
              <span className="text-[#00ff41]">XTreinos</span> e faça parte de
              <span className="text-[#ff00ff]"> comunidades</span> incríveis.
              Eleve seu jogo ao próximo nível.
            </p>
            <div className="flex flex-col sm:flex-row gap-4 justify-center items-center">
              <button className="px-8 py-4 bg-[#00f5ff] text-[#0a0a0a] font-bold text-lg rounded-lg hover:shadow-[0_0_30px_rgba(0,245,255,0.6)] transition-all transform hover:scale-105">
                Começar Agora
              </button>
              <button className="px-8 py-4 border-2 border-[#ff00ff] text-[#ff00ff] font-bold text-lg rounded-lg hover:bg-[#ff00ff]/10 hover:shadow-[0_0_30px_rgba(255,0,255,0.5)] transition-all">
                Saiba Mais
              </button>
            </div>
          </div>

          {/* Animated Grid Background */}
          <div className="absolute inset-0 -z-10 opacity-20">
            <div className="absolute inset-0 bg-[linear-gradient(to_right,#00f5ff_1px,transparent_1px),linear-gradient(to_bottom,#00f5ff_1px,transparent_1px)] bg-[size:4rem_4rem]"></div>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section
        id="features"
        className="py-20 px-4 sm:px-6 lg:px-8 bg-[#1a1a2e]/30"
      >
        <div className="max-w-7xl mx-auto">
          <div className="text-center mb-16">
            <h2 className="text-4xl md:text-5xl font-bold mb-4">
              <span className="text-[#00f5ff]">Recursos Poderosos</span>
            </h2>
            <p className="text-gray-400 text-lg max-w-2xl mx-auto">
              Tudo que você precisa para dominar o jogo e construir sua legenda
            </p>
          </div>

          <div className="grid md:grid-cols-2 lg:grid-cols-3 gap-8">
            {/* Feature 1 - Comunidades */}
            <div className="group relative p-6 border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg hover:border-[#00f5ff] hover:shadow-[0_0_20px_rgba(0,245,255,0.3)] transition-all">
              <div className="absolute inset-0 bg-gradient-to-r from-[#00f5ff]/0 to-[#00f5ff]/0 group-hover:from-[#00f5ff]/5 group-hover:to-transparent rounded-lg transition-all"></div>
              <div className="relative">
                <div className="w-12 h-12 mb-4 bg-[#00f5ff]/20 rounded-lg flex items-center justify-center">
                  <svg
                    className="w-6 h-6 text-[#00f5ff]"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z"
                    />
                  </svg>
                </div>
                <h3 className="text-xl font-bold mb-2 text-[#00f5ff]">
                  Comunidades
                </h3>
                <p className="text-gray-400">
                  Junte-se a comunidades de jogadores apaixonados. Compartilhe
                  estratégias, organize eventos e construa sua rede.
                </p>
              </div>
            </div>

            {/* Feature 2 - XTreinos */}
            <div className="group relative p-6 border border-[#00ff41]/30 bg-[#1a1a2e]/50 rounded-lg hover:border-[#00ff41] hover:shadow-[0_0_20px_rgba(0,255,65,0.3)] transition-all">
              <div className="absolute inset-0 bg-gradient-to-r from-[#00ff41]/0 to-[#00ff41]/0 group-hover:from-[#00ff41]/5 group-hover:to-transparent rounded-lg transition-all"></div>
              <div className="relative">
                <div className="w-12 h-12 mb-4 bg-[#00ff41]/20 rounded-lg flex items-center justify-center">
                  <svg
                    className="w-6 h-6 text-[#00ff41]"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M13 10V3L4 14h7v7l9-11h-7z"
                    />
                  </svg>
                </div>
                <h3 className="text-xl font-bold mb-2 text-[#00ff41]">
                  XTreinos
                </h3>
                <p className="text-gray-400">
                  Crie e participe de sessões de treino personalizadas. Melhore
                  suas habilidades com outros jogadores dedicados.
                </p>
              </div>
            </div>

            {/* Feature 3 - Progresso */}
            <div className="group relative p-6 border border-[#ff00ff]/30 bg-[#1a1a2e]/50 rounded-lg hover:border-[#ff00ff] hover:shadow-[0_0_20px_rgba(255,0,255,0.3)] transition-all">
              <div className="absolute inset-0 bg-gradient-to-r from-[#ff00ff]/0 to-[#ff00ff]/0 group-hover:from-[#ff00ff]/5 group-hover:to-transparent rounded-lg transition-all"></div>
              <div className="relative">
                <div className="w-12 h-12 mb-4 bg-[#ff00ff]/20 rounded-lg flex items-center justify-center">
                  <svg
                    className="w-6 h-6 text-[#ff00ff]"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
                    />
                  </svg>
                </div>
                <h3 className="text-xl font-bold mb-2 text-[#ff00ff]">
                  Acompanhe Progresso
                </h3>
                <p className="text-gray-400">
                  Visualize suas estatísticas, conquistas e evolução. Veja seu
                  crescimento em tempo real.
                </p>
              </div>
            </div>

            {/* Feature 4 - Competições */}
            <div className="group relative p-6 border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg hover:border-[#00f5ff] hover:shadow-[0_0_20px_rgba(0,245,255,0.3)] transition-all">
              <div className="absolute inset-0 bg-gradient-to-r from-[#00f5ff]/0 to-[#00f5ff]/0 group-hover:from-[#00f5ff]/5 group-hover:to-transparent rounded-lg transition-all"></div>
              <div className="relative">
                <div className="w-12 h-12 mb-4 bg-[#00f5ff]/20 rounded-lg flex items-center justify-center">
                  <svg
                    className="w-6 h-6 text-[#00f5ff]"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"
                    />
                  </svg>
                </div>
                <h3 className="text-xl font-bold mb-2 text-[#00f5ff]">
                  Competições
                </h3>
                <p className="text-gray-400">
                  Participe de torneios e eventos. Prove seu valor e suba nos
                  rankings.
                </p>
              </div>
            </div>

            {/* Feature 5 - Networking */}
            <div className="group relative p-6 border border-[#00ff41]/30 bg-[#1a1a2e]/50 rounded-lg hover:border-[#00ff41] hover:shadow-[0_0_20px_rgba(0,255,65,0.3)] transition-all">
              <div className="absolute inset-0 bg-gradient-to-r from-[#00ff41]/0 to-[#00ff41]/0 group-hover:from-[#00ff41]/5 group-hover:to-transparent rounded-lg transition-all"></div>
              <div className="relative">
                <div className="w-12 h-12 mb-4 bg-[#00ff41]/20 rounded-lg flex items-center justify-center">
                  <svg
                    className="w-6 h-6 text-[#00ff41]"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M8.684 13.342C8.886 12.938 9 12.482 9 12c0-.482-.114-.938-.316-1.342m0 2.684a3 3 0 110-2.684m0 2.684l6.632 3.316m-6.632-6l6.632-3.316m0 0a3 3 0 105.367-2.684 3 3 0 00-5.367 2.684zm0 9.316a3 3 0 105.368 2.684 3 3 0 00-5.368-2.684z"
                    />
                  </svg>
                </div>
                <h3 className="text-xl font-bold mb-2 text-[#00ff41]">
                  Networking
                </h3>
                <p className="text-gray-400">
                  Conecte-se com jogadores profissionais, coaches e equipes.
                  Expanda sua rede.
                </p>
              </div>
            </div>

            {/* Feature 6 - Tempo Real */}
            <div className="group relative p-6 border border-[#ff00ff]/30 bg-[#1a1a2e]/50 rounded-lg hover:border-[#ff00ff] hover:shadow-[0_0_20px_rgba(255,0,255,0.3)] transition-all">
              <div className="absolute inset-0 bg-gradient-to-r from-[#ff00ff]/0 to-[#ff00ff]/0 group-hover:from-[#ff00ff]/5 group-hover:to-transparent rounded-lg transition-all"></div>
              <div className="relative">
                <div className="w-12 h-12 mb-4 bg-[#ff00ff]/20 rounded-lg flex items-center justify-center">
                  <svg
                    className="w-6 h-6 text-[#ff00ff]"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M13 10V3L4 14h7v7l9-11h-7z"
                    />
                  </svg>
                </div>
                <h3 className="text-xl font-bold mb-2 text-[#ff00ff]">
                  Tempo Real
                </h3>
                <p className="text-gray-400">
                  Notificações instantâneas, chat em tempo real e atualizações
                  ao vivo. Nunca perca nada.
                </p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* CTA Section */}
      <section className="py-20 px-4 sm:px-6 lg:px-8 relative overflow-hidden">
        <div className="absolute inset-0 bg-[#00f5ff]/10"></div>
        <div className="max-w-4xl mx-auto text-center relative z-10">
          <h2 className="text-4xl md:text-5xl font-bold mb-6">
            <span className="text-[#00f5ff]">Pronto para Começar?</span>
          </h2>
          <p className="text-xl text-gray-400 mb-8">
            Junte-se a milhares de jogadores que já estão elevando seu nível de
            jogo
          </p>
          <button className="px-10 py-5 bg-[#00f5ff] text-[#0a0a0a] font-bold text-xl rounded-lg hover:shadow-[0_0_40px_rgba(0,245,255,0.6)] transition-all transform hover:scale-105">
            Criar Conta Gratuita
          </button>
        </div>
      </section>

      {/* Footer */}
      <footer className="border-t border-[#00f5ff]/20 bg-[#1a1a2e]/30 py-12 px-4 sm:px-6 lg:px-8">
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
                  <a
                    href="#"
                    className="hover:text-[#00f5ff] transition-colors"
                  >
                    Features
                  </a>
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
              <h4 className="text-[#00ff41] font-semibold mb-4">Comunidade</h4>
              <ul className="space-y-2 text-gray-400 text-sm">
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00ff41] transition-colors"
                  >
                    Discord
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00ff41] transition-colors"
                  >
                    Twitter
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#00ff41] transition-colors"
                  >
                    GitHub
                  </a>
                </li>
              </ul>
            </div>
            <div>
              <h4 className="text-[#ff00ff] font-semibold mb-4">Suporte</h4>
              <ul className="space-y-2 text-gray-400 text-sm">
                <li>
                  <a
                    href="#"
                    className="hover:text-[#ff00ff] transition-colors"
                  >
                    Documentação
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#ff00ff] transition-colors"
                  >
                    Contato
                  </a>
                </li>
                <li>
                  <a
                    href="#"
                    className="hover:text-[#ff00ff] transition-colors"
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
