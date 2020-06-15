package tools

type DataType string

var (
	String   DataType = "string"
	Master   DataType = "master"
	UInteger DataType = "uinteger"
	Binary   DataType = "binary"
)

var Schema = map[string]ElementData{
	"1a45dfa3": ElementData{
		Name: "EBML",
		Typ:  Master,
		Hex:  "0x1a45dfa3",
	},
	"4286": ElementData{
		Name: "EBMLVersion",
		Typ:  UInteger,
		Hex:  "0x4286",
	},
	"4282": ElementData{
		Name: "DocType",
		Typ:  String,
		Hex:  "0x4282",
	},
	"18538067": ElementData{
		Name: "Segment",
		Typ:  Master,
		Hex:  "0x18538067",
	},
	"114d9b74": ElementData{
		Name: "SeekHead",
		Typ:  Master,
		Hex:  "0x114D9B74",
	},
	"4dbb": ElementData{
		Name: "Seek",
		Typ:  Master,
		Hex:  "0x4DBB",
	},
	"1549a966": ElementData{
		Name: "Info",
		Typ:  Master,
		Hex:  "0x1549A966",
	},
	"1f43b675": ElementData{
		Name: "Cluster",
		Typ:  Master,
		Hex:  "0x1F43B675",
	},
	"1654ae6b": ElementData{
		Name: "Tracks",
		Typ:  Master,
		Hex:  "0x1654AE6B",
	},
	"ae": ElementData{
		Name: "TrackEntry",
		Typ:  Master,
		Hex:  "0xAE",
	},
	"e0": ElementData{
		Name: "Video",
		Typ:  Master,
		Hex:  "0xE0",
	},
	"a3": ElementData{
		Name: "SimpleBlock",
		Typ:  Binary,
		Hex:  "0xA3",
	},
}

type ElementData struct {
	Name string
	Hex  string
	Typ  DataType
}
