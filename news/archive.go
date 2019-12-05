package news

//Topic ...
type Topic struct {
	Titel  string `json:"titel"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

//Archive type ...
type Archive map[string][]Topic

//New ...
func New() Archive  {
	return make(map[string][]Topic,0)
}

func (a *Archive) collect(cat string) {
	a[cat] = topics
}

func (a *Archive) result(cat string) []Topic {
	return a[cat]
}