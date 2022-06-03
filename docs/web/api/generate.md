---
lang: pt-BR
title: Sua Primeira API REST
description: Construindo um serviço de REST API para Web em diversos tecnologias
head:
  - - meta
    - name: foo
      content: bar
  - - link
    - rel: canonical
      href: foobar
prev:
  text: Início
  link: /
next:
  text: Conectar ao Banco de Dados
  link: /web/api/connect-db.html
---

## Gerar API REST


O **Barca CLI** propoem uma diversidade em opções de tecnologias para podermos montar projetos eficientes em pouco tempo, este guia vai nos orientar em como você pode criar muito rápido uma **API REST** utilizando linguagens como *Python*, *Rust*, *Go*, *JavaScript*, *Julia*, *Scala*, PHP, *JAVA*, Elixir e etc.


<br>

::: tip Dica
Além do **Barca CLI**, existe mais ferramentas na plataforma [Barca Web Cloud](https://cloud.project-barca.com/tools/overview) que pode ajudar a encontrar algum recurso que você esteja procurando 
:::
<br>
<br>

Para obter mais detalhes sobre as ferramentas do Barca, consulte no guia [BWC Tools](https://cloud.barca.com/bwcloud/overview)

## Primeiros passos

<br/>

Ao começar a usar o `Barca CLI`, você irá desenvolver com uma configuração `padrão` e poderá definir as propriedades para montar uma **API** de duas formas:

<br/>

1. A primeira opção seria através do comando `barca setconfig <arquivo> <projeto>` carregar um arquivo com as configurações ideais para seu projeto.

---
2. A segunda opção seria usando flags:
  - **--framework**: | selecione o framework
  - **--port**: | define uma porta para seu servidor
  - **--lang**: | define o idioma
---

<br>
<br>
<br>

Vamos montar um projeto básico com [**Barca CLI**](https://docs.project-barca.com/cli/overview) escolhendo uma das linguagens mais utilizadas na Web, o famoso JavaScript.

Nesta **API** decidimos específicar a porta **5000** e o framework **Express.js**, por `padrão` é utilizado a porta **3000**.

<br>

Abra o terminal e digite:

<br>

```bash
barca --framework express --port 5000 init api <nome-de-sua-api>
```

<br>
<br>
<br>

Especifique a linguagem de idioma que você deseja usar em seu projeto da seguinte forma:

`barca --lang francais --framework express --port 5000 init api <nome-de-sua-api>`

<br>

Pronto! você já tem uma **API em Node.js** funcionando perfeitamente em seu ambiente
<br>
Entre em seu projeto e execute:
<br>
<br>

```sh
npm run start
```
<br>
Seu servidor será executado no endereço exibido em seu terminal

<br>
<br>
<br>
<br>
<br>
<br>

## API com Laravel PHP

Nesta **API** decidimos específicar a porta **5000** e o framework **Laravel PHP**, por `padrão` é utilizado a porta **3000**.

<br>

Abra o terminal e digite:

<br>

```bash
barca --framework laravel --port 5000 init api <nome-de-sua-api>
```


<br>

Pronto! você já tem uma **API em Laravel PHP +8** funcionando perfeitamente em seu ambiente
<br>
Entre em seu projeto e execute:
<br>
<br>

```sh
php artisan server
```

<br>

## API com Flask Python

---

---

<br>


## API com Django Python

---

---

<br>

## API com Symfony PHP

---

---

<br>

## API com Gorilla GO

## API com Bottle Python

## API com Gin GO

## API com Beego GO
## API com Fiber GO

## API com Echo GO

## API com Cake PHP



<!-- relative path -->
[Início](../../README.md)    
[Adicionar Entidade](../api/add-domain.md#adicionar-entidade)  
[Adicionar Função](../api/add-domain.md#adicionar-recurso)  
[Adicionar Docker](../api/docker.md)  
[Adicionar Segurança](../proj/security.md)  
[Escanear Projeto](../proj/scan.md#api)  
[Impulsionar Projeto](../proj/promote.md)  
[Fazer Deploy ](../proj/deploy.md#api)  

[Ecossistema](https://eco.project-barca.io)
<!-- absolute path -->
<!-- URL -->



[[toc]]