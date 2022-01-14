package entities

// **************************************************************************************
// **************************        MONGODB      ***************************************
// **************************************************************************************
type AboutGPUMongo struct {
	Name                 string `json:"name,omitempty" bson:"name,omitempty"`
	Vendor               string `json:"vendor,omitempty" bson:"vendor,omitempty"`
	Addrs                string `json:"addrs,omitempty" bson:"addrs,omitempty"`
	SerialNumber         string `json:"serial_number,omitempty" bson:"serial_number,omitempty"`
	Subsystem            string `json:"subsystem,omitempty" bson:"subsystem,omitempty"`
	Class                string `json:"class,omitempty" bson:"class,omitempty"`
	Subclass             bool   `json:"subclass,omitempty" bson:"subclass,omitempty"`
	ProgrammingInterface string `json:"programming_interface,omitempty" bson:"programming_interface,omitempty"`
}

type GPUMongo struct {
	AboutGPUMongo AboutGPUMongo `json:"about_gpu,omitempty" bson:"about_gpu,omitempty"`
}

// **************************************************************************************
// **************************************************************************************
type AboutGPU struct {
	Name                 string
	Vendor               string
	Addrs                string
	SerialNumber         string
	Subsystem            string
	Class                string
	Subclass             string
	ProgrammingInterface string
}

type GPU struct {
	AboutGPU AboutGPU
}
