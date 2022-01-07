package entities

// **************************************************************************************
// **************************       XML DATA      ***************************************
// **************************************************************************************
type AboutXml struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Size        string `xml:"size"`
}
type RepositoryXml struct {
	Url   string `xml:"url"`
	Stars int    `xml:"stars"`
	Fork  int    `xml:"fork"`
}
type AuthorsXml struct {
	Name   string `xml:"name"`
	Email  string `xml:"email"`
	Github string `xml:"github"`
}
type LanguagesXml struct {
	Name string `xml:"name"`
	Ext  string `xml:"ext"`
	Docs string `xml:"docs"`
}
type FrameworkXml struct {
	Name     string `xml:"name"`
	Version  string `xml:"version"`
	Language string `xml:"language"`
}

type ProjectXml struct {
	AboutMongo      AboutMongo       `xml:"about"`
	LanguagesMongo  []LanguagesMongo `xml:"languages"`
	AuthorsMongo    []AuthorsMongo   `xml:"authors"`
	RepositoryMongo RepositoryMongo  `xml:"repository"`
	FrameworkMongo  FrameworkMongo   `xml:"framework"`
}

// **************************************************************************************
// **************************        MONGODB      ***************************************
// **************************************************************************************
type AboutMongo struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Size        string `json:"size,omitempty" bson:"size,omitempty"`
}
type RepositoryMongo struct {
	Url   string `json:"url,omitempty" bson:"url,omitempty"`
	Stars int    `json:"stars,omitempty" bson:"stars,omitempty"`
	Fork  int    `json:"fork,omitempty" bson:"fork,omitempty"`
}
type AuthorsMongo struct {
	Name   string `json:"name,omitempty" bson:"name,omitempty"`
	Email  string `json:"email,omitempty" bson:"email,omitempty"`
	Github string `json:"github,omitempty" bson:"github,omitempty"`
}
type LanguagesMongo struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Ext  string `json:"ext,omitempty" bson:"ext,omitempty"`
	Docs string `json:"docs,omitempty" bson:"docs,omitempty"`
}
type FrameworkMongo struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Version  string `json:"version,omitempty" bson:"version,omitempty"`
	Language string `json:"language,omitempty" bson:"language,omitempty"`
}

type ProjectMongo struct {
	AboutMongo      AboutMongo       `json:"about,omitempty" bson:"about,omitempty"`
	LanguagesMongo  []LanguagesMongo `json:"languages,omitempty" bson:"languages,omitempty"`
	AuthorsMongo    []AuthorsMongo   `json:"authors,omitempty" bson:"authors,omitempty"`
	RepositoryMongo RepositoryMongo  `json:"repository,omitempty" bson:"repository,omitempty"`
	FrameworkMongo  FrameworkMongo   `json:"framework,omitempty" bson:"framework,omitempty"`
}

// **************************************************************************************
// **************************************************************************************
type About struct {
	Name        string
	Description string
	Size        string
}
type Repository struct {
	Url   string
	Stars int
	Fork  int
}
type Authors struct {
	Name   string
	Email  string
	Github string
}
type Languages struct {
	Name string
	Ext  string
	Docs string
}
type Framework struct {
	Name     string
	Version  string
	Language string
}

type Project struct {
	About      About
	Languages  []Languages
	Authors    []Authors
	Repository Repository
	Framework  Framework
}
