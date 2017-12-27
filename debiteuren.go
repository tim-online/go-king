package king

import (
	"encoding/xml"

	"github.com/tim-online/go-king/omitempty"
)

type Debiteuren []Debiteur

func (d Debiteuren) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias Debiteuren
	type Document struct {
		XMLName xml.Name `xml:"KING_DEBITEUREN" json:"-"`

		Debiteuren alias `xml:"DEBITEUREN>DEBITEUR"`
	}

	doc := Document{Debiteuren: alias(d)}
	return e.Encode(doc)
}

type DeelleveringToegestaan string
type AanmaningsType string
type AparteFacturen string
type FacturenInExBtw string

type FactuuradresSoort string
type VerzendadresSoort string
type OrderkortingSoort string

type Debiteur struct {
	XMLName xml.Name `xml:"DEBITEUR" json:"-"`

	Nummer                    int                    `xml:"NAW_NUMMER,omitempty" json:",omitempty"`
	Zoekcode                  string                 `xml:"NAW_ZOEKCODE"`
	BetalingsConditie         int                    `xml:"NAW_BETALINGSCONDITIE,omitempty" json:",omitempty"`
	LeveringsConditie         string                 `xml:"NAW_LEVERINGSCONDITIE,omitempty" json:",omitempty"`
	Valutacode                int                    `xml:"NAW_VALUTACODE,omitempty" json:",omitempty"`
	Btwcode                   int                    `xml:"NAW_BTWCODE,omitempty" json:",omitempty"`
	Taalcode                  int                    `xml:"NAW_TAALCODE,omitempty" json:",omitempty"`
	Kredietlimiet             Decimal                `xml:"NAW_KREDIETLIMIET,omitempty" json:",omitempty"`
	Tegenrekening             string                 `xml:"NAW_TEGENREKENING,omitempty" json:",omitempty"`
	Aanmaningstype            AanmaningsType         `xml:"NAW_AANMANINGSTYPE,omitempty" json:",omitempty"`
	Btwnummer                 string                 `xml:"NAW_BTWNUMMER,omitempty" json:",omitempty"`
	DeelleveringToegestaan    DeelleveringToegestaan `xml:"NAW_DEELLEVERINGTOEGESTAAN,omitempty" json:",omitempty"`
	Apartefacturen            AparteFacturen         `xml:"NAW_APARTEFACTUREN,omitempty" json:",omitempty"`
	Facturereninexbtw         FacturenInExBtw        `xml:"NAW_FACTURERENINEXBTW,omitempty" json:",omitempty"`
	Factuuradressoort         FactuuradresSoort      `xml:"NAW_FACTUURADRESSOORT,omitempty" json:",omitempty"`
	Verzendadressoort         VerzendadresSoort      `xml:"NAW_VERZENDADRESSOORT,omitempty" json:",omitempty"`
	Verzendadresnummer        int                    `xml:"NAW_VERZENDADRESNUMMER,omitempty" json:",omitempty"`
	Eindbestemmingsoort       string                 `xml:"NAW_EINDBESTEMMINGSOORT,omitempty" json:",omitempty"`
	Eindbestemmingadresnummer string                 `xml:"NAW_EINDBESTEMMINGADRESNUMMER,omitempty" json:",omitempty"`
	Vertegenwoordiger         int                    `xml:"NAW_VERTEGENWOORDIGER,omitempty" json:",omitempty"`
	Debiteurgroep             int                    `xml:"NAW_DEBITEURGROEP,omitempty" json:",omitempty"`
	Aantalkopiefacturen       int                    `xml:"NAW_AANTALKOPIEFACTUREN,omitempty" json:",omitempty"`
	Orderkortingsoort         OrderkortingSoort      `xml:"NAW_ORDERKORTINGSOORT,omitempty" json:",omitempty"`
	Orderkorting              int                    `xml:"NAW_ORDERKORTING,omitempty" json:",omitempty"`
	Contributiecode           string                 `xml:"NAW_CONTRIBUTIECODE" json:",omitempty"`
	Blokkeerorderinvoer       string                 `xml:"NAW_BLOKKEERORDERINVOER,omitempty" json:",omitempty"`
	DefaultVervoerder         string                 `xml:"NAW_DEFAULT_VERVOERDER,omitempty" json:",omitempty"`
	Webklant                  bool                   `xml:"NAW_WEBKLANT,omitempty" json:",omitempty"`
	Webwinkels                []Webwinkel            `xml:"NAW_WEBWINKELS,omitempty" json:",omitempty"`
	Kvknummer                 string                 `xml:"NAW_KVKNUMMER,omitempty" json:",omitempty"`
	Inkoopcombinatie          bool                   `xml:"NAW_INKOOPCOMBINATIE"`
	Debiteurnummerfactuur     int                    `xml:"NAW_DEBITEURNUMMERFACTUUR,omitempty" json:",omitempty"`
	Debiteurnummeromzet       int                    `xml:"NAW_DEBITEURNUMMEROMZET,omitempty" json:",omitempty"`
	Debiteurnummerprijzen     int                    `xml:"NAW_DEBITEURNUMMERPRIJZEN,omitempty" json:",omitempty"`
	Opmerking                 string                 `xml:"NAW_OPMERKING"`
	Website                   string                 `xml:"NAW_WEBSITE,omitempty" json:",omitempty"`
	Vestigingadres            Adres                  `xml:"NAW_VESTIGINGADRES,omitempty" json:",omitempty"`
	Correspondentieadres      Adres                  `xml:"NAW_CORRESPONDENTIEADRES,omitempty" json:",omitempty"`
	Contactpersonen           []ContactPersoon       `xml:"NAW_CONTACTPERSONEN,omitempty" json:",omitempty"`
	Verzendadressen           []Adres                `xml:"NAW_VERZENDADRESSEN,omitempty" json:",omitempty"`
	// Selecties                 []Selectie             `xml:"NAW_SELECTIES"`
	// Bankrekeningen            []Bankrekening         `xml:"NAW_BANKREKENINGEN"`
	// nawDocumenten                []Documenten           `xml:"NAW_DOCUMENTEN"`
	// Vrijerubrieken            []Rubriek              `xml:"NAW_VRIJERUBRIEKEN"`
}

func (d Debiteur) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return omitempty.MarshalXML(d, e, start)
}

func (d Debiteur) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(d)
}

type Webwinkel struct {
}

type Adres struct {
	Naam1      string `xml:"ADR_NAAM1"`
	Naam2      string `xml:"ADR_NAAM2"`
	Straat     string `xml:"ADR_STRAAT"`
	Huisnummer string `xml:"ADR_HUISNUMMER"`
	Postcode   string `xml:"ADR_POSTCODE"`
	Woonplaats string `xml:"ADR_WOONPLAATS"`
	Land       string `xml:"ADR_LAND"`
	Email      string `xml:"ADR_EMAIL"`
	Telefoon   string `xml:"ADR_TELEFOON"`
	Telefoon2  string `xml:"ADR_TELEFOON2"`
	Telefax    string `xml:"ADR_TELEFAX"`
	Ean        string `xml:"ADR_EAN,omitempty" json:",omitempty"`
}

func (a Adres) IsEmpty() bool {
	return a.Naam1 == "" &&
		a.Naam2 == "" &&
		a.Straat == "" &&
		a.Huisnummer == "" &&
		a.Postcode == "" &&
		a.Woonplaats == "" &&
		a.Land == "" &&
		a.Email == "" &&
		a.Telefoon == "" &&
		a.Telefoon2 == "" &&
		a.Telefax == "" &&
		a.Ean == ""
}

type ContactPersoon struct {
}
