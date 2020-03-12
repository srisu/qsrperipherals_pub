package models

type Values struct {
	Configs Configs `yaml:"configs" json:"configs"`
}
type Configs struct {
	Weighingscale Weighingscale `yaml:"weighingscale" json:"weighingScale"`
	Poledisplay       Poledisplay       `yaml:"poledisplay" json:"poleDisplay"` 
}

type Weighingscale struct
{
	Port string `yaml:"port" json:"port"` 
	Baud int `yaml:"baud" json:"baud"`
	Parity int `yaml:"parity" json:"parity"`
	DataBits int `yaml:"dataBits" json:"dataBits"`
	StopBits int `yaml:"stopBits" json:"stopBits"`
	Position Position `yaml:"position" json:"position"`
}

type Poledisplay struct
{
	Port string `yaml:"port" json:"port"` 
	Baud int `yaml:"baud" json:"baud"`
}

type Position struct {
	Start int `yaml:"start" json:"start"`
	Length int `yaml:"length" json:"length"`
	Prefix string `yaml:"prefix" json:"prefix"`
	Suffix string `yaml:"suffix" json:"suffix"`
}
