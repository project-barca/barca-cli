package entities

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
	AboutMongo      AboutMongo
	LanguagesMongo  []LanguagesMongo
	AuthorsMongo    []AuthorsMongo
	RepositoryMongo RepositoryMongo
	FrameworkMongo  FrameworkMongo
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
