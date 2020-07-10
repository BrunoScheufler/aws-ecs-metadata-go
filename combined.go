package metadata

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

// Will retrieve the task metadata
// for your current Fargate environment (either V3 or V4)
// based on the environment variables that are present
func GetTaskMetadata(ctx context.Context, client *http.Client) (interface{}, error){
	isV3 := os.Getenv(ecsMetadataUriEnvV3) != ""
	if isV3 {
		return GetTaskV3(ctx, client)
	}

	isV4 := os.Getenv(ecsMetadataUriEnvV4) != ""
	if isV4 {
		return GetTaskV4(ctx, client)
	}

	return nil, fmt.Errorf("could not resolve ECS Task metadata")
}

func HasMetadata() bool {
	if os.Getenv(ecsMetadataUriEnvV3) != ""{
		return true
	}

	if os.Getenv(ecsMetadataUriEnvV4) != "" {
		return true
	}

	return false
}
