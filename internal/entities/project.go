package entities

// **************************************************************************************
// **************************       XML DATA      ***************************************
// **************************************************************************************
type AboutXml struct {
	Name        string `xml:"Name"`
	Description string `xml:"Description"`
	Size        string `xml:"Size"`
}
type RepositoryXml struct {
	Url   string `xml:"Url"`
	Stars int    `xml:"Stars"`
	Fork  int    `xml:"Fork"`
}
type AuthorsXml struct {
	Name   string `xml:"Name"`
	Email  string `xml:"Email"`
	Github string `xml:"Github"`
}
type LanguagesXml struct {
	Name string `xml:"Name"`
	Ext  string `xml:"Ext"`
	Docs string `xml:"Docs"`
}
type FrameworkXml struct {
	Name     string `xml:"Name"`
	Version  string `xml:"Version"`
	Language string `xml:"Language"`
}

type ProjectXml struct {
	AboutXml      AboutXml       `xml:"About"`
	LanguagesXml  []LanguagesXml `xml:"Languages"`
	AuthorsXml    []AuthorsXml   `xml:"Authors"`
	RepositoryXml RepositoryXml  `xml:"Repository"`
	FrameworkXml  FrameworkXml   `xml:"Framework"`
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
