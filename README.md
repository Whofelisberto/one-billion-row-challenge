# One Billion Row Challenge

Este projeto é desafio feito em `java` ,mas recebeu muitas estrelas em `golang` então eu quiz fazer para praticar meus estudos. Ele lê um arquivo de 1 bilhão de linhas mostrando medições de temperatura por localização, processa os dados usando **goroutines** e **WaitGroups** para paralelismo, e calcula a temperatura mínima, máxima e média para cada local.

---

## Funcionalidades

- Lê medições de um arquivo `measurements.txt` no formato:


Exemplo:

São Paulo; 23.5, 
Rio de Janeiro; 28.0, 
Brasília; 20.1


- Utiliza **goroutines** para processar as linhas do arquivo em paralelo.
- Calcula para cada local:
  - Temperatura mínima
  - Temperatura máxima
  - Temperatura média
- Exibe os resultados em ordem alfabética dos locais.
- Mostra o tempo total de execução do programa.

---

## Estrutura do Código

- `Measurement`: Estrutura que armazena os dados agregados de cada local.
- `Partial`: Estrutura que representa uma medição individual.
- Funções auxiliares:
  - `min(a, b float64) float64` – Retorna o menor valor.
  - `max(a, b float64) float64` – Retorna o maior valor.
- Uso de **channels** (`lines` e `partials`) para comunicação entre goroutines.
- Proteção do mapa de resultados com **Mutex** (`sync.Mutex`) para concorrência segura.
- Processamento paralelo com `numWorkers` goroutines.

---

## Como usar

1. Certifique-se de ter o Go instalado: [https://golang.org/doc/install](https://golang.org/doc/install)
2. Crie um arquivo `measurements.txt` com os dados de temperaturas.
3. Compile e rode o programa:

```bash
go run main.go
```

## Tecnologias utilizadas

1. Golang
2. Goroutines e WaitGroups para paralelismo
3. Channels para comunicação entre goroutines
4. Mutex para concorrência segura

## Observações

O número de goroutines (numWorkers) pode ser ajustado conforme o número de núcleos da CPU.
O programa ignora linhas inválidas ou mal formatadas no arquivo de entrada.
