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

=========== rascunho...

# Binary Search Tree (BST)

Podemos pensar em uma BST como uma **árvore de decisão**, que tem uma regra muito simples:

- **Números menores**: sempre vão para **esquerda**.
- **Números maiores**: sempre vão para a **direita**.

Ao Inserir 4, 2, 1, 10, teremos a seguinte árvore:

```bash
    4      ← (raiz)
   / \
  2   10   ← (2 é menor que 4, vai para esquerda; 10 é maior, vai para direita)
 /
1          ← (1 é menor que 2, vai para esquerda de 2)
```

## Entendendo as funções feitas:

### PrintInorder

O ponto chave dessa função é na **ordem** que a função visita os nós:

```go
func (t *bst) PrintInorder(n *Node) {
    if n == nil {
        return
    }
    
    t.PrintInorder(n.left)    // 1º: Visita TUDO da esquerda primeiro
    fmt.Printf("%d ", n.data) // 2º: Imprime o nó atual
    t.PrintInorder(n.right)   // 3º: Visita TUDO da direita por último
}
```

Tendo em vista a nossa árvore declarada, teremos o seguinte passo a passo:

- Chamada 1

```go
PrintInorder(4) {
    // Chamada 1: n = nó(4)
    if n == nil { return }  // falso, continua
    
    PrintInorder(n.left)    // Chama PrintInorder(2) - VAI PARA STACK MEMORY
    // ⚠️ PAUSA AQUI! Não executa as próximas linhas ainda
    fmt.Printf("%d ", n.data)  // ← Ainda não executou
    PrintInorder(n.right)      // ← Ainda não executou
}

- Chamada 2

PrintInorder(2) {
    // Chamada 2: n = nó(2)  
    if n == nil { return }  // falso, continua
    
    PrintInorder(n.left)    // Chama PrintInorder(1) - VAI PARA STACK MEMORY
    // ⚠️ PAUSA AQUI!
    fmt.Printf("%d ", n.data)  // ← Ainda não executou
    PrintInorder(n.right)      // ← Ainda não executou  
}

- Chamada 3

PrintInorder(1) {
    // Chamada 3: n = nó(1)
    if n == nil { return }  // falso, continua
    
    PrintInorder(n.left)    // Chama PrintInorder(nil) 
}

- Chamada 4

PrintInorder(nil) {
    // Chamada 4: n = nil
    if n == nil { return }  // ✅ BASE CASE! RETORNA
}
```

Depois do base case, a **stack** memory se "desenrola" todas as chamdas anteriores. Ou seja, ele só vai printar realmente, depois que chegar no último elemento.

```go
// ↩️ VOLTA para PrintInorder(1)
PrintInorder(1) {
    PrintInorder(n.left)    // ✅ JÁ TERMINOU (retornou nil)
    fmt.Printf("%d ", n.data)  // 🎯 IMPRIME: "1 "
    PrintInorder(n.right)      // Chama PrintInorder(nil) → retorna
} // ✅ TERMINA PrintInorder(1)

// ↩️ VOLTA para PrintInorder(2) 
PrintInorder(2) {
    PrintInorder(n.left)    // ✅ JÁ TERMINOU 
    fmt.Printf("%d ", n.data)  // 🎯 IMPRIME: "2 "
    PrintInorder(n.right)      // Chama PrintInorder(nil) → retorna
} // ✅ TERMINA PrintInorder(2)

// ↩️ VOLTA para PrintInorder(4)
PrintInorder(4) {
    PrintInorder(n.left)    // ✅ JÁ TERMINOU
    fmt.Printf("%d ", n.data)  // 🎯 IMPRIME: "4 "
    PrintInorder(n.right)      // Chama PrintInorder(10)
}

PrintInorder(10) {
    PrintInorder(n.left)    // PrintInorder(nil) → retorna
    fmt.Printf("%d ", n.data)  // 🎯 IMPRIME: "10 "
    PrintInorder(n.right)      // PrintInorder(nil) → retorna
} // ✅ TERMINA PrintInorder(10)
```

Uma forma de visualizarmos isso, é colocar um **contador** e ver quando de fato está imprimindo:

```go
var count = 0

func (t *bst) PrintInorder(n *Node) {
	if n == nil {
		return
	}

	fmt.Printf("chamando novamente: %d\n", n.data)
	count++
	t.PrintInorder(n.left)
	fmt.Printf("printing data: %d, count: %d\n", n.data, count)
	fmt.Printf("%d \n", n.data)
	t.PrintInorder(n.right)
}
```

Podemos visualizar que irá imprimir:

```
chamando novamente: 4
chamando novamente: 2
chamando novamente: 1
printing data: 1, count: 3
1 
printing data: 2, count: 3
2 
printing data: 4, count: 3
4 
chamando novamente: 10
printing data: 10, count: 4
10 
```

#### Vamos isualizar isso na **Stack Memory**

**Fase 1**: Acumulando na Stack:

