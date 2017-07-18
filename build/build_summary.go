// Copyright 2017 Google, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package build

// StartStep is a build step WaitFor dependency that is always satisfied.
const StartStep = "-"

// BuildStatus is a pseudo-enum of valid build states.
type BuildStatus string

const (
	// StatusFetchSource - Fetching source.
	StatusFetchSource BuildStatus = "FETCHSOURCE"
	// StatusBuild - Executing the build step images on the source.
	StatusBuild BuildStatus = "BUILD"
	// StatusPush - Pushing the resultant image to GCR.
	StatusPush BuildStatus = "PUSH"
	// StatusDone - Build completed successfully.
	StatusDone BuildStatus = "DONE"
	// StatusError - Build failed.
	StatusError BuildStatus = "ERROR"
)

// HashType is a pseudo-enum of valid SourceHashes.
type HashType string

// These constants match those in pb cloudbuild.proto pb.Hash.HashType.
const (
	// MD5 indicates hash type md5.
	MD5 HashType = "MD5"
	// SHA256 indicates hash type sha256.
	SHA256 HashType = "SHA256"
	// No hash requested.
	NONE HashType = "NONE"
)

// Hash captures a hash by HashType and Value.
type Hash struct {
	// Type of the hash.
	Type HashType
	// Value of the hash.
	Value []byte
}

// BuildSummary is the data returned by the blocking /build/status endpoint.
type BuildSummary struct {
	Status          BuildStatus
	BuiltImages     []BuiltImage
	BuildStepImages []string // index of build step -> digest, else empty string
	FileHashes      map[string][]Hash
}

// BuiltImage is information about an image that resulted from this build and
// was pushed to an image registry.
type BuiltImage struct {
	// Name is the full name of an image, as given to 'docker pull $NAME'.
	Name string
	// Digest is the digest reported by registry2.0, if available. If not
	// available, it will be the empty string.
	Digest string
}
