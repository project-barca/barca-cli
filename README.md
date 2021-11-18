
<p align="center">
 <img width="200px" height="200" src="/docs/assets/logo/barca-logo.jpeg" align="center" alt="GitHub Readme Stats" />
 <h2 align="center">Barca CLI</h2>
 <p align="center">
  Barca CLI is a project generator written in GO and its purpose is to build and configure HTTP servers, web proxy, SPA/PWA, Blog and custom landing page. It's easy, fast and productive.
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
    <a href="/docs/lang/readme_fr.md">Français </a>
    ·
    <a href="/docs/lang/readme_cn.md">简体中文</a>
    ·
    <a href="/docs/lang/readme_es.md">Español</a>
    ·
    <a href="/docs/lang/readme_ru.md">русский</a>
    .
    <a href="/docs/lang/readme_ja.md">日本語</a>
    ·
    <a href="/docs/lang/readme_de.md">Deutsch</a>
    ·
    <a href="/docs/lang/readme_it.md">Italiano</a>
    ·
    <a href="/docs/lang/readme_kr.md">한국어</a>
    .
    <a href="/docs/lang/readme_ro.md">Română</a>
    .
    <a href="/docs/lang/readme_uk.md">Українська</a>
    .
    <a href="/docs/lang/readme_hi.md">हिंदी</a>
    .
    <a href="/docs/lang/readme_th.md">ไทย</a>
    .
    <a href="/docs/lang/readme_vi.md">Tiếng Việt</a>
    .
    <a href="/docs/lang/readme_pt-BR.md">Português Brasileiro</a>
    .
  </p>
</p>

### Using the Tool

  **CLI** (Command-Line Interface) is an environment where users can enter command lines to perform tasks in operating systems or in programming jobs.



Downloads all modules in the file to the local cache

```sh
go mod download
```


Run `go run main.go` with argument `--help` to list all CLI commands and flags

Example: 
```sh
go run main.go --help
```

Output:
![alt text](docs/assets/output/cli-help.png)



To build some HTTP web server, we will mention the Framework Gorilla Mux and name of your project with the following command

```sh
go run main.go init api <project-name>
```

Example: 
```sh
go run main.go init api my-project
```

<p align="center">
  <img align="center" alt="barca-cli" src="docs/assets/gif/barca-cli-1.gif" />
</p>