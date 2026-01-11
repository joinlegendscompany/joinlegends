import Link from "next/link";
import Image from "next/image";
import { xtreinosData } from "@/app/mock/xtreinos";

export default function XTreinoDetailPage({
  params,
}: {
  params: { id: string };
}) {
  const xtreino = xtreinosData;

  //Deixar para implementar esse trecho depois.
  // if (!xtreino) {
  //   return (
  //     <div className="min-h-screen bg-[#0a0a0a] text-[#ededed] flex items-center justify-center">
  //       <div className="text-center">
  //         <h1 className="text-4xl font-bold text-[#00f5ff] mb-4">
  //           XTreino não encontrado
  //         </h1>
  //         <Link href="/xtreinos" className="text-[#00f5ff] hover:underline">
  //           Voltar para XTreinos
  //         </Link>
  //       </div>
  //     </div>
  //   );
  // }

  const vagasRestantes = 10;
  const porcentagemVagas = "50%";

  console.log(JSON.stringify(xtreino, null, 2));

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

      {/* Hero Banner */}
      <section className="relative pt-24 pb-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-7xl mx-auto">
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

          <div className="relative w-full h-64 md:h-96 rounded-lg overflow-hidden border border-[#00f5ff]/30 mb-8">
            {xtreino?.imagem && (
              <Image
                src={xtreino.imagem}
                alt={xtreino.titulo}
                fill
                className="object-cover"
              />
            )}
            <div className="absolute inset-0 bg-gradient-to-t from-[#0a0a0a] via-[#0a0a0a]/50 to-transparent"></div>
            <div className="absolute bottom-6 left-6 right-6">
              <div className="flex items-center gap-4 mb-4">
                <div className="w-16 h-16 bg-[#0a0a0a]/90 backdrop-blur-sm rounded-lg border border-[#00f5ff]/30 flex items-center justify-center p-2">
                  <Image
                    src={xtreino.logo}
                    alt={xtreino.jogo}
                    width={48}
                    height={48}
                    className="object-contain"
                  />
                </div>
                <div>
                  <h1 className="text-3xl md:text-4xl font-bold text-white mb-2">
                    {xtreino.titulo}
                  </h1>
                  <p className="text-lg text-gray-300">{xtreino.jogo}</p>
                </div>
              </div>
              <div className="flex flex-wrap gap-2">
                {xtreino.tags.map((tag: string, index: number) => (
                  <span
                    key={index}
                    className="px-3 py-1 text-sm bg-[#00f5ff]/20 text-[#00f5ff] rounded-full border border-[#00f5ff]/30"
                  >
                    {tag}
                  </span>
                ))}
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Main Content */}
      <section className="px-4 sm:px-6 lg:px-8 pb-20">
        <div className="max-w-7xl mx-auto">
          <div className="grid lg:grid-cols-3 gap-8">
            {/* Conteúdo Principal */}
            <div className="lg:col-span-2 space-y-8">
              {/* Informações do Criador */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
                <h2 className="text-xl font-bold text-[#00f5ff] mb-4">
                  Criado por
                </h2>
                <div className="flex items-center gap-4">
                  <div className="w-16 h-16 rounded-full overflow-hidden border-2 border-[#00f5ff]/30">
                    <Image
                      src={xtreino.avatarCriador}
                      alt={xtreino.criador}
                      width={64}
                      height={64}
                      className="object-cover w-full h-full"
                    />
                  </div>
                  <div>
                    <h3 className="text-lg font-semibold text-white">
                      {xtreino.criador}
                    </h3>
                    <p className="text-sm text-gray-400">Organizador</p>
                  </div>
                </div>
              </div>

              {/* Descrição Completa */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
                <h2 className="text-xl font-bold text-[#00f5ff] mb-4">
                  Sobre este XTreino
                </h2>
                <div 
                  className="wysiwyg-content"
                  dangerouslySetInnerHTML={{ __html: xtreino.descricaoCompleta || xtreino.descricao }}
                />
              </div>

              {/* Participantes */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
                <h2 className="text-xl font-bold text-[#00f5ff] mb-4">
                  Participantes ({vagasRestantes}/{xtreino.vagasTotal})
                </h2>
                <div className="grid grid-cols-2 md:grid-cols-3 gap-4">
                  {xtreino.participantes.map((participante, index: number) => (
                    <div
                      key={index}
                      className="flex items-center gap-3 p-3 bg-[#0a0a0a]/50 rounded-lg border border-[#00f5ff]/10"
                    >
                      <div className="w-10 h-10 rounded-full overflow-hidden border border-[#00f5ff]/30">
                        <Image
                          src={participante.avatar}
                          alt={participante.nome}
                          width={40}
                          height={40}
                          className="object-cover w-full h-full"
                        />
                      </div>
                      <div className="flex-1 min-w-0">
                        <p className="text-sm font-semibold text-white truncate">
                          {participante.nome}
                        </p>
                        <p className="text-xs text-gray-400">
                          {participante.rank}
                        </p>
                      </div>
                    </div>
                  ))}
                  {/* Vagas vazias */}
                  {Array.from({ length: xtreino.vagas }).map((_, index) => (
                    <div
                      key={`vaga-${index}`}
                      className="flex items-center justify-center p-3 bg-[#0a0a0a]/30 rounded-lg border border-dashed border-[#00f5ff]/20"
                    >
                      <span className="text-sm text-gray-500">
                        Vaga disponível
                      </span>
                    </div>
                  ))}
                </div>
              </div>
            </div>

            {/* Sidebar */}
            <div className="lg:col-span-1 space-y-6">
              {/* Card de Inscrição */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6 sticky top-24">
                <div className="text-center mb-6">
                  <div className="inline-block px-4 py-2 bg-[#00f5ff]/20 text-[#00f5ff] rounded-full border border-[#00f5ff]/30 mb-4">
                    <span className="text-sm font-semibold">
                      {xtreino.nivel}
                    </span>
                  </div>
                  <h3 className="text-2xl font-bold text-white mb-2">
                    {xtreino.vagas} vagas disponíveis
                  </h3>
                  <div className="w-full bg-[#0a0a0a] rounded-full h-2 mb-2">
                    <div
                      className="bg-[#00f5ff] h-2 rounded-full transition-all"
                      style={{ width: `${porcentagemVagas}%` }}
                    ></div>
                  </div>
                  <p className="text-sm text-gray-400">
                    {vagasRestantes} de {xtreino.vagasTotal} vagas preenchidas
                  </p>
                </div>

                <div className="space-y-4 mb-6">
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-gray-400">Data:</span>
                    <span className="text-white font-semibold">
                      {xtreino.data}
                    </span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-gray-400">Horário:</span>
                    <span className="text-white font-semibold">
                      {xtreino.horario}
                    </span>
                  </div>
                  <div className="flex items-center justify-between text-sm">
                    <span className="text-gray-400">Duração:</span>
                    <span className="text-white font-semibold">
                      {xtreino.duracao}
                    </span>
                  </div>
                  <div className="pt-4 border-t border-[#00f5ff]/20">
                    <div className="flex items-center justify-between mb-3">
                      <span className="text-gray-400">Valor da Inscrição:</span>
                      <span className="text-xl font-bold text-[#00f5ff]">
                        R$ {typeof xtreino.valorInscricao === 'number' 
                          ? xtreino.valorInscricao.toFixed(2).replace(".", ",") 
                          : "0,00"}
                      </span>
                    </div>
                    <div className="flex items-center justify-between">
                      <span className="text-gray-400">Premiação Total:</span>
                      <span className="text-xl font-bold text-[#00ff41]">
                        R$ {typeof xtreino.valorPremiacao === 'number' 
                          ? xtreino.valorPremiacao.toFixed(2).replace(".", ",") 
                          : "0,00"}
                      </span>
                    </div>
                  </div>
                </div>

                <button className="w-full px-6 py-3 bg-[#00f5ff] text-[#0a0a0a] font-bold rounded-lg hover:shadow-[0_0_30px_rgba(0,245,255,0.6)] transition-all transform hover:scale-105">
                  Inscrever-se Agora
                </button>

                <p className="text-xs text-gray-500 text-center mt-4">
                  Ao se inscrever, você receberá um link de acesso por email
                </p>
              </div>

              {/* Informações Adicionais */}
              <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-6">
                <h3 className="text-lg font-bold text-[#00f5ff] mb-4">
                  Informações
                </h3>
                <div className="space-y-3 text-sm">
                  <div className="flex items-start gap-3">
                    <svg
                      className="w-5 h-5 text-[#00f5ff] mt-0.5"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth={2}
                        d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                      />
                    </svg>
                    <div>
                      <p className="text-white font-semibold">Requisitos</p>
                      <p className="text-gray-400">
                        Nível {xtreino.nivel.toLowerCase()} ou superior
                      </p>
                    </div>
                  </div>
                  <div className="flex items-start gap-3">
                    <svg
                      className="w-5 h-5 text-[#00f5ff] mt-0.5"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        strokeLinecap="round"
                        strokeLinejoin="round"
                        strokeWidth={2}
                        d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
                      />
                    </svg>
                    <div>
                      <p className="text-white font-semibold">Inclui</p>
                      <p className="text-gray-400">
                        Material de estudo e análise pós-treino
                      </p>
                    </div>
                  </div>
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
