package repositories

type artistDto struct {
	id                string
	name              string
	description       string
	website           string
	mounthlyListeners uint64
	email             string
	status            string
}

type artistIdDto struct {
	id string
}
