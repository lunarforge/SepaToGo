// Code generated by ___go_build_github_com_dminGod_sepaToGo. DO NOT EDIT.

package iso20022_camt_051_001_05

import (
	"bytes"
	"encoding/xml"
	"time"
)

type AccountIdentification4Choice struct {
	IBAN IBAN2007Identifier            `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 IBAN"`
	Othr GenericAccountIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Othr"`
}

type AccountSchemeName1Choice struct {
	Cd    ExternalAccountIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Cd"`
	Prtry Max35Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Prtry"`
}

type ActiveCurrencyAndAmount struct {
	Value float64            `xml:",chardata"`
	Ccy   ActiveCurrencyCode `xml:"Ccy,attr"`
}

// Must match the pattern [A-Z]{3,3}
type ActiveCurrencyCode string

// Must match the pattern [A-Z]{3,3}
type ActiveOrHistoricCurrencyCode string

// May be one of ADDR, PBOX, HOME, BIZZ, MLTO, DLVY
type AddressType2Code string

type AddressType3Choice struct {
	Cd    AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Cd"`
	Prtry GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Prtry"`
}

type Amount2Choice struct {
	AmtWthtCcy float64                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 AmtWthtCcy"`
	AmtWthCcy  ActiveCurrencyAndAmount `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 AmtWthCcy"`
}

// Must match the pattern [A-Z0-9]{4,4}[A-Z]{2,2}[A-Z0-9]{2,2}([A-Z0-9]{3,3}){0,1}
type BICFIDec2014Identifier string

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 FinInstnId"`
	BrnchId    BranchData3                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Id,omitempty"`
	LEI     LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 LEI,omitempty"`
	Nm      Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Nm,omitempty"`
	PstlAdr PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 PstlAdr,omitempty"`
}

type CashAccount38 struct {
	Id   AccountIdentification4Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Id"`
	Tp   CashAccountType2Choice       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Tp,omitempty"`
	Ccy  ActiveOrHistoricCurrencyCode `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Ccy,omitempty"`
	Nm   Max70Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Nm,omitempty"`
	Prxy ProxyAccountIdentification1  `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Prxy,omitempty"`
}

type CashAccountType2Choice struct {
	Cd    ExternalCashAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Cd"`
	Prtry Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Prtry"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Cd"`
	Prtry Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Prtry"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 ClrSysId,omitempty"`
	MmbId    Max35Text                           `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 MmbId"`
}

// Must match the pattern [A-Z]{2,2}
type CountryCode string

type Document struct {
	LqdtyDbtTrf LiquidityDebitTransferV05 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 LqdtyDbtTrf"`
}

// Must match the pattern [a-zA-Z0-9]{4}
type Exact4AlphaNumericText string

// Must be at least 1 items long
type ExternalAccountIdentification1Code string

// Must be at least 1 items long
type ExternalCashAccountType1Code string

// Must be at least 1 items long
type ExternalClearingSystemIdentification1Code string

// Must be at least 1 items long
type ExternalFinancialInstitutionIdentification1Code string

// Must be at least 1 items long
type ExternalProxyAccountType1Code string

type FinancialIdentificationSchemeName1Choice struct {
	Cd    ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Cd"`
	Prtry Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Prtry"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 BICFI,omitempty"`
	ClrSysMmbId ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 ClrSysMmbId,omitempty"`
	LEI         LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 LEI,omitempty"`
	Nm          Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Nm,omitempty"`
	PstlAdr     PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 PstlAdr,omitempty"`
	Othr        GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Othr,omitempty"`
}

type GenericAccountIdentification1 struct {
	Id      Max34Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Id"`
	SchmeNm AccountSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Issr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Id"`
	SchmeNm FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 SchmeNm,omitempty"`
	Issr    Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Issr"`
	SchmeNm Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 SchmeNm,omitempty"`
}

// Must match the pattern [A-Z]{2,2}[0-9]{2,2}[a-zA-Z0-9]{1,30}
type IBAN2007Identifier string

type ISODate time.Time

