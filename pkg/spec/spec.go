package spec

import (
	"sync"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	// DefaultVersion is the default version of the specification.
	DefaultVersion = 1
	// DefaultFilename is the default filename of the specification.
	DefaultFilename = ".g.yml"
)

// Spec is the specification for the scaffolding tool.
type Spec struct {
	// Version is the version of the specification.
	Version int `validate:"required" yaml:"version"`
	// Name is the name of the project.
	Name string `validate:"required" yaml:"name"`
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
		Version   int        `yaml:"version"`
		Name      string     `yaml:"name"`
		Templates []Template `yaml:"templates"`
	}{}

	if err := yaml.Unmarshal(data, &s); err != nil {
		return errors.WithStack(err)
	}

	s.Version = ss.Version
	s.Name = ss.Name
	s.Templates = ss.Templates

	return nil
}

// Default returns the default specification.
func Default() *Spec {
	return &Spec{
		Version: DefaultVersion,
	}
}
