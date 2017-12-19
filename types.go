package king

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/aodin/date"
)

type Int10 int

type Decimal float64

type Money float64

func (m Money) String() string {
	s := fmt.Sprintf("%.2f", m)
	return s
	// return strings.Replace(s, ".", ",", -1)
}

func (m Money) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

type Time struct {
	time.Time
}

func (t Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.Encode(nil)
	}

	return e.EncodeElement(t, start)
}

type Date struct {
	date.Date
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d.IsZero() {
		return e.Encode(nil)
	}

	return e.EncodeElement(d.String(), start)
}
