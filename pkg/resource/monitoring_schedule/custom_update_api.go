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

// Use this file to add custom implementation for any operation of intercept
// the autogenerated code that trigger an update

package monitoring_schedule

import (
	"context"
	"errors"

	"github.com/aws-controllers-k8s/runtime/pkg/requeue"
)

// statusAllowUpdates is a helper method to determine if monitoring schedule status allows modification
// Modifications to monitoring schedule are only allowed if:
//   1. The schedule is in a terminal state i.e. Status != Pending
//   2. There are no pending or in-progress jobs/executions
func (rm *resourceManager) statusAllowUpdates(
	ctx context.Context,
	r *resource,
) error {
	scheduleStatus := r.ko.Status.MonitoringScheduleStatus
	if scheduleStatus == nil {
		return nil
	}

	inProgressExecutions := false
	// It is possible that schedule doesn't have any execution yet
	if r.ko.Status.LastMonitoringExecutionSummary != nil {
		executionStatus := r.ko.Status.LastMonitoringExecutionSummary.MonitoringExecutionStatus
		if *executionStatus == "Pending" || *executionStatus == "InProgress" || *executionStatus == "Stopping" {
			inProgressExecutions = true
		}
	}
	if *scheduleStatus == "Pending" || inProgressExecutions {
		return requeue.NeededAfter(
			errors.New("monitoring schedule status does not allow modification, it is in-progress"),
			requeue.DefaultRequeueAfterDuration)
	}

	return nil
}
