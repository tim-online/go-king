package king_test

import (
	"bytes"
	"encoding/xml"
	"errors"
	"io"
	"os/exec"
	"testing"

	king "github.com/tim-online/go-king"
)

func TestDing(t *testing.T) {
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

func validate(xml io.Reader, xsd string) error {
	cmd := exec.Command("xmllint", "-schema", xsd, "--noout", "-")

	// cmdOutput := &bytes.Buffer{}
	// cmd.Stdout = cmdOutput

	cmdErr := &bytes.Buffer{}
	cmd.Stderr = cmdErr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	go func() error {
		defer stdin.Close()
		_, err := io.Copy(stdin, xml)
		return err
	}()

	err = cmd.Wait()
	if err != nil {
		return errors.New(cmdErr.String())
	}
	return nil
}
