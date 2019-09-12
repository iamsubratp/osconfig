//  Copyright 2019 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package recipes

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"path/filepath"

	"cloud.google.com/go/storage"
	osconfigpb "github.com/GoogleCloudPlatform/osconfig/_internal/gapi-cloud-osconfig-go/google.golang.org/genproto/googleapis/cloud/osconfig/v1alpha2"
	"github.com/GoogleCloudPlatform/osconfig/external"
)

// FetchArtifacts takes in a slice of artifacts and downloads them into the specified directory,
// Returns a map of artifact names to their new locations on the local disk.
func FetchArtifacts(ctx context.Context, artifacts []*osconfigpb.SoftwareRecipe_Artifact, directory string) (map[string]string, error) {
	localNames := make(map[string]string)

	for _, a := range artifacts {
		path, err := fetchArtifact(ctx, a, directory)
		if err != nil {
			return nil, err
		}
		localNames[a.Id] = path
	}

	return localNames, nil
}

func fetchArtifact(ctx context.Context, artifact *osconfigpb.SoftwareRecipe_Artifact, directory string) (string, error) {
	var checksum, extension string
	var reader io.ReadCloser

	switch v := artifact.Artifact.(type) {
	case *osconfigpb.SoftwareRecipe_Artifact_Gcs_:
		extension = path.Ext(v.Gcs.Object)
		cl, err := storage.NewClient(ctx)
		if err != nil {
			return "", fmt.Errorf("error creating gcs client: %v", err)
		}
		reader, err = getGCSArtifact(ctx, cl, v.Gcs.Object, v.Gcs.Bucket, v.Gcs.Generation)
		if err != nil {
			return "", fmt.Errorf("error fetching artifact %q from GCS: %v", artifact.Id, err)
		}

		defer reader.Close()
	case *osconfigpb.SoftwareRecipe_Artifact_Remote_:
		uri, err := url.Parse(v.Remote.Uri)
		extension = path.Ext(uri.Path)
		if err != nil {
			return "", fmt.Errorf("Could not parse url: %q; got error: %v", uri, err)
		}
		cl := &http.Client{}
		reader, err := getHTTPArtifact(cl, *uri)
		if err != nil {
			return "", fmt.Errorf("error fetching artifact %q: %v", artifact.Id, err)
		}
		defer reader.Close()
	default:
		return "", fmt.Errorf("unknown artifact type %T", v)
	}

	localPath := getStoragePath(directory, artifact.Id, extension)
	if err := external.DownloadStream(reader, checksum, localPath); err != nil {
		return "", fmt.Errorf("Error downloading stream: %v", err)
	}

	return localPath, nil
}

func getGCSArtifact(ctx context.Context, cl *storage.Client, object, bucket string, generation int64) (io.ReadCloser, error) {
	reader, err := external.FetchGCSObject(ctx, cl, object, bucket, generation)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func getHTTPArtifact(client *http.Client, uri url.URL) (io.ReadCloser, error) {
	if !isSupportedURL(uri) {
		return nil, fmt.Errorf("error, unsupported protocol scheme %s", uri.Scheme)
	}
	reader, err := external.FetchRemoteObjectHTTP(client, uri.String())
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func isSupportedURL(uri url.URL) bool {
	return (uri.Scheme == "http") || (uri.Scheme == "https")
}

func getStoragePath(directory, fname, extension string) string {
	localpath := filepath.Join(directory, fname)
	if extension != "" {
		localpath = localpath + extension
	}
	return localpath
}
