package template

import (
	"io"
	"text/template"
)

// Generate writes the template to the writer, interpolating the data.
func Generate(t *template.Template, wr io.Writer, name string, data any) error {
	_ = "STUB: not implemented"
	return nil
}

// GenerateFile generates a file from a template. It will create the directory if it does not exist,
// finally calling the Generate function to perform the interpolation.
func GenerateFile(t *template.Template, exampleFilePath string, name string, data any) error {
	_ = "STUB: not implemented"
	return nil
}
