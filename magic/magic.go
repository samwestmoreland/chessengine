package magic

import (
	_ "embed"
)

type Entry struct {
	Square string `json:"square"`
	Magic  string `json:"magic"` // as hex string
	Shift  int    `json:"shift"`
	Mask   string `json:"mask"` // as hex string
}

type Data struct {
	Rook     RookData   `json:"rook"`
	Bishop   BishopData `json:"bishop"`
	Metadata Metadata   `json:"metadata"`
}

type RookData struct {
	Magics         []Entry `json:"magics"`
	TotalTableSize string  `json:"total_table_size"`
}

type BishopData struct {
	Magics         []Entry `json:"magics"`
	TotalTableSize string  `json:"total_table_size"`
}

type Metadata struct {
	TotalTableSize string `json:"total_table_size"`
	Generated      string `json:"generated"`
	Version        string `json:"version"`
}

var PrecalculatedData Data

//go:embed version.txt
var VersionString string

//go:embed magics.json
var JSONData []byte // The raw JSON data
