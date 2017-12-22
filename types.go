package king

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/aodin/date"
	"gopkg.in/guregu/null.v3/zero"
)

type Int10 int

type Decimal float64

type Money float64

type NullBool struct {
	zero.Bool
}

func (b NullBool) IsEmpty() bool {
	return b.IsZero()
}

func (m Money) String() string {
	s := fmt.Sprintf("%.2f", m)
	return s
	// return strings.Replace(s, ".", ",", -1)
}

func (m Money) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

type Time struct {
	time.Time
}

func (t Time) IsEmpty() bool {
	return t.IsZero()
}

func (t Time) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t.IsZero() {
		return e.Encode(nil)
	}

	return e.EncodeElement(t, start)
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(t.String())
}

type Date struct {
	date.Date
}

func (d Date) IsEmpty() bool {
	return d.IsZero()
}

func (d Date) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if d.IsZero() {
		return e.Encode(nil)
	}

	return e.EncodeElement(d.String(), start)
}

func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.String())
}
