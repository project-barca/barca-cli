const version = require("./../../package.json").version;

module.exports = {
  head: [
    ['link', { rel: 'manifest', href: '/public/manifest.webmanifest' }],
    ['meta', { name: 'theme-color', content: '#fff' }],
    // ...other tags
  ],
  lang: 'pt-BR',
  title: `Barca CLI ${version}`,
  description: 'This is documentation for the Barca CLI command line tool',
  theme: '@vuepress/theme-default',
  markdown: {
    lineNumbers: true
  },
  locales: {
    '/': {
      lang: 'Português-Brasileiro',
      title: 'Barca CLI',
      description: 'Ferramenta para criar e personalizar qualquer projeto'
    },
    '/it/': {
      lang: 'Italiano',
      title: 'Barca CLI',
      description: 'Strumento per creare e personalizzare qualsiasi progetto'
    },
    '/fr/': {
      lang: 'Francês',
      title: 'Barca CLI',
      description: "Outil pour créer et personnaliser n'importe quel projet"
    },
    '/en/': {
      lang: 'Inglês',
      title: 'Barca CLI',
      description: 'Tool to create and customize any project'
    },
    '/al/': {
      lang: 'Alemão',
      title: 'Barca CLI',
      description: 'Tool zum Erstellen und Anpassen von Projekten'
    },
    '/jp/': {
      lang: 'Japonês',
      title: 'Barca CLI',
      description: 'プロジェクトを作成およびカスタマイズするためのツール'
    },
  },
  plugins: [
    '@vuepress/back-to-top', 
    '@vuepress/medium-zoom', 
    '@vuepress/plugin-search',
    {
      getExtraFields: (page) => page.frontmatter.tags ?? [],
    },
    [
      '@vuepress/plugin-google-analytics',
      {
        id: 'G-XXXXXXXXXX',
      },
    ],
    [
      '@vuepress/pwa',
      {
        skipWaiting: true,
      },
    ],
  ],
  themeConfig: {
    docsRepo: 'https://github.com/project-barca/quickstart',
    docsBranch: 'main',
    navbar: [
      // DOCS
      {
        text: 'Documentação',
        children: [
          {
            text: 'Barca',
            children: [
              {text: 'Introdução', link: '/barca/introduction/introduction.html'},
              {text: 'CLI', link: '/'},
              {text: 'Cliente API', link: '/web/api/generate.html'},
              {text: 'Desktop App', link: '/barca/install/generate.html'},
            ]
          },
          {
            text: 'Gerador',
            children: [
              {text: 'API', link: '/web/api/generate.html'},
              {text: 'Blog', link: '/web/blog/generate.html'},
              {text: 'Interfaces', link: '/web/interfaces/generate.html'},
              {text: 'Servidor Web', link: '/infra/webserver/generate.html'},
              {text: 'Infraestrutura', link: '/web/server/generate.html'},
            ]
          },
          {
            text: 'Referências',
            children: [
              {text: 'Base de Dados', link: '/'},
              {text: 'Temas', link: '/'},
              {text: 'Componentes', link: '/'},
            ]
          },
        ],
      },
      // LICENSE
      {
        text: 'Licença',
        link: '/barca/license',
      },
      // SUPORT
      {
        text: 'Suporte',
        children: [
          {
            text: 'Social',
            children: [
              {text: 'Twitter', link: 'https://twitter.com/projetobarca'},
              {text: 'Instagram', link: 'https://www.instagram.com/barca.project_br/'},
              {text: 'Youtube', link: '/barca/install/generate.html'},
              {text: 'Facebook', link: '/barca/install/generate.html'},
              {text: 'LinkedIn', link: 'https://www.linkedin.com/company/barca-project/'},
            ]
          },
          {
            text: 'Contato',
            children: [
              {text: 'Discord', link: '/barca/introduction/introduction.html'},
              {text: 'Line', link: '/'},
              {text: 'Status', link: 'https://status.im/'},
              {text: 'Telegram', link: '/barca/install/generate.html'},
              {text: 'Viber', link: '/barca/install/generate.html'},
            ]
          },
          {
            text: 'Ajuda',
            children: [
              {text: 'Issues', link: 'https://github.com/project-barca/barca-cli/issues'},
              {text: 'Pull Request CLI', link: 'https://github.com/project-barca/barca-cli/pulls'},
              {text: 'Pull Request API', link: 'https://github.com/project-barca/barca-api/pulls'}
            ]
          },
        ],
      },
      // VERSIONS
      {
        text: 'v0.2.x-beta',
        children: ['v0.1.0'],
        link: '/barca/versions/releases'
      },
    ],
    backToHome: 'Voltar ao inicio',
    //logoDark: '/public/img/logo/logo.jpeg',
    locales: {
      '/': {
        selectLanguageName: 'Português-Brasileiro',
      },
      '/en/': {
        selectLanguageName: 'English',
      },
      '/zh/': {
        selectLanguageName: '简体中文',
      },
    },
    logo: 'https://vuejs.org/images/logo.png',
    logoDark: 'https://vuejs.org/images/logo.png',
    lastUpdated: false,







    sidebar: {
      '/web/api': [
        {
          text: 'Crie sua API REST',
          children: [
            {
              text: 'Arquitetura',
              children: [
                {text: 'Client Server', link: '/web/api/architecture.html#client-server-architecture'},
                {text: 'Event Driven', link: '/web/api/architecture.html#event-driven-architecture'},
                {text: 'Model View Controller', link: '/web/api/architecture.html#model-view-controller-architecture'},
                {text: 'Microservice', link: '/web/api/architecture.html#microservice-architecture'},
                {text: 'Pipe and Filter', link: '/web/api/architecture.html#pipe-and-filter-architecture'},
              ]
            },
            {
              text: 'Linguagem',
              children: [
                {text: 'C#', link: '/web/api/languages.html#c#'},
                {text: 'GO', link: '/web/api/languages.html#golang'},
                {text: 'JAVA', link: '/web/api/languages.html#java'},
                {text: 'Kotlin', link: '/web/api/languages.html#kotlin'},
                {text: 'JavaScript', link: '/web/api/languages.html#javascript'},
                {text: 'Python', link: '/web/api/languages.html#python'},
                {text: 'PHP', link: '/web/api/languages.html#php'},
                {text: 'Rust', link: '/web/api/languages.html#rust'},
                {text: 'Swift', link: '/web/api/languages.html#swift'},
                {text: 'Scala', link: '/web/api/languages.html#scala'},
                {text: 'Swift', link: '/web/api/languages.html#swift'},
                {text: 'TypeScript', link: '/web/api/languages.html#typescript'},
              ]
            },
          ],
        },
        {
          text: 'Banco de Dados',
          children: ['Criar Modelos', 'Conexão', 'Variáveis de Ambiente', 'Queries'],
        },
        {
          text: 'Controles',
          children: ['Criar Recursos', 'Alterar', 'Excluír', 'Clonar'],
        },
        {
          text: 'Rotas',
          children: ['Nova Rota', 'Controllers em Rotas', 'Privilégios para Usuários', 'Dependências', 'Tradução'],
        },
        {
          text: 'Autenticações',
          children: ['JWT', 'Passport.js'],
        },
        {
          text: 'Segurança na API',
          children: ['Proxy', 'Ant-SQL Injection',],
        },
        {
          text: 'Deploy',
          children: ['Linux', 'MacOS', 'AWS', 'Google Cloud Platform', 'Heroku', 'Netlify', 'Github', 'OpenShift4', 'Windows'],
        },
      ],
      '/environment/install': [
        {
          text: 'Ambiente',
          children: ['Instalação no Windows', 'Instalação no Linux', 'Instalação no MacOS', 'Instalação com Docker', 'Instalação com Kubernets', 'Instalação com Vagrant'],
        },
      ],
      '/': [
        {
          text: 'Introdução',
          link: '/barca/introduction/introduction',
          children: [
            {
              text: 'Instalação',
              link: '/barca/install/so',
              children: [
                {text: 'Windows', link: '/barca/install/so.html#instalar-barca-no-windows'},
                {text: 'Linux', link: '/barca/install/so.html#instalar-barca-no-linux'},
                {text: 'MacOS', link: '/barca/install/so.html#instalar-barca-no-macos'},
              ]
            },
            {
              text: 'Versões',
              link: '/barca/versions/releases.html',
            },
            {text: 'Contribuições', link: '/barca/contributors'},
            {text: 'Patrocinadores', link: '/barca/sponsors'},
            {text: 'Licença', link: '/barca/introduction/introduction.html#license'},
          ],
        },
        {
          text: 'Início Rápido',
          children: [
          {
            text: 'Gerar API',
            link: '/web/api/generate.html',
            children: [
              {text: 'Express.js', link: '/web/api/generate.html#api-com-express-js'},
              {text: 'Laravel PHP', link: '/web/api/generate.html#api-com-laravel-php'},
              {text: 'GO', link: '/web/api/generate.html#api-com-go-lang'},
              {text: 'Rust', link: '/web/api/generate.html#api-com-rocket-rust'},
              {text: 'Symfony', link: '/web/api/generate.html#api-com-symfony-php'},
              {text: 'Spring', link: '/web/api/generate.html#api-com-spring-java'},
              {text: 'Play', link: '/web/api/generate.html#api-com-play-java'},
              {text: 'Django', link: '/web/api/generate.html#api-com-django-python'},
              {text: 'Flask', link: '/web/api/generate.html#api-com-flask-python'},
              {text: 'Bottle', link: '/web/api/generate.html#api-com-bottle-python'},
            ],
          },
          {
            text: 'Gerar Interfaces',
            link: '/web/interfaces/generate.html',
            children: [ 
              {text: 'Angular.js', link: '/web/angular'}, 
              {text: 'React.js', link: '/web/react'}, 
              {text: 'Vue.js', link: '/web/vue'}
            ],
          },
          {
            text: 'Servidor',
            link: '/machine/server',
            children: [ {text: 'Linux', link: '/so/server/linux'}, {text: 'Windows', link: '/so/server/windows'}],
          },
        ]
        },
      ],
    },
    nav: [
      { text: "Home", link: "/" }
    ],
    repoLabel: "Contribute an Example !",
    displayAllHeaders: true
  },
}