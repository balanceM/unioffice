// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package chart

import (
	"encoding/xml"

	"github.com/unidoc/unioffice"
)

type CT_AxDataSourceChoice struct {
	MultiLvlStrRef *CT_MultiLvlStrRef
	NumRef         *CT_NumRef
	NumLit         *CT_NumData
	StrRef         *CT_StrRef
	StrLit         *CT_StrData
}

func NewCT_AxDataSourceChoice() *CT_AxDataSourceChoice {
	ret := &CT_AxDataSourceChoice{}
	return ret
}

func (m *CT_AxDataSourceChoice) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MultiLvlStrRef != nil {
		semultiLvlStrRef := xml.StartElement{Name: xml.Name{Local: "c:multiLvlStrRef"}}
		e.EncodeElement(m.MultiLvlStrRef, semultiLvlStrRef)
	}
	if m.NumRef != nil {
		senumRef := xml.StartElement{Name: xml.Name{Local: "c:numRef"}}
		e.EncodeElement(m.NumRef, senumRef)
	}
	if m.NumLit != nil {
		senumLit := xml.StartElement{Name: xml.Name{Local: "c:numLit"}}
		e.EncodeElement(m.NumLit, senumLit)
	}
	if m.StrRef != nil {
		sestrRef := xml.StartElement{Name: xml.Name{Local: "c:strRef"}}
		e.EncodeElement(m.StrRef, sestrRef)
	}
	if m.StrLit != nil {
		sestrLit := xml.StartElement{Name: xml.Name{Local: "c:strLit"}}
		e.EncodeElement(m.StrLit, sestrLit)
	}
	return nil
}

func (m *CT_AxDataSourceChoice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_AxDataSourceChoice:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "multiLvlStrRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "multiLvlStrRef"}:
				m.MultiLvlStrRef = NewCT_MultiLvlStrRef()
				if err := d.DecodeElement(m.MultiLvlStrRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numRef"}:
				m.NumRef = NewCT_NumRef()
				if err := d.DecodeElement(m.NumRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "numLit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "numLit"}:
				m.NumLit = NewCT_NumData()
				if err := d.DecodeElement(m.NumLit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "strRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "strRef"}:
				m.StrRef = NewCT_StrRef()
				if err := d.DecodeElement(m.StrRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "strLit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "strLit"}:
				m.StrLit = NewCT_StrData()
				if err := d.DecodeElement(m.StrLit, &el); err != nil {
					return err
				}
			default:
				unioffice.Log("skipping unsupported element on CT_AxDataSourceChoice %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AxDataSourceChoice
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AxDataSourceChoice and its children
func (m *CT_AxDataSourceChoice) Validate() error {
	return m.ValidateWithPath("CT_AxDataSourceChoice")
}

// ValidateWithPath validates the CT_AxDataSourceChoice and its children, prefixing error messages with path
func (m *CT_AxDataSourceChoice) ValidateWithPath(path string) error {
	if m.MultiLvlStrRef != nil {
		if err := m.MultiLvlStrRef.ValidateWithPath(path + "/MultiLvlStrRef"); err != nil {
			return err
		}
	}
	if m.NumRef != nil {
		if err := m.NumRef.ValidateWithPath(path + "/NumRef"); err != nil {
			return err
		}
	}
	if m.NumLit != nil {
		if err := m.NumLit.ValidateWithPath(path + "/NumLit"); err != nil {
			return err
		}
	}
	if m.StrRef != nil {
		if err := m.StrRef.ValidateWithPath(path + "/StrRef"); err != nil {
			return err
		}
	}
	if m.StrLit != nil {
		if err := m.StrLit.ValidateWithPath(path + "/StrLit"); err != nil {
			return err
		}
	}
	return nil
}
