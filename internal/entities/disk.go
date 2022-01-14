package entities

// **************************************************************************************
// **************************        MONGODB      ***************************************
// **************************************************************************************
type PartitionsMongo struct {
	Name       string `json:"name,omitempty" bson:"name,omitempty" `
	SizeBytes  string `json:"size_bytes,omitempty" bson:"size_bytes,omitempty"`
	MountPoint string `json:"mount_point,omitempty" bson:"mount_point,omitempty"`
	Type       string `json:"type,omitempty" bson:"type,omitempty"`
	IsReadOnly bool   `json:"is_readonly,omitempty" bson:"is_readonly,omitempty"`
	Disk       string `json:"disk,omitempty" bson:"disk,omitempty"`
	UUID       string `json:"uuid,omitempty" bson:"uuid,omitempty"`
}
type AboutDiskMongo struct {
	Name                   string `json:"name,omitempty" bson:"name,omitempty"`
	SizeBytes              string `json:"size_bytes,omitempty" bson:"size_bytes,omitempty"`
	PhysicalBlockSizeBytes string `json:"physical_block_size_bytes,omitempty" bson:"physical_block_size_bytes,omitempty"`
	IsRemovable            bool   `json:"is_removable,omitempty" bson:"is_removable,omitempty"`
	DriveType              string `json:"drive_type,omitempty" bson:"drive_type,omitempty"`
	StorageController      string `json:"storage_controller,omitempty" bson:"storage_controller,omitempty"`
	Vendor                 string `json:"vendor,omitempty" bson:"vendor,omitempty"`
	Model                  string `json:"model,omitempty" bson:"model,omitempty"`
	SerialNumber           string `json:"serial_number,omitempty" bson:"serial_number,omitempty"`
	WWN                    string `json:"wwn,omitempty" bson:"wwn,omitempty"`
	Partitions             string `json:"partitions,omitempty" bson:"partitions,omitempty"`
}

type DiskMongo struct {
	AboutCPUMongo  AboutCPUMongo  `json:"about_cpu,omitempty" bson:"about_cpu,omitempty"`
	ProcessorMongo ProcessorMongo `json:"processor,omitempty" bson:"processor,omitempty"`
}

// **************************************************************************************
// **************************************************************************************
type Partitions struct {
	Name       string
	SizeBytes  string
	MountPoint string
	Type       string
	IsReadOnly bool
	Disk       string
	UUID       string
}

type AboutDisk struct {
	Name                   string
	SizeBytes              string
	PhysicalBlockSizeBytes string
	IsRemovable            bool
	DriveType              string
	StorageController      string
	Vendor                 string
	Model                  string
	SerialNumber           string
	WWN                    string
	Partitions             string
}

type Disk struct {
	AboutDisk  AboutDisk
	Partitions []Partitions
}
