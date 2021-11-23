package definition

type Service struct {
	Host      string
	Port      int
	Endpoints []Endpoint
}

type Endpoint struct {
	Name    string
	Path    string
	Actions []Action
}

type Action struct {
	Method      string
	Description string
}
