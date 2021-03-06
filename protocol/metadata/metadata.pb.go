// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metadata.proto

/*
Package metadata is a generated protocol buffer package.

It is generated from these files:
	metadata.proto

It has these top-level messages:
	TopTracks
	ActivityPeriod
	Artist
	AlbumGroup
	Date
	Album
	Track
	Image
	ImageGroup
	Biography
	Disc
	Copyright
	Restriction
	SalePeriod
	ExternalId
	AudioFile
*/
package metadata

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Album_Type int32

const (
	Album_ALBUM       Album_Type = 1
	Album_SINGLE      Album_Type = 2
	Album_COMPILATION Album_Type = 3
	Album_EP          Album_Type = 4
)

var Album_Type_name = map[int32]string{
	1: "ALBUM",
	2: "SINGLE",
	3: "COMPILATION",
	4: "EP",
}
var Album_Type_value = map[string]int32{
	"ALBUM":       1,
	"SINGLE":      2,
	"COMPILATION": 3,
	"EP":          4,
}

func (x Album_Type) Enum() *Album_Type {
	p := new(Album_Type)
	*p = x
	return p
}
func (x Album_Type) String() string {
	return proto.EnumName(Album_Type_name, int32(x))
}
func (x *Album_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Album_Type_value, data, "Album_Type")
	if err != nil {
		return err
	}
	*x = Album_Type(value)
	return nil
}
func (Album_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Image_Size int32

const (
	Image_DEFAULT Image_Size = 0
	Image_SMALL   Image_Size = 1
	Image_LARGE   Image_Size = 2
	Image_XLARGE  Image_Size = 3
)

var Image_Size_name = map[int32]string{
	0: "DEFAULT",
	1: "SMALL",
	2: "LARGE",
	3: "XLARGE",
}
var Image_Size_value = map[string]int32{
	"DEFAULT": 0,
	"SMALL":   1,
	"LARGE":   2,
	"XLARGE":  3,
}

func (x Image_Size) Enum() *Image_Size {
	p := new(Image_Size)
	*p = x
	return p
}
func (x Image_Size) String() string {
	return proto.EnumName(Image_Size_name, int32(x))
}
func (x *Image_Size) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Image_Size_value, data, "Image_Size")
	if err != nil {
		return err
	}
	*x = Image_Size(value)
	return nil
}
func (Image_Size) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type Copyright_Type int32

const (
	Copyright_P Copyright_Type = 0
	Copyright_C Copyright_Type = 1
)

var Copyright_Type_name = map[int32]string{
	0: "P",
	1: "C",
}
var Copyright_Type_value = map[string]int32{
	"P": 0,
	"C": 1,
}

func (x Copyright_Type) Enum() *Copyright_Type {
	p := new(Copyright_Type)
	*p = x
	return p
}
func (x Copyright_Type) String() string {
	return proto.EnumName(Copyright_Type_name, int32(x))
}
func (x *Copyright_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Copyright_Type_value, data, "Copyright_Type")
	if err != nil {
		return err
	}
	*x = Copyright_Type(value)
	return nil
}
func (Copyright_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{11, 0} }

type Restriction_Type int32

const (
	Restriction_STREAMING Restriction_Type = 0
)

var Restriction_Type_name = map[int32]string{
	0: "STREAMING",
}
var Restriction_Type_value = map[string]int32{
	"STREAMING": 0,
}

func (x Restriction_Type) Enum() *Restriction_Type {
	p := new(Restriction_Type)
	*p = x
	return p
}
func (x Restriction_Type) String() string {
	return proto.EnumName(Restriction_Type_name, int32(x))
}
func (x *Restriction_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Restriction_Type_value, data, "Restriction_Type")
	if err != nil {
		return err
	}
	*x = Restriction_Type(value)
	return nil
}
func (Restriction_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{12, 0} }

type AudioFile_Format int32

const (
	AudioFile_OGG_VORBIS_96  AudioFile_Format = 0
	AudioFile_OGG_VORBIS_160 AudioFile_Format = 1
	AudioFile_OGG_VORBIS_320 AudioFile_Format = 2
	AudioFile_MP3_256        AudioFile_Format = 3
	AudioFile_MP3_320        AudioFile_Format = 4
	AudioFile_MP3_160        AudioFile_Format = 5
	AudioFile_MP3_96         AudioFile_Format = 6
	AudioFile_MP3_160_ENC    AudioFile_Format = 7
	AudioFile_OTHER2         AudioFile_Format = 8
	AudioFile_OTHER3         AudioFile_Format = 9
	AudioFile_AAC_160        AudioFile_Format = 10
	AudioFile_AAC_320        AudioFile_Format = 11
	AudioFile_OTHER4         AudioFile_Format = 12
	AudioFile_OTHER5         AudioFile_Format = 13
)

var AudioFile_Format_name = map[int32]string{
	0:  "OGG_VORBIS_96",
	1:  "OGG_VORBIS_160",
	2:  "OGG_VORBIS_320",
	3:  "MP3_256",
	4:  "MP3_320",
	5:  "MP3_160",
	6:  "MP3_96",
	7:  "MP3_160_ENC",
	8:  "OTHER2",
	9:  "OTHER3",
	10: "AAC_160",
	11: "AAC_320",
	12: "OTHER4",
	13: "OTHER5",
}
var AudioFile_Format_value = map[string]int32{
	"OGG_VORBIS_96":  0,
	"OGG_VORBIS_160": 1,
	"OGG_VORBIS_320": 2,
	"MP3_256":        3,
	"MP3_320":        4,
	"MP3_160":        5,
	"MP3_96":         6,
	"MP3_160_ENC":    7,
	"OTHER2":         8,
	"OTHER3":         9,
	"AAC_160":        10,
	"AAC_320":        11,
	"OTHER4":         12,
	"OTHER5":         13,
}

