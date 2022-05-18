---
lang: pt-BR
title: Conectar ao Banco de Dados
description: Conectando com serviços de Banco de Dados
head:
  - - meta
    - name: foo
      content: bar
  - - link
    - rel: canonical
      href: foobar
prev:
  text: Conectar ao Banco de Dados
  link: /web/api/connect-db.html
next:
  text: Adicionar Docker
  link: /web/api/docker.html
---


## Adicionar Entidade


Adicione modelos de entidade para seu projeto, vamos configurar primeiro com as opções abaixo:

---
Configurando:
  - **--collection**: | especificar a entidade que vai ser adicionada ao projeto
  - **--database**: | definir o banco de dados
---

<br>
<br>

Neste exemplo, vamos adicionar um modelo **MySQL** nomeado de **Escola**:

```sh
barca --collection escola --database mysql add model <diretorio>
```


<br>
<br>
Pronto!
<br>
Agora temos a entidade Escola em nosso projeto

<br>
<br>
<br>
<br>
<br>
<br>


## Adicionar Recurso


Adicione funções em seu projeto:
<br>
<br>

---
Configurando:
  - **--method**: | especificar o método
  - **--collection**: | definir nome da função
  - **--path**: | especificar o caminho do projeto
---

<br>
<br>

As funções por padrão não vem com parâmetros nelas, nós passamos parâmetros para solucionar algum tipo de problema.
<br><br>
Sendo assim, temos as variedades nos métodos de criação de uma função.
<br><br>
O método **create** é obrigatório mencionar parâmetros para solucionarmos algo, assim também para o método **update**.

<br>

No exemplo a seguir vamos adicionar uma função com nome de **Escola** usando método  **create**, pois queremos usar está função para salvar os dados da Escola no Banco de Dados que já tinhamos configurado:

<br>

```
barca --method create --collection escola --path <diretorio> add function <param1> <param2>
```


<br>
<br>
Ok!
<br>
Foi adicionado uma função no projeto

<br>
<br>





<!-- relative path -->
[Início](../../README.md)  
[Aplicar Segurança](../sec/websecurity.md)  
[Alterar Arquitetura](./architecture.md)  
[Adicionar Interfaces](../components/template.md)  


[[toc]]