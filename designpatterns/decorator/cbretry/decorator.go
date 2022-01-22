package decorator

type Args map[string]interface{}

type Response interface{}

type Fetcher interface {
	Fetch(args Args) (Response, error)
}
