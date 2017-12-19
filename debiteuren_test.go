package king_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	king "github.com/tim-online/go-king"
)

func TestDebiteurenXsd(t *testing.T) {
	debiteuren := king.Debiteuren{
		king.Debiteur{
			NawNummer: 18524515,
		},
		king.Debiteur{
			NawNummer: 12,
		},
	}
	xml, err := xml.MarshalIndent(debiteuren, "", "\t")
	if err != nil {
		t.Error(err)
	}

	xmlReader := bytes.NewReader(xml)
	err = validate(xmlReader, "xsds/KingDebiteuren.xsd")
	if err != nil {
		t.Error(err)
	}
}
