package metadata

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseContainerMetadataV4(t *testing.T) {
	content, err := os.ReadFile("testdata/metadatav4-response-container.json")

	assert.NoError(t, err, "FAILED")

	containerMetadata := &ContainerMetadataV4{}
	json.Unmarshal(content, containerMetadata)

	assert.Equal(t, "/ecs/metadata", containerMetadata.LogOptions.AwsLogsGroup)
	assert.Equal(t, "ecs/curl/8f03e41243824aea923aca126495f665", containerMetadata.LogOptions.AwsLogsStream)
	assert.Equal(t, "us-west-2", containerMetadata.LogOptions.AwsRegion)
	assert.Equal(t, "true", containerMetadata.LogOptions.AwsLogsCreateGroup)
	assert.Equal(t, "arn:aws:ecs:us-west-2:111122223333:task/default/8f03e41243824aea923aca126495f665", containerMetadata.Labels.EcsTaskArn)
}

func TestParseTaskMetadataV4(t *testing.T) {
	content, err := os.ReadFile("testdata/metadatav4-response-task.json")

	assert.NoError(t, err, "FAILED")

	taskMetadata := &TaskMetadataV4{}
	json.Unmarshal(content, taskMetadata)

	assert.Equal(t, "/ecs/metadata", taskMetadata.Containers[1].LogOptions.AwsLogsGroup)
	assert.Equal(t, "ecs/curl/158d1c8083dd49d6b527399fd6414f5c", taskMetadata.Containers[1].LogOptions.AwsLogsStream)
	assert.Equal(t, "us-west-2", taskMetadata.Containers[1].LogOptions.AwsRegion)
	assert.Equal(t, "true", taskMetadata.Containers[1].LogOptions.AwsLogsCreateGroup)
	assert.Equal(t, "arn:aws:ecs:us-west-2:111122223333:task/default/158d1c8083dd49d6b527399fd6414f5c", taskMetadata.Containers[1].Labels.EcsTaskArn)
}
