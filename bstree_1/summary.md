
# Binary Search Tree (BST)

```bash
		   6
	      / \
	     2   10
	    / \  /
	   1   4 8
	      /
	     3
```

Arvores binarias, sao arvores onde cada no pode ter nos maximo 2 filhos. No caso da Binary Search Tree, o filho do lado **esquerdo**, deve ter um valor **menor** do que o pai, e o lado **direito** deve ter um valor **maior** que seu pai.

Numa arvore binaria, a busca quando a arvore esta **balanceada** tera tempo O(log n), enquanto em uma lista, com uma busca linear gastaria O(n).

**Pontos Importantes ao Implementar**:

Existem diversas formas de se implementar uma árvore binária, e isso vai depender muito do seu use-case específico. Um ponto importante a se considere é; sua árvore admite valores repetidos ou não?

- Se sua árvore admite valores repetidos:
	- As inserções serão mais fáceis. Porém, você vai precisar **especificar** suas regras de *busca* e de *remoção*. Ao buscar um elemento, vai retornar sempre a primeira apariçao ou retornará todos existentes na árvore? Ao remover, vai remover o primeiro encontrado, ou todos?

- Se sua árvore não pode ter valores repetidos:

	- Nesse caso, o processo de buscar e remover será mais rápido, porém, muito *provavelmente* seus **nós** precisarão de um **mutex de escrita** já que processos / requisições simultâneas podem tentar inserir em um mesmo nó simultâneamente, que pode causar *problemas de race conditions e deadlocks*.

Vamos implementar os dois.

## Declarando sua árvore

Podemos declarar nossa árvore de diversas formas diferentes, muitos gostam de declarar apenas o *node* para ser a árvore e o nó, outros gostam de separar a estrutura do node com a da árvore. Pessoalmente, prefiro separar, isso me permite ver mais claramente e declarar métodos separados para ambos, caso necessário, além de usar o poder do *method / pointer receiver* em Go com mais poder. Além disso, seus nós podem ter uma *key* além do dado em si, mas não vamos usar nesse exemplo.

- Tree:

```go
type BSTree struct {
	root *Node
}
```

- Node (sem locks):

```go
type Node struct {
	data int
	left *Node
	right *Node
}
```

- Node (com mutex):

```go
type Node struct {
	data int
	left *Node
	right *Node

	mu *sync.RWMutex
}
```

## Atravessando uma Binary Search Tree

Como a BST e um **tipo de dados nao linear**, existem varias formas de percorrer isso. Existem **duas formas populares** de se fazer isso: *Percorrer em ordem* (inorder traversal) e *Percorrer em Niveis* (Level Order Traversal).

No caso do nosso exemplo, o output seria:

```bash
1 2 3 4 6 8 10
```

### Percorrendo em Ordem (Inorder Traversal)

Quando percorremos na ordem da arvore binaria, estamos fazendo uma travessia **depth-first** (profundidade primeiro), que e um metodo recursivo onde vamos na ordem **left node > root node > right node**. No depth-first, buscamos a ultima folha a esquerda de uma subarvore antes de mover para a proxima subarvore. 

```bash
1 2 3 4 6 8 10
```

### Percorrer em Niveis (Level Order Traversal)

Nesse caso, a abordagem toma como premissa a **largura da arvore**, e vai iterar atraves de toda a arvore, camada por camada.

```bash
6 2 10 1 4 8 3
```

### 