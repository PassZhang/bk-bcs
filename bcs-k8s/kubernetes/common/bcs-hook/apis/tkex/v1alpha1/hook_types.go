/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	HookPhasePending      HookPhase = "Pending"
	HookPhaseRunning      HookPhase = "Running"
	HookPhaseSuccessful   HookPhase = "Successful"
	HookPhaseFailed       HookPhase = "Failed"
	HookPhaseError        HookPhase = "Error"
	HookPhaseInconclusive HookPhase = "Inconclusive"
)

func (hp HookPhase) Completed() bool {
	switch hp {
	case HookPhaseSuccessful, HookPhaseFailed, HookPhaseError, HookPhaseInconclusive:
		return true
	}
	return false
}

type HookPhase string

// HookTemplate is the Schema for the hooktemplates API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HookTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec HookTemplateSpec `json:"spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HookTemplateList contains a list of HookTemplate
type HookTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []HookTemplate `json:"items"`
}

type HookTemplateSpec struct {
	Metrics []Metric   `json:"metrics"`
	Args    []Argument `json:"args,omitempty"`
}

// Argument is an argument to an AnalysisRun
type Argument struct {
	// Name is the name of the argument
	Name string `json:"name"`
	// Value is the value of the argument
	// +optional
	Value *string `json:"value,omitempty"`
}

type DurationString string

// Duration converts DurationString into a time.Duration
func (d DurationString) Duration() (time.Duration, error) {
	return time.ParseDuration(string(d))
}

type Metric struct {
	Name string `json:"name"`
	// Interval defines an interval string (e.g. 30s, 5m, 1h) between each measurement.
	// If omitted, will perform a single measurement
	Interval DurationString `json:"interval,omitempty"`
	// InitialDelay how long the AnalysisRun should wait before starting this metric
	InitialDelay DurationString `json:"initialDelay,omitempty"`
	// Count is the number of times to run the measurement. If both interval and count are omitted,
	// the effective count is 1. If only interval is specified, metric runs indefinitely.
	// If count > 1, interval must be specified.
	Count int32 `json:"count,omitempty"`
	// SuccessCondition is an expression which determines if a measurement is considered successful
	// Expression is a goevaluate expression. The keyword `result` is a variable reference to the
	// value of measurement. Results can be both structured data or primitive.
	// Examples:
	//   result > 10
	//   (result.requests_made * result.requests_succeeded / 100) >= 90
	SuccessCondition string `json:"successCondition,omitempty"`
	// FailureCondition is an expression which determines if a measurement is considered failed
	// If both success and failure conditions are specified, and the measurement does not fall into
	// either condition, the measurement is considered Inconclusive
	FailureCondition string `json:"failureCondition,omitempty"`
	// FailureLimit is the maximum number of times the measurement is allowed to fail, before the
	// entire metric is considered Failed (default: 0)
	FailureLimit int32 `json:"failureLimit,omitempty"`
	// InconclusiveLimit is the maximum number of times the measurement is allowed to measure
	// Inconclusive, before the entire metric is considered Inconclusive (default: 0)
	InconclusiveLimit int32 `json:"inconclusiveLimit,omitempty"`
	// ConsecutiveErrorLimit is the maximum number of times the measurement is allowed to error in
	// succession, before the metric is considered error (default: 4)
	ConsecutiveErrorLimit *int32 `json:"consecutiveErrorLimit,omitempty"`
	// Provider configuration to the external system to use to verify the analysis
	Provider MetricProvider `json:"provider"`
}

func (m *Metric) EffectiveCount() *int32 {
	if m.Count == 0 {
		if m.Interval == "" {
			one := int32(1)
			return &one
		}
		return nil
	}
	return &m.Count
}

type MetricProvider struct {
	Web *WebMetric `json:"web,omitempty"`
	// Prometheus specifies the prometheus metric to query
	Prometheus *PrometheusMetric `json:"prometheus,omitempty"`
}

type PrometheusMetric struct {
	// Address is the HTTP address and port of the prometheus server
	Address string `json:"address,omitempty"`
	// Query is a raw prometheus query to perform
	Query string `json:"query,omitempty"`
}

type WebMetric struct {
	URL            string            `json:"url"`
	Headers        []WebMetricHeader `json:"headers,omitempty"`
	TimeoutSeconds int               `json:"timeoutSeconds,omitempty"`
	JsonPath       string            `json:"jsonPath"`
}

type WebMetricHeader struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// HookRun is the Schema for the hookruns API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type HookRun struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HookRunSpec   `json:"spec"`
	Status            HookRunStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HookRunList contains a list of HookRun
type HookRunList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []HookRun `json:"items"`
}

type HookRunSpec struct {
	Metrics   []Metric   `json:"metrics"`
	Args      []Argument `json:"args,omitempty"`
	Terminate bool       `json:"terminate,omitempty"`
}

type HookRunStatus struct {
	Phase         HookPhase      `json:"phase"`
	Message       string         `json:"message,omitempty"`
	MetricResults []MetricResult `json:"metricResults,omitempty"`
	StartedAt     *metav1.Time   `json:"startedAt,omitempty"`
}

type MetricResult struct {
	Name             string        `json:"name"`
	Phase            HookPhase     `json:"phase"`
	Measurements     []Measurement `json:"measurements,omitempty"`
	Message          string        `json:"message,omitempty"`
	Count            int32         `json:"count,omitempty"`
	Successful       int32         `json:"successful,omitempty"`
	Failed           int32         `json:"failed,omitempty"`
	Inconclusive     int32         `json:"inconclusive,omitempty"`
	Error            int32         `json:"error,omitempty"`
	ConsecutiveError int32         `json:"consecutiveError,omitempty"`
}

type Measurement struct {
	Phase      HookPhase    `json:"phase"`
	Message    string       `json:"message,omitempty"`
	StartedAt  *metav1.Time `json:"startedAt,omitempty"`
	FinishedAt *metav1.Time `json:"finishedAt,omitempty"`
	ResumeAt   *metav1.Time `json:"resumeAt,omitempty"`
	Value      string       `json:"value,omitempty"`
}

type PreDeleteHookCondition struct {
	PodName   string      `json:"podName"`
	StartTime metav1.Time `json:"startTime"`
	HookPhase HookPhase   `json:"phase"`
}

type HookStep struct {
	TemplateName string            `json:"templateName"`
	Args         []HookRunArgument `json:"args,omitempty"`
}

type HookRunArgument struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

type PauseReason string

const (
	PauseReasonCanaryPauseStep PauseReason = "PausedByCanaryPauseStep"
	PauseReasonStepBasedHook   PauseReason = "PausedByStepBasedHook"
)

type PauseCondition struct {
	Reason    PauseReason `json:"reason"`
	StartTime metav1.Time `json:"startTime"`
}
