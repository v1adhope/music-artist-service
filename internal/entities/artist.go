package entities

import (
	"github.com/v1adhope/music-artist-service/internal/objectvalues"
)

// INFO: mounthlyListeners is concept with cron and cache
type Artist struct {
	id                string
	name              objectvalues.ArtistName
	description       objectvalues.ArtistDescription
	website           objectvalues.Website
	mounthlyListeners uint64
	email             objectvalues.Email
	status            objectvalues.ArtistStatus
}

func ParseArtist(name, description, website, email string, mounthlyListeners uint64) (Artist, error) {
	validName, err := objectvalues.ParseArtistName(name)
	if err != nil {
		return Artist{}, err
	}

	validDesc, err := objectvalues.ParseArtistDescription(description)
	if err != nil {
		return Artist{}, err
	}

	validWebsite, err := objectvalues.ParseWebsite(website)
	if err != nil {
		return Artist{}, err
	}

	validEmail, err := objectvalues.ParseEmail(email)
	if err != nil {
		return Artist{}, err
	}

	status := objectvalues.ParseArtistStatus(mounthlyListeners)

	return Artist{
		name:              validName,
		description:       validDesc,
		website:           validWebsite,
		mounthlyListeners: mounthlyListeners,
		email:             validEmail,
		status:            status,
	}, nil
}

func (e *Artist) SetID(id string) {
	e.id = id
}

func (e *Artist) GetID() string {
	return e.id
}

type ArtistId struct {
	value string
}

func (e *ArtistId) String() string {
	return e.value
}
