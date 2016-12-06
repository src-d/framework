package email

import "strings"

// IsIgnored returns true for emails that should be
// ignored because they do not correspond to individuals
// such as example emails or local machine addresses.
func IsIgnored(s string) bool {
	return isIgnoredEmail(s) ||
		isIgnoredDomain(s) ||
		isIgnoredTLD(s) ||
		isSingleLabelDomain(s)
}

var ignoredEmails = map[string]bool{
	"nobody@android.com": true,
	"badger@gitter.im":   true,
}

func isIgnoredEmail(s string) bool {
	_, found := ignoredEmails[s]
	return found
}

var ignoredDomains = map[string]bool{
	"localhost.localdomain": true,
	"example.com":           true,
	"test.com":              true,
}

func isIgnoredDomain(s string) bool {
	parts := strings.Split(s, "@")
	_, found := ignoredDomains[parts[len(parts)-1]]
	return found
}

var ignoredTLD = map[string]bool{
	"localhost":   true,
	"localdomain": true,
	"local":       true,
	"test":        true,
}

func isIgnoredTLD(s string) bool {
	parts := strings.Split(s, ".")
	_, found := ignoredTLD[parts[len(parts)-1]]
	return found
}

func isSingleLabelDomain(s string) bool {
	parts := strings.Split(s, "@")
	last := parts[len(parts)-1]
	labels := strings.Split(last, ".")
	return len(labels) == 1
}
