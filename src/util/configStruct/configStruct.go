package configStruct

type ConfigStruct struct {
	Version     string `json:"version"`
	Downloadurl string `json:"downloadUrl"`
}

type ConfigFileStruct struct {
	id     string       `json:"id"`
	Config ConfigStruct `json:"config"`
}
