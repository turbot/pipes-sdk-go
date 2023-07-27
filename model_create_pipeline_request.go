/*
Turbot Pipes

Turbot Pipes is an intelligence, automation & security platform built specifically for DevOps.

API version: {{OPEN_API_VERSION}}
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipes

import (
	"encoding/json"
)

// CreatePipelineRequest struct for CreatePipelineRequest
type CreatePipelineRequest struct {
	Args      interface{}       `json:"args"`
	Frequency PipelineFrequency `json:"frequency"`
	// The name of the pipeline to be executed.
	Pipeline string      `json:"pipeline"`
	Tags     interface{} `json:"tags,omitempty"`
	// The title of the pipeline.
	Title string `json:"title"`
}

// NewCreatePipelineRequest instantiates a new CreatePipelineRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePipelineRequest(args interface{}, frequency PipelineFrequency, pipeline string, title string) *CreatePipelineRequest {
	this := CreatePipelineRequest{}
	this.Args = args
	this.Frequency = frequency
	this.Pipeline = pipeline
	this.Title = title
	return &this
}

// NewCreatePipelineRequestWithDefaults instantiates a new CreatePipelineRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePipelineRequestWithDefaults() *CreatePipelineRequest {
	this := CreatePipelineRequest{}
	return &this
}

// GetArgs returns the Args field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CreatePipelineRequest) GetArgs() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePipelineRequest) GetArgsOk() (*interface{}, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *CreatePipelineRequest) SetArgs(v interface{}) {
	o.Args = v
}

// GetFrequency returns the Frequency field value
func (o *CreatePipelineRequest) GetFrequency() PipelineFrequency {
	if o == nil {
		var ret PipelineFrequency
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *CreatePipelineRequest) GetFrequencyOk() (*PipelineFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *CreatePipelineRequest) SetFrequency(v PipelineFrequency) {
	o.Frequency = v
}

// GetPipeline returns the Pipeline field value
func (o *CreatePipelineRequest) GetPipeline() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value
// and a boolean to check if the value has been set.
func (o *CreatePipelineRequest) GetPipelineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// SetPipeline sets field value
func (o *CreatePipelineRequest) SetPipeline(v string) {
	o.Pipeline = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreatePipelineRequest) GetTags() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreatePipelineRequest) GetTagsOk() (*interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreatePipelineRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given interface{} and assigns it to the Tags field.
func (o *CreatePipelineRequest) SetTags(v interface{}) {
	o.Tags = v
}

// GetTitle returns the Title field value
func (o *CreatePipelineRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreatePipelineRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreatePipelineRequest) SetTitle(v string) {
	o.Title = v
}

func (o CreatePipelineRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Args != nil {
		toSerialize["args"] = o.Args
	}
	if true {
		toSerialize["frequency"] = o.Frequency
	}
	if true {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableCreatePipelineRequest struct {
	value *CreatePipelineRequest
	isSet bool
}

func (v NullableCreatePipelineRequest) Get() *CreatePipelineRequest {
	return v.value
}

func (v *NullableCreatePipelineRequest) Set(val *CreatePipelineRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePipelineRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePipelineRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePipelineRequest(val *CreatePipelineRequest) *NullableCreatePipelineRequest {
	return &NullableCreatePipelineRequest{value: val, isSet: true}
}

func (v NullableCreatePipelineRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePipelineRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
