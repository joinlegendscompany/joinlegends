"use client";

import Link from "next/link";
import { useState } from "react";

export default function LoginPage() {
  const [formData, setFormData] = useState({
    email: "",
    senha: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // TODO: Implementar autenticação via API
    console.log("Dados de login:", formData);
    alert("Login realizado com sucesso! (Funcionalidade será implementada)");
  };

  const handleGoogleLogin = () => {
    // TODO: Implementar autenticação com Google
    alert("Autenticação com Google será implementada");
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
              <Link
                href="/auth/login"
                className="px-4 py-2 border border-[#00f5ff] text-[#00f5ff] hover:bg-[#00f5ff] hover:text-[#0a0a0a] transition-all"
              >
                Entrar
              </Link>
              <Link
                href="/auth/cadastro"
                className="px-4 py-2 bg-[#00f5ff] text-[#0a0a0a] font-semibold hover:shadow-[0_0_20px_rgba(0,245,255,0.5)] transition-all"
              >
                Cadastrar
              </Link>
            </div>
          </div>
        </div>
      </nav>

      {/* Main Content */}
      <section className="pt-32 pb-20 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md mx-auto">
          <div className="text-center mb-8">
            <h1 className="text-4xl md:text-5xl font-bold mb-4">
              <span className="text-[#00f5ff]">Entrar</span>
            </h1>
            <p className="text-gray-400">
              Acesse sua conta e continue sua jornada
            </p>
          </div>

          <div className="border border-[#00f5ff]/30 bg-[#1a1a2e]/50 rounded-lg p-8">
            <form onSubmit={handleSubmit} className="space-y-6">
              {/* Email */}
              <div>
                <label
                  htmlFor="email"
                  className="block text-sm font-semibold text-gray-300 mb-2"
                >
                  Email *
                </label>
                <input
                  type="email"
                  id="email"
                  name="email"
                  value={formData.email}
                  onChange={handleChange}
                  required
                  className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                  placeholder="seu@email.com"
                />
              </div>

              {/* Senha */}
              <div>
                <label
                  htmlFor="senha"
                  className="block text-sm font-semibold text-gray-300 mb-2"
                >
                  Senha *
                </label>
                <input
                  type="password"
                  id="senha"
                  name="senha"
                  value={formData.senha}
                  onChange={handleChange}
                  required
                  className="w-full px-4 py-3 bg-[#0a0a0a]/50 border border-[#00f5ff]/30 rounded-lg text-white placeholder-gray-500 focus:outline-none focus:border-[#00f5ff] focus:shadow-[0_0_20px_rgba(0,245,255,0.2)] transition-all"
                  placeholder="••••••••"
                />
              </div>

              {/* Esqueci a senha */}
              <div className="flex justify-end">
                <Link
                  href="/auth/recuperar-senha"
                  className="text-sm text-[#00f5ff] hover:underline"
                >
                  Esqueci minha senha
                </Link>
              </div>

              {/* Botão de Login */}
              <button
                type="submit"
                className="w-full px-6 py-3 bg-[#00f5ff] text-[#0a0a0a] font-bold rounded-lg hover:shadow-[0_0_30px_rgba(0,245,255,0.6)] transition-all transform hover:scale-105"
              >
                Entrar
              </button>

              {/* Divisor */}
              <div className="relative">
                <div className="absolute inset-0 flex items-center">
                  <div className="w-full border-t border-[#00f5ff]/20"></div>
                </div>
                <div className="relative flex justify-center text-sm">
                  <span className="px-2 bg-[#1a1a2e]/50 text-gray-400">
                    Ou continue com
                  </span>
                </div>
              </div>

              {/* Google Login */}
              <button
                type="button"
                onClick={handleGoogleLogin}
                className="w-full px-6 py-3 border-2 border-[#00f5ff]/30 bg-[#0a0a0a]/50 text-white font-semibold rounded-lg hover:border-[#00f5ff] hover:bg-[#00f5ff]/10 transition-all flex items-center justify-center gap-3"
              >
                <svg
                  className="w-5 h-5"
                  viewBox="0 0 24 24"
                  fill="currentColor"
                >
                  <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4" />
                  <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853" />
                  <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05" />
                  <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335" />
                </svg>
                Continuar com Google
              </button>

              {/* Link para Cadastro */}
              <div className="text-center text-sm text-gray-400">
                Não tem uma conta?{" "}
                <Link
                  href="/auth/cadastro"
                  className="text-[#00f5ff] hover:underline font-semibold"
                >
                  Cadastre-se
                </Link>
              </div>
            </form>
          </div>
        </div>
      </section>
    </div>
  );
}
