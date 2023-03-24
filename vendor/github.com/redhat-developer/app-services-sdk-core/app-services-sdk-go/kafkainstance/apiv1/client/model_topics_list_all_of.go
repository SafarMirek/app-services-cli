/*
 * Kafka Instance API
 *
 * API for interacting with Kafka Instance. Includes Produce, Consume and Admin APIs
 *
 * API version: 0.13.1-SNAPSHOT
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// TopicsListAllOf A list of topics.
type TopicsListAllOf struct {
	Items []Topic `json:"items"`
}

// NewTopicsListAllOf instantiates a new TopicsListAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopicsListAllOf(items []Topic) *TopicsListAllOf {
	this := TopicsListAllOf{}
	this.Items = items
	return &this
}

// NewTopicsListAllOfWithDefaults instantiates a new TopicsListAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopicsListAllOfWithDefaults() *TopicsListAllOf {
	this := TopicsListAllOf{}
	return &this
}

// GetItems returns the Items field value
func (o *TopicsListAllOf) GetItems() []Topic {
	if o == nil {
		var ret []Topic
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TopicsListAllOf) GetItemsOk() (*[]Topic, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *TopicsListAllOf) SetItems(v []Topic) {
	o.Items = v
}

func (o TopicsListAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableTopicsListAllOf struct {
	value *TopicsListAllOf
	isSet bool
}

func (v NullableTopicsListAllOf) Get() *TopicsListAllOf {
	return v.value
}

func (v *NullableTopicsListAllOf) Set(val *TopicsListAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTopicsListAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTopicsListAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopicsListAllOf(val *TopicsListAllOf) *NullableTopicsListAllOf {
	return &NullableTopicsListAllOf{value: val, isSet: true}
}

func (v NullableTopicsListAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopicsListAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

