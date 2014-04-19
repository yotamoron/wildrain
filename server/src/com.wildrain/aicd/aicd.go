package aicd

type Param struct {
	ParamName string
	ParamType string
	Required  bool
}

type ParametrizedEndpoint struct {
	Name   string
	Params []Param
}

type Query struct {
	Name   string
	Params []Param
	Return []Param
}

type Aicd struct {
	ApplicationName string
	Version         string
	Events          []ParametrizedEndpoint
	Commands        []ParametrizedEndpoint
	Queries         []Query
}