func (x AudioFile_Format) Enum() *AudioFile_Format {
	p := new(AudioFile_Format)
	*p = x
	return p
}
func (x AudioFile_Format) String() string {
	return proto.EnumName(AudioFile_Format_name, int32(x))
}
func (x *AudioFile_Format) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(AudioFile_Format_value, data, "AudioFile_Format")
	if err != nil {
		return err
	}
	*x = AudioFile_Format(value)
	return nil
}
func (AudioFile_Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{15, 0} }

type TopTracks struct {
	Country          *string  `protobuf:"bytes,1,opt,name=country" json:"country,omitempty"`
	Track            []*Track `protobuf:"bytes,2,rep,name=track" json:"track,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *TopTracks) Reset()                    { *m = TopTracks{} }
func (m *TopTracks) String() string            { return proto.CompactTextString(m) }
func (*TopTracks) ProtoMessage()               {}
func (*TopTracks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TopTracks) GetCountry() string {
	if m != nil && m.Country != nil {
		return *m.Country
	}
	return ""
}

func (m *TopTracks) GetTrack() []*Track {
	if m != nil {
		return m.Track
	}
	return nil
}

type ActivityPeriod struct {
	StartYear        *int32 `protobuf:"zigzag32,1,opt,name=start_year,json=startYear" json:"start_year,omitempty"`
	EndYear          *int32 `protobuf:"zigzag32,2,opt,name=end_year,json=endYear" json:"end_year,omitempty"`
	Decade           *int32 `protobuf:"zigzag32,3,opt,name=decade" json:"decade,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ActivityPeriod) Reset()                    { *m = ActivityPeriod{} }
func (m *ActivityPeriod) String() string            { return proto.CompactTextString(m) }
func (*ActivityPeriod) ProtoMessage()               {}
func (*ActivityPeriod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ActivityPeriod) GetStartYear() int32 {
	if m != nil && m.StartYear != nil {
		return *m.StartYear
	}
	return 0
}

func (m *ActivityPeriod) GetEndYear() int32 {
	if m != nil && m.EndYear != nil {
		return *m.EndYear
	}
	return 0
}

func (m *ActivityPeriod) GetDecade() int32 {
	if m != nil && m.Decade != nil {
		return *m.Decade
	}
	return 0
}

type Artist struct {
	Gid                  []byte            `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name                 *string           `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Popularity           *int32            `protobuf:"zigzag32,3,opt,name=popularity" json:"popularity,omitempty"`
	TopTrack             []*TopTracks      `protobuf:"bytes,4,rep,name=top_track,json=topTrack" json:"top_track,omitempty"`
	AlbumGroup           []*AlbumGroup     `protobuf:"bytes,5,rep,name=album_group,json=albumGroup" json:"album_group,omitempty"`
	SingleGroup          []*AlbumGroup     `protobuf:"bytes,6,rep,name=single_group,json=singleGroup" json:"single_group,omitempty"`
	CompilationGroup     []*AlbumGroup     `protobuf:"bytes,7,rep,name=compilation_group,json=compilationGroup" json:"compilation_group,omitempty"`
	AppearsOnGroup       []*AlbumGroup     `protobuf:"bytes,8,rep,name=appears_on_group,json=appearsOnGroup" json:"appears_on_group,omitempty"`
	Genre                []string          `protobuf:"bytes,9,rep,name=genre" json:"genre,omitempty"`
	ExternalId           []*ExternalId     `protobuf:"bytes,10,rep,name=external_id,json=externalId" json:"external_id,omitempty"`
	Portrait             []*Image          `protobuf:"bytes,11,rep,name=portrait" json:"portrait,omitempty"`
	Biography            []*Biography      `protobuf:"bytes,12,rep,name=biography" json:"biography,omitempty"`
	ActivityPeriod       []*ActivityPeriod `protobuf:"bytes,13,rep,name=activity_period,json=activityPeriod" json:"activity_period,omitempty"`
	Restriction          []*Restriction    `protobuf:"bytes,14,rep,name=restriction" json:"restriction,omitempty"`
	Related              []*Artist         `protobuf:"bytes,15,rep,name=related" json:"related,omitempty"`
	IsPortraitAlbumCover *bool             `protobuf:"varint,16,opt,name=is_portrait_album_cover,json=isPortraitAlbumCover" json:"is_portrait_album_cover,omitempty"`
	PortraitGroup        *ImageGroup       `protobuf:"bytes,17,opt,name=portrait_group,json=portraitGroup" json:"portrait_group,omitempty"`
	XXX_unrecognized     []byte            `json:"-"`
}

func (m *Artist) Reset()                    { *m = Artist{} }
func (m *Artist) String() string            { return proto.CompactTextString(m) }
func (*Artist) ProtoMessage()               {}
func (*Artist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Artist) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Artist) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Artist) GetPopularity() int32 {
	if m != nil && m.Popularity != nil {
		return *m.Popularity
	}
	return 0
}

func (m *Artist) GetTopTrack() []*TopTracks {
	if m != nil {
		return m.TopTrack
	}
	return nil
}

func (m *Artist) GetAlbumGroup() []*AlbumGroup {
	if m != nil {
		return m.AlbumGroup
	}
	return nil
}

func (m *Artist) GetSingleGroup() []*AlbumGroup {
	if m != nil {
		return m.SingleGroup
	}
	return nil
}

func (m *Artist) GetCompilationGroup() []*AlbumGroup {
	if m != nil {
		return m.CompilationGroup
	}
	return nil
}

func (m *Artist) GetAppearsOnGroup() []*AlbumGroup {
	if m != nil {
		return m.AppearsOnGroup
	}
	return nil
}

func (m *Artist) GetGenre() []string {
	if m != nil {
		return m.Genre
	}
	return nil
}

func (m *Artist) GetExternalId() []*ExternalId {
	if m != nil {
		return m.ExternalId
	}
	return nil
}

func (m *Artist) GetPortrait() []*Image {
	if m != nil {
		return m.Portrait
	}
	return nil
}

func (m *Artist) GetBiography() []*Biography {
	if m != nil {
		return m.Biography
	}
	return nil
}

