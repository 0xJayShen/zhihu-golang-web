
package helpers

import (
	"fmt"
)

// KafkaproducerVersionNumber represents the current build version.
// This should be the only one
const (
	// Major and minor version.
	ServerVersionNumber = 1.0

	// Increment this for bug releases
	ServerPatchVersion = 0
)

// KafkaproducerVersionSuffix is the suffix used in the Kafkaproducer version string.
// It will be blank for release versions.
// const KafkaproducerVersionSuffix = "-DEV" // use this when not doing a release
const ServerVersionSuffix = "-Release" // use this line when doing a release

// KafkaproducerVersion returns the current Kafkaproducer version. It will include
// a suffix, typically '-DEV', if it's development version.
func ServerVersion() string {
	return serverVersion(ServerVersionNumber, ServerPatchVersion, ServerVersionSuffix)
}

// KafkaproducerReleaseVersion is same as KafkaproducerVersion, but no suffix.
func ServerReleaseVersion() string {
	return serverVersionNoSuffix(ServerVersionNumber, ServerPatchVersion)
}

// NextKafkaproducerReleaseVersion returns the next Kafkaproducer release version.
func NextKafkaproducerReleaseVersion() string {
	return serverVersionNoSuffix(ServerVersionNumber+0.01, 0)
}

func serverVersion(version float32, patchVersion int, suffix string) string {
	if patchVersion > 0 {
		return fmt.Sprintf("%.2g.%d%s", version, patchVersion, suffix)
	}
	return fmt.Sprintf("%.2g%s", version, suffix)
}

func serverVersionNoSuffix(version float32, patchVersion int) string {
	if patchVersion > 0 {
		return fmt.Sprintf("%.2g.%d", version, patchVersion)
	}
	return fmt.Sprintf("%.2g", version)
}
