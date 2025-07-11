package FedwireFundsPaymentStatus

type PACS_002_001_VERSION string

const (
	PACS_002_001_03 PACS_002_001_VERSION = "PACS.002.001.03"
	PACS_002_001_04 PACS_002_001_VERSION = "PACS.002.001.04"
	PACS_002_001_05 PACS_002_001_VERSION = "PACS.002.001.05"
	PACS_002_001_06 PACS_002_001_VERSION = "PACS.002.001.06"
	PACS_002_001_07 PACS_002_001_VERSION = "PACS.002.001.07"
	PACS_002_001_08 PACS_002_001_VERSION = "PACS.002.001.08"
	PACS_002_001_09 PACS_002_001_VERSION = "PACS.002.001.09"
	PACS_002_001_10 PACS_002_001_VERSION = "PACS.002.001.10"
	PACS_002_001_11 PACS_002_001_VERSION = "PACS.002.001.11"
	PACS_002_001_12 PACS_002_001_VERSION = "PACS.002.001.12"
	PACS_002_001_13 PACS_002_001_VERSION = "PACS.002.001.13"
	PACS_002_001_14 PACS_002_001_VERSION = "PACS.002.001.14"
)

var VersionNameSpaceMap = map[PACS_002_001_VERSION]string{
	PACS_002_001_03: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.03",
	PACS_002_001_04: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.04",
	PACS_002_001_05: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.05",
	PACS_002_001_06: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.06",
	PACS_002_001_07: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.07",
	PACS_002_001_08: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08",
	PACS_002_001_09: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.09",
	PACS_002_001_10: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10",
	PACS_002_001_11: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.11",
	PACS_002_001_12: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12",
	PACS_002_001_13: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.13",
	PACS_002_001_14: "urn:iso:std:iso:20022:tech:xsd:pacs.002.001.14",
}
var NameSpaceVersionMap = map[string]PACS_002_001_VERSION{
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.03": PACS_002_001_03,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.04": PACS_002_001_04,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.05": PACS_002_001_05,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.06": PACS_002_001_06,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.07": PACS_002_001_07,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08": PACS_002_001_08,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.09": PACS_002_001_09,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.10": PACS_002_001_10,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.11": PACS_002_001_11,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.12": PACS_002_001_12,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.13": PACS_002_001_13,
	"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.14": PACS_002_001_14,
}
var VersionPathMap = map[PACS_002_001_VERSION]map[string]any{
	PACS_002_001_03: pathMapV3(),
	PACS_002_001_04: pathMapV4(),
	PACS_002_001_05: pathMapV5(),
	PACS_002_001_06: pathMapV6(),
	PACS_002_001_07: pathMapV7(),
	PACS_002_001_08: pathMapV8(),
	PACS_002_001_09: pathMapV9(),
	PACS_002_001_10: pathMapV10(),
	PACS_002_001_11: pathMapV11(),
	PACS_002_001_12: pathMapV12(),
	PACS_002_001_13: pathMapV13(),
	PACS_002_001_14: pathMapV14(),
}
