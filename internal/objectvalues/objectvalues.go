package objectvalues

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
)

const (
	StatusPlatinum = "Platinum"
	StatusGold     = "Gold"
	StatusSilver   = "Silver"
	StatusBronse   = "Bronse"
)

var (
	ErrNotValidEmail             = errors.New("Not valide email")
	ErrNotValidWebsite           = errors.New("Not valide website")
	ErrNotValidArtistName        = errors.New("Not valide name")
	ErrNotValidArtistDescription = errors.New("Not valid description")
)

type ArtistStatus string

func (ov ArtistStatus) String() string {
	return string(ov)
}

func ParseArtistStatus(mounthlyListeners uint64) ArtistStatus {
	status := ArtistStatus(StatusBronse)

	switch {
	case mounthlyListeners >= 100000 && mounthlyListeners < 1000000:
		status = StatusSilver
	case mounthlyListeners >= 1000000 && mounthlyListeners < 10000000:
		status = StatusGold
	case mounthlyListeners >= 10000000:
		status = StatusPlatinum
	}

	return status
}

type Email string

func (ov Email) String() string {
	return string(ov)
}

func ParseEmail(target string) (Email, error) {
	_, err := mail.ParseAddress(target)
	if err != nil {
		return "", ErrNotValidEmail
	}

	return Email(target), nil
}

type Website string

func (ov Website) String() string {
	return string(ov)
}

func ParseWebsite(target string) (Website, error) {
	url, err := url.ParseRequestURI(target)
	if err != nil {
		return "", ErrNotValidWebsite
	}

	if url.Scheme != "http" && url.Scheme != "https" {
		return "", ErrNotValidWebsite
	}

	return Website(target), nil
}

type ArtistName string

func (ov ArtistName) String() string {
	return string(ov)
}

func ParseArtistName(target string) (ArtistName, error) {
	if len(target) > 255 {
		return "", ErrNotValidArtistName
	}

	isMatched, err := regexp.MatchString("^[A-Za-z ,.'-]+$", target)
	if !isMatched {
		return "", ErrNotValidArtistName
	}
	if err != nil {
		return "", err
	}

	return ArtistName(target), nil
}

type ArtistDescription string

func (ov ArtistDescription) String() string {
	return string(ov)
}

func ParseArtistDescription(target string) (ArtistDescription, error) {
	if len(target) > 255 {
		return "", ErrNotValidArtistDescription
	}

	return ArtistDescription(target), nil
}
