package entities

// **************************************************************************************
// **************************       XML DATA      ***************************************
// **************************************************************************************
type ProcessorCPUXml struct {
	ID           string `xml:"id"`
	NumCores     string `xml:"num_cores"`
	NumThreads   string `xml:"num_threads"`
	Vendor       string `xml:"vendor"`
	Model        string `xml:"model"`
	Capabilities string `xml:"capabilities"`
	Cores        string `xml:"cores"`
}
type AboutCPUXml struct {
	Name         string `xml:"name"`
	Version      string `xml:"version"`
	SerialNumber string `xml:"serial_number"`
	Vendor       string `xml:"vendor"`
	BIOSversion  string `xml:"bios_version"`
}
type CPUXml struct {
	AboutCPUXml     AboutCPUXml     `xml:"about"`
	ProcessorCPUXml ProcessorCPUXml `xml:"processor"`
}

// **************************************************************************************
// **************************        MONGODB      ***************************************
// **************************************************************************************
type ProcessorMongo struct {
	ID           string `json:"id,omitempty" bson:"id,omitempty"`
	NumCores     string `json:"num_cores,omitempty" bson:"num_cores,omitempty"`
	NumThreads   string `json:"num_threads,omitempty" bson:"num_threads,omitempty"`
	Vendor       string `json:"vendor,omitempty" bson:"vendor,omitempty"`
	Model        string `json:"model,omitempty" bson:"model,omitempty"`
	Capabilities string `json:"capabilities,omitempty" bson:"capabilities,omitempty"`
	Cores        string `json:"cores,omitempty" bson:"cores,omitempty"`
}
type AboutCPUMongo struct {
	Name         string `json:"name,omitempty" bson:"name,omitempty"`
	Version      string `json:"version,omitempty" bson:"version,omitempty"`
	SerialNumber string `json:"serial_number,omitempty" bson:"serial_number,omitempty"`
	Vendor       string `json:"vendor,omitempty" bson:"vendor,omitempty"`
	BIOSversion  string `json:"bios_version,omitempty" bson:"bios_version,omitempty"`
}

type CPUMongo struct {
	AboutCPUMongo  AboutCPUMongo  `json:"about_cpu,omitempty" bson:"about_cpu,omitempty"`
	ProcessorMongo ProcessorMongo `json:"processor,omitempty" bson:"processor,omitempty"`
}

// **************************************************************************************
// **************************************************************************************
type Processor struct {
	ID           string
	NumCores     string
	NumThreads   string
	Vendor       string
	Model        string
	Capabilities string
	Cores        string
}
type AboutCPU struct {
	TotalCores   string
	TotalThreads string
	Processors   string
}

type CPU struct {
	AboutCPU  AboutCPU
	Processor Processor
}
