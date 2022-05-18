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
  text: Gerar sua primeira API
  link: /web/api/generate.html
next:
  text: Adicionar Entidade
  link: /web/api/add-domain.html#entidade
---


# Conectar Banco de Dados


Para conectar sua API com um serviço de Banco de Dados é bem simples, vamos precisar de especificar algumas das opções abaixo:

---
Configurando:
  - **--user**: | especificar o usuário
  - **--password**: | específicar porta para o servidor de escutar
  - **--dbname**: | definir tradução de linguagem
---

<br>
<br>

Agora vamos digitar o comando abaixo para conectar:

```sh
barca --host 127.0.0.1 --user root --password 12345 --dbname "test" integrate <database> <diretorio>
```

<br>
<br>
Pronto!
<br>
Seu projeto adquiriu conexão

<br>
<br>






<!-- relative path -->
[Início](../README.md)  
[Criar Instância MySQL](../cloud/mysql.md)  
[Criar Instância Maria DB](../cloud/mariadb.md)  
[Criar Instância MongoDB](../cloud/mongodb.md)  
[Criar Instância PostgreSQL](../cloud/postgres.md)  
[Criar Instância Oracle DB](../cloud/oracledb.md)  
[Criar Instância BarcaDB](../cloud/barcadb.md)  
[veja mais](../cloud/bwc-intro.md)  
<!-- absolute path -->
<!-- URL -->
[BWC - Barca Web Cloud](https://cloud.project-barca.io)  



[[toc]]