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
	"strconv"

	"github.com/unidoc/unioffice"
)

type CT_GroupTransform2D struct {
	RotAttr   *int32
	FlipHAttr *bool
	FlipVAttr *bool
	Off       *CT_Point2D
	Ext       *CT_PositiveSize2D
	ChOff     *CT_Point2D
	ChExt     *CT_PositiveSize2D
}

func NewCT_GroupTransform2D() *CT_GroupTransform2D {
	ret := &CT_GroupTransform2D{}
	return ret
}

func (m *CT_GroupTransform2D) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rot"},
			Value: fmt.Sprintf("%v", *m.RotAttr)})
	}
	if m.FlipHAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "flipH"},
			Value: fmt.Sprintf("%d", b2i(*m.FlipHAttr))})
	}
	if m.FlipVAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "flipV"},
			Value: fmt.Sprintf("%d", b2i(*m.FlipVAttr))})
	}
	e.EncodeToken(start)
	if m.Off != nil {
		seoff := xml.StartElement{Name: xml.Name{Local: "a:off"}}
		e.EncodeElement(m.Off, seoff)
	}
	if m.Ext != nil {
		seext := xml.StartElement{Name: xml.Name{Local: "a:ext"}}
		e.EncodeElement(m.Ext, seext)
	}
	if m.ChOff != nil {
		sechOff := xml.StartElement{Name: xml.Name{Local: "a:chOff"}}
		e.EncodeElement(m.ChOff, sechOff)
	}
	if m.ChExt != nil {
		sechExt := xml.StartElement{Name: xml.Name{Local: "a:chExt"}}
		e.EncodeElement(m.ChExt, sechExt)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GroupTransform2D) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "rot" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.RotAttr = &pt
			continue
		}
		if attr.Name.Local == "flipH" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FlipHAttr = &parsed
			continue
		}
		if attr.Name.Local == "flipV" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FlipVAttr = &parsed
			continue
		}
	}
lCT_GroupTransform2D:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "off"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "off"}:
				m.Off = NewCT_Point2D()
				if err := d.DecodeElement(m.Off, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "ext"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "ext"}:
				m.Ext = NewCT_PositiveSize2D()
				if err := d.DecodeElement(m.Ext, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "chOff"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "chOff"}:
				m.ChOff = NewCT_Point2D()
				if err := d.DecodeElement(m.ChOff, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "chExt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "chExt"}:
				m.ChExt = NewCT_PositiveSize2D()
				if err := d.DecodeElement(m.ChExt, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_GroupTransform2D %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GroupTransform2D
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GroupTransform2D and its children
func (m *CT_GroupTransform2D) Validate() error {
	return m.ValidateWithPath("CT_GroupTransform2D")
}

// ValidateWithPath validates the CT_GroupTransform2D and its children, prefixing error messages with path
func (m *CT_GroupTransform2D) ValidateWithPath(path string) error {
	if m.Off != nil {
		if err := m.Off.ValidateWithPath(path + "/Off"); err != nil {
			return err
		}
	}
	if m.Ext != nil {
		if err := m.Ext.ValidateWithPath(path + "/Ext"); err != nil {
			return err
		}
	}
	if m.ChOff != nil {
		if err := m.ChOff.ValidateWithPath(path + "/ChOff"); err != nil {
			return err
		}
	}
	if m.ChExt != nil {
		if err := m.ChExt.ValidateWithPath(path + "/ChExt"); err != nil {
			return err
		}
	}
	return nil
}
