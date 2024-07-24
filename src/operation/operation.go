package operation

type Operation interface {
	Execute([]string) ([]string, error)
}

type AggregatingOperation interface {
	Operation
	Result([]string) []string
}
