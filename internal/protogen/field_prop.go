package protogen

import (
	"github.com/tableauio/tableau/proto/tableaupb"
	"google.golang.org/protobuf/proto"
)

var emptyFieldProp = &tableaupb.FieldProp{}

func ExtractMapFieldProp(prop *tableaupb.FieldProp) *tableaupb.FieldProp {
	if prop == nil {
		return nil
	}
	p := &tableaupb.FieldProp{
		Unique:   prop.Unique,
		Sequence: prop.Sequence,
		Fixed:    prop.Fixed,
		Length:   prop.Length,
		Range:    prop.Range,
	}
	if proto.Equal(emptyFieldProp, p) {
		return nil
	}
	return p
}

func ExtractListFieldProp(prop *tableaupb.FieldProp) *tableaupb.FieldProp {
	if prop == nil {
		return nil
	}
	p := &tableaupb.FieldProp{
		Unique:   prop.Unique, // only for keyed list ?
		Sequence: prop.Sequence,
		Fixed:    prop.Fixed,
		Length:   prop.Length,
	}
	if proto.Equal(emptyFieldProp, p) {
		return nil
	}
	return p
}

func ExtractMapKeyFieldProp(prop *tableaupb.FieldProp) *tableaupb.FieldProp {
	if prop == nil {
		return nil
	}
	p := &tableaupb.FieldProp{
		Refer: prop.Refer,
	}
	if proto.Equal(emptyFieldProp, p) {
		return nil
	}
	return p
}

func ExtractScalarFieldProp(prop *tableaupb.FieldProp) *tableaupb.FieldProp {
	if prop == nil {
		return nil
	}
	p := &tableaupb.FieldProp{
		Range:   prop.Range,
		Refer:   prop.Refer,
		Default: prop.Default,
	}
	if proto.Equal(emptyFieldProp, p) {
		return nil
	}
	return p
}