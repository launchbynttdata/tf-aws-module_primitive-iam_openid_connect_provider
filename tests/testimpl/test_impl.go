package testimpl

import (
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/launchbynttdata/lcaf-component-terratest/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	errFailedToGetOIDCProvider = "Failed to get OIDC provider"
)

func TestComposableComplete(t *testing.T, ctx types.TestContext) {
	iamClient := GetAWSIAMClient(t)

	oidcProviderArn := terraform.Output(t, ctx.TerratestTerraformOptions(), "arn")
	oidcProviderURL := terraform.Output(t, ctx.TerratestTerraformOptions(), "url")

	t.Run("TestOIDCProviderExists", func(t *testing.T) {
		provider, err := iamClient.GetOpenIDConnectProvider(context.TODO(), &iam.GetOpenIDConnectProviderInput{
			OpenIDConnectProviderArn: &oidcProviderArn,
		})
		require.NoError(t, err, errFailedToGetOIDCProvider)
		assert.NotNil(t, provider, "OIDC provider should not be nil")
		assert.NotEmpty(t, provider.Url, "OIDC provider URL should not be empty")
		assert.NotEmpty(t, provider.ClientIDList, "OIDC provider should have at least one client ID")
		assert.NotEmpty(t, provider.ThumbprintList, "OIDC provider should have at least one thumbprint")
	})

	t.Run("TestOIDCProviderConfiguration", func(t *testing.T) {
		provider, err := iamClient.GetOpenIDConnectProvider(context.TODO(), &iam.GetOpenIDConnectProviderInput{
			OpenIDConnectProviderArn: &oidcProviderArn,
		})
		require.NoError(t, err, errFailedToGetOIDCProvider)

		// Verify the URL is present and properly formatted
		// Note: AWS stores OIDC provider URLs without the https:// prefix
		assert.NotEmpty(t, provider.Url, "OIDC provider URL should not be empty")
		assert.NotContains(t, *provider.Url, "://", "OIDC provider URL should not contain protocol prefix")

		// Verify client IDs are present
		assert.Greater(t, len(provider.ClientIDList), 0, "OIDC provider should have at least one client ID")
		for _, clientID := range provider.ClientIDList {
			assert.NotEmpty(t, clientID, "Client ID should not be empty")
		}

		// Verify thumbprints are present and properly formatted
		assert.Greater(t, len(provider.ThumbprintList), 0, "OIDC provider should have at least one thumbprint")
		for _, thumbprint := range provider.ThumbprintList {
			assert.NotEmpty(t, thumbprint, "Thumbprint should not be empty")
			assert.Len(t, thumbprint, 40, "Thumbprint should be 40 characters long (SHA-1 hash)")
		}
	})

	t.Run("TestOIDCProviderTags", func(t *testing.T) {
		provider, err := iamClient.GetOpenIDConnectProvider(context.TODO(), &iam.GetOpenIDConnectProviderInput{
			OpenIDConnectProviderArn: &oidcProviderArn,
		})
		require.NoError(t, err, errFailedToGetOIDCProvider)

		// Verify tags are present (if any were set in the test configuration)
		if len(provider.Tags) > 0 {
			for _, tag := range provider.Tags {
				assert.NotEmpty(t, *tag.Key, "Tag key should not be empty")
				assert.NotNil(t, tag.Value, "Tag value should not be nil")
			}
		}
	})

	t.Run("TestOIDCProviderURLOutput", func(t *testing.T) {
		// Verify the URL output from Terraform is present and correctly formatted
		provider, err := iamClient.GetOpenIDConnectProvider(context.TODO(), &iam.GetOpenIDConnectProviderInput{
			OpenIDConnectProviderArn: &oidcProviderArn,
		})
		require.NoError(t, err, errFailedToGetOIDCProvider)

		// Verify the URL output is not empty
		assert.NotEmpty(t, oidcProviderURL, "URL output should not be empty")

		// Verify AWS stored the URL correctly (without protocol prefix)
		assert.NotContains(t, *provider.Url, "://", "AWS OIDC provider URL should not contain protocol prefix")
		assert.Equal(t, "example.com/oidc", *provider.Url, "AWS stored URL should match the expected value without protocol")

		// The Terraform output should ALWAYS include the https:// prefix for usability
		assert.True(t, strings.HasPrefix(oidcProviderURL, "https://"),
			"Terraform URL output should always include https:// prefix")
		assert.Equal(t, "https://example.com/oidc", oidcProviderURL,
			"URL output should be https://example.com/oidc")
	})
}

func GetAWSIAMClient(t *testing.T) *iam.Client {
	awsIAMClient := iam.NewFromConfig(GetAWSConfig(t))
	return awsIAMClient
}

func GetAWSConfig(t *testing.T) (cfg aws.Config) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	require.NoErrorf(t, err, "unable to load SDK config, %v", err)
	return cfg
}
