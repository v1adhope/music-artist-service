package objectvalues_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/v1adhope/music-artist-service/internal/objectvalues"
)

func TestParseArtistStatus(t *testing.T) {
	tcs := []struct {
		key      string
		input    uint64
		expected objectvalues.ArtistStatus
	}{
		{
			key:      "Case 1",
			input:    0,
			expected: objectvalues.StatusBronse,
		},
		{
			key:      "Case 2",
			input:    100000,
			expected: objectvalues.StatusSilver,
		},
		{
			key:      "Case 3",
			input:    1000000,
			expected: objectvalues.StatusGold,
		},
		{
			key:      "Case 4",
			input:    10000000,
			expected: objectvalues.StatusPlatinum,
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			status := objectvalues.ParseArtistStatus(tc.input)

			assert.Equal(t, tc.expected, status, tc.key)
		})
	}
}

func TestParseEmailPositive(t *testing.T) {
	tcs := []struct {
		key      string
		input    string
		expected objectvalues.Email
	}{
		{
			key:      "Case 1",
			input:    "spam@eggs.com",
			expected: "spam@eggs.com",
		},
		{
			key:      "Case 2",
			input:    "imaginedragons@gmail.com",
			expected: "imaginedragons@gmail.com",
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			email, err := objectvalues.ParseEmail(tc.input)

			assert.NoError(t, err, tc.key)
			assert.Equal(t, tc.expected, email)
		})
	}
}

func TestParseEmailNegative(t *testing.T) {
	tcs := []struct {
		key   string
		input string
	}{
		{
			key:   "Case 1",
			input: "spam-eggs.com",
		},
		{
			key:   "Case 2",
			input: "imaginedragonsgmail.com",
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			email, err := objectvalues.ParseEmail(tc.input)

			assert.ErrorIs(t, err, objectvalues.ErrNotValidEmail, tc.key)
			assert.Equal(t, objectvalues.Email(""), email, tc.key)
		})
	}
}

func TestParseWebsitePositive(t *testing.T) {
	tcs := []struct {
		key    string
		input  string
		expect objectvalues.Website
	}{
		{
			key:    "Case 1",
			input:  "http://spam.eggs",
			expect: "http://spam.eggs",
		},
		{
			key:    "Case 2",
			input:  "https://www.facebook.com/ImagineDragons",
			expect: "https://www.facebook.com/ImagineDragons",
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			website, err := objectvalues.ParseWebsite(tc.input)

			assert.NoError(t, err, tc.key)
			assert.Equal(t, tc.expect, website, tc.key)
		})
	}
}

func TestParseWebsiteNegative(t *testing.T) {
	tcs := []struct {
		key   string
		input string
	}{
		{
			key:   "Case 1",
			input: "ftp://spam.eggs",
		},
		{
			key:   "Case 2",
			input: "www.facebook.com/ImagineDragons",
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			website, err := objectvalues.ParseWebsite(tc.input)

			assert.ErrorIs(t, err, objectvalues.ErrNotValidWebsite, tc.key)
			assert.Equal(t, objectvalues.Website(""), website, tc.key)
		})
	}
}

func TestParseArtistNamePositive(t *testing.T) {
	tcs := []struct {
		key      string
		input    string
		expected objectvalues.ArtistName
	}{
		{
			key:      "Case 1",
			input:    "Spa-m, eg.g's",
			expected: "Spa-m, eg.g's",
		},
		{
			key:      "Case 2",
			input:    "Imagine Dragons",
			expected: "Imagine Dragons",
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			name, err := objectvalues.ParseArtistName(tc.input)

			assert.NoError(t, err, tc.key)
			assert.Equal(t, tc.expected, name, tc.key)
		})
	}
}

func TestParseArtistNameNegative(t *testing.T) {
	tcs := []struct {
		key   string
		input string
	}{
		{
			key:   "Case 1",
			input: "Spam1 eggs",
		},
		{
			key:   "Case 2",
			input: "Imagine&Dragons",
		},
		{
			key:   "Case 3",
			input: strings.Repeat("A", 256),
		},
		{
			key:   "Case 4",
			input: "",
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			name, err := objectvalues.ParseArtistName(tc.input)

			assert.ErrorIs(t, err, objectvalues.ErrNotValidArtistName, tc.key)
			assert.Equal(t, objectvalues.ArtistName(""), name, tc.key)
		})
	}
}

func TestParseArtistDescriptionPositive(t *testing.T) {
	tcs := []struct {
		key      string
		input    string
		expected objectvalues.ArtistDescription
	}{
		{
			key:      "Case 1",
			input:    "Some description",
			expected: "Some description",
		},
		{
			key:      "Case 2",
			input:    "Formed in 2009, Imagine Dragons first revealed their emotionally charged and inventive sensibilities with a series of independently released EPs that earned them grassroots following.",
			expected: "Formed in 2009, Imagine Dragons first revealed their emotionally charged and inventive sensibilities with a series of independently released EPs that earned them grassroots following.",
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			desc, err := objectvalues.ParseArtistDescription(tc.input)

			assert.NoError(t, err, tc.key)
			assert.Equal(t, tc.expected, desc, tc.key)
		})
	}
}

func TestParseArtistDescriptionNegative(t *testing.T) {
	tcs := []struct {
		key   string
		input string
	}{
		{
			key:   "Case 1",
			input: strings.Repeat("A", 256),
		},
		{
			key:   "Case 2",
			input: strings.Repeat("A", 11),
		},
	}

	for _, tc := range tcs {
		t.Run("", func(t *testing.T) {
			desc, err := objectvalues.ParseArtistDescription(tc.input)

			assert.ErrorIs(t, err, objectvalues.ErrNotValidArtistDescription, tc.key)
			assert.Equal(t, objectvalues.ArtistDescription(""), desc, tc.key)
		})
	}
}
