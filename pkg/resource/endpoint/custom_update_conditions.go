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

// Use this file if conditions need to be updated based on the latest status
// of endpoint which is not evident from API response

package endpoint

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	svcapitypes "github.com/aws-controllers-k8s/sagemaker-controller/apis/v1alpha1"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/sagemaker"
	corev1 "k8s.io/api/core/v1"
)

// CustomUpdateConditions sets conditions (terminal) on supplied endpoint
// it examines supplied resource to determine conditions.
// It returns true if conditions are updated
func (rm *resourceManager) customUpdateConditions(
	ko *svcapitypes.Endpoint,
	r *resource,
	err error,
) bool {
	latestStatus := r.ko.Status.EndpointStatus
	if latestStatus == nil || *latestStatus != svcsdk.EndpointStatusFailed {
		return false
	}
	var terminalCondition *ackv1alpha1.Condition = nil
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	} else {
		for _, condition := range ko.Status.Conditions {
			if condition.Type == ackv1alpha1.ConditionTypeTerminal {
				terminalCondition = condition
				break
			}
		}
		if terminalCondition != nil && terminalCondition.Status == corev1.ConditionTrue {
			// some other exception already put the resource in terminal condition
			return false
		}
	}
	if terminalCondition == nil {
		terminalCondition = &ackv1alpha1.Condition{
			Type:   ackv1alpha1.ConditionTypeTerminal,
			Status: corev1.ConditionTrue,
			Reason: aws.String("Endpoint status: Failed. Cannot be updated"),
		}

		failureReason := r.ko.Status.FailureReason
		if failureReason != nil {
			terminalCondition.Message = failureReason
		}

		ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)

	}
	return true
}
