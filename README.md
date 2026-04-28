# 💰 Sistema de Planejamento Financeiro - Go Expert Session

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![CircleCI](https://img.shields.io/circleci/build/github/ricardovalehrio/dio-expert-session-finance?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Finalizado-green?style=for-the-badge)

Esta é uma API REST desenvolvida em **GoLang** que une conceitos de **Ciência da Computação** e **Contabilidade**. O projeto foi criado para gerenciar transações financeiras, permitindo o registro e a consulta de entradas e saídas de capital.

## 🚀 Tecnologias e Ferramentas
* **Go (Golang)**: Linguagem de alta performance para o backend.
* **Net/HTTP**: Servidor web nativo do Go.
* **JSON**: Formato de comunicação para a API.
* **CircleCI**: Esteira de integração contínua (CI/CD) para automação de testes.
* **Postman**: Utilizado para validar as rotas e payloads.

## 📂 Organização do Projeto
* `main.go`: Ponto de entrada, configuração das rotas e inicialização do servidor.
* `model/transaction/`: Definição das estruturas de dados (`Transaction` e `Transactions`).
* `.circleci/`: Configurações de automação para build e testes.

## 🛠️ Como Executar e Testar

1. **Clone o repositório:**
   ```bash
   git clone [https://github.com/ricardovalehrio/dio-expert-session-finance.git](https://github.com/ricardovalehrio/dio-expert-session-finance.git)