func (m *Artist) GetActivityPeriod() []*ActivityPeriod {
	if m != nil {
		return m.ActivityPeriod
	}
	return nil
}

func (m *Artist) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *Artist) GetRelated() []*Artist {
	if m != nil {
		return m.Related
	}
	return nil
}

func (m *Artist) GetIsPortraitAlbumCover() bool {
	if m != nil && m.IsPortraitAlbumCover != nil {
		return *m.IsPortraitAlbumCover
	}
	return false
}

func (m *Artist) GetPortraitGroup() *ImageGroup {
	if m != nil {
		return m.PortraitGroup
	}
	return nil
}

type AlbumGroup struct {
	Album            []*Album `protobuf:"bytes,1,rep,name=album" json:"album,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *AlbumGroup) Reset()                    { *m = AlbumGroup{} }
func (m *AlbumGroup) String() string            { return proto.CompactTextString(m) }
func (*AlbumGroup) ProtoMessage()               {}
func (*AlbumGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AlbumGroup) GetAlbum() []*Album {
	if m != nil {
		return m.Album
	}
	return nil
}

type Date struct {
	Year             *int32 `protobuf:"zigzag32,1,opt,name=year" json:"year,omitempty"`
	Month            *int32 `protobuf:"zigzag32,2,opt,name=month" json:"month,omitempty"`
	Day              *int32 `protobuf:"zigzag32,3,opt,name=day" json:"day,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Date) Reset()                    { *m = Date{} }
func (m *Date) String() string            { return proto.CompactTextString(m) }
func (*Date) ProtoMessage()               {}
func (*Date) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Date) GetYear() int32 {
	if m != nil && m.Year != nil {
		return *m.Year
	}
	return 0
}

func (m *Date) GetMonth() int32 {
	if m != nil && m.Month != nil {
		return *m.Month
	}
	return 0
}

func (m *Date) GetDay() int32 {
	if m != nil && m.Day != nil {
		return *m.Day
	}
	return 0
}

type Album struct {
	Gid              []byte         `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Artist           []*Artist      `protobuf:"bytes,3,rep,name=artist" json:"artist,omitempty"`
	Typ              *Album_Type    `protobuf:"varint,4,opt,name=typ,enum=Album_Type" json:"typ,omitempty"`
	Label            *string        `protobuf:"bytes,5,opt,name=label" json:"label,omitempty"`
	Date             *Date          `protobuf:"bytes,6,opt,name=date" json:"date,omitempty"`
	Popularity       *int32         `protobuf:"zigzag32,7,opt,name=popularity" json:"popularity,omitempty"`
	Genre            []string       `protobuf:"bytes,8,rep,name=genre" json:"genre,omitempty"`
	Cover            []*Image       `protobuf:"bytes,9,rep,name=cover" json:"cover,omitempty"`
	ExternalId       []*ExternalId  `protobuf:"bytes,10,rep,name=external_id,json=externalId" json:"external_id,omitempty"`
	Disc             []*Disc        `protobuf:"bytes,11,rep,name=disc" json:"disc,omitempty"`
	Review           []string       `protobuf:"bytes,12,rep,name=review" json:"review,omitempty"`
	Copyright        []*Copyright   `protobuf:"bytes,13,rep,name=copyright" json:"copyright,omitempty"`
	Restriction      []*Restriction `protobuf:"bytes,14,rep,name=restriction" json:"restriction,omitempty"`
	Related          []*Album       `protobuf:"bytes,15,rep,name=related" json:"related,omitempty"`
	SalePeriod       []*SalePeriod  `protobuf:"bytes,16,rep,name=sale_period,json=salePeriod" json:"sale_period,omitempty"`
	CoverGroup       *ImageGroup    `protobuf:"bytes,17,opt,name=cover_group,json=coverGroup" json:"cover_group,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Album) Reset()                    { *m = Album{} }
func (m *Album) String() string            { return proto.CompactTextString(m) }
func (*Album) ProtoMessage()               {}
func (*Album) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Album) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Album) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Album) GetArtist() []*Artist {
	if m != nil {
		return m.Artist
	}
	return nil
}

func (m *Album) GetTyp() Album_Type {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return Album_ALBUM
}

func (m *Album) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Album) GetDate() *Date {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Album) GetPopularity() int32 {
	if m != nil && m.Popularity != nil {
		return *m.Popularity
	}
	return 0
}

func (m *Album) GetGenre() []string {
	if m != nil {
		return m.Genre
	}
	return nil
}

func (m *Album) GetCover() []*Image {
	if m != nil {
		return m.Cover
	}
	return nil
}

func (m *Album) GetExternalId() []*ExternalId {
	if m != nil {
		return m.ExternalId
	}
	return nil
}

func (m *Album) GetDisc() []*Disc {
	if m != nil {
		return m.Disc
	}
	return nil
}

func (m *Album) GetReview() []string {
	if m != nil {
		return m.Review
	}
	return nil
}

func (m *Album) GetCopyright() []*Copyright {
	if m != nil {
		return m.Copyright
	}
	return nil
}

func (m *Album) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *Album) GetRelated() []*Album {
	if m != nil {
		return m.Related
	}
	return nil
}

func (m *Album) GetSalePeriod() []*SalePeriod {
	if m != nil {
		return m.SalePeriod
	}
	return nil
}

func (m *Album) GetCoverGroup() *ImageGroup {
	if m != nil {
		return m.CoverGroup
	}
	return nil
}

