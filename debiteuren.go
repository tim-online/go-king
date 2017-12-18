package king

import "encoding/xml"

type Debiteuren []Debiteur

func (d Debiteuren) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type alias Debiteuren
	type Document struct {
		XMLName xml.Name `xml:"KING_DEBITEUREN"`

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
	XMLName xml.Name `xml:"DEBITEUR"`

	NawNummer                    int                    `xml:"NAW_NUMMER,omitempty"`
	NawZoekcode                  string                 `xml:"NAW_ZOEKCODE"`
	NawBetalingsConditie         int                    `xml:"NAW_BETALINGSCONDITIE,omitempty"`
	NawLeveringsConditie         string                 `xml:"NAW_LEVERINGSCONDITIE,omitempty"`
	NawValutacode                int                    `xml:"NAW_VALUTACODE,omitempty"`
	NawBtwcode                   int                    `xml:"NAW_BTWCODE,omitempty"`
	NawTaalcode                  int                    `xml:"NAW_TAALCODE,omitempty"`
	NawKredietlimiet             Decimal                `xml:"NAW_KREDIETLIMIET,omitempty"`
	NawTegenrekening             string                 `xml:"NAW_TEGENREKENING"`
	NawAanmaningstype            AanmaningsType         `xml:"NAW_AANMANINGSTYPE"`
	NawBtwnummer                 string                 `xml:"NAW_BTWNUMMER"`
	NawDeelleveringToegestaan    DeelleveringToegestaan `xml:"NAW_DEELLEVERINGTOEGESTAAN,omitempty"`
	NawApartefacturen            AparteFacturen         `xml:"NAW_APARTEFACTUREN,omitempty"`
	NawFacturereninexbtw         FacturenInExBtw        `xml:"NAW_FACTURERENINEXBTW"`
	NawFactuuradressoort         FactuuradresSoort      `xml:"NAW_FACTUURADRESSOORT,omitempty"`
	NawVerzendadressoort         VerzendadresSoort      `xml:"NAW_VERZENDADRESSOORT,omitempty"`
	NawVerzendadresnummer        int                    `xml:"NAW_VERZENDADRESNUMMER,omitempty"`
	NawEindbestemmingsoort       string                 `xml:"NAW_EINDBESTEMMINGSOORT"`
	NawEindbestemmingadresnummer string                 `xml:"NAW_EINDBESTEMMINGADRESNUMMER"`
	NawVertegenwoordiger         int                    `xml:"NAW_VERTEGENWOORDIGER,omitempty"`
	NawDebiteurgroep             int                    `xml:"NAW_DEBITEURGROEP,omitempty"`
	NawAantalkopiefacturen       int                    `xml:"NAW_AANTALKOPIEFACTUREN,omitempty"`
	NawOrderkortingsoort         OrderkortingSoort      `xml:"NAW_ORDERKORTINGSOORT"`
	NawOrderkorting              int                    `xml:"NAW_ORDERKORTING,omitempty"`
	NawContributiecode           string                 `xml:"NAW_CONTRIBUTIECODE"`
	NawBlokkeerorderinvoer       string                 `xml:"NAW_BLOKKEERORDERINVOER"`
	NawDefaultVervoerder         string                 `xml:"NAW_DEFAULT_VERVOERDER,omitempty"`
	NawWebklant                  bool                   `xml:"NAW_WEBKLANT,omitempty"`
	NawWebwinkels                []Webwinkel            `xml:"NAW_WEBWINKELS"`
	NawKvknummer                 string                 `xml:"NAW_KVKNUMMER"`
	NawInkoopcombinatie          bool                   `xml:"NAW_INKOOPCOMBINATIE"`
	NawDebiteurnummerfactuur     int                    `xml:"NAW_DEBITEURNUMMERFACTUUR,omitempty"`
	NawDebiteurnummeromzet       int                    `xml:"NAW_DEBITEURNUMMEROMZET,omitempty"`
	NawDebiteurnummerprijzen     int                    `xml:"NAW_DEBITEURNUMMERPRIJZEN,omitempty"`
	NawOpmerking                 string                 `xml:"NAW_OPMERKING"`
	NawWebsite                   string                 `xml:"NAW_WEBSITE"`
	NawVestigingadres            Adres                  `xml:"NAW_VESTIGINGADRES"`
	NawCorrespondentieadres      Adres                  `xml:"NAW_CORRESPONDENTIEADRES"`
	NawContactpersonen           []ContactPersoon       `xml:"NAW_CONTACTPERSONEN"`
	NawVerzendadressen           []VerzendAdres         `xml:"NAW_VERZENDADRESSEN"`
	// NawSelecties                 []Selectie             `xml:"NAW_SELECTIES"`
	// NawBankrekeningen            []Bankrekening         `xml:"NAW_BANKREKENINGEN"`
	// nawDocumenten                []Documenten           `xml:"NAW_DOCUMENTEN"`
	// NawVrijerubrieken            []Rubriek              `xml:"NAW_VRIJERUBRIEKEN"`
}

type Webwinkel struct {
}

type Adres struct {
	AdrNaam1      string `xml:"ADR_NAAM1"`
	AdrNaam2      string `xml:"ADR_NAAM2"`
	AdrStraat     string `xml:"ADR_STRAAT"`
	AdrHuisnummer string `xml:"ADR_HUISNUMMER"`
	AdrPostcode   string `xml:"ADR_POSTCODE"`
	AdrWoonplaats string `xml:"ADR_WOONPLAATS"`
	AdrLand       string `xml:"ADR_LAND"`
	AdrEmail      string `xml:"ADR_EMAIL"`
	AdrTelefoon   string `xml:"ADR_TELEFOON"`
	AdrTelefoon2  string `xml:"ADR_TELEFOON2"`
	AdrTelefax    string `xml:"ADR_TELEFAX"`
	AdrEan        string `xml:"ADR_EAN"`
}

type ContactPersoon struct {
}

type VerzendAdres struct {
}
