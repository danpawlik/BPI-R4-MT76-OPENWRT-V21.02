// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package webrisk aliases all exported identifiers in package
// "cloud.google.com/go/webrisk/apiv1/webriskpb".
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package webrisk

import (
	src "cloud.google.com/go/webrisk/apiv1/webriskpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/webrisk/apiv1/webriskpb
const (
	CompressionType_COMPRESSION_TYPE_UNSPECIFIED            = src.CompressionType_COMPRESSION_TYPE_UNSPECIFIED
	CompressionType_RAW                                     = src.CompressionType_RAW
	CompressionType_RICE                                    = src.CompressionType_RICE
	ComputeThreatListDiffResponse_DIFF                      = src.ComputeThreatListDiffResponse_DIFF
	ComputeThreatListDiffResponse_RESET                     = src.ComputeThreatListDiffResponse_RESET
	ComputeThreatListDiffResponse_RESPONSE_TYPE_UNSPECIFIED = src.ComputeThreatListDiffResponse_RESPONSE_TYPE_UNSPECIFIED
	ThreatType_MALWARE                                      = src.ThreatType_MALWARE
	ThreatType_SOCIAL_ENGINEERING                           = src.ThreatType_SOCIAL_ENGINEERING
	ThreatType_SOCIAL_ENGINEERING_EXTENDED_COVERAGE         = src.ThreatType_SOCIAL_ENGINEERING_EXTENDED_COVERAGE
	ThreatType_THREAT_TYPE_UNSPECIFIED                      = src.ThreatType_THREAT_TYPE_UNSPECIFIED
	ThreatType_UNWANTED_SOFTWARE                            = src.ThreatType_UNWANTED_SOFTWARE
)

// Deprecated: Please use vars in: cloud.google.com/go/webrisk/apiv1/webriskpb
var (
	CompressionType_name                             = src.CompressionType_name
	CompressionType_value                            = src.CompressionType_value
	ComputeThreatListDiffResponse_ResponseType_name  = src.ComputeThreatListDiffResponse_ResponseType_name
	ComputeThreatListDiffResponse_ResponseType_value = src.ComputeThreatListDiffResponse_ResponseType_value
	File_google_cloud_webrisk_v1_webrisk_proto       = src.File_google_cloud_webrisk_v1_webrisk_proto
	ThreatType_name                                  = src.ThreatType_name
	ThreatType_value                                 = src.ThreatType_value
)

// The ways in which threat entry sets can be compressed.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type CompressionType = src.CompressionType

// Describes an API diff request.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type ComputeThreatListDiffRequest = src.ComputeThreatListDiffRequest

// The constraints for this diff.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type ComputeThreatListDiffRequest_Constraints = src.ComputeThreatListDiffRequest_Constraints
type ComputeThreatListDiffResponse = src.ComputeThreatListDiffResponse

// The expected state of a client's local database.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type ComputeThreatListDiffResponse_Checksum = src.ComputeThreatListDiffResponse_Checksum

// The type of response sent to the client.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type ComputeThreatListDiffResponse_ResponseType = src.ComputeThreatListDiffResponse_ResponseType

// Request to send a potentially phishy URI to WebRisk.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type CreateSubmissionRequest = src.CreateSubmissionRequest

// The uncompressed threat entries in hash format. Hashes can be anywhere from
// 4 to 32 bytes in size. A large majority are 4 bytes, but some hashes are
// lengthened if they collide with the hash of a popular URI. Used for sending
// ThreatEntryAdditons to clients that do not support compression, or when
// sending non-4-byte hashes to clients that do support compression.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type RawHashes = src.RawHashes

// A set of raw indices to remove from a local list.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type RawIndices = src.RawIndices

// The Rice-Golomb encoded data. Used for sending compressed 4-byte hashes or
// compressed removal indices.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type RiceDeltaEncoding = src.RiceDeltaEncoding

// Request to return full hashes matched by the provided hash prefixes.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type SearchHashesRequest = src.SearchHashesRequest
type SearchHashesResponse = src.SearchHashesResponse

// Contains threat information on a matching hash.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type SearchHashesResponse_ThreatHash = src.SearchHashesResponse_ThreatHash

// Request to check URI entries against threatLists.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type SearchUrisRequest = src.SearchUrisRequest
type SearchUrisResponse = src.SearchUrisResponse

// Contains threat information on a matching uri.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type SearchUrisResponse_ThreatUri = src.SearchUrisResponse_ThreatUri

// Wraps a URI that might be displaying malicious content.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type Submission = src.Submission

// Contains the set of entries to add to a local database. May contain a
// combination of compressed and raw data in a single response.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type ThreatEntryAdditions = src.ThreatEntryAdditions

// Contains the set of entries to remove from a local database.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type ThreatEntryRemovals = src.ThreatEntryRemovals

// The type of threat. This maps directly to the threat list a threat may
// belong to.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type ThreatType = src.ThreatType

// UnimplementedWebRiskServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type UnimplementedWebRiskServiceServer = src.UnimplementedWebRiskServiceServer

// WebRiskServiceClient is the client API for WebRiskService service. For
// semantics around ctx use and closing/ending streaming RPCs, please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type WebRiskServiceClient = src.WebRiskServiceClient

// WebRiskServiceServer is the server API for WebRiskService service.
//
// Deprecated: Please use types in: cloud.google.com/go/webrisk/apiv1/webriskpb
type WebRiskServiceServer = src.WebRiskServiceServer

// Deprecated: Please use funcs in: cloud.google.com/go/webrisk/apiv1/webriskpb
func NewWebRiskServiceClient(cc grpc.ClientConnInterface) WebRiskServiceClient {
	return src.NewWebRiskServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/webrisk/apiv1/webriskpb
func RegisterWebRiskServiceServer(s *grpc.Server, srv WebRiskServiceServer) {
	src.RegisterWebRiskServiceServer(s, srv)
}
