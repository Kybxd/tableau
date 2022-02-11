package protogen

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/emirpasic/gods/sets/treeset"
	"github.com/pkg/errors"
	"github.com/tableauio/tableau/internal/atom"
	"github.com/tableauio/tableau/internal/fs"
	"github.com/tableauio/tableau/internal/printer"
	"github.com/tableauio/tableau/proto/tableaupb"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var protoTypeRegexp *regexp.Regexp

func init() {
	protoTypeRegexp = regexp.MustCompile(`^(message|enum) +(.+?) +(\{)?`) // message or enum type
}

type bookExporter struct {
	ProtoPackage   string
	GoPackage      string
	OutputDir      string
	FilenameSuffix string
	wb             *tableaupb.Workbook
	customImports  []string
}

func newBookExporter(protoPackage, goPackage, outputDir, filenameSuffix string, imports []string, wb *tableaupb.Workbook) *bookExporter {
	return &bookExporter{
		ProtoPackage:   protoPackage,
		GoPackage:      goPackage,
		OutputDir:      outputDir,
		FilenameSuffix: filenameSuffix,
		wb:             wb,
		customImports:  imports,
	}
}

func parseCustomImports(dir string, filenames []string) (map[string]string, error) {
	type2import := make(map[string]string)
	for _, filename := range filenames {
		fpath := filepath.Join(dir, filename)
		file, err := os.Open(fpath)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to open custom import proto file: %s", fpath)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			if matches := protoTypeRegexp.FindStringSubmatch(scanner.Text()); len(matches) > 0 {
				type2import[matches[2]] = filename
			}
		}

		if err := scanner.Err(); err != nil {
			return nil, errors.Wrapf(err, "failed to read custom import proto file: %s", fpath)
		}
	}
	return type2import, nil
}

func (x *bookExporter) export() error {
	type2import, err := parseCustomImports(x.OutputDir, x.customImports)
	if err != nil {
		return errors.Wrapf(err, "failed to parse custom imports")
	}

	// atom.Log.Debug(proto.MarshalTextString(wb))
	g1 := NewGeneratedBuf()
	g1.P("// Generated by ", App, " ", Version, ". DO NOT EDIT.")
	g1.P(`syntax = "proto3";`)
	g1.P("package ", x.ProtoPackage, ";")
	g1.P(`option go_package = "`, x.GoPackage, `";`)
	g1.P("")

	// keep the elements ordered by import path
	set := treeset.NewWithStringComparator()
	set.Add(tableauProtoPath) // default must import path

	g3 := NewGeneratedBuf()
	for i, ws := range x.wb.Worksheets {
		x := &sheetExporter{
			ws:             ws,
			g:              g3,
			isLastSheet:    i == len(x.wb.Worksheets)-1,
			nestedMessages: make(map[string]*tableaupb.Field),
			type2import:    type2import,
			Imports:        make(map[string]bool),
		}
		if err := x.export(); err != nil {
			return err
		}
		for key := range x.Imports {
			set.Add(key)
		}
	}

	// generate imports
	g2 := NewGeneratedBuf()
	for _, key := range set.Values() {
		g2.P(`import "`, key, `";`)
	}
	g2.P("")
	g2.P("option (tableau.workbook) = {", marshalToText(x.wb.Options), "};")
	g2.P("")

	path := filepath.Join(x.OutputDir, x.wb.Name+x.FilenameSuffix+".proto")
	atom.Log.Infof("output: %s", path)

	if existed, err := fs.Exists(path); err != nil {
		return errors.WithMessagef(err, "failed to check if file exists: %s", path)
	} else {
		if existed {
			return errors.Errorf("file already exists: %s", path)
		}
	}

	if f, err := os.Create(path); err != nil {
		return errors.Wrapf(err, "failed to create output file: %s", path)
	} else {
		defer f.Close()
		if _, err = f.Write(g1.Content()); err != nil {
			return errors.Wrapf(err, "failed to write output file: %s", path)
		}
		if _, err = f.Write(g2.Content()); err != nil {
			return errors.Wrapf(err, "failed to write output file: %s", path)
		}
		if _, err = f.Write(g3.Content()); err != nil {
			return errors.Wrapf(err, "failed to write output file: %s", path)
		}
	}

	return nil
}

