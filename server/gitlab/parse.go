package gitlab

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sganon/slack-dat-changelog/slack"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Changelog struct {
	Version string
	Project string
	Added   []string
	Changed []string
	Removed []string
	Found   bool
}

var (
	ErrNoChangelog = errors.New("no changelog found")
)

func ParseChangelog(content []byte, version string) (changelog Changelog) {
	version = strings.Trim(version, "v")
	md := blackfriday.New()
	node := md.Parse(content)
	node = node.FirstChild
	found := false
	var currentKind *[]string
	for node.Next != nil {
		// Parse Heading block to determine block of the version or the kind
		if node.Type == blackfriday.Heading {
			if node.HeadingData.Level == 2 && !found {
				// We found an ## and it contains our version
				if i := strings.Index(string(node.FirstChild.Literal), version); i > -1 {
					found = true
				}
			} else if node.HeadingData.Level == 2 && found {
				// Our block version has already be found and now we have found next one so break
				break
			} else if node.HeadingData.Level == 3 && found {
				// Determine which changelog slice to complete
				kind := strings.Trim(string(node.FirstChild.Literal), " ")
				switch kind {
				case "Added":
					currentKind = &changelog.Added
				case "Changed":
					currentKind = &changelog.Changed
				case "Removed":
					currentKind = &changelog.Removed
				}
			}
		}

		if node.Type == blackfriday.List && found && currentKind != nil {
			item := node.FirstChild
			for item != nil {
				text := string(item.FirstChild.FirstChild.Literal)
				*currentKind = append(*currentKind, text)
				item = item.Next
			}
		}
		node = node.Next
	}
	changelog.Version = version
	changelog.Found = found
	return changelog
}

func (c Changelog) GeneratePayload() (payload slack.Payload) {
	var color string
	var Fields []slack.Field
	text := fmt.Sprintf("New release *v%s* of project %s", c.Version, c.Project)
	if !c.Found {
		color = "danger"
		Fields = []slack.Field{
			slack.Field{
				Title: "Changelog not found",
				Value: "No changelog for this version is present on CHANGELOG.md",
			},
		}
	} else if c.Found && len(c.Added) == 0 && len(c.Changed) == 0 && len(c.Removed) == 0 {
		color = "warning"
		Fields = []slack.Field{
			slack.Field{
				Title: "Changelog empty",
				Value: "A changelog has been found for this version but it seems empty, maybe it's just malformed",
			},
		}
	} else {
		color = "good"
		added := getField("Added", c.Added)
		changed := getField("Changed", c.Changed)
		removed := getField("Removed", c.Removed)
		Fields = append(Fields, added...)
		Fields = append(Fields, changed...)
		Fields = append(Fields, removed...)
	}
	payload.Attachments = []slack.Attachment{
		{
			Fallback: text,
			Pretext:  text,
			Color:    color,
			Fields:   Fields,
		},
	}
	return payload
}

func getField(kind string, values []string) (fields []slack.Field) {
	for _, v := range values {
		fields = append(fields, slack.Field{
			Title: kind,
			Value: v,
		})
	}
	return fields
}
