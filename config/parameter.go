package config

import "fmt"

const (
	// PGVersion is the default PostgreSQL version
	PGVersion float32 = 12.0
)

// Parameter contains a exame
type Parameter struct {
	input Input
	Name  string `json:"guc"`
	value interface{}
	Type  OutputType `json:"type"`
}

// ToSQL exports the parameter using the `ALTER SYSTEM` syntax
func (p *Parameter) ToSQL() string {
	return fmt.Sprintf("ALTER SYSTEM SET %s TO '%s';", p.Name, format(p.Type, p.value))
}

// DocURL compute the documentation url
func (p *Parameter) DocURL() string {
	return fmt.Sprintf("https://postgresqlco.nf/en/doc/param/%s/%.0f/", p.Name, p.input.PostgresVersion)
}
