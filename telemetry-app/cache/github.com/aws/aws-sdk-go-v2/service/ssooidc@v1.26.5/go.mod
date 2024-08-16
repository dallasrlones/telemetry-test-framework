module github.com/aws/aws-sdk-go-v2/service/ssooidc

go 1.21

require (
	github.com/aws/aws-sdk-go-v2 v1.30.4
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.16
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.16
	github.com/aws/smithy-go v1.20.4
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/

replace github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 => ../../internal/endpoints/v2/
