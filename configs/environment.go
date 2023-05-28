package configs

type Environment string

const (
	Testing    Environment = "testing"
	Production Environment = "production"
	Staging    Environment = "staging"
	Local      Environment = "local"
)

func (e Environment) IsTesting() bool {
	return e == Testing
}

func (e Environment) IsProduction() bool {
	return e == Production
}

func (e Environment) IsStaging() bool {
	return e == Staging
}

func (e Environment) IsLocal() bool {
	return e == Local
}
