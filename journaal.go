package king

import (
	"encoding/xml"
	"fmt"

	"github.com/tim-online/go-king/omitempty"
)

type Journaal struct {
	XMLName xml.Name `xml:"KING_JOURNAAL" json:"-"`

	Boekingsgangen Boekingsgangen `xml:"BOEKINGSGANGEN>BOEKINGSGANG"`
}

type Boekingsgangen []Boekingsgang

type Boekingsgang struct {
	Omschrijving   string         `xml:"BG_OMSCHRIJVING"`
	Definitief     bool           `xml:"BG_DEFINITIEF"`
	Journaalposten Journaalposten `xml:"JOURNAALPOSTEN>JOURNAALPOST,omitempty"`
}

type Journaalposten []Journaalpost

type Journaalpost struct {
	DagboekCode    string         `xml:"JP_DAGBOEKCODE"`
	Boekdatum      Date           `xml:"JP_BOEKDATUM"`
	Stuknummer     string         `xml:"JP_STUKNUMMER"`
	Omschrijving   string         `xml:"JP_OMSCHRIJVING"`
	Journaalregels Journaalregels `xml:"JOURNAALREGELS>JOURNAALREGEL,omitempty"`
}

type Journaalregels []Journaalregel

type Journaalregel struct {
	Volgnummer                             int          `xml:"JR_VOLGNUMMER,omitempty"`
	Rekeningnummer                         string       `xml:"JR_REKENINGNUMMER"`
	Boekdatum                              Date         `xml:"JR_BOEKDATUM"`
	Boekzijde                              Boekzijde    `xml:"JR_BOEKZIJDE"`
	Valutacode                             string       `xml:"JR_VALUTACODE"`
	Valutabedrag                           Money        `xml:"JR_VALUTABEDRAG"`
	Omschrijving                           string       `xml:"JR_OMSCHRIJVING"`
	Factuurnummer                          string       `xml:"JR_FACTUURNUMMER,omitempty"`
	Factuurdatum                           Date         `xml:"JR_FACTUURDATUM"`
	Vervaldatum                            Date         `xml:"JR_VERVALDATUM"`
	Betalingskenmerk                       string       `xml:"JR_BETALINGSKENMERK,omitempty"`
	Aantal                                 float64      `xml:"JR_AANTAL,omitempty"`
	ArchiefstukNummer                      string       `xml:"JR_ARCHIEFSTUK_NUMMER,omitempty"`
	ArchiefstukExternID                    string       `xml:"JR_ARCHIEFSTUK_EXTERN_ID,omitempty"`
	OpenstaandePostGefiatteerd             bool         `xml:"JR_OPENSTAANDE_POST_GEFIATTEERD,omitempty"`
	OpenstaandePostFiatteringGewijzigdDoor string       `xml:"JR_OPENSTAANDE_POST_FIATTERING_GEWIJZIGD_DOOR,omitempty"`
	OpenstaandePostFiatteringGewijzigdOp   Time         `xml:"JR_OPENSTAANDE_POST_FIATTERING_GEWIJZIGD_OP,omitempty"`
	Hulprekening                           Hulprekening `xml:"HULPREKENING,omitempty"`
}

func (j Journaalregel) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(j, e, start)
}

const (
	BoekzijdeCredit Boekzijde = "CRED"
	BoekzijdeDebet  Boekzijde = "DEB"
)

type Boekzijde string

type Hulprekening struct {
	Soort          HulpSoort `xml:"HULP_SOORT"`
	Btwcode        Btwcode   `xml:"HULP_BTWCODE"`
	Rekeningnummer string    `xml:"HULP_REKENINGNUMMER"`
	Boekzijde      Boekzijde `xml:"HULP_BOEKZIJDE"`
	Valutacode     string    `xml:"HULP_VALUTACODE"`
	Valutabedrag   Money     `xml:"HULP_VALUTABEDRAG"`
}

func (h Hulprekening) IsEmpty() bool {
	return h.Soort == "" && h.Btwcode == 0 && h.Rekeningnummer == "" && h.Boekzijde == "" && h.Valutacode == "" && h.Valutabedrag == 0.0
}

type Btwcode int

func (b Btwcode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	s := fmt.Sprintf("%03d", b)
	return e.EncodeElement(s, start)
}

const (
	HulpSoortBtw     HulpSoort = "BTW"
	HulpSoortBetvs   HulpSoort = "BETVS"
	HulpSoortKrsvs   HulpSoort = "KRSVS"
	HulpSoortKredbep HulpSoort = "KREDBEP"
	HulpSoortBetkort HulpSoort = "BETKORT"
)

type HulpSoort string
