package entities

// **************************************************************************************
// **************************        MONGODB      ***************************************
// **************************************************************************************
type AboutCacheMemoryMongo struct {
	Type              string `json:"type,omitempty" bson:"type,omitempty"`
	UNIFIED           string `json:"unified,omitempty" bson:"unified,omitempty"`
	Level             string `json:"level,omitempty" bson:"level,omitempty"`
	SizeBytes         bool   `json:"size_bytes,omitempty" bson:"size_bytes,omitempty"`
	LogicalProcessors string `json:"logical_processors,omitempty" bson:"logical_processors,omitempty"`
}

type CacheMemoryMongo struct {
	AboutCacheMemoryMongo AboutCacheMemoryMongo `json:"about_cache,omitempty" bson:"about_cache,omitempty"`
}

// **************************************************************************************
// **************************************************************************************
type AboutCacheMemory struct {
	Type              string
	UNIFIED           string
	Level             string
	SizeBytes         bool
	LogicalProcessors string
}

type CacheMemory struct {
	AboutCacheMemory AboutCacheMemory
}
