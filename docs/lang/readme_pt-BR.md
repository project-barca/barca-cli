
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



Baixe todos módulos necessário no projeto:

```sh
go mod download
```


Execute o comando `go run main.go` com o argumento `--help` para listar todos comandos e flags da CLI

Exemplo:

```sh
go run main.go --help
```



Para construir um Servidor Web HTTP, nós vamos citar os argumentos  `init` e `api` para iniciar uma API Rest simples.

Exemplo:

```sh
go run main.go init api <seu-projeto>
```

<br>

<p align="center">
  <img align="center" alt="barca-cli" src="docs/assets/gif/barca-cli-1.gif" />
</p>