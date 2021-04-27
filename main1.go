package cloudinary

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"gopkg.in/mgo.v2"
)

type ResourceType int

const (
	ImageType ResourceType = iota
	PdfType
	VideoType
	RawType
)

type Service struct {
	cloudName        string
	apiKey           string
	apiSecret        string
	uploadURI        *url.URL     // To upload resources
	adminURI         *url.URL     // To use the admin API
	uploadResType    ResourceType // Upload resource type
	basePathDir      string       // Base path directory
	prependPath      string       // Remote prepend path
	verbose          bool
	simulate         bool // Dry run (NOP)
	keepFilesPattern *regexp.Regexp

	mongoDbURI *url.URL // Can be nil: checksum checks are disabled
	dbSession  *mgo.Session
	col        *mgo.Collection
}

// Resource holds information about an image or a raw file.
type Resource struct {
	PublicId     string `json:"public_id"`
	Version      int    `json:"version"`
	ResourceType string `json:"resource_type"` // image or raw
	Size         int    `json:"bytes"`         // In bytes
	Url          string `json:"url"`           // Remote url
	SecureUrl    string `json:"secure_url"`    // Over https
}

type pagination struct {
	NextCursor int64 `json: "next_cursor"`
}

type resourceList struct {
	pagination
	Resources []*Resource `json: "resources"`
}

type ResourceDetails struct {
	PublicId     string     `json:"public_id"`
	Format       string     `json:"format"`
	Version      int        `json:"version"`
	ResourceType string     `json:"resource_type"` // image or raw
	Size         int        `json:"bytes"`         // In bytes
	Width        int        `json:"width"`         // Width
	Height       int        `json:"height"`        // Height
	Url          string     `json:"url"`           // Remote url
	SecureUrl    string     `json:"secure_url"`    // Over https
	Derived      []*Derived `json:"derived"`       // Derived
}

type Derived struct {
	Transformation string `json:"transformation"` // Transformation
	Size           int    `json:"bytes"`          // In bytes
	Url            string `json:"url"`            // Remote url
}

// Upload response after uploading a file.
type uploadResponse struct {
	Id           string `bson:"_id"`
	PublicId     string `json:"public_id"`
	Version      uint   `json:"version"`
	Format       string `json:"format"`
	ResourceType string `json:"resource_type"` // "image" or "raw"
	Size         int    `json:"bytes"`         // In bytes
	Checksum     string // SHA1 Checksum
}

func Dial(uri string) (*Service, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if u.Scheme != "cloudinary" {
		return nil, errors.New("Missing cloudinary:// scheme in URI")
	}
	secret, exists := u.User.Password()
	if !exists {
		return nil, errors.New("No API secret provided in URI.")
	}
	s := &Service{
		cloudName:     u.Host,
		apiKey:        u.User.Username(),
		apiSecret:     secret,
		uploadResType: ImageType,
		simulate:      false,
		verbose:       false,
	}
	// Default upload URI to the service. Can change at runtime in the
	// Upload() function for raw file uploading.
	up, err := url.Parse(fmt.Sprintf("%s/%s/image/upload/", baseUploadUrl, s.cloudName))
	if err != nil {
		return nil, err
	}
	s.uploadURI = up

	// Admin API url
	adm, err := url.Parse(fmt.Sprintf("%s/%s", baseAdminUrl, s.cloudName))
	if err != nil {
		return nil, err
	}
	adm.User = url.UserPassword(s.apiKey, s.apiSecret)
	s.adminURI = adm
	return s, nil
}

