package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Based on https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-metadata-endpoint-v4.html

const (
	ecsMetadataUriEnvV4 = "ECS_CONTAINER_METADATA_URI_V4"
)

type Limits struct {
	CPU    float64 `json:"CPU"`
	Memory int     `json:"Memory"`
}

type ContainerMetadataV4 struct {
	DockerID   string `json:"DockerId"`
	Name       string `json:"Name"`
	DockerName string `json:"DockerName"`
	Image      string `json:"Image"`
	ImageID    string `json:"ImageID"`
	Labels     struct {
		EcsCluster               string `json:"com.amazonaws.ecs.cluster"`
		EcsContainerName         string `json:"com.amazonaws.ecs.container-name"`
		EcsTaskArn               string `json:"com.amazonaws.ecs.task-arn"`
		EcsTaskDefinitionFamily  string `json:"com.amazonaws.ecs.task-definition-family"`
		EcsTaskDefinitionVersion string `json:"com.amazonaws.ecs.task-definition-version"`
	} `json:"Labels"`
	DesiredStatus string    `json:"DesiredStatus"`
	KnownStatus   string    `json:"KnownStatus"`
	Limits        Limits    `json:"Limits"`
	CreatedAt     time.Time `json:"CreatedAt"`
	StartedAt     time.Time `json:"StartedAt"`
	Type          string    `json:"Type"`
	ContainerARN  string    `json:"ContainerARN"`
	LogDriver     string    `json:"LogDriver"`
	LogOptions    struct {
		AwsLogsCreateGroup bool   `json:"awslogs-create-group"`
		AwsLogsGroup       string `json:"awslogs-group"`
		AwsLogsStream      string `json:"awslogs-stream"`
		AwsRegion          string `json:"awslogs-region"`
	} `json:"LogOptions"`
	Networks []struct {
		NetworkMode              string   `json:"NetworkMode"`
		IPv4Addresses            []string `json:"IPv4Addresses"`
		AttachmentIndex          int      `json:"AttachmentIndex"`
		IPv4SubnetCIDRBlock      string   `json:"IPv4SubnetCIDRBlock"`
		MACAddress               string   `json:"MACAddress"`
		DomainNameServers        []string `json:"DomainNameServers"`
		DomainNameSearchList     []string `json:"DomainNameSearchList"`
		PrivateDNSName           string   `json:"PrivateDNSName"`
		SubnetGatewayIpv4Address string   `json:"SubnetGatewayIpv4Address"`
	} `json:"Networks"`
}

type TaskMetadataV4 struct {
	Cluster          string                `json:"Cluster"`
	TaskARN          string                `json:"TaskARN"`
	Family           string                `json:"Family"`
	Revision         string                `json:"Revision"`
	DesiredStatus    string                `json:"DesiredStatus"`
	KnownStatus      string                `json:"KnownStatus"`
	Limits           Limits                `json:"Limits"`
	PullStartedAt    time.Time             `json:"PullStartedAt"`
	PullStoppedAt    time.Time             `json:"PullStoppedAt"`
	AvailabilityZone string                `json:"AvailabilityZone"`
	LaunchType       string                `json:"LaunchType"`
	Containers       []ContainerMetadataV4 `json:"Containers"`
}

// Retrieve ECS Task Metadata in V4 format
func GetTaskV4(ctx context.Context, client *http.Client) (*TaskMetadataV4, error) {
	metadataUrl := os.Getenv(ecsMetadataUriEnvV4)
	if metadataUrl == "" {
		return nil, fmt.Errorf("missing metadata uri in environment (%s)", ecsMetadataUriEnvV4)
	}

	taskMetadata := &TaskMetadataV4{}
	body, err := fetch(ctx, client, fmt.Sprintf("%s/task", metadataUrl))

	err = json.Unmarshal(body, &taskMetadata)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal into task metadata v4: %w", err)
	}

	return taskMetadata, nil
}

// Retrieve ECS Container Metadata in V4 format
func GetContainerV4(ctx context.Context, client *http.Client) (*ContainerMetadataV4, error) {
	metadataUrl := os.Getenv(ecsMetadataUriEnvV4)
	if metadataUrl == "" {
		return nil, fmt.Errorf("missing metadata uri in environment (%s)", ecsMetadataUriEnvV4)
	}

	contaienrMetadata := &ContainerMetadataV4{}
	body, err := fetch(ctx, client, fmt.Sprintf("%s", metadataUrl))

	err = json.Unmarshal(body, &contaienrMetadata)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal into container metadata v4: %w", err)
	}

	return contaienrMetadata, nil
}