type sheetExporter struct {
	ws             *tableaupb.Worksheet
	g              *GeneratedBuf
	isLastSheet    bool
	nestedMessages map[string]*tableaupb.Field // type name -> field
	type2import    map[string]string           // type name -> import path
	Imports        map[string]bool             // import name -> defined
}

func (x *sheetExporter) export() error {
	x.g.P("message ", x.ws.Name, " {")
	x.g.P("  option (tableau.worksheet) = {", marshalToText(x.ws.Options), "};")
	x.g.P("")
	// generate the fields
	depth := 1
	for i, field := range x.ws.Fields {
		tagid := i + 1
		if err := x.exportField(depth, tagid, field, x.ws.Name); err != nil {
			return err
		}
	}
	x.g.P("}")
	if !x.isLastSheet {
		x.g.P("")
	}
	return nil
}

func (x *sheetExporter) exportField(depth int, tagid int, field *tableaupb.Field, prefix string) error {
	// head := "%x%x"
	cardTypeSep := ""
	if field.Card != "" {
		// head += " " // cardinality exists
		cardTypeSep = " "
	}
	x.g.P(printer.Indent(depth), field.Card, cardTypeSep, field.Type, " ", field.Name, " = ", tagid, " [(tableau.field) = {", marshalToText(field.Options), "}];")

	if field.Type == "google.protobuf.Timestamp" {
		x.Imports[timestampProtoPath] = true
	} else if field.Type == "google.protobuf.Duration" {
		x.Imports[durationProtoPath] = true
	}

	if field.TypeDefined {
		// NOTE: import corresponding message's custom defined proto file
		if path, ok := x.type2import[field.Type]; ok {
			x.Imports[path] = true
		}
	}

	if !field.TypeDefined && field.Fields != nil {
		// iff field is a map or list and message type is not imported.
		msgName := field.Type
		if field.MapEntry != nil {
			msgName = field.MapEntry.ValueType
		}
		nestedMsgName := prefix + "." + msgName

		if isSameFieldMessageType(field, x.nestedMessages[nestedMsgName]) {
			// if the nested message is the same as the previous one,
			// just use the previous one, and don't generate a new one.
			return nil
		}

		// bookkeeping this nested msessage, so we can check if we can reuse it later.
		x.nestedMessages[nestedMsgName] = field

		// x.g.P("")
		x.g.P(printer.Indent(depth), "message ", msgName, " {")
		for i, f := range field.Fields {
			tagid := i + 1
			if err := x.exportField(depth+1, tagid, f, nestedMsgName); err != nil {
				return err
			}
		}
		x.g.P(printer.Indent(depth), "}")
	}
	return nil
}

func marshalToText(m protoreflect.ProtoMessage) string {
	// text := proto.CompactTextString(field.Options)
	bin, err := prototext.Marshal(m)
	if err != nil {
		panic(err)
	}
	// NOTE: remove redundant spaces/whitespace from a string
	// refer: https://stackoverflow.com/questions/37290693/how-to-remove-redundant-spaces-whitespace-from-a-string-in-golang
	text := strings.Join(strings.Fields(string(bin)), " ")
	return text
}

func isSameFieldMessageType(left, right *tableaupb.Field) bool {
	if left == nil || right == nil {
		return false
	}
	if left.Fields == nil || right.Fields == nil {
		return false
	}
	if len(left.Fields) != len(right.Fields) ||
		left.Type != right.Type ||
		left.Card != right.Card {
		return false
	}

	for i, l := range left.Fields {
		r := right.Fields[i]
		if !proto.Equal(l, r) {
			return false
		}
	}
	return true
}
