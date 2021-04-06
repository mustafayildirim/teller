// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches the contents of the specified resource-based permission policy to a
// secret. A resource-based policy is optional. Alternatively, you can use IAM
// identity-based policies that specify the secret's Amazon Resource Name (ARN) in
// the policy statement's Resources element. You can also use a combination of both
// identity-based and resource-based policies. The affected users and roles receive
// the permissions that are permitted by all of the relevant policies. For more
// information, see Using Resource-Based Policies for AWS Secrets Manager
// (http://docs.aws.amazon.com/secretsmanager/latest/userguide/auth-and-access_resource-based-policies.html).
// For the complete description of the AWS policy syntax and grammar, see IAM JSON
// Policy Reference
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies.html) in
// the IAM User Guide. Minimum permissions To run this command, you must have the
// following permissions:
//
// * secretsmanager:PutResourcePolicy
//
// Related
// operations
//
// * To retrieve the resource policy attached to a secret, use
// GetResourcePolicy.
//
// * To delete the resource-based policy that's attached to a
// secret, use DeleteResourcePolicy.
//
// * To list all of the currently available
// secrets, use ListSecrets.
func (c *Client) PutResourcePolicy(ctx context.Context, params *PutResourcePolicyInput, optFns ...func(*Options)) (*PutResourcePolicyOutput, error) {
	if params == nil {
		params = &PutResourcePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResourcePolicy", params, optFns, addOperationPutResourcePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResourcePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResourcePolicyInput struct {

	// A JSON-formatted string that's constructed according to the grammar and syntax
	// for an AWS resource-based policy. The policy in the string identifies who can
	// access or manage this secret and its versions. For information on how to format
	// a JSON parameter for the various command line tool environments, see Using JSON
	// for Parameters
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json)
	// in the AWS CLI User Guide.
	//
	// This member is required.
	ResourcePolicy *string

	// Specifies the secret that you want to attach the resource-based policy to. You
	// can specify either the ARN or the friendly name of the secret. If you specify an
	// ARN, we generally recommend that you specify a complete ARN. You can specify a
	// partial ARN too—for example, if you don’t include the final hyphen and six
	// random characters that Secrets Manager adds at the end of the ARN when you
	// created the secret. A partial ARN match can work as long as it uniquely matches
	// only one secret. However, if your secret has a name that ends in a hyphen
	// followed by six characters (before Secrets Manager adds the hyphen and six
	// characters to the ARN) and you try to use that as a partial ARN, then those
	// characters cause Secrets Manager to assume that you’re specifying a complete
	// ARN. This confusion can cause unexpected results. To avoid this situation, we
	// recommend that you don’t create secret names ending with a hyphen followed by
	// six characters. If you specify an incomplete ARN without the random suffix, and
	// instead provide the 'friendly name', you must not include the random suffix. If
	// you do include the random suffix added by Secrets Manager, you receive either a
	// ResourceNotFoundException or an AccessDeniedException error, depending on your
	// permissions.
	//
	// This member is required.
	SecretId *string

	// Makes an optional API call to Zelkova to validate the Resource Policy to prevent
	// broad access to your secret.
	BlockPublicPolicy bool
}

type PutResourcePolicyOutput struct {

	// The ARN of the secret retrieved by the resource-based policy.
	ARN *string

	// The friendly name of the secret that the retrieved by the resource-based policy.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutResourcePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResourcePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutResourcePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutResourcePolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutResourcePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "PutResourcePolicy",
	}
}