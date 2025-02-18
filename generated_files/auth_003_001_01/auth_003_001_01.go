// Code generated by ___go_build_github_com_dminGod_sepaToGo. DO NOT EDIT.

package iso20022_auth_003_001_01

type Document struct {
	InfReqStsChngNtfctn InformationRequestStatusChangeNotificationV01 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.003.001.01 InfReqStsChngNtfctn"`
}

type InformationRequestStatusChangeNotificationV01 struct {
	OrgnlBizQry Max35Text            `xml:"urn:iso:std:iso:20022:tech:xsd:auth.003.001.01 OrgnlBizQry"`
	CnfdtltySts bool                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.003.001.01 CnfdtltySts"`
	SplmtryData []SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.003.001.01 SplmtryData,omitempty"`
}

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.003.001.01 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:auth.003.001.01 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}
