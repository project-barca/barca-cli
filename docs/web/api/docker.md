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
---


# Adicionar Docker

Para integrar [Docker](https://github.com/project-barca) em seu projeto é muito simples, basta digitar o comando seguinte com o nome do diretório aonde o aplicativo se encontra:

<br/>
<br/>

```sh
 barca integrate docker <nome-do-projeto>
```

<br/>
<br/>
<br/>

Você pode utilizar algumas das configurações abaixo para ter uma melhor experiência:

<br/>

| Comando        | Descrição                                          |
| ------------- |:---------------------------------------------------:|
| **--user**      |especificar o usuário root do container (opcional) |
| **--password**      |específicar a senha do usuário (opcional)|
| **--port** |específicar porta do container Docker (opcional)|
| **--version**      |definir a versão para o ambiente (opcional) |
| **--remote**      |denifir recurso remoto (opcional) |
| **--cert**      |envio de certificado digital para comunicação segura (opcional) |
| **--dbname** |definir banco de dados (opcional)|


<br>
<br>


<br>
<br>






<!-- relative path -->
[Home](../README.md)  
[Config Reference](../reference/config.md)  
[Getting Started](./getting-started.md)  
<!-- absolute path -->
[Guide](/guide/README.md)  
[Config Reference > markdown.links](/reference/config.md#links)  
<!-- URL -->
[GitHub](https://github.com)  



[[toc]]