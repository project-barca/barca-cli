
<p align="center">
 <img width="200px" height="200" src="../assets/logo/barca-logo.jpeg" align="center" alt="GitHub Readme Stats" />
 <h2 align="center">Barca CLI</h2>
 <p align="center">
  Barca CLI é uma ferramenta escrita em GO que tem como objetivo construir e configurar servidores HTTP, proxy web, SPA/PWA, Blog e landing Page customizada. É fácil, rápido e produtivo.
  </p>
 </p>
  <p align="center">
    <a href="https://github.com/anuraghazra/github-readme-stats/actions">
      <img alt="GitHub issues" src="https://img.shields.io/github/issues/project-barca/barca-cli">
    </a>
    <a href="https://codecov.io/gh/anuraghazra/github-readme-stats">
      <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/project-barca/barca-cli">
    </a>
    <a href="https://a.paddle.com/v2/click/16413/119403?link=1227">
      <img alt="GitHub Release Date" src="https://img.shields.io/github/release-date/project-barca/barca-cli">
    </a>
    <a href="https://a.paddle.com/v2/click/16413/119403?link=2345">
      <img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/project-barca/barca-cli">
    </a>
  </p>
   
  <p align="center">
    <a href="/readme_fr.md">Français </a>
    ·
    <a href="/readme_cn.md">简体中文</a>
    ·
    <a href="/readme_es.md">Español</a>
    ·
    <a href="/readme_ru.md">русский</a>
    .
    <a href="/readme_ja.md">日本語</a>
    ·
    <a href="/readme_de.md">Deutsch</a>
    ·
    <a href="/readme_it.md">Italiano</a>
    ·
    <a href="/readme_kr.md">한국어</a>
    .
    <a href="/readme_uk.md">Українська</a>
    .
    <a href="/readme_hi.md">हिंदी</a>
    .
    <a href="/readme_vi.md">Tiếng Việt</a>
    .
    <a href="../../README.md">English</a>
    .
  </p>
</p>

### Usando a Ferramenta

  **CLI** (Command-Line Interface) é um ambiente onde os usuários podem inserir linhas de comando para executar tarefas em sistemas operacionais ou em trabalhos de programação.

<br>
<br>
<br>

Baixe todos módulos necessário no projeto:

```sh
go mod download
```
<br>
<br>

Execute o comando `go run main.go` com o argumento `--help` para listar todos comandos e flags da CLI
<br>
Exemplo:

```sh
go run main.go --help
```


<br>
<br>

Para construir um Servidor Web HTTP, nós vamos citar os argumentos  `init` e `api` para iniciar uma API Rest simples.

Exemplo:
<br>

```sh
go run main.go init api <seu-projeto>
```

<br>
<br>

<p align="center">
  <img align="center" alt="barca-cli" src="../assets/gif/barca-cli-1.gif" />
</p>

<br>
<br>

#### Especificando Idioma


Você pode traduzir as mensagens que a ferramenta **barca-cli** retorna especificando por uma flag `--lang`

Exemplo:

```sh
go run main.go --lang francais init api <seu-projeto>
```

<br>
<br>

<p align="center">
  <img align="center" alt="barca-cli" src="../assets/gif/barca-cli-2.gif" />
</p>

<br>
<br>

## Gerar API REST

É muito simples gerar uma API com **Barca CLI**, com poucos comandos você vai permitir a consrução do projeto da maneira que você quer.

A seguir está um exemplo de como gerar uma **API REST** em *Node.js* e configurar seu ambiente.

<br>
1. Servidor Web
2. Controladores, Rotas & Modelos
3. Banco de Dados
<br>

#### Configurar Servidor

Vamos mencionar **flags** para configurar o projeto, este é um exemplo utilizando o framework *Express.js* como servidor sendo executado na porta *4200*.

<br>

```sh
go run main.go --framework express --port 4200 init api <seu-projeto>
```

<br>
<br>

<p align="center">
  <img align="center" alt="barca-cli-express-api" src="docs/assets/gif/barca-cli-04-express-api.gif" />
</p>

<br>
<br>


#### Especificando seu idioma:

<br>

```sh
go run main.go --language portugues-brasileiro --framework express --port 4200 init api <your-project>
```

<br>
<br>

<p align="center">
  <img align="center" alt="barca-cli-api-language-i18n" src="../assets/gif/barca-cli-05-express-api-language.gif" />
</p>

<br>
<br>


#### Adicionar Modelos DB

<br>

Após a construção do servidor, vamos adicionar modelos para nosso projeto, especificamos o tipo de Banco de Dados que vai ser usado com a flag `--database` e a tabela/coleção com `--collection`. No exemplo a seguir está mostrando como inserir um modelo de usuários para Banco de dados **MySQL**

<br>
<br>

```sh
go run main.go --database mysql --collection usuarios  add model <seu-projeto>
```


<br>
<br>

<p align="center">
  <img align="center" alt="barca-cli-insert-models" src="../assets/gif/barca-cli-07-insert-models.gif" />
</p>

<br>
<br>

#### Integrar MySQL

Agora precisamos informar por **flags** as configurações para acessar e conectar o nosso servidor **MySQL**

<br>

```sh
go run main.go --dbname testdb --host 127.0.0.1 --user root --password 12345 integrate mysql <seu-projeto>
```

<br>
<br>

<p align="center">
  <img align="center" alt="barca-cli-mysql-integrate" src="../assets/gif/barca-cli-08-integrate-mysql.gif" />
</p>

<br>
<br>