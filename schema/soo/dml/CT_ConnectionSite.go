// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml

import (
	"encoding/xml"
	"fmt"

	"github.com/unidoc/unioffice"
)

type CT_ConnectionSite struct {
	AngAttr ST_AdjAngle
	Pos     *CT_AdjPoint2D
}

func NewCT_ConnectionSite() *CT_ConnectionSite {
	ret := &CT_ConnectionSite{}
	ret.Pos = NewCT_AdjPoint2D()
	return ret
}

func (m *CT_ConnectionSite) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "ang"},
		Value: fmt.Sprintf("%v", m.AngAttr)})
	e.EncodeToken(start)
	sepos := xml.StartElement{Name: xml.Name{Local: "a:pos"}}
	e.EncodeElement(m.Pos, sepos)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ConnectionSite) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Pos = NewCT_AdjPoint2D()
	for _, attr := range start.Attr {
		if attr.Name.Local == "ang" {
			parsed, err := ParseUnionST_AdjAngle(attr.Value)
			if err != nil {
				return err
			}
			m.AngAttr = parsed
			continue
		}
	}
lCT_ConnectionSite:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "pos"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "pos"}:
				if err := d.DecodeElement(m.Pos, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_ConnectionSite %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ConnectionSite
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ConnectionSite and its children
func (m *CT_ConnectionSite) Validate() error {
	return m.ValidateWithPath("CT_ConnectionSite")
}

// ValidateWithPath validates the CT_ConnectionSite and its children, prefixing error messages with path
func (m *CT_ConnectionSite) ValidateWithPath(path string) error {
	if err := m.AngAttr.ValidateWithPath(path + "/AngAttr"); err != nil {
		return err
	}
	if err := m.Pos.ValidateWithPath(path + "/Pos"); err != nil {
		return err
	}
	return nil
}