- *Primeiro estado*:

```bash
Stack:
┌─────────────────────────────┐
│ PrintInorder(4)             │ ← pausado antes de n.left
│ • vai chamar n.left (nó 2)  │
└─────────────────────────────┘
```

- *Segundo estado*:

```bash
Stack:
┌─────────────────────────────┐
│ PrintInorder(2)             │ ← pausado antes de n.left  
│ • vai chamar n.left (nó 1)  │
├─────────────────────────────┤
│ PrintInorder(4)             │ ← pausado em n.left
│ • esperando 2 terminar      │
└─────────────────────────────┘
```

- *Terceiro estado*: Chegou no **nó folha**

```bash
Stack:
┌─────────────────────────────┐
│ PrintInorder(1)             │ ← pausado antes de n.left
│ • vai chamar n.left (nil)   │
├─────────────────────────────┤
│ PrintInorder(2)             │ ← pausado em n.left
│ • esperando 1 terminar      │
├─────────────────────────────┤
│ PrintInorder(4)             │ ← pausado em n.left  
│ • esperando 2 terminar      │
└─────────────────────────────┘
```

- *Quarto estado*: **Base case** - fundo da recursão

Nesse 

```bash
Stack:
┌─────────────────────────────┐
│ PrintInorder(nil) ⚡BASE CASE│ ← return imediato!
├─────────────────────────────┤
│ PrintInorder(1)             │ ← pausado em n.left
│ • esperando nil terminar    │
├─────────────────────────────┤
│ PrintInorder(2)             │ ← pausado em n.left
│ • esperando 1 terminar      │
├─────────────────────────────┤
│ PrintInorder(4)             │ ← pausado em n.left
│ • esperando 2 terminar      │
└─────────────────────────────┘
```

**Fase 2**: Desenrolando a Stack (executa os prints).

- *Quinto estado*: nil retornou

Veja aqui, essa é a **mágica** da Stack Memory :). A stack funciona na forma onde: **last-in first out**, ou seja, o último a entrar. Dessa forma o **nó folha** será chamado primeiro, ao desenrolar a stack após chegar no base case.

```bash
Stack:
┌─────────────────────────────┐
│ PrintInorder(1)             │ ← n.left terminou!
│ • EXECUTA: print "1 "       │ 🎯 OUTPUT: "1 "
│ • vai chamar n.right (nil)  │
├─────────────────────────────┤
│ PrintInorder(2)             │ ← pausado em n.left
├─────────────────────────────┤
│ PrintInorder(4)             │ ← pausado em n.left
└─────────────────────────────┘
```

- *Sexto estado*: PrintInorder(1) terminou

```bash
Stack:
┌─────────────────────────────┐
│ PrintInorder(2)             │ ← n.left terminou!
│ • EXECUTA: print "2 "       │ 🎯 OUTPUT: "1 2 "
│ • vai chamar n.right (nil)  │
├─────────────────────────────┤
│ PrintInorder(4)             │ ← pausado em n.left
└─────────────────────────────┘
```

- *Sétimo estado*: PrintInorder(2) terminou

```bash
Stack:
┌─────────────────────────────┐
│ PrintInorder(4)             │ ← n.left terminou!
│ • EXECUTA: print "4 "       │ 🎯 OUTPUT: "1 2 4 "
│ • vai chamar n.right (nó 10)│
└─────────────────────────────┘
```

**Fase 3**: Lado direito. (vai acumulando novamente, mas do lado direito).

- *Estado 8*:

```bash
Stack:
┌─────────────────────────────┐
│ PrintInorder(10)            │ ← pausado antes de n.left
│ • vai chamar n.left (nil)   │
├─────────────────────────────┤
│ PrintInorder(4)             │ ← pausado em n.right
│ • esperando 10 terminar     │
└─────────────────────────────┘
```

- *Estado 9*: Base case.

```bash
Stack:
┌─────────────────────────────┐
│ PrintInorder(nil) ⚡BASE CASE│ ← return imediato!
├─────────────────────────────┤
│ PrintInorder(10)            │ ← pausado em n.left
├─────────────────────────────┤
│ PrintInorder(4)             │ ← pausado em n.right
└─────────────────────────────┘
```

- *Estado 10*: Printando o 10.

Stack:
┌─────────────────────────────┐
│ PrintInorder(10)            │ ← n.left terminou!
│ • EXECUTA: print "10 "      │ 🎯 OUTPUT: "1 2 4 10 "
│ • vai chamar n.right (nil)  │
├─────────────────────────────┤
│ PrintInorder(4)             │ ← pausado em n.right
└─────────────────────────────┘

- *Estado 11*: Fim

Stack:
┌─────────────────────────────┐
│ PrintInorder(4)             │ ← n.right terminou!
│ • TERMINA função completa   │ ✅ TUDO DONE!
└─────────────────────────────┘

🎯 OUTPUT FINAL: "1 2 4 10 "



