# Huffman

A árvore de Huffman é muito usada como algoritmo de comrpessão.

- Com perdar (Lossy Compression): audio e video
- Sem perdas (Lossless): Uso geral

## Visão geral

### Leaf Node

A leaf, é um nó sem filhos. A soma do peso (freqência) vai ser a de seu próprio nó.

### Node intermediário

O Nó intermediário vai ter um nó a esquerda e/ou a direita. Seu peso (frequência) vai ser a soma da frequência dos filhos.


## Codificação (Encode)

- Huffman Table

Geralmente, a Huffman Table é um hash map, que mapeia cada caracatere para sua freqência respectiva (char -> freq).

- Fila de Prioridade (Min Heap)

Para garantirmos que sempre vamos ter os nós em forma ascendente baseado em seus **pesos/frequência**, podemos simplesmente usar a estrutura de dados **min-heap**. Em Go, podemos implementar a *interface heap*. 



## Decodificação (Decode)

É uma árvore binária, mais fácil de implementar.