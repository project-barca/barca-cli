package entities

// **************************************************************************************
// **************************        MONGODB      ***************************************
// **************************************************************************************
type AboutChassiMongo struct {
	Tag             string `json:"tag,omitempty" bson:"tag,omitempty"`
	SerialNumber    string `json:"serial_number,omitempty" bson:"serial_number,omitempty"`
	Type            string `json:"type,omitempty" bson:"type,omitempty"`
	TypeDescription string `json:"type_description,omitempty" bson:"type_description,omitempty"`
	Vendor          bool   `json:"vendor,omitempty" bson:"vendor,omitempty"`
	Version         string `json:"version,omitempty" bson:"version,omitempty"`
}

type ChassiMongo struct {
	AboutChassiMongo AboutChassiMongo `json:"about_chassi,omitempty" bson:"about_chassi,omitempty"`
}

// **************************************************************************************
// **************************************************************************************
type AboutChassi struct {
	Tag             string
	SerialNumber    string
	Type            string
	TypeDescription bool
	Vendor          string
	Version         string
}

type Chassi struct {
	AboutChassi AboutChassi
}
