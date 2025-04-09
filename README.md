# goauth

# Go Auth Kit

![Go Version](https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat-square&logo=go)
![GitHub License](https://img.shields.io/github/license/josevitorrodriguess/goauth)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)

A modular, ready-to-use authentication kit for your Go APIs.
This project aims to eliminate the need to rewrite authentication logic across multiple services. Instead of reimplementing auth every time, you can integrate a robust, reusable, and flexible solution in just a few steps.

## 🌟 Features

- **Multiple authentication methods:** JWT, Cookies, and Sessions.
- **Compatible with net/http and Chi:** Easy integration with the standard library and the [Chi router](https://github.com/go-chi/chi).
- **Ready-to-use middlewares:** Plug-and-play middlewares for seamless integration.
- **Predefined tables:** SQL scripts ready for different databases.
- **Password recovery:** Full password reset flow via email.
- **Flexible and extensible:** Easy to customize to fit your needs.


<!-- ## 📦 Instalação

```bash
go get -u github.com/seu-usuario/go-auth-kit
```

## 🚀 Início Rápido

### Autenticação JWT

```go
package main

import (
    "database/sql"
    "log"
    "net/http"
    
    "github.com/go-chi/chi/v5"
    "github.com/seu-usuario/go-auth-kit/auth/jwt"
    "github.com/seu-usuario/go-auth-kit/storage/postgres"
    
    _ "github.com/lib/pq"
)

func main() {
    // Conectar ao banco de dados
    db, err := sql.Open("postgres", "postgres://user:password@localhost/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    
    // Inicializar armazenamento
    storage := postgres.New(db)
    
    // Configurar autenticação JWT
    jwtAuth := jwt.New(jwt.Config{
        Secret: "chave-super-secreta",
        Expiration: 24*time.Hour,
        Storage: storage,
    })
    
    // Criar router Chi
    r := chi.NewRouter()
    
    // Rota pública - login
    r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
        username := r.FormValue("username")
        password := r.FormValue("password")
        
        token, err := jwtAuth.Login(username, password)
        if err != nil {
            http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
            return
        }
        
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"token": "` + token + `"}`))
    })
    
    // Grupo de rotas protegidas
    r.Group(func(r chi.Router) {
        // Aplicar middleware de autenticação
        r.Use(jwtAuth.Middleware())
        
        r.Get("/perfil", func(w http.ResponseWriter, r *http.Request) {
            // Obter ID do usuário do contexto (adicionado pelo middleware)
            userID := jwt.GetUserID(r.Context())
            
            // ... buscar dados do perfil no banco
            w.Write([]byte(`{"id": "` + userID + `", "nome": "Usuário Exemplo"}`))
        })
    })
    
    log.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", r)
}
```

### Autenticação por Sessão

```go
// Configurar autenticação por sessão
sessionAuth := session.New(session.Config{
    Secret: "chave-secreta-sessao",
    Expiration: 24*time.Hour,
    Storage: storage,
    CookieName: "auth_session",
})

// Usar middleware
r.Use(sessionAuth.Middleware())
```

## 🗄️ Estrutura do Banco de Dados

Antes de usar a biblioteca, você precisa criar as tabelas necessárias no seu banco de dados:

### Para JWT (PostgreSQL)

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Opcional: para armazenar tokens revogados (blacklist)
CREATE TABLE jwt_blacklist (
    token VARCHAR(500) PRIMARY KEY,
    expiry TIMESTAMP NOT NULL,
    revoked_at TIMESTAMP NOT NULL DEFAULT NOW()
);
```

### Para Sessões (PostgreSQL)

```sql
CREATE TABLE sessions (
    id VARCHAR(64) PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    data JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    last_activity TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Índice para limpar sessões expiradas
CREATE INDEX idx_sessions_expires_at ON sessions(expires_at);
```

Scripts para outros bancos de dados estão disponíveis na pasta `docs/schemas/`.

## 🔧 Opções de Configuração

### JWT

| Opção | Descrição | Padrão |
|-------|-----------|--------|
| `Secret` | Chave secreta para assinatura dos tokens | (obrigatório) |
| `Expiration` | Tempo de validade do token | 24h |
| `Issuer` | Emissor do token | "go-auth-kit" |
| `Algorithm` | Algoritmo de assinatura | HS256 |
| `Storage` | Interface de armazenamento | (obrigatório) |

### Sessão

| Opção | Descrição | Padrão |
|-------|-----------|--------|
| `Secret` | Chave para encriptar IDs de sessão | (obrigatório) |
| `Expiration` | Duração da sessão | 24h |
| `CookieName` | Nome do cookie | "session_id" |
| `CookieSecure` | Requer HTTPS | false |
| `CookieHTTPOnly` | Impede acesso via JavaScript | true |
| `Storage` | Interface de armazenamento | (obrigatório) |

## 🔄 Recuperação de Senha

```go
recoveryHandler := recovery.New(recovery.Config{
    Storage: storage,
    EmailSender: emailSender,
    TokenExpiration: 1*time.Hour,
    URLTemplate: "https://seusite.com/reset-password?token={token}",
})

// Rota para solicitar recuperação
r.Post("/forgot-password", recoveryHandler.RequestHandler)

// Rota para redefinir senha
r.Post("/reset-password", recoveryHandler.ResetHandler)
```

## 📋 Lista de Funcionalidades

- [x] Autenticação JWT
- [x] Autenticação por Sessão
- [x] Autenticação por Cookies
- [x] Middlewares para net/http e Chi
- [x] Recuperação de senha
- [ ] Autenticação OAuth (em desenvolvimento)
- [ ] Autenticação 2FA (em desenvolvimento)

## 🧪 Testes

```bash
go test ./...
``` -->

## 🤝 Contributing

Contributions are welcome! Please read the [contribution guidelines](CONTRIBUTING.md) before submitting a PR.


## 📄 License

This project is licensed under the [MIT License](LICENSE).
<!-- 
## 🙏 Agradecimentos

- [jwt-go](https://github.com/golang-jwt/jwt) - Biblioteca JWT para Go
- [Chi](https://github.com/go-chi/chi) - Router leve para Go -->