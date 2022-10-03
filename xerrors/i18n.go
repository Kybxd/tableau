package xerrors

import (
	"bytes"
	"embed"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"text/template"

	"gopkg.in/yaml.v2"
)

const (
	LangEn = "en"
	LangZh = "zh"
)

var lang = "zh"

func SetLang(language string) {
	lang = language
}

//go:embed i18n
var localeFS embed.FS

// init tmpl once
var once sync.Once
var tmpl *template.Template
var ecodes ecodeMap

func initI18N(lang string) {
	// init ecode
	filename := "i18n/" + lang + ".yaml"
	data, err := localeFS.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, &ecodes); err != nil {
		panic(err)
	}
	// init summary
	// refer: https://stackoverflow.com/questions/22367337/last-item-in-a-template-range
	tplFuncs := template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
		"notlast": func(x int, a interface{}) bool {
			return x != reflect.ValueOf(a).Len()-1
		},
		"hasprefix": strings.HasPrefix,
	}
	t := template.New("ERR-TMPL")
	t, err = t.Funcs(tplFuncs).ParseFS(localeFS, "i18n/summary/*.tmpl")
	if err != nil {
		panic(err)
	}
	tmpl = t
}

// See https://rustc-dev-guide.rust-lang.org/diagnostics.html
type ErrDetail struct {
	ecode string // error code like: EXXXX
	Desc  string `yaml:"desc"` // basic description
	Text  string `yaml:"text"` // error text
	Help  string `yaml:"help"` // fix suggestion
}

// error code -> error detail
type ecodeMap map[string]ErrDetail

// E2001 describes field prop "refer" not configured correctly.
func E2001(refer string, messageName string) error {
	detail := renderI18N("E2001", map[string]interface{}{
		"Refer":       refer,
		"MessageName": messageName,
	})

	return ErrorKV(detail.Text,
		keyErrCode, detail.ecode,
		keyErrDesc, detail.Desc,
		keyHelp, detail.Help)
}

// E2002 describes field value not in referred space.
func E2002(value, refer string) error {
	detail := renderI18N("E2002", map[string]interface{}{
		"Value": value,
		"Refer": refer,
	}) 

	return ErrorKV(detail.Text,
		keyErrCode, detail.ecode,
		keyErrDesc, detail.Desc,
		keyHelp, detail.Help)
}

func renderI18N(ecode string, data interface{}) *ErrDetail {
	once.Do(func() {
		initI18N(lang)
	})
	detail, ok := ecodes[ecode]
	if !ok {
		panic(fmt.Sprintf("ecode %s not found", ecode))
	}
	return &ErrDetail{
		ecode: ecode,
		Desc:  detail.Desc,
		Text:  render(detail.Text, data),
		Help:  render(detail.Help, data),
	}
}

func render(text string, data interface{}) string {
	tmpl, err := template.New("ERROR").Parse(text)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBufferString("")
	err = tmpl.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func renderSummary(ecode string, data interface{}) string {
	once.Do(func() {
		initI18N(lang)
	})
	buf := bytes.NewBufferString("")
	name := fmt.Sprintf("%s_%s.tmpl", ecode, lang)
	if err := tmpl.ExecuteTemplate(buf, name, data); err != nil {
		panic(err)
	}
	return buf.String()
}