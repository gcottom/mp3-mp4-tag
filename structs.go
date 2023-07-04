package mp3mp4tag

import (
	"image"
)

type IDTag struct {
	artist      string
	albumArtist string
	album       string
	albumArt    *image.Image
	comments    string
	composer    string
	genre       string
	title       string
	year        int
	bpm         int
	id3         ID3Frames
	fileUrl     string
}
type ID3Frames struct {
	contentType   string //Content Type
	copyrightMsg  string //Copyright Message
	date          string //Date
	encodedBy     string //Endcoded By
	lyricist      string //Lyricist
	fileType      string //File Type
	language      string //Language
	length        string //Length
	partOfSet     string //Part of a set
	publisher     string //Publisher
}