func (t *ISODate) UnmarshalText(text []byte) error {
	return (*xsdDate)(t).UnmarshalText(text)
}
func (t ISODate) MarshalText() ([]byte, error) {
	return xsdDate(t).MarshalText()
}

type ISODateTime time.Time

func (t *ISODateTime) UnmarshalText(text []byte) error {
	return (*xsdDateTime)(t).UnmarshalText(text)
}
func (t ISODateTime) MarshalText() ([]byte, error) {
	return xsdDateTime(t).MarshalText()
}

// Must match the pattern [A-Z0-9]{18,18}[0-9]{2,2}
type LEIIdentifier string

type LiquidityDebitTransfer2 struct {
	LqdtyTrfId PaymentIdentification8                       `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 LqdtyTrfId,omitempty"`
	Cdtr       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Cdtr,omitempty"`
	CdtrAcct   CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 CdtrAcct,omitempty"`
	TrfdAmt    Amount2Choice                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 TrfdAmt"`
	Dbtr       BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Dbtr,omitempty"`
	DbtrAcct   CashAccount38                                `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 DbtrAcct,omitempty"`
	SttlmDt    ISODate                                      `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 SttlmDt,omitempty"`
}

type LiquidityDebitTransferV05 struct {
	MsgHdr      MessageHeader1          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 MsgHdr"`
	LqdtyDbtTrf LiquidityDebitTransfer2 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 LqdtyDbtTrf"`
	SplmtryData []SupplementaryData1    `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 SplmtryData,omitempty"`
}

// Must be at least 1 items long
type Max140Text string

// Must be at least 1 items long
type Max16Text string

// Must be at least 1 items long
type Max2048Text string

// Must be at least 1 items long
type Max34Text string

// Must be at least 1 items long
type Max350Text string

// Must be at least 1 items long
type Max35Text string

// Must be at least 1 items long
type Max70Text string

type MessageHeader1 struct {
	MsgId   Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 MsgId"`
	CreDtTm ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 CreDtTm,omitempty"`
}

type PaymentIdentification8 struct {
	InstrId    Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 InstrId,omitempty"`
	EndToEndId Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 EndToEndId"`
	TxId       Max35Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 TxId,omitempty"`
	UETR       UUIDv4Identifier `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 UETR,omitempty"`
}

type PostalAddress24 struct {
	AdrTp       AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 AdrTp,omitempty"`
	Dept        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Dept,omitempty"`
	SubDept     Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 SubDept,omitempty"`
	StrtNm      Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 StrtNm,omitempty"`
	BldgNb      Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 BldgNb,omitempty"`
	BldgNm      Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 BldgNm,omitempty"`
	Flr         Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Flr,omitempty"`
	PstBx       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 PstBx,omitempty"`
	Room        Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Room,omitempty"`
	PstCd       Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 PstCd,omitempty"`
	TwnNm       Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 TwnNm,omitempty"`
	TwnLctnNm   Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 TwnLctnNm,omitempty"`
	DstrctNm    Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 DstrctNm,omitempty"`
	CtrySubDvsn Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 CtrySubDvsn,omitempty"`
	Ctry        CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Ctry,omitempty"`
	AdrLine     []Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 AdrLine,omitempty"`
}

type ProxyAccountIdentification1 struct {
	Tp ProxyAccountType1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Tp,omitempty"`
	Id Max2048Text             `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Id"`
}

type ProxyAccountType1Choice struct {
	Cd    ExternalProxyAccountType1Code `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Cd"`
	Prtry Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Prtry"`
}

type SupplementaryData1 struct {
	PlcAndNm Max350Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 PlcAndNm,omitempty"`
	Envlp    SupplementaryDataEnvelope1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.051.001.05 Envlp"`
}

type SupplementaryDataEnvelope1 struct {
	Item string `xml:",any"`
}

// Must match the pattern [a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}
type UUIDv4Identifier string

type xsdDate time.Time

func (t *xsdDate) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02")
}
func (t xsdDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
func _marshalTime(t time.Time, format string) ([]byte, error) {
	return []byte(t.Format(format + "Z07:00")), nil
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return _marshalTime((time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
