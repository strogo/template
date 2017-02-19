package template_test

import (
	"bytes"
	"testing"

	"github.com/go-qbit/template"

	"github.com/stretchr/testify/assert"
)

func TestTemplate_Parse(t *testing.T) {
	tpl := template.New()

	assert.NoError(t, tpl.ParseFile("test/templates/index.gtt"))

	buf := &bytes.Buffer{}
	tpl.WriteGo(buf, "test", "Test")

	//fmt.Println(buf.String())
}
