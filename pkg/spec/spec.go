package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/katallaxie/pkg/utils/files"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	// DefaultVersion is the default version of the specification.
	DefaultVersion = 1
	// DefaultFilename is the default filename of the specification.
	DefaultFilename = ".g.yml"
)

var validate = validator.New()

// Spec is the specification for the scaffolding tool.
type Spec struct {
	// Version is the version of the specification.
	Version int `validate:"required" yaml:"version"`
	// Name is the name of the template.
	Name string `validate:"required" yaml:"name"`
	// Description is the description of the template.
	Description string `yaml:"description"`
	// Templates is the list of templates to use.
	Templates []Template `validate:"required" yaml:"templates"`

	sync.Mutex `yaml:"-"`
}

// Template is a template to use.
type Template struct {
	// Source is the source of the template.
	Source string `validate:"required" yaml:"source"`
	// Destination is the destination of the template.
	Destination string `validate:"required" yaml:"destination"`
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (s *Spec) UnmarshalYAML(data []byte) error {
	ss := struct {
		Version     int        `yaml:"version"`
		Name        string     `yaml:"name"`
		Description string     `yaml:"description"`
		Templates   []Template `yaml:"templates"`
	}{}

	if err := yaml.Unmarshal(data, &s); err != nil {
		return errors.WithStack(err)
	}

	s.Version = ss.Version
	s.Name = ss.Name
	s.Description = ss.Description
	s.Templates = ss.Templates

	err := validate.Struct(s)
	if err != nil {
		return err
	}

	return err
}

// Default returns the default specification.
func Default() *Spec {
	return &Spec{
		Version: DefaultVersion,
	}
}

// Write writes the specification to the given file.
func Write(s *Spec, file string, force bool) error {
	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	ok, _ := files.FileExists(filepath.Clean(file))
	if ok && !force {
		return fmt.Errorf("%s already exists, use --force to overwrite", file)
	}

	f, err := os.Create(filepath.Clean(file))
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
