/**
 * Go SDK for OpenFGA
 *
 * API version: 0.1
 * Website: https://openfga.dev
 * Documentation: https://openfga.dev/docs
 * Support: https://discord.gg/8naAwJfWN6
 * License: [Apache-2.0](https://github.com/openfga/go-sdk/blob/main/LICENSE)
 *
 * NOTE: This file was auto generated by OpenAPI Generator (https://openapi-generator.tech). DO NOT EDIT.
 */

package openfga

import (
	"encoding/json"
)

// TupleToUserset struct for TupleToUserset
type TupleToUserset struct {
	Tupleset        *ObjectRelation `json:"tupleset,omitempty"`
	ComputedUserset *ObjectRelation `json:"computedUserset,omitempty"`
}

// NewTupleToUserset instantiates a new TupleToUserset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTupleToUserset() *TupleToUserset {
	this := TupleToUserset{}
	return &this
}

// NewTupleToUsersetWithDefaults instantiates a new TupleToUserset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTupleToUsersetWithDefaults() *TupleToUserset {
	this := TupleToUserset{}
	return &this
}

// GetTupleset returns the Tupleset field value if set, zero value otherwise.
func (o *TupleToUserset) GetTupleset() ObjectRelation {
	if o == nil || o.Tupleset == nil {
		var ret ObjectRelation
		return ret
	}
	return *o.Tupleset
}

// GetTuplesetOk returns a tuple with the Tupleset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TupleToUserset) GetTuplesetOk() (*ObjectRelation, bool) {
	if o == nil || o.Tupleset == nil {
		return nil, false
	}
	return o.Tupleset, true
}

// HasTupleset returns a boolean if a field has been set.
func (o *TupleToUserset) HasTupleset() bool {
	if o != nil && o.Tupleset != nil {
		return true
	}

	return false
}

// SetTupleset gets a reference to the given ObjectRelation and assigns it to the Tupleset field.
func (o *TupleToUserset) SetTupleset(v ObjectRelation) {
	o.Tupleset = &v
}

// GetComputedUserset returns the ComputedUserset field value if set, zero value otherwise.
func (o *TupleToUserset) GetComputedUserset() ObjectRelation {
	if o == nil || o.ComputedUserset == nil {
		var ret ObjectRelation
		return ret
	}
	return *o.ComputedUserset
}

// GetComputedUsersetOk returns a tuple with the ComputedUserset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TupleToUserset) GetComputedUsersetOk() (*ObjectRelation, bool) {
	if o == nil || o.ComputedUserset == nil {
		return nil, false
	}
	return o.ComputedUserset, true
}

// HasComputedUserset returns a boolean if a field has been set.
func (o *TupleToUserset) HasComputedUserset() bool {
	if o != nil && o.ComputedUserset != nil {
		return true
	}

	return false
}

// SetComputedUserset gets a reference to the given ObjectRelation and assigns it to the ComputedUserset field.
func (o *TupleToUserset) SetComputedUserset(v ObjectRelation) {
	o.ComputedUserset = &v
}

func (o TupleToUserset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tupleset != nil {
		toSerialize["tupleset"] = o.Tupleset
	}
	if o.ComputedUserset != nil {
		toSerialize["computedUserset"] = o.ComputedUserset
	}
	return json.Marshal(toSerialize)
}

type NullableTupleToUserset struct {
	value *TupleToUserset
	isSet bool
}

func (v NullableTupleToUserset) Get() *TupleToUserset {
	return v.value
}

func (v *NullableTupleToUserset) Set(val *TupleToUserset) {
	v.value = val
	v.isSet = true
}

func (v NullableTupleToUserset) IsSet() bool {
	return v.isSet
}

func (v *NullableTupleToUserset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTupleToUserset(val *TupleToUserset) *NullableTupleToUserset {
	return &NullableTupleToUserset{value: val, isSet: true}
}

func (v NullableTupleToUserset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTupleToUserset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}