package entities

// **************************************************************************************
// **************************       XML DATA      ***************************************
// **************************************************************************************
type DriversXml struct {
	URL  string `xml:"url"`
	Name string `xml:"name"`
	Size string `xml:"size"`
}
type AboutBaseboardXml struct {
	Name         string `xml:"name"`
	Version      string `xml:"version"`
	SerialNumber string `xml:"serial_number"`
	Vendor       string `xml:"vendor"`
	BIOSversion  string `xml:"bios_version"`
}
type BaseboardXml struct {
	AboutBaseboardXml AboutBaseboardXml `xml:"about"`
	DriversXml        []DriversXml      `xml:"drivers"`
}

// **************************************************************************************
// **************************        MONGODB      ***************************************
// **************************************************************************************
type DriversMongo struct {
	URL  string `json:"url,omitempty" bson:"url,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Size int    `json:"size,omitempty" bson:"size,omitempty"`
}
type AboutBaseboardMongo struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Version      string `json:"version,omitempty" bson:"version,omitempty"`
	SerialNumber string `json:"serial_number,omitempty" bson:"serial_number,omitempty"`
	Vendor       string `json:"vendor,omitempty" bson:"vendor,omitempty"`
	BIOSversion  string `json:"bios_version,omitempty" bson:"bios_version,omitempty"`
}

type BaseboardMongo struct {
	AboutBaseboardMongo AboutBaseboardMongo `json:"about,omitempty" bson:"about,omitempty"`
	DriversMongo        []DriversMongo      `json:"drivers,omitempty" bson:"drivers,omitempty"`
}

// **************************************************************************************
// **************************************************************************************
type Drivers struct {
	URL  string
	Name string
	Size int
}
type AboutBaseboard struct {
	Name         string
	Version      string
	SerialNumber string
	Vendor       string
	BIOSversion  string
}

type Baseboard struct {
	AboutBaseboard AboutBaseboard
	Drivers        []Drivers
}
