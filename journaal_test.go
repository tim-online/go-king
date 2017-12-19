package king_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/aodin/date"
	king "github.com/tim-online/go-king"
)

func TestJournaalXsd(t *testing.T) {
	journaal := king.Journaal{
		Boekingsgangen: king.Boekingsgangen{
			king.Boekingsgang{
				Journaalposten: king.Journaalposten{
					king.Journaalpost{
						DagboekCode:  "Verkoo",
						Boekdatum:    king.Date{date.New(1983, 4, 12)},
						Omschrijving: "",
						Journaalregels: king.Journaalregels{
							king.Journaalregel{
								Volgnummer:       0,
								Rekeningnummer:   "8010",
								Boekzijde:        king.BoekzijdeCredit,
								Valutacode:       "EUR",
								Valutabedrag:     king.Money(5.12),
								Omschrijving:     "Overnachting",
								Factuurnummer:    "2017-11-19",
								Factuurdatum:     king.Date{},
								Vervaldatum:      king.Date{},
								Betalingskenmerk: "",
								Hulprekening: king.Hulprekening{
									Soort:        king.HulpSoortBtw,
									Btwcode:      king.Btwcode(99),
									Boekzijde:    king.BoekzijdeDebet,
									Valutacode:   "EUR",
									Valutabedrag: king.Money(12.0),
								},
							},
						},
					},
				},
			},
		},
	}
	xml, err := xml.MarshalIndent(journaal, "", "\t")
	if err != nil {
		t.Error(err)
	}

	xmlReader := bytes.NewReader(xml)
	err = validate(xmlReader, "xsds/KingJournaal.xsd")
	if err != nil {
		t.Error(err)
	}
}