type Track struct {
	Gid              []byte         `protobuf:"bytes,1,opt,name=gid" json:"gid,omitempty"`
	Name             *string        `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Album            *Album         `protobuf:"bytes,3,opt,name=album" json:"album,omitempty"`
	Artist           []*Artist      `protobuf:"bytes,4,rep,name=artist" json:"artist,omitempty"`
	Number           *int32         `protobuf:"zigzag32,5,opt,name=number" json:"number,omitempty"`
	DiscNumber       *int32         `protobuf:"zigzag32,6,opt,name=disc_number,json=discNumber" json:"disc_number,omitempty"`
	Duration         *int32         `protobuf:"zigzag32,7,opt,name=duration" json:"duration,omitempty"`
	Popularity       *int32         `protobuf:"zigzag32,8,opt,name=popularity" json:"popularity,omitempty"`
	Explicit         *bool          `protobuf:"varint,9,opt,name=explicit" json:"explicit,omitempty"`
	ExternalId       []*ExternalId  `protobuf:"bytes,10,rep,name=external_id,json=externalId" json:"external_id,omitempty"`
	Restriction      []*Restriction `protobuf:"bytes,11,rep,name=restriction" json:"restriction,omitempty"`
	File             []*AudioFile   `protobuf:"bytes,12,rep,name=file" json:"file,omitempty"`
	Alternative      []*Track       `protobuf:"bytes,13,rep,name=alternative" json:"alternative,omitempty"`
	SalePeriod       []*SalePeriod  `protobuf:"bytes,14,rep,name=sale_period,json=salePeriod" json:"sale_period,omitempty"`
	Preview          []*AudioFile   `protobuf:"bytes,15,rep,name=preview" json:"preview,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Track) Reset()                    { *m = Track{} }
func (m *Track) String() string            { return proto.CompactTextString(m) }
func (*Track) ProtoMessage()               {}
func (*Track) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Track) GetGid() []byte {
	if m != nil {
		return m.Gid
	}
	return nil
}

func (m *Track) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Track) GetAlbum() *Album {
	if m != nil {
		return m.Album
	}
	return nil
}

func (m *Track) GetArtist() []*Artist {
	if m != nil {
		return m.Artist
	}
	return nil
}

func (m *Track) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *Track) GetDiscNumber() int32 {
	if m != nil && m.DiscNumber != nil {
		return *m.DiscNumber
	}
	return 0
}

func (m *Track) GetDuration() int32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Track) GetPopularity() int32 {
	if m != nil && m.Popularity != nil {
		return *m.Popularity
	}
	return 0
}

func (m *Track) GetExplicit() bool {
	if m != nil && m.Explicit != nil {
		return *m.Explicit
	}
	return false
}

func (m *Track) GetExternalId() []*ExternalId {
	if m != nil {
		return m.ExternalId
	}
	return nil
}

func (m *Track) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *Track) GetFile() []*AudioFile {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Track) GetAlternative() []*Track {
	if m != nil {
		return m.Alternative
	}
	return nil
}

func (m *Track) GetSalePeriod() []*SalePeriod {
	if m != nil {
		return m.SalePeriod
	}
	return nil
}

func (m *Track) GetPreview() []*AudioFile {
	if m != nil {
		return m.Preview
	}
	return nil
}

type Image struct {
	FileId           []byte      `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	Size             *Image_Size `protobuf:"varint,2,opt,name=size,enum=Image_Size" json:"size,omitempty"`
	Width            *int32      `protobuf:"zigzag32,3,opt,name=width" json:"width,omitempty"`
	Height           *int32      `protobuf:"zigzag32,4,opt,name=height" json:"height,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Image) GetFileId() []byte {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *Image) GetSize() Image_Size {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return Image_DEFAULT
}

func (m *Image) GetWidth() int32 {
	if m != nil && m.Width != nil {
		return *m.Width
	}
	return 0
}

func (m *Image) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

type ImageGroup struct {
	Image            []*Image `protobuf:"bytes,1,rep,name=image" json:"image,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *ImageGroup) Reset()                    { *m = ImageGroup{} }
func (m *ImageGroup) String() string            { return proto.CompactTextString(m) }
func (*ImageGroup) ProtoMessage()               {}
func (*ImageGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ImageGroup) GetImage() []*Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type Biography struct {
	Text             *string       `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	Portrait         []*Image      `protobuf:"bytes,2,rep,name=portrait" json:"portrait,omitempty"`
	PortraitGroup    []*ImageGroup `protobuf:"bytes,3,rep,name=portrait_group,json=portraitGroup" json:"portrait_group,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *Biography) Reset()                    { *m = Biography{} }
func (m *Biography) String() string            { return proto.CompactTextString(m) }
func (*Biography) ProtoMessage()               {}
func (*Biography) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Biography) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *Biography) GetPortrait() []*Image {
	if m != nil {
		return m.Portrait
	}
	return nil
}

func (m *Biography) GetPortraitGroup() []*ImageGroup {
	if m != nil {
		return m.PortraitGroup
	}
	return nil
}

type Disc struct {
	Number           *int32   `protobuf:"zigzag32,1,opt,name=number" json:"number,omitempty"`
	Name             *string  `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Track            []*Track `protobuf:"bytes,3,rep,name=track" json:"track,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Disc) Reset()                    { *m = Disc{} }
func (m *Disc) String() string            { return proto.CompactTextString(m) }
func (*Disc) ProtoMessage()               {}
func (*Disc) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Disc) GetNumber() int32 {
	if m != nil && m.Number != nil {
		return *m.Number
	}
	return 0
}

func (m *Disc) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Disc) GetTrack() []*Track {
	if m != nil {
		return m.Track
	}
	return nil
}

type Copyright struct {
	Typ              *Copyright_Type `protobuf:"varint,1,opt,name=typ,enum=Copyright_Type" json:"typ,omitempty"`
	Text             *string         `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Copyright) Reset()                    { *m = Copyright{} }
func (m *Copyright) String() string            { return proto.CompactTextString(m) }
func (*Copyright) ProtoMessage()               {}
func (*Copyright) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Copyright) GetTyp() Copyright_Type {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return Copyright_P
}

func (m *Copyright) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

type Restriction struct {
	CountriesAllowed   *string           `protobuf:"bytes,2,opt,name=countries_allowed,json=countriesAllowed" json:"countries_allowed,omitempty"`
	CountriesForbidden *string           `protobuf:"bytes,3,opt,name=countries_forbidden,json=countriesForbidden" json:"countries_forbidden,omitempty"`
	Typ                *Restriction_Type `protobuf:"varint,4,opt,name=typ,enum=Restriction_Type" json:"typ,omitempty"`
	CatalogueStr       []string          `protobuf:"bytes,5,rep,name=catalogue_str,json=catalogueStr" json:"catalogue_str,omitempty"`
	XXX_unrecognized   []byte            `json:"-"`
}

