package util

import (
	"regexp"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Convert a string to kebab-case
func KebabCase(str string) string {
	// Replace non-alphanumeric characters with hyphens
	reAlnumdash := regexp.MustCompile(`[^a-zA-Z0-9-]`)
	str = reAlnumdash.ReplaceAllString(str, "-")
	// Insert hyphen between lowercase and uppercase letters
	reHyphenBetween := regexp.MustCompile(`([a-z])([A-Z])`)
	str = reHyphenBetween.ReplaceAllString(str, "$1-$2")
	// Replace whitespace with hyphens
	reReplaceSpaces := regexp.MustCompile(`[\s]+`)
	str = reReplaceSpaces.ReplaceAllString(str, "-")
	// Replace multiple hyphens with a single hyphen
	reStripHyphens := regexp.MustCompile(`[-]+`)
	str = reStripHyphens.ReplaceAllString(str, "-")
	// Convert to lowercase
	str = strings.ToLower(str)
	// Trim leading hyphens
	str = strings.TrimPrefix(str, "-")
	// Trim trailing hyphens
	str = strings.TrimSuffix(str, "-")
	return str
}

// Convert a string to camelCase
func CamelCase(str string) string {
	// Replace non-alphanumeric characters with spaces
	reAlnumSpace := regexp.MustCompile(`[^a-zA-Z0-9]`)
	str = reAlnumSpace.ReplaceAllString(str, " ")
	// Convert to lowercase
	str = strings.ToLower(str)
	// Split string into words
	words := strings.Fields(str)
	// Capitalize first letter of each word except the first word
	for i, word := range words {
		if i > 0 {
			words[i] = cases.Title(language.Tag{}).String(word)
		}
	}
	// Join words into a single string
	str = strings.Join(words, "")
	return str
}

// Output a formatted resource name, kebab-case, 140 characters or less, prefixed with project's name
// Should be used for most resources
func FormatResourceName(ctx *pulumi.Context, name string) string {
	resourceName := KebabCase(ctx.Project() + "-" + name)
	if len(resourceName) > 140 {
		resourceName = resourceName[:140]
	}
	return resourceName
}

// Output a formatted resource name, kebab-case, 63 characters or less, prefixed with project's name
// Should be used resources with fewwer allowed characters in their names such as S3, RDS, ...
func FormatResourceNameShort(ctx *pulumi.Context, name string) string {
	resourceNameLong := FormatResourceName(ctx, name)
	if len(resourceNameLong) > 63 {
		return resourceNameLong[:63]
	}
	return resourceNameLong
}
