# 🧠 Goblockchain

Um projeto educacional de blockchain desenvolvido em Go, que implementa conceitos fundamentais como criação de blocos, mineração, validação de transações e geração de carteiras. Inclui também uma API REST para interação com a blockchain e um servidor de carteira com interface web.

## 📦 Estrutura do Projeto

```
├── block                # Implementação da blockchain e lógica dos blocos
│   └── blockchain.go
├── blockchain_server    # API para interagir com a blockchain
│   ├── blockchain_server.go
│   └── main.go
├── wallet               # Geração e gerenciamento de carteiras (Wallets)
│   └── wallet.go
├── wallet_server        # Servidor web para interação com carteiras
│   ├── main.go
│   ├── templates
│   │   └── index.html
│   └── wallet_server.go
├── utils                # Funções utilitárias (ECDSA, JSON, etc.)
│   ├── ecdsa.go
│   └── json.go
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

---

## 🚀 Funcionalidades

* 🔗 **Blockchain:** Criação de blocos, prova de trabalho (Proof of Work) e verificação de integridade.
* ⛏️ **Mineração:** Sistema de recompensa para mineradores e validação das transações.
* 🪙 **Carteira:** Geração de chaves públicas e privadas utilizando ECDSA, endereço blockchain derivado com hashing SHA-256 e RIPEMD-160.
* 🔐 **Assinatura Digital:** Assinatura de transações com segurança criptográfica.
* 🌐 **API REST:** Permite interagir com a blockchain via HTTP.
* 🖥️ **Interface Web:** Criação e gerenciamento de carteiras através de uma interface HTML simples.

---

## 🏗️ Tecnologias Utilizadas

* Golang
* Criptografia ECDSA
* SHA-256 e RIPEMD-160
* HTTP/REST (net/http)
* HTML (para wallet interface)

---

## ⚙️ Como Executar

### Pré-requisitos

* Go 1.18 ou superior instalado na sua máquina.

### Clone o repositório

```bash
git clone https://github.com/seu-usuario/goblockchain.git
cd goblockchain
```

### Instale as dependências

```bash
go mod tidy
```

### Execute o servidor da blockchain

```bash
go run blockchain_server/main.go --port=5000
```

A API estará disponível em:
👉 `http://localhost:5000`

### Execute o servidor da wallet (opcional)

```bash
go run wallet_server/main.go --port=8080
```

A interface estará disponível em:
👉 `http://localhost:8080`

---

## 📑 Endpoints da Blockchain API

| Método | Endpoint | Descrição                        |
| ------ | -------- | -------------------------------- |
| GET    | `/`      | Retorna a cadeia de blocos atual |

*(Outros endpoints podem ser implementados para enviar transações, minerar, etc.)*

---

## 🏛️ Conceitos Implementados

* **Blockchain:** Lista encadeada de blocos, onde cada bloco possui hash do anterior.
* **Prova de Trabalho (PoW):** Algoritmo de consenso simples com dificuldade configurável.
* **Mineração:** Geração de blocos válidos e recompensa para mineradores.
* **Assinatura Digital:** Verificação da autenticidade das transações.
* **Carteiras:** Endereços públicos, privados e codificação utilizando Base58.

---

## 🔧 Melhorias Futuras

* Implementar smart contracts simples.
* Suporte para múltiplos nodes e peer-to-peer.
* Validação de saldo antes da criação da transação (já previsto no código, mas comentado).
* Melhorias na interface web.
* Adicionar endpoints REST para envio de transações e mineração.

---

## 📜 Licença

Este projeto está licenciado sob a licença MIT - veja o arquivo [LICENSE](./LICENSE) para detalhes.

---

## 🙌 Créditos

Projeto com fins educacionais, inspirado nos conceitos de criptografia, blockchain e desenvolvimento backend em Go, desenvolvido como parte do curso de desenvolvimento de blockchain em Golang da Udemy.

---