func (m *Restriction) Reset()                    { *m = Restriction{} }
func (m *Restriction) String() string            { return proto.CompactTextString(m) }
func (*Restriction) ProtoMessage()               {}
func (*Restriction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Restriction) GetCountriesAllowed() string {
	if m != nil && m.CountriesAllowed != nil {
		return *m.CountriesAllowed
	}
	return ""
}

func (m *Restriction) GetCountriesForbidden() string {
	if m != nil && m.CountriesForbidden != nil {
		return *m.CountriesForbidden
	}
	return ""
}

func (m *Restriction) GetTyp() Restriction_Type {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return Restriction_STREAMING
}

func (m *Restriction) GetCatalogueStr() []string {
	if m != nil {
		return m.CatalogueStr
	}
	return nil
}

type SalePeriod struct {
	Restriction      []*Restriction `protobuf:"bytes,1,rep,name=restriction" json:"restriction,omitempty"`
	Start            *Date          `protobuf:"bytes,2,opt,name=start" json:"start,omitempty"`
	End              *Date          `protobuf:"bytes,3,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *SalePeriod) Reset()                    { *m = SalePeriod{} }
func (m *SalePeriod) String() string            { return proto.CompactTextString(m) }
func (*SalePeriod) ProtoMessage()               {}
func (*SalePeriod) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *SalePeriod) GetRestriction() []*Restriction {
	if m != nil {
		return m.Restriction
	}
	return nil
}

func (m *SalePeriod) GetStart() *Date {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *SalePeriod) GetEnd() *Date {
	if m != nil {
		return m.End
	}
	return nil
}

type ExternalId struct {
	Typ              *string `protobuf:"bytes,1,opt,name=typ" json:"typ,omitempty"`
	Id               *string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ExternalId) Reset()                    { *m = ExternalId{} }
func (m *ExternalId) String() string            { return proto.CompactTextString(m) }
func (*ExternalId) ProtoMessage()               {}
func (*ExternalId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *ExternalId) GetTyp() string {
	if m != nil && m.Typ != nil {
		return *m.Typ
	}
	return ""
}

func (m *ExternalId) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type AudioFile struct {
	FileId           []byte            `protobuf:"bytes,1,opt,name=file_id,json=fileId" json:"file_id,omitempty"`
	Format           *AudioFile_Format `protobuf:"varint,2,opt,name=format,enum=AudioFile_Format" json:"format,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *AudioFile) Reset()                    { *m = AudioFile{} }
func (m *AudioFile) String() string            { return proto.CompactTextString(m) }
func (*AudioFile) ProtoMessage()               {}
func (*AudioFile) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *AudioFile) GetFileId() []byte {
	if m != nil {
		return m.FileId
	}
	return nil
}

func (m *AudioFile) GetFormat() AudioFile_Format {
	if m != nil && m.Format != nil {
		return *m.Format
	}
	return AudioFile_OGG_VORBIS_96
}

func init() {
	proto.RegisterType((*TopTracks)(nil), "TopTracks")
	proto.RegisterType((*ActivityPeriod)(nil), "ActivityPeriod")
	proto.RegisterType((*Artist)(nil), "Artist")
	proto.RegisterType((*AlbumGroup)(nil), "AlbumGroup")
	proto.RegisterType((*Date)(nil), "Date")
	proto.RegisterType((*Album)(nil), "Album")
	proto.RegisterType((*Track)(nil), "Track")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*ImageGroup)(nil), "ImageGroup")
	proto.RegisterType((*Biography)(nil), "Biography")
	proto.RegisterType((*Disc)(nil), "Disc")
	proto.RegisterType((*Copyright)(nil), "Copyright")
	proto.RegisterType((*Restriction)(nil), "Restriction")
	proto.RegisterType((*SalePeriod)(nil), "SalePeriod")
	proto.RegisterType((*ExternalId)(nil), "ExternalId")
	proto.RegisterType((*AudioFile)(nil), "AudioFile")
	proto.RegisterEnum("Album_Type", Album_Type_name, Album_Type_value)
	proto.RegisterEnum("Image_Size", Image_Size_name, Image_Size_value)
	proto.RegisterEnum("Copyright_Type", Copyright_Type_name, Copyright_Type_value)
	proto.RegisterEnum("Restriction_Type", Restriction_Type_name, Restriction_Type_value)
	proto.RegisterEnum("AudioFile_Format", AudioFile_Format_name, AudioFile_Format_value)
}

