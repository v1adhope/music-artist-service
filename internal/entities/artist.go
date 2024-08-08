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

// INFO: validate depends of service
func (e *Artist) SetId(target string) {
	e.id = target
}

func (e *Artist) SetName(target string) error {
	var err error

	e.name, err = objectvalues.ParseArtistName(target)
	if err != nil {
		return err
	}

	return nil
}

func (e *Artist) SetDescription(target string) error {
	var err error

	e.description, err = objectvalues.ParseArtistDescription(target)
	if err != nil {
		return err
	}

	return nil
}

func (e *Artist) SetWebsite(target string) error {
	var err error

	e.website, err = objectvalues.ParseWebsite(target)
	if err != nil {
		return err
	}

	return nil
}

func (e *Artist) SetMounthlyListeners(target uint64) {
	e.mounthlyListeners = target
}

func (e *Artist) SetEmaiil(target string) error {
	var err error

	e.email, err = objectvalues.ParseEmail(target)
	if err != nil {
		return err
	}

	return nil
}

func (e *Artist) SetStatus(mounthlyListeners uint64) {
	e.status = objectvalues.ParseArtistStatus(mounthlyListeners)
}

func (e *Artist) GetId() string {
	return e.id
}

func (e *Artist) GetName() objectvalues.ArtistName {
	return e.name
}

func (e *Artist) GetDescription() objectvalues.ArtistDescription {
	return e.description
}
func (e *Artist) GetWebsite() objectvalues.Website {
	return e.website
}

func (e *Artist) GetMounthlyListeners() uint64 {
	return e.mounthlyListeners
}

func (e *Artist) GetEmail() objectvalues.Email {
	return e.email
}

func (e *Artist) GetStatus() objectvalues.ArtistStatus {
	return e.status
}

type ArtistId struct {
	value string
}

// INFO: validate depends of service
func (e *ArtistId) Set(id string) {
	e.value = id
}

func (e *ArtistId) Get() string {
	return e.value
}
