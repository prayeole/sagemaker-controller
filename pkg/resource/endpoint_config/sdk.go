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

package endpoint_config

import (
	"context"
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
	_ = &svcapitypes.EndpointConfig{}
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
	defer exit(err)
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

	var resp *svcsdk.DescribeEndpointConfigOutput
	resp, err = rm.sdkapi.DescribeEndpointConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribeEndpointConfig", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "ValidationException" && strings.HasPrefix(awsErr.Message(), "Could not find endpoint configuration") {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if resp.DataCaptureConfig != nil {
		f1 := &svcapitypes.DataCaptureConfig{}
		if resp.DataCaptureConfig.CaptureContentTypeHeader != nil {
			f1f0 := &svcapitypes.CaptureContentTypeHeader{}
			if resp.DataCaptureConfig.CaptureContentTypeHeader.CsvContentTypes != nil {
				f1f0f0 := []*string{}
				for _, f1f0f0iter := range resp.DataCaptureConfig.CaptureContentTypeHeader.CsvContentTypes {
					var f1f0f0elem string
					f1f0f0elem = *f1f0f0iter
					f1f0f0 = append(f1f0f0, &f1f0f0elem)
				}
				f1f0.CsvContentTypes = f1f0f0
			}
			if resp.DataCaptureConfig.CaptureContentTypeHeader.JsonContentTypes != nil {
				f1f0f1 := []*string{}
				for _, f1f0f1iter := range resp.DataCaptureConfig.CaptureContentTypeHeader.JsonContentTypes {
					var f1f0f1elem string
					f1f0f1elem = *f1f0f1iter
					f1f0f1 = append(f1f0f1, &f1f0f1elem)
				}
				f1f0.JSONContentTypes = f1f0f1
			}
			f1.CaptureContentTypeHeader = f1f0
		}
		if resp.DataCaptureConfig.CaptureOptions != nil {
			f1f1 := []*svcapitypes.CaptureOption{}
			for _, f1f1iter := range resp.DataCaptureConfig.CaptureOptions {
				f1f1elem := &svcapitypes.CaptureOption{}
				if f1f1iter.CaptureMode != nil {
					f1f1elem.CaptureMode = f1f1iter.CaptureMode
				}
				f1f1 = append(f1f1, f1f1elem)
			}
			f1.CaptureOptions = f1f1
		}
		if resp.DataCaptureConfig.DestinationS3Uri != nil {
			f1.DestinationS3URI = resp.DataCaptureConfig.DestinationS3Uri
		}
		if resp.DataCaptureConfig.EnableCapture != nil {
			f1.EnableCapture = resp.DataCaptureConfig.EnableCapture
		}
		if resp.DataCaptureConfig.InitialSamplingPercentage != nil {
			f1.InitialSamplingPercentage = resp.DataCaptureConfig.InitialSamplingPercentage
		}
		if resp.DataCaptureConfig.KmsKeyId != nil {
			f1.KMSKeyID = resp.DataCaptureConfig.KmsKeyId
		}
		ko.Spec.DataCaptureConfig = f1
	} else {
		ko.Spec.DataCaptureConfig = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.EndpointConfigArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.EndpointConfigArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.EndpointConfigName != nil {
		ko.Spec.EndpointConfigName = resp.EndpointConfigName
	} else {
		ko.Spec.EndpointConfigName = nil
	}
	if resp.KmsKeyId != nil {
		ko.Spec.KMSKeyID = resp.KmsKeyId
	} else {
		ko.Spec.KMSKeyID = nil
	}
	if resp.ProductionVariants != nil {
		f5 := []*svcapitypes.ProductionVariant{}
		for _, f5iter := range resp.ProductionVariants {
			f5elem := &svcapitypes.ProductionVariant{}
			if f5iter.AcceleratorType != nil {
				f5elem.AcceleratorType = f5iter.AcceleratorType
			}
			if f5iter.CoreDumpConfig != nil {
				f5elemf1 := &svcapitypes.ProductionVariantCoreDumpConfig{}
				if f5iter.CoreDumpConfig.DestinationS3Uri != nil {
					f5elemf1.DestinationS3URI = f5iter.CoreDumpConfig.DestinationS3Uri
				}
				if f5iter.CoreDumpConfig.KmsKeyId != nil {
					f5elemf1.KMSKeyID = f5iter.CoreDumpConfig.KmsKeyId
				}
				f5elem.CoreDumpConfig = f5elemf1
			}
			if f5iter.InitialInstanceCount != nil {
				f5elem.InitialInstanceCount = f5iter.InitialInstanceCount
			}
			if f5iter.InitialVariantWeight != nil {
				f5elem.InitialVariantWeight = f5iter.InitialVariantWeight
			}
			if f5iter.InstanceType != nil {
				f5elem.InstanceType = f5iter.InstanceType
			}
			if f5iter.ModelName != nil {
				f5elem.ModelName = f5iter.ModelName
			}
			if f5iter.VariantName != nil {
				f5elem.VariantName = f5iter.VariantName
			}
			f5 = append(f5, f5elem)
		}
		ko.Spec.ProductionVariants = f5
	} else {
		ko.Spec.ProductionVariants = nil
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
	return r.ko.Spec.EndpointConfigName == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribeEndpointConfigInput, error) {
	res := &svcsdk.DescribeEndpointConfigInput{}

	if r.ko.Spec.EndpointConfigName != nil {
		res.SetEndpointConfigName(*r.ko.Spec.EndpointConfigName)
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
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateEndpointConfigOutput
	_ = resp
	resp, err = rm.sdkapi.CreateEndpointConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateEndpointConfig", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.EndpointConfigArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.EndpointConfigArn)
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
) (*svcsdk.CreateEndpointConfigInput, error) {
	res := &svcsdk.CreateEndpointConfigInput{}

	if r.ko.Spec.DataCaptureConfig != nil {
		f0 := &svcsdk.DataCaptureConfig{}
		if r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader != nil {
			f0f0 := &svcsdk.CaptureContentTypeHeader{}
			if r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader.CsvContentTypes != nil {
				f0f0f0 := []*string{}
				for _, f0f0f0iter := range r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader.CsvContentTypes {
					var f0f0f0elem string
					f0f0f0elem = *f0f0f0iter
					f0f0f0 = append(f0f0f0, &f0f0f0elem)
				}
				f0f0.SetCsvContentTypes(f0f0f0)
			}
			if r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader.JSONContentTypes != nil {
				f0f0f1 := []*string{}
				for _, f0f0f1iter := range r.ko.Spec.DataCaptureConfig.CaptureContentTypeHeader.JSONContentTypes {
					var f0f0f1elem string
					f0f0f1elem = *f0f0f1iter
					f0f0f1 = append(f0f0f1, &f0f0f1elem)
				}
				f0f0.SetJsonContentTypes(f0f0f1)
			}
			f0.SetCaptureContentTypeHeader(f0f0)
		}
		if r.ko.Spec.DataCaptureConfig.CaptureOptions != nil {
			f0f1 := []*svcsdk.CaptureOption{}
			for _, f0f1iter := range r.ko.Spec.DataCaptureConfig.CaptureOptions {
				f0f1elem := &svcsdk.CaptureOption{}
				if f0f1iter.CaptureMode != nil {
					f0f1elem.SetCaptureMode(*f0f1iter.CaptureMode)
				}
				f0f1 = append(f0f1, f0f1elem)
			}
			f0.SetCaptureOptions(f0f1)
		}
		if r.ko.Spec.DataCaptureConfig.DestinationS3URI != nil {
			f0.SetDestinationS3Uri(*r.ko.Spec.DataCaptureConfig.DestinationS3URI)
		}
		if r.ko.Spec.DataCaptureConfig.EnableCapture != nil {
			f0.SetEnableCapture(*r.ko.Spec.DataCaptureConfig.EnableCapture)
		}
		if r.ko.Spec.DataCaptureConfig.InitialSamplingPercentage != nil {
			f0.SetInitialSamplingPercentage(*r.ko.Spec.DataCaptureConfig.InitialSamplingPercentage)
		}
		if r.ko.Spec.DataCaptureConfig.KMSKeyID != nil {
			f0.SetKmsKeyId(*r.ko.Spec.DataCaptureConfig.KMSKeyID)
		}
		res.SetDataCaptureConfig(f0)
	}
	if r.ko.Spec.EndpointConfigName != nil {
		res.SetEndpointConfigName(*r.ko.Spec.EndpointConfigName)
	}
	if r.ko.Spec.KMSKeyID != nil {
		res.SetKmsKeyId(*r.ko.Spec.KMSKeyID)
	}
	if r.ko.Spec.ProductionVariants != nil {
		f3 := []*svcsdk.ProductionVariant{}
		for _, f3iter := range r.ko.Spec.ProductionVariants {
			f3elem := &svcsdk.ProductionVariant{}
			if f3iter.AcceleratorType != nil {
				f3elem.SetAcceleratorType(*f3iter.AcceleratorType)
			}
			if f3iter.CoreDumpConfig != nil {
				f3elemf1 := &svcsdk.ProductionVariantCoreDumpConfig{}
				if f3iter.CoreDumpConfig.DestinationS3URI != nil {
					f3elemf1.SetDestinationS3Uri(*f3iter.CoreDumpConfig.DestinationS3URI)
				}
				if f3iter.CoreDumpConfig.KMSKeyID != nil {
					f3elemf1.SetKmsKeyId(*f3iter.CoreDumpConfig.KMSKeyID)
				}
				f3elem.SetCoreDumpConfig(f3elemf1)
			}
			if f3iter.InitialInstanceCount != nil {
				f3elem.SetInitialInstanceCount(*f3iter.InitialInstanceCount)
			}
			if f3iter.InitialVariantWeight != nil {
				f3elem.SetInitialVariantWeight(*f3iter.InitialVariantWeight)
			}
			if f3iter.InstanceType != nil {
				f3elem.SetInstanceType(*f3iter.InstanceType)
			}
			if f3iter.ModelName != nil {
				f3elem.SetModelName(*f3iter.ModelName)
			}
			if f3iter.VariantName != nil {
				f3elem.SetVariantName(*f3iter.VariantName)
			}
			f3 = append(f3, f3elem)
		}
		res.SetProductionVariants(f3)
	}
	if r.ko.Spec.Tags != nil {
		f4 := []*svcsdk.Tag{}
		for _, f4iter := range r.ko.Spec.Tags {
			f4elem := &svcsdk.Tag{}
			if f4iter.Key != nil {
				f4elem.SetKey(*f4iter.Key)
			}
			if f4iter.Value != nil {
				f4elem.SetValue(*f4iter.Value)
			}
			f4 = append(f4, f4elem)
		}
		res.SetTags(f4)
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
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteEndpointConfigOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteEndpointConfigWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteEndpointConfig", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteEndpointConfigInput, error) {
	res := &svcsdk.DeleteEndpointConfigInput{}

	if r.ko.Spec.EndpointConfigName != nil {
		res.SetEndpointConfigName(*r.ko.Spec.EndpointConfigName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.EndpointConfig,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
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

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
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