func init() { proto.RegisterFile("metadata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0x8d, 0x24, 0xea, 0xc2, 0x91, 0x2d, 0xd3, 0xdb, 0xb4, 0x51, 0x6f, 0x49, 0xca, 0x14, 0xa8,
	0x7b, 0x01, 0x9b, 0x3a, 0xb5, 0xd1, 0x3e, 0xca, 0x8a, 0xec, 0x1a, 0x90, 0x6d, 0x61, 0xa5, 0x14,
	0xe9, 0x13, 0x41, 0x8b, 0x1b, 0x79, 0x51, 0x49, 0x24, 0x48, 0x2a, 0x8e, 0xfa, 0x3b, 0x05, 0xfa,
	0x35, 0x7d, 0x68, 0xbf, 0xa1, 0x5f, 0xd1, 0xb7, 0xce, 0x5e, 0xb8, 0xa2, 0xac, 0xb8, 0x49, 0x9e,
	0xb4, 0x73, 0x66, 0x76, 0xc5, 0x9d, 0x39, 0x67, 0x66, 0xa1, 0x35, 0x63, 0x59, 0x10, 0x06, 0x59,
	0xe0, 0xc5, 0x49, 0x94, 0x45, 0x6e, 0x17, 0xec, 0x51, 0x14, 0x8f, 0x92, 0x60, 0xfc, 0x6b, 0x4a,
	0xda, 0x50, 0x1f, 0x47, 0x8b, 0x79, 0x96, 0x2c, 0xdb, 0xa5, 0x87, 0xa5, 0x3d, 0x9b, 0xe6, 0x26,
	0xf9, 0x04, 0xaa, 0x99, 0x88, 0x69, 0x97, 0x1f, 0x56, 0xf6, 0x9a, 0xfb, 0x35, 0x4f, 0xee, 0xa0,
	0x0a, 0x74, 0x2f, 0xa1, 0xd5, 0x19, 0x67, 0xfc, 0x25, 0xcf, 0x96, 0x03, 0x96, 0xf0, 0x28, 0x24,
	0x9f, 0x02, 0xa4, 0x59, 0x90, 0x64, 0xfe, 0x92, 0x05, 0x89, 0x3c, 0x6c, 0x97, 0xda, 0x12, 0xf9,
	0x05, 0x01, 0xf2, 0x21, 0x34, 0xd8, 0x3c, 0x54, 0xce, 0xb2, 0x74, 0xd6, 0xd1, 0x96, 0xae, 0x0f,
	0xa0, 0x16, 0xb2, 0x71, 0x10, 0xb2, 0x76, 0x45, 0x3a, 0xb4, 0xe5, 0xfe, 0x55, 0x85, 0x5a, 0x27,
	0xc9, 0x78, 0x9a, 0x11, 0x07, 0x2a, 0x13, 0x1e, 0xca, 0x53, 0xb7, 0xa8, 0x58, 0x12, 0x02, 0xd6,
	0x3c, 0x98, 0x31, 0x79, 0x96, 0x4d, 0xe5, 0x9a, 0xdc, 0x07, 0x88, 0xa3, 0x78, 0x31, 0x0d, 0x12,
	0xfc, 0x2c, 0x7d, 0x58, 0x01, 0x21, 0x5f, 0x80, 0x9d, 0x45, 0xb1, 0xaf, 0xae, 0x65, 0xc9, 0x6b,
	0x81, 0x67, 0x72, 0x41, 0x1b, 0x99, 0x5e, 0x92, 0x6f, 0xa0, 0x19, 0x4c, 0x2f, 0x17, 0x33, 0x7f,
	0x92, 0x44, 0x8b, 0xb8, 0x5d, 0x95, 0xa1, 0x4d, 0xaf, 0x23, 0xb0, 0x13, 0x01, 0x51, 0x08, 0xcc,
	0x9a, 0x78, 0xb0, 0x95, 0xf2, 0xf9, 0x64, 0xca, 0x74, 0x78, 0x6d, 0x33, 0xbc, 0xa9, 0x02, 0x54,
	0xfc, 0x0f, 0xb0, 0x3b, 0x8e, 0x66, 0x31, 0x9f, 0x06, 0x19, 0x8f, 0xe6, 0x7a, 0x53, 0x7d, 0x73,
	0x93, 0x53, 0x88, 0x52, 0x3b, 0x0f, 0xc0, 0x09, 0xe2, 0x18, 0x73, 0x96, 0xfa, 0x66, 0x63, 0x63,
	0x73, 0x63, 0x4b, 0x07, 0x5d, 0xe8, 0x6d, 0x77, 0xa1, 0x3a, 0x61, 0xf3, 0x84, 0xb5, 0x6d, 0x8c,
	0xb5, 0xa9, 0x32, 0xc4, 0x25, 0xd9, 0xab, 0x8c, 0x25, 0xf3, 0x60, 0xea, 0x63, 0x6e, 0x41, 0x9f,
	0xd3, 0xd3, 0xd8, 0x69, 0x48, 0x81, 0x99, 0x35, 0x71, 0xa1, 0x11, 0x47, 0x09, 0xa6, 0x8e, 0x67,
	0xed, 0xa6, 0x66, 0xc4, 0xe9, 0x2c, 0x98, 0x30, 0x6a, 0x70, 0xb2, 0x07, 0xf6, 0x25, 0x8f, 0x26,
	0x49, 0x10, 0x5f, 0x2d, 0xdb, 0x5b, 0x3a, 0xbf, 0x47, 0x39, 0x42, 0x57, 0x4e, 0x4c, 0xc1, 0x4e,
	0xa0, 0xe9, 0xe3, 0xc7, 0x92, 0x3f, 0xed, 0x6d, 0x19, 0xbf, 0xe3, 0xad, 0xd3, 0x0a, 0xef, 0xb2,
	0x4e, 0x33, 0x0f, 0x9a, 0x09, 0x4b, 0xb3, 0x84, 0x8f, 0x45, 0x5a, 0xda, 0x2d, 0xb9, 0x6b, 0xcb,
	0xa3, 0x2b, 0x8c, 0x16, 0x03, 0xc8, 0x67, 0x50, 0x4f, 0x18, 0xe6, 0x90, 0x85, 0xed, 0x1d, 0x19,
	0x5b, 0xf7, 0x14, 0xa7, 0x68, 0x8e, 0x63, 0x56, 0xef, 0xf1, 0xd4, 0xcf, 0x6f, 0xe1, 0xab, 0xca,
	0x8f, 0xa3, 0x97, 0x2c, 0x69, 0x3b, 0xc8, 0xa1, 0x06, 0xbd, 0xcb, 0xd3, 0x81, 0xf6, 0xca, 0x2c,
	0x77, 0x85, 0x8f, 0xec, 0x43, 0xcb, 0xec, 0x51, 0xa5, 0xd8, 0xc5, 0x68, 0x91, 0x42, 0x99, 0x17,
	0x55, 0x8a, 0xed, 0x3c, 0x44, 0x9a, 0xee, 0x57, 0x00, 0xab, 0x3a, 0x09, 0x89, 0xc9, 0x3f, 0x43,
	0x5e, 0xab, 0x84, 0x4a, 0x1f, 0x55, 0xa0, 0x7b, 0x04, 0xd6, 0x53, 0xfc, 0x3e, 0xc1, 0xf4, 0x82,
	0xa4, 0xe4, 0x5a, 0x54, 0x74, 0x16, 0xcd, 0xb3, 0x2b, 0x2d, 0x25, 0x65, 0x08, 0x95, 0x84, 0x41,
	0x4e, 0x7c, 0xb1, 0x74, 0xff, 0xb1, 0xa0, 0x2a, 0x0f, 0x7d, 0x4b, 0x05, 0x3d, 0x80, 0x5a, 0x20,
	0xb3, 0x83, 0x87, 0xac, 0x25, 0x4b, 0xc3, 0xa8, 0xf2, 0x4a, 0xb6, 0x8c, 0x51, 0x3c, 0xa5, 0xbd,
	0x56, 0x4e, 0x3a, 0x6f, 0xb4, 0x8c, 0x19, 0x15, 0xb8, 0xf8, 0xae, 0x69, 0x70, 0xc9, 0xa6, 0x28,
	0x19, 0x71, 0xa8, 0x32, 0x50, 0xfb, 0x16, 0xf6, 0x1f, 0x86, 0xc2, 0x10, 0xf9, 0xa9, 0x7a, 0xe2,
	0x5a, 0x54, 0x42, 0x37, 0x24, 0x5b, 0xdf, 0x90, 0xac, 0xa1, 0x6e, 0xa3, 0x48, 0x5d, 0x4c, 0x9c,
	0xaa, 0x8f, 0xbd, 0xc6, 0x44, 0x05, 0xbe, 0x23, 0xb1, 0xc5, 0xc7, 0xf1, 0x74, 0xac, 0x49, 0x8d,
	0x1f, 0x87, 0x06, 0x95, 0x90, 0x68, 0x4c, 0x09, 0x7b, 0xc9, 0xd9, 0xb5, 0x24, 0xb3, 0x4d, 0xb5,
	0x25, 0x78, 0x3e, 0x8e, 0xe2, 0x65, 0xc2, 0x27, 0x57, 0x99, 0xe6, 0x2d, 0x78, 0xdd, 0x1c, 0xa1,
	0x2b, 0xe7, 0x3b, 0xb3, 0xf5, 0xe1, 0x4d, 0xb6, 0xe6, 0x9c, 0x30, 0x64, 0xc5, 0xcb, 0xa5, 0x01,
	0xb6, 0x1a, 0xad, 0x1a, 0x47, 0x5f, 0x6e, 0x88, 0x98, 0x56, 0x0c, 0xa4, 0x66, 0x2d, 0xa2, 0x65,
	0x4e, 0x6e, 0x27, 0x28, 0x48, 0xbf, 0x62, 0xe7, 0x21, 0x58, 0xa2, 0x94, 0xc4, 0x46, 0xd2, 0xf4,
	0x8f, 0x9e, 0x9d, 0x39, 0x25, 0x02, 0x50, 0x1b, 0x9e, 0x9e, 0x9f, 0xf4, 0x7b, 0x4e, 0x99, 0xec,
	0x40, 0xb3, 0x7b, 0x71, 0x36, 0x38, 0xed, 0x77, 0x46, 0xa7, 0x17, 0xe7, 0x4e, 0x85, 0xd4, 0xa0,
	0xdc, 0x1b, 0x38, 0x96, 0xfb, 0x6f, 0x05, 0xaa, 0xaa, 0x71, 0xbe, 0x1d, 0xcb, 0x0c, 0xef, 0x2b,
	0xf2, 0x7b, 0xd6, 0x79, 0x5f, 0xe0, 0xa0, 0xf5, 0x7a, 0x0e, 0x62, 0x59, 0xe6, 0x8b, 0xd9, 0x25,
	0x96, 0xbf, 0xaa, 0xe6, 0x85, 0xb2, 0x70, 0x63, 0x53, 0x94, 0xcd, 0xd7, 0xce, 0x9a, 0x22, 0x93,
	0x80, 0xce, 0x55, 0xc0, 0x47, 0xd0, 0x08, 0x17, 0x89, 0xec, 0xa7, 0x9a, 0x6a, 0xc6, 0xbe, 0x41,
	0xc4, 0xc6, 0x06, 0x11, 0x71, 0x2f, 0x7b, 0x15, 0x4f, 0xf9, 0x18, 0xfb, 0x9f, 0x2d, 0xbb, 0x82,
	0xb1, 0xdf, 0x91, 0x70, 0x37, 0x38, 0xd1, 0x7c, 0x13, 0x27, 0xee, 0x83, 0xf5, 0x82, 0x4f, 0x99,
	0x69, 0xa8, 0x9d, 0x45, 0xc8, 0xa3, 0x63, 0x44, 0xa8, 0xc4, 0x91, 0x8d, 0x38, 0xac, 0xe4, 0xe9,
	0xd8, 0x28, 0x99, 0xe6, 0x63, 0x3e, 0xae, 0x8b, 0xae, 0x9b, 0xdc, 0x69, 0xfd, 0x3f, 0x77, 0x3e,
	0x87, 0x7a, 0xac, 0xe9, 0xbf, 0xb3, 0xf1, 0xd7, 0xb9, 0xcb, 0xfd, 0xa3, 0x04, 0x55, 0x49, 0x27,
	0x72, 0x0f, 0xea, 0xe2, 0x7b, 0x7c, 0x53, 0xff, 0x9a, 0x30, 0xf1, 0xc2, 0x0f, 0xc0, 0x4a, 0xf9,
	0x6f, 0x8a, 0x02, 0xad, 0x9c, 0x7d, 0xde, 0x10, 0x21, 0x2a, 0x1d, 0x42, 0xe4, 0xd7, 0x3c, 0xc4,
	0x6e, 0xa6, 0x3a, 0x97, 0x32, 0x44, 0x99, 0xaf, 0x98, 0x94, 0x98, 0xa5, 0xca, 0xac, 0x2c, 0xf7,
	0x00, 0x2c, 0xb1, 0x97, 0x34, 0xa1, 0xfe, 0xb4, 0x77, 0xdc, 0x79, 0xd6, 0x1f, 0x39, 0x77, 0x04,
	0x65, 0x87, 0x67, 0x9d, 0x7e, 0x1f, 0x29, 0x8b, 0xcb, 0x7e, 0x87, 0x9e, 0x08, 0xc6, 0x22, 0x7b,
	0x9f, 0xab, 0x75, 0x45, 0xb4, 0xde, 0x15, 0xed, 0x05, 0x05, 0xb9, 0xb0, 0x4c, 0xeb, 0xd5, 0x1d,
	0x44, 0x82, 0xee, 0x35, 0xd8, 0x66, 0x6c, 0x09, 0x06, 0x67, 0x58, 0x3e, 0xfd, 0x3e, 0x92, 0xeb,
	0xb5, 0x69, 0x58, 0xbe, 0x65, 0x1a, 0x6e, 0xce, 0x87, 0x8a, 0x4e, 0xf8, 0xed, 0xf3, 0x61, 0x80,
	0x3d, 0x5f, 0x77, 0x1e, 0xcd, 0xe2, 0xd2, 0x1a, 0xc5, 0x6f, 0x51, 0x93, 0x7a, 0xd1, 0x54, 0x5e,
	0xf7, 0x50, 0x7b, 0x0e, 0xb6, 0xe9, 0x4c, 0x38, 0x0c, 0x65, 0xf7, 0x2e, 0xc9, 0x42, 0xec, 0xac,
	0x5a, 0x56, 0xa1, 0x83, 0xe7, 0xb7, 0x2d, 0xaf, 0x6e, 0xeb, 0xde, 0xd5, 0x7d, 0xa1, 0x0a, 0xa5,
	0x01, 0xe6, 0x1a, 0x7f, 0xba, 0x4e, 0xc9, 0xfd, 0xb3, 0x04, 0xcd, 0x02, 0x69, 0xc9, 0xd7, 0xe2,
	0x59, 0x23, 0xde, 0x8e, 0x9c, 0xa5, 0x38, 0x44, 0xa7, 0xd1, 0x35, 0x76, 0x31, 0x75, 0x8c, 0x63,
	0x1c, 0x1d, 0x85, 0x93, 0x6f, 0xe1, 0xbd, 0x55, 0xf0, 0x8b, 0x28, 0xb9, 0xe4, 0x61, 0xc8, 0xe6,
	0x92, 0x00, 0x36, 0x25, 0xc6, 0x75, 0x9c, 0x7b, 0xc8, 0xa3, 0xe2, 0xe0, 0xd9, 0x2d, 0xaa, 0xa5,
	0xf0, 0xf1, 0x8f, 0x60, 0x7b, 0x8c, 0x0f, 0xdd, 0x69, 0x34, 0x59, 0x30, 0x1f, 0x03, 0xe4, 0xcb,
	0xcd, 0xa6, 0x5b, 0x06, 0x1c, 0x66, 0x89, 0xfb, 0xbe, 0xbe, 0xcd, 0x36, 0xd8, 0xc3, 0x11, 0xed,
	0x75, 0xce, 0xb0, 0xbf, 0x39, 0x77, 0xdc, 0x04, 0x60, 0x25, 0x84, 0x9b, 0x22, 0x2d, 0xbd, 0x49,
	0xa4, 0x1f, 0x43, 0x55, 0xbe, 0x75, 0xe5, 0x85, 0xcd, 0x8c, 0x53, 0x18, 0x2a, 0xa3, 0x82, 0x6f,
	0x5d, 0xdd, 0xed, 0xb4, 0x4b, 0x20, 0xae, 0x07, 0xb0, 0x6a, 0x12, 0xa2, 0x79, 0xe6, 0xd5, 0xb1,
	0xd5, 0x7d, 0x5a, 0x50, 0xe6, 0x79, 0x0e, 0x71, 0xe5, 0xfe, 0x5e, 0x06, 0xdb, 0x68, 0xf0, 0x76,
	0xc1, 0x7d, 0x09, 0x35, 0x4c, 0xe9, 0x2c, 0xc8, 0xb4, 0xe4, 0x76, 0x57, 0xc2, 0xf5, 0x8e, 0xa5,
	0x83, 0xea, 0x00, 0xf7, 0xef, 0x12, 0xd4, 0x14, 0x44, 0x76, 0x61, 0xfb, 0xe2, 0xe4, 0xc4, 0xff,
	0xf9, 0x82, 0x1e, 0x9d, 0x0e, 0xfd, 0x1f, 0x0f, 0xb1, 0xd2, 0x04, 0x5a, 0x05, 0xe8, 0xbb, 0xc3,
	0xc7, 0x28, 0xaf, 0x75, 0xec, 0xc9, 0xfe, 0x63, 0xd4, 0x19, 0x4a, 0xf1, 0x6c, 0xf0, 0xc4, 0xdf,
	0x3f, 0x38, 0xc4, 0xa9, 0xa0, 0x0d, 0xe1, 0xb1, 0x72, 0x43, 0x6c, 0xad, 0x0a, 0x39, 0x0a, 0x03,
	0x8f, 0xae, 0x89, 0x61, 0xa2, 0x1d, 0x7e, 0xef, 0xbc, 0xeb, 0xd4, 0x85, 0xf3, 0x62, 0xf4, 0x53,
	0x8f, 0xee, 0x3b, 0x0d, 0xb3, 0x7e, 0xe2, 0xd8, 0xe2, 0x84, 0x4e, 0xa7, 0x2b, 0x4f, 0x80, 0xdc,
	0x10, 0x67, 0x37, 0x4d, 0xd4, 0xf7, 0xce, 0x96, 0x59, 0x1f, 0x38, 0xdb, 0xff, 0x05, 0x00, 0x00,
	0xff, 0xff, 0xe2, 0x37, 0x14, 0x02, 0xf1, 0x0c, 0x00, 0x00,
}
