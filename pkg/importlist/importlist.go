package importlist

type Item struct {
	Title  string
	ImdbID string
	TvdbID string
	TmdbID string
}

type Response struct {
	Items []Item
}
