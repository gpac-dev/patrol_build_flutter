package build_parameters

import (
	"errors"
	"fmt"
	"strings"
)

// SetPlatform sets the build platform. Accepted: "android", "iOS".
func SetPlatform(bp *BuildParameters, value string) error {
	switch strings.ToLower(value) {
	case "android", "ios":
		bp.Platform = value
		return nil
	default:
		return errors.New("invalid platform: expected 'android' or 'iOS'")
	}
}

// SetTargets sets the target list. Required and must not be empty.
func SetTargets(bp *BuildParameters, value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("targets cannot be empty")
	}
	bp.Targets = value
	return nil
}

func SetBuildType(bp *BuildParameters, value string) error {
	switch value {
	case "emulator", "physical":
		bp.BuildType = value
		return nil
	default:
		return errors.New("invalid build type: expected 'emulator' or 'physical'")
	}
}

func SetTags(bp *BuildParameters, value string) error {
	bp.Tags = formatTags(value)
	return nil
}

func SetExcludedTags(bp *BuildParameters, value string) error {
	bp.ExcludedTags = formatTags(value)
	return nil
}

func SetVerbose(bp *BuildParameters, value string) error {
	return setFlag(value, "--verbose", &bp.IsVerbose, "verbose")
}

func SetCoverage(bp *BuildParameters, value string) error {
	return setFlag(value, "--covered", &bp.IsCoverage, "isCoverage")
}

// formatTags converts comma-separated values to '( tag1 && tag2 )' format.
func formatTags(input string) string {
	tags := strings.Split(input, ",")
	var trimmed []string
	for _, t := range tags {
		if tag := strings.TrimSpace(t); tag != "" {
			trimmed = append(trimmed, tag)
		}
	}
	if len(trimmed) == 0 {
		return ""
	}
	return "'( " + strings.Join(trimmed, " && ") + " )'"
}

func setFlag(value, flag string, target *string, name string) error {
	switch strings.ToLower(value) {
	case "true":
		*target = flag
	case "false":
		*target = ""
	default:
		return fmt.Errorf("invalid value for %s: expected 'true' or 'false'", name)
	}
	return nil
}
