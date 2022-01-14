package entities

// **************************************************************************************
// **************************       XML DATA      ***************************************
// **************************************************************************************
type InterfaceXml struct {
	Name      string `xml:"name"`
	HAddr     string `xml:"hardware_address"`
	MTU       string `xml:"mtu"`
	Flag      string `xml:"flag"`
	IsVirtual string `xml:"is_virtual"`
}
type BluetoothXml struct {
	Name    string `xml:"name"`
	Version string `xml:"version"`
	MAC     string `xml:"mac"`
}

type NetworkXml struct {
	InterfaceXml []LanguagesXml `xml:"interface"`
	BluetoothXml []AuthorsXml   `xml:"bluetooth"`
}

// **************************************************************************************
// **************************        MONGODB      ***************************************
// **************************************************************************************
type BluetoothMongo struct {
	Name    string `json:"name,omitempty" bson:"name,omitempty"`
	Version string `json:"version,omitempty" bson:"version,omitempty"`
	MAC     string `json:"mac,omitempty" bson:"mac,omitempty"`
}
type InterfaceMongo struct {
	Name      string `json:"name,omitempty" bson:"name,omitempty"`
	HAddr     string `json:"hardware_address,omitempty" bson:"hardware_address,omitempty"`
	MTU       string `json:"mtu,omitempty" bson:"mtu,omitempty"`
	Flag      string `json:"flag,omitempty" bson:"flag,omitempty"`
	IsVirtual bool   `json:"is_virtual,omitempty" bson:"is_virtual,omitempty"`
}

type NetworkMongo struct {
	InterfaceMongo []InterfaceMongo `json:"interface_mongo,omitempty" bson:"interface_mongo,omitempty"`
	BluetoothMongo []BluetoothMongo `json:"bluetooth_mongo,omitempty" bson:"bluetooth_mongo,omitempty"`
}

// **************************************************************************************
// **************************************************************************************
type Interface struct {
	Name      string
	HAddr     string
	MTU       string
	Flag      string
	IsVirtual bool
}
type Bluetooth struct {
	Name    string
	Version string
	MAC     string
}

type Network struct {
	Interface []Interface
	Bluetooth []Bluetooth
}
