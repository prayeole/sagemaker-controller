// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.SageMaker{}
	_ = &svcapitypes.Model{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribeModelOutput
	resp, err = rm.sdkapi.DescribeModelWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeModel", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ValidationException" && strings.HasPrefix(awsErr.Message(), "Could not find model") {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.Containers != nil {
		f0 := []*svcapitypes.ContainerDefinition{}
		for _, f0iter := range resp.Containers {
			f0elem := &svcapitypes.ContainerDefinition{}
			if f0iter.ContainerHostname != nil {
				f0elem.ContainerHostname = f0iter.ContainerHostname
			}
			if f0iter.Environment != nil {
				f0elemf1 := map[string]*string{}
				for f0elemf1key, f0elemf1valiter := range f0iter.Environment {
					var f0elemf1val string
					f0elemf1val = *f0elemf1valiter
					f0elemf1[f0elemf1key] = &f0elemf1val
				}
				f0elem.Environment = f0elemf1
			}
			if f0iter.Image != nil {
				f0elem.Image = f0iter.Image
			}
			if f0iter.ImageConfig != nil {
				f0elemf3 := &svcapitypes.ImageConfig{}
				if f0iter.ImageConfig.RepositoryAccessMode != nil {
					f0elemf3.RepositoryAccessMode = f0iter.ImageConfig.RepositoryAccessMode
				}
				if f0iter.ImageConfig.RepositoryAuthConfig != nil {
					f0elemf3f1 := &svcapitypes.RepositoryAuthConfig{}
					if f0iter.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderArn != nil {
						f0elemf3f1.RepositoryCredentialsProviderARN = f0iter.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderArn
					}
					f0elemf3.RepositoryAuthConfig = f0elemf3f1
				}
				f0elem.ImageConfig = f0elemf3
			}
			if f0iter.Mode != nil {
				f0elem.Mode = f0iter.Mode
			}
			if f0iter.ModelDataUrl != nil {
				f0elem.ModelDataURL = f0iter.ModelDataUrl
			}
			if f0iter.ModelPackageName != nil {
				f0elem.ModelPackageName = f0iter.ModelPackageName
			}
			if f0iter.MultiModelConfig != nil {
				f0elemf7 := &svcapitypes.MultiModelConfig{}
				if f0iter.MultiModelConfig.ModelCacheSetting != nil {
					f0elemf7.ModelCacheSetting = f0iter.MultiModelConfig.ModelCacheSetting
				}
				f0elem.MultiModelConfig = f0elemf7
			}
			f0 = append(f0, f0elem)
		}
		ko.Spec.Containers = f0
	} else {
		ko.Spec.Containers = nil
	}
	if resp.EnableNetworkIsolation != nil {
		ko.Spec.EnableNetworkIsolation = resp.EnableNetworkIsolation
	} else {
		ko.Spec.EnableNetworkIsolation = nil
	}
	if resp.ExecutionRoleArn != nil {
		ko.Spec.ExecutionRoleARN = resp.ExecutionRoleArn
	} else {
		ko.Spec.ExecutionRoleARN = nil
	}
	if resp.InferenceExecutionConfig != nil {
		f4 := &svcapitypes.InferenceExecutionConfig{}
		if resp.InferenceExecutionConfig.Mode != nil {
			f4.Mode = resp.InferenceExecutionConfig.Mode
		}
		ko.Spec.InferenceExecutionConfig = f4
	} else {
		ko.Spec.InferenceExecutionConfig = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ModelArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ModelArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.ModelName != nil {
		ko.Spec.ModelName = resp.ModelName
	} else {
		ko.Spec.ModelName = nil
	}
	if resp.PrimaryContainer != nil {
		f7 := &svcapitypes.ContainerDefinition{}
		if resp.PrimaryContainer.ContainerHostname != nil {
			f7.ContainerHostname = resp.PrimaryContainer.ContainerHostname
		}
		if resp.PrimaryContainer.Environment != nil {
			f7f1 := map[string]*string{}
			for f7f1key, f7f1valiter := range resp.PrimaryContainer.Environment {
				var f7f1val string
				f7f1val = *f7f1valiter
				f7f1[f7f1key] = &f7f1val
			}
			f7.Environment = f7f1
		}
		if resp.PrimaryContainer.Image != nil {
			f7.Image = resp.PrimaryContainer.Image
		}
		if resp.PrimaryContainer.ImageConfig != nil {
			f7f3 := &svcapitypes.ImageConfig{}
			if resp.PrimaryContainer.ImageConfig.RepositoryAccessMode != nil {
				f7f3.RepositoryAccessMode = resp.PrimaryContainer.ImageConfig.RepositoryAccessMode
			}
			if resp.PrimaryContainer.ImageConfig.RepositoryAuthConfig != nil {
				f7f3f1 := &svcapitypes.RepositoryAuthConfig{}
				if resp.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderArn != nil {
					f7f3f1.RepositoryCredentialsProviderARN = resp.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderArn
				}
				f7f3.RepositoryAuthConfig = f7f3f1
			}
			f7.ImageConfig = f7f3
		}
		if resp.PrimaryContainer.Mode != nil {
			f7.Mode = resp.PrimaryContainer.Mode
		}
		if resp.PrimaryContainer.ModelDataUrl != nil {
			f7.ModelDataURL = resp.PrimaryContainer.ModelDataUrl
		}
		if resp.PrimaryContainer.ModelPackageName != nil {
			f7.ModelPackageName = resp.PrimaryContainer.ModelPackageName
		}
		if resp.PrimaryContainer.MultiModelConfig != nil {
			f7f7 := &svcapitypes.MultiModelConfig{}
			if resp.PrimaryContainer.MultiModelConfig.ModelCacheSetting != nil {
				f7f7.ModelCacheSetting = resp.PrimaryContainer.MultiModelConfig.ModelCacheSetting
			}
			f7.MultiModelConfig = f7f7
		}
		ko.Spec.PrimaryContainer = f7
	} else {
		ko.Spec.PrimaryContainer = nil
	}
	if resp.VpcConfig != nil {
		f8 := &svcapitypes.VPCConfig{}
		if resp.VpcConfig.SecurityGroupIds != nil {
			f8f0 := []*string{}
			for _, f8f0iter := range resp.VpcConfig.SecurityGroupIds {
				var f8f0elem string
				f8f0elem = *f8f0iter
				f8f0 = append(f8f0, &f8f0elem)
			}
			f8.SecurityGroupIDs = f8f0
		}
		if resp.VpcConfig.Subnets != nil {
			f8f1 := []*string{}
			for _, f8f1iter := range resp.VpcConfig.Subnets {
				var f8f1elem string
				f8f1elem = *f8f1iter
				f8f1 = append(f8f1, &f8f1elem)
			}
			f8.Subnets = f8f1
		}
		ko.Spec.VPCConfig = f8
	} else {
		ko.Spec.VPCConfig = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.ModelName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeModelInput, error) {
	res := &svcsdk.DescribeModelInput{}

	if r.ko.Spec.ModelName != nil {
		res.SetModelName(*r.ko.Spec.ModelName)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateModelOutput
	_ = resp
	resp, err = rm.sdkapi.CreateModelWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateModel", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.ModelArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.ModelArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateModelInput, error) {
	res := &svcsdk.CreateModelInput{}

	if r.ko.Spec.Containers != nil {
		f0 := []*svcsdk.ContainerDefinition{}
		for _, f0iter := range r.ko.Spec.Containers {
			f0elem := &svcsdk.ContainerDefinition{}
			if f0iter.ContainerHostname != nil {
				f0elem.SetContainerHostname(*f0iter.ContainerHostname)
			}
			if f0iter.Environment != nil {
				f0elemf1 := map[string]*string{}
				for f0elemf1key, f0elemf1valiter := range f0iter.Environment {
					var f0elemf1val string
					f0elemf1val = *f0elemf1valiter
					f0elemf1[f0elemf1key] = &f0elemf1val
				}
				f0elem.SetEnvironment(f0elemf1)
			}
			if f0iter.Image != nil {
				f0elem.SetImage(*f0iter.Image)
			}
			if f0iter.ImageConfig != nil {
				f0elemf3 := &svcsdk.ImageConfig{}
				if f0iter.ImageConfig.RepositoryAccessMode != nil {
					f0elemf3.SetRepositoryAccessMode(*f0iter.ImageConfig.RepositoryAccessMode)
				}
				if f0iter.ImageConfig.RepositoryAuthConfig != nil {
					f0elemf3f1 := &svcsdk.RepositoryAuthConfig{}
					if f0iter.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN != nil {
						f0elemf3f1.SetRepositoryCredentialsProviderArn(*f0iter.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN)
					}
					f0elemf3.SetRepositoryAuthConfig(f0elemf3f1)
				}
				f0elem.SetImageConfig(f0elemf3)
			}
			if f0iter.Mode != nil {
				f0elem.SetMode(*f0iter.Mode)
			}
			if f0iter.ModelDataURL != nil {
				f0elem.SetModelDataUrl(*f0iter.ModelDataURL)
			}
			if f0iter.ModelPackageName != nil {
				f0elem.SetModelPackageName(*f0iter.ModelPackageName)
			}
			if f0iter.MultiModelConfig != nil {
				f0elemf7 := &svcsdk.MultiModelConfig{}
				if f0iter.MultiModelConfig.ModelCacheSetting != nil {
					f0elemf7.SetModelCacheSetting(*f0iter.MultiModelConfig.ModelCacheSetting)
				}
				f0elem.SetMultiModelConfig(f0elemf7)
			}
			f0 = append(f0, f0elem)
		}
		res.SetContainers(f0)
	}
	if r.ko.Spec.EnableNetworkIsolation != nil {
		res.SetEnableNetworkIsolation(*r.ko.Spec.EnableNetworkIsolation)
	}
	if r.ko.Spec.ExecutionRoleARN != nil {
		res.SetExecutionRoleArn(*r.ko.Spec.ExecutionRoleARN)
	}
	if r.ko.Spec.InferenceExecutionConfig != nil {
		f3 := &svcsdk.InferenceExecutionConfig{}
		if r.ko.Spec.InferenceExecutionConfig.Mode != nil {
			f3.SetMode(*r.ko.Spec.InferenceExecutionConfig.Mode)
		}
		res.SetInferenceExecutionConfig(f3)
	}
	if r.ko.Spec.ModelName != nil {
		res.SetModelName(*r.ko.Spec.ModelName)
	}
	if r.ko.Spec.PrimaryContainer != nil {
		f5 := &svcsdk.ContainerDefinition{}
		if r.ko.Spec.PrimaryContainer.ContainerHostname != nil {
			f5.SetContainerHostname(*r.ko.Spec.PrimaryContainer.ContainerHostname)
		}
		if r.ko.Spec.PrimaryContainer.Environment != nil {
			f5f1 := map[string]*string{}
			for f5f1key, f5f1valiter := range r.ko.Spec.PrimaryContainer.Environment {
				var f5f1val string
				f5f1val = *f5f1valiter
				f5f1[f5f1key] = &f5f1val
			}
			f5.SetEnvironment(f5f1)
		}
		if r.ko.Spec.PrimaryContainer.Image != nil {
			f5.SetImage(*r.ko.Spec.PrimaryContainer.Image)
		}
		if r.ko.Spec.PrimaryContainer.ImageConfig != nil {
			f5f3 := &svcsdk.ImageConfig{}
			if r.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode != nil {
				f5f3.SetRepositoryAccessMode(*r.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAccessMode)
			}
			if r.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig != nil {
				f5f3f1 := &svcsdk.RepositoryAuthConfig{}
				if r.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN != nil {
					f5f3f1.SetRepositoryCredentialsProviderArn(*r.ko.Spec.PrimaryContainer.ImageConfig.RepositoryAuthConfig.RepositoryCredentialsProviderARN)
				}
				f5f3.SetRepositoryAuthConfig(f5f3f1)
			}
			f5.SetImageConfig(f5f3)
		}
		if r.ko.Spec.PrimaryContainer.Mode != nil {
			f5.SetMode(*r.ko.Spec.PrimaryContainer.Mode)
		}
		if r.ko.Spec.PrimaryContainer.ModelDataURL != nil {
			f5.SetModelDataUrl(*r.ko.Spec.PrimaryContainer.ModelDataURL)
		}
		if r.ko.Spec.PrimaryContainer.ModelPackageName != nil {
			f5.SetModelPackageName(*r.ko.Spec.PrimaryContainer.ModelPackageName)
		}
		if r.ko.Spec.PrimaryContainer.MultiModelConfig != nil {
			f5f7 := &svcsdk.MultiModelConfig{}
			if r.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting != nil {
				f5f7.SetModelCacheSetting(*r.ko.Spec.PrimaryContainer.MultiModelConfig.ModelCacheSetting)
			}
			f5.SetMultiModelConfig(f5f7)
		}
		res.SetPrimaryContainer(f5)
	}
	if r.ko.Spec.Tags != nil {
		f6 := []*svcsdk.Tag{}
		for _, f6iter := range r.ko.Spec.Tags {
			f6elem := &svcsdk.Tag{}
			if f6iter.Key != nil {
				f6elem.SetKey(*f6iter.Key)
			}
			if f6iter.Value != nil {
				f6elem.SetValue(*f6iter.Value)
			}
			f6 = append(f6, f6elem)
		}
		res.SetTags(f6)
	}
	if r.ko.Spec.VPCConfig != nil {
		f7 := &svcsdk.VpcConfig{}
		if r.ko.Spec.VPCConfig.SecurityGroupIDs != nil {
			f7f0 := []*string{}
			for _, f7f0iter := range r.ko.Spec.VPCConfig.SecurityGroupIDs {
				var f7f0elem string
				f7f0elem = *f7f0iter
				f7f0 = append(f7f0, &f7f0elem)
			}
			f7.SetSecurityGroupIds(f7f0)
		}
		if r.ko.Spec.VPCConfig.Subnets != nil {
			f7f1 := []*string{}
			for _, f7f1iter := range r.ko.Spec.VPCConfig.Subnets {
				var f7f1elem string
				f7f1elem = *f7f1iter
				f7f1 = append(f7f1, &f7f1elem)
			}
			f7.SetSubnets(f7f1)
		}
		res.SetVpcConfig(f7)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (*resource, error) {
	// TODO(jaypipes): Figure this out...
	return nil, ackerr.NotImplemented
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteModelOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteModelWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteModel", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteModelInput, error) {
	res := &svcsdk.DeleteModelInput{}

	if r.ko.Spec.ModelName != nil {
		res.SetModelName(*r.ko.Spec.ModelName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.Model,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "InvalidParameterCombination",
		"InvalidParameterValue",
		"MissingParameter":
		return true
	default:
		return false
	}
}
