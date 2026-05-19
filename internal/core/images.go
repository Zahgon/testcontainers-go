package core

import (
	"io"
	"regexp"
)

const (
	IndexDockerIO   = "https://index.docker.io/v1/"
	maxURLRuneCount = 2083
	minURLRuneCount = 3
	URLSchema       = `((ftp|tcp|udp|wss?|https?):\/\/)`
	URLUsername     = `(\S+(:\S*)?@)`
	URLIP           = `([1-9]\d?|1\d\d|2[01]\d|22[0-3]|24\d|25[0-5])(\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-5]))`
	IP              = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	URLSubdomain    = `((www\.)|([a-zA-Z0-9]+([-_\.]?[a-zA-Z0-9])*[a-zA-Z0-9]\.[a-zA-Z0-9]+))`
	URLPath         = `((\/|\?|#)[^\s]*)`
	URLPort         = `(:(\d{1,5}))`
	URL             = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-_]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
)

var rxURL = regexp.MustCompile(URL)

// ExtractImagesFromDockerfile extracts images from the Dockerfile sourced from dockerfile.
func ExtractImagesFromDockerfile(dockerfile string, buildArgs map[string]*string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractImagesFromReader extracts images from the Dockerfile sourced from r.
func ExtractImagesFromReader(r io.Reader, buildArgs map[string]*string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extract images from dockerfile

// remove FROM

// interpolate build args

// ExtractRegistry extracts the registry from the image name, using a regular expression to extract the registry from the image name.
// regular expression to extract the registry from the image name
// the regular expression is based on the grammar defined in
// - image:tag
// - image
// - repository/image:tag
// - repository/image
// - registry/image:tag
// - registry/image
// - registry/repository/image:tag
// - registry/repository/image
// - registry:port/repository/image:tag
// - registry:port/repository/image
// - registry:port/image:tag
// - registry:port/image
// Once extracted the registry, it is validated to check if it is a valid URL or an IP address.
func ExtractRegistry(image string, fallback string) string { _ = "STUB: not implemented"; return "" }

// docker.io is an implicit reference, return fallback for normalization

// registry.hub.docker.com is an explicit registry reference, preserve it

// IsURL checks if the string is a URL.
// Extracted from https://github.com/asaskevich/govalidator/blob/f21760c49a8d/validator.go#L104
func IsURL(str string) bool { _ = "STUB: not implemented"; return false }

// support no indicated urlscheme but with colon for port number
// http:// is appended so url.Parse will succeed, strTemp used so it does not impact rxURL.MatchString
