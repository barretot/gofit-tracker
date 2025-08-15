# 🏋️ GoFit Tracker

GoFit Tracker é um projeto escrito em Go que processa treinos de múltiplos usuários de forma **concorrente**, calcula pontuações baseadas em seus treinos e exibe um ranking com os melhores colocados.


## ⚙️ Tecnologias Utilizadas

- **Go 1.22+**

## 🚀 Como Executar

1. **Clone o repositório**:

```bash
git clone https://github.com/seu-usuario/gofit-tracker.git
cd gofit-tracker
```

2. **Adicione os dados de treino** (já existe um exemplo em `testdata/workouts.json`). Você pode usar [este arquivo com 53 usuários](./workouts.json).

3. **Execute o projeto**:

```bash
go run ./cmd/main.go
```

---

## 💡 Como Funciona

Cada usuário é processado em uma **goroutine separada**, que calcula a pontuação total (reps × weight). Um **contexto com timeout** limita o tempo máximo por usuário. Todas as pontuações são inseridas com segurança em uma estrutura protegida por **RWMutex**.

Exemplo de saída:

```
🏋️ Iniciando processamento dos treinos...
🏆 Top 3 usuários:
1. maria@example.com - 975 pontos
2. joao@example.com - 900 pontos
3. ana@example.com - 860 pontos
```

## 📈 Próximos Passos (Ideias)

* ✅ Adicionar leitura de JSON via API
* ⏳ Exportar ranking para CSV
* ⏳ Adicionar filtros por tipo de exercício


## 🧑‍💻 Autor

**Ruan Barreto**
Desenvolvedor Backend | Explorador de Go 🐹
[linkedin.com/in/ruanbarreto](https://www.linkedin.com/in/ruan-barreto-6253b1181/)


