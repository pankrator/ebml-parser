package tools

import "strings"

type DataType string

var (
	String   DataType = "string"
	Master   DataType = "master"
	UInteger DataType = "uinteger"
	Binary   DataType = "binary"
)

func init() {
	for hexID, el := range Schema {
		el.Hex = "0x" + strings.ToUpper(hexID)
	}
}

var Schema = map[string]*ElementData{
	"1a45dfa3": &ElementData{
		Name: "EBML",
		Typ:  Master,
	},
	"42f2": &ElementData{
		Name: "EBMLMaxIDLength",
		Typ:  UInteger,
	},
	"42f3": &ElementData{
		Name: "EBMLMaxSizeLength",
		Typ:  UInteger,
	},
	"4286": &ElementData{
		Name: "EBMLVersion",
		Typ:  UInteger,
	},
	"4282": &ElementData{
		Name: "DocType",
		Typ:  String,
	},
	"18538067": &ElementData{
		Name: "Segment",
		Typ:  Master,
	},
	"114d9b74": &ElementData{
		Name: "SeekHead",
		Typ:  Master,
	},
	"4dbb": &ElementData{
		Name: "Seek",
		Typ:  Master,
	},
	"53ab": &ElementData{
		Name: "SeekID",
		Typ:  Binary,
	},
	"53ac": &ElementData{
		Name: "SeekPosition",
		Typ:  UInteger,
	},
	"2ad7b1": &ElementData{
		Name: "TimestampScale",
		Typ:  UInteger,
	},
	"4489": &ElementData{
		Name: "Duration",
		// TODO: This is float!
		Typ: Binary,
	},
	"4461": &ElementData{
		Name: "DateUTC",
		// TODO: This is date!
		Typ: Binary,
	},
	"7ba9": &ElementData{
		Name: "Title",
		Typ:  String,
	},
	"4d80": &ElementData{
		Name: "MuxingApp",
		Typ:  String,
	},
	"55ee": &ElementData{
		Name: "MaxBlockAdditionID",
		Typ:  UInteger,
	},
	"5741": &ElementData{
		Name: "WritingApp",
		Typ:  String,
	},
	"e7": &ElementData{
		Name: "Timestamp",
		Typ:  UInteger,
	},
	"ab": &ElementData{
		Name: "PrevSize",
		Typ:  UInteger,
	},
	"a0": &ElementData{
		Name: "BlockGroup",
		Typ:  Master,
	},
	"a1": &ElementData{
		Name: "Block",
		Typ:  UInteger,
	},
	"75a1": &ElementData{
		Name: "BlockAdditions",
		Typ:  Master,
	},
	"a6": &ElementData{
		Name: "BlockMore",
		Typ:  Master,
	},
	"ee": &ElementData{
		Name: "BlockAddID",
		Typ:  UInteger,
	},
	"a5": &ElementData{
		Name: "BlockAdditional",
		Typ:  Binary,
	},
	"9b": &ElementData{
		Name: "BlockDuration",
		Typ:  UInteger,
	},
	"fb": &ElementData{
		Name: "ReferenceBlock",
		Typ:  UInteger,
	},
	"75a2": &ElementData{
		Name: "DiscardPadding",
		Typ:  UInteger,
	},
	"d7": &ElementData{
		Name: "TrackNumber",
		Typ:  UInteger,
	},
	"73c5": &ElementData{
		Name: "TrackUID",
		Typ:  UInteger,
	},
	"83": &ElementData{
		Name: "TrackType",
		Typ:  UInteger,
	},
	"b9": &ElementData{
		Name: "FlagEnabled",
		Typ:  UInteger,
	},
	"88": &ElementData{
		Name: "FlagDefault",
		Typ:  UInteger,
	},
	"55aa": &ElementData{
		Name: "FlagForced",
		Typ:  UInteger,
	},
	"9c": &ElementData{
		Name: "FlagLacing",
		Typ:  UInteger,
	},
	"23e383": &ElementData{
		Name: "DefaultDuration",
		Typ:  UInteger,
	},
	"536e": &ElementData{
		Name: "Name",
		Typ:  String,
	},
	"22b59c": &ElementData{
		Name: "Language",
		Typ:  String,
	},
	"86": &ElementData{
		Name: "CodecID",
		Typ:  String,
	},
	"63a2": &ElementData{
		Name: "CodecPrivate",
		Typ:  Binary,
	},
	"258688": &ElementData{
		Name: "CodecName",
		Typ:  String,
	},
	"56aa": &ElementData{
		Name: "CodecDelay",
		Typ:  UInteger,
	},
	"56bb": &ElementData{
		Name: "SeekPreRoll",
		Typ:  UInteger,
	},
	"9a": &ElementData{
		Name: "FlagInterlaced",
		Typ:  UInteger,
	},
	"53b8": &ElementData{
		Name: "StereoMode",
		Typ:  UInteger,
	},
	"53c0": &ElementData{
		Name: "AlphaMode",
		Typ:  UInteger,
	},
	"b0": &ElementData{
		Name: "PixelWidth",
		Typ:  UInteger,
	},
	"ba": &ElementData{
		Name: "PixelHeight",
		Typ:  UInteger,
	},
	"54aa": &ElementData{
		Name: "PixelCropBottom",
		Typ:  UInteger,
	},
	"54bb": &ElementData{
		Name: "PixelCropTop",
		Typ:  UInteger,
	},
	"54cc": &ElementData{
		Name: "PixelCropLeft",
		Typ:  UInteger,
	},
	"54dd": &ElementData{
		Name: "PixelCropRight",
		Typ:  UInteger,
	},
	"54b0": &ElementData{
		Name: "DisplayWidth",
		Typ:  UInteger,
	},
	"54ba": &ElementData{
		Name: "DisplayHeight",
		Typ:  UInteger,
	},
	"54b2": &ElementData{
		Name: "DisplayUnit",
		Typ:  UInteger,
	},
	"54b3": &ElementData{
		Name: "AspectRatioType",
		Typ:  UInteger,
	},
	"7670": &ElementData{
		Name: "Projection",
		Typ:  Master,
	},
	"7671": &ElementData{
		Name: "ProjectionType",
		Typ:  UInteger,
	},
	"7672": &ElementData{
		Name: "ProjectionPrivate",
		Typ:  Binary,
	},
	"7673": &ElementData{
		Name: "ProjectionPoseYaw",
		// TODO: Should be float
		Typ: Binary,
	},
	"7674": &ElementData{
		Name: "ProjectionPosePitch",
		// TODO: Should be float
		Typ: Binary,
	},
	"7675": &ElementData{
		Name: "ProjectionPoseRoll",
		// TODO: Should be float
		Typ: Binary,
	},
	"e1": &ElementData{
		Name: "Audio",
		Typ:  Master,
	},
	"b5": &ElementData{
		Name: "SamplingFrequency",
		// TODO: Should be float
		Typ: Binary,
	},
	"78b5": &ElementData{
		Name: "OutputSamplingFrequency",
		// TODO: Should be float
		Typ: Binary,
	},
	"9f": &ElementData{
		Name: "Channels",
		Typ:  UInteger,
	},
	"6264": &ElementData{
		Name: "BitDepth",
		Typ:  UInteger,
	},
	"6d80": &ElementData{
		Name: "ContentEncodings",
		Typ:  Master,
	},
	"6240": &ElementData{
		Name: "ContentEncoding",
		Typ:  Master,
	},
	"5031": &ElementData{
		Name: "ContentEncodingOrder",
		Typ:  UInteger,
	},
	"5032": &ElementData{
		Name: "ContentEncodingScope",
		Typ:  UInteger,
	},
	"5033": &ElementData{
		Name: "ContentEncodingType",
		Typ:  UInteger,
	},
	"5035": &ElementData{
		Name: "ContentEncryption",
		Typ:  Master,
	},
	"47e1": &ElementData{
		Name: "ContentEncAlgo",
		Typ:  UInteger,
	},
	"47e2": &ElementData{
		Name: "ContentEncKeyID",
		Typ:  Binary,
	},
	"47e7": &ElementData{
		Name: "ContentEncAESSettings",
		Typ:  Master,
	},
	"47e8": &ElementData{
		Name: "AESSettingsCipherMode",
		Typ:  UInteger,
	},
	"1549a966": &ElementData{
		Name: "Info",
		Typ:  Master,
	},
	"1f43b675": &ElementData{
		Name: "Cluster",
		Typ:  Master,
	},
	"1654ae6b": &ElementData{
		Name: "Tracks",
		Typ:  Master,
	},
	"ae": &ElementData{
		Name: "TrackEntry",
		Typ:  Master,
	},
	"e0": &ElementData{
		Name: "Video",
		Typ:  Master,
	},
	"a3": &ElementData{
		Name: "SimpleBlock",
		Typ:  Binary,
	},
}

type ElementData struct {
	Name string
	Hex  string
	Typ  DataType
}
