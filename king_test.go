package king_test

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
)

func validate(xml io.Reader, xsd string) error {
	cmd := exec.Command("xmllint", "-schema", xsd, "--noout", "-")

	// capture stderr to generate error message
	cmdErr := &bytes.Buffer{}
	cmd.Stderr = cmdErr

	// connect stdin to io.Writer so we can push the xml
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	// push xml to stdin
	go func() error {
		defer stdin.Close()
		_, err := io.Copy(stdin, xml)
		return err
	}()

	// wait for command to finish
	err = cmd.Wait()
	if err != nil {
		return errors.New(cmdErr.String())
	}

	return nil
}
