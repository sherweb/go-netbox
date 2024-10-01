/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.1.2 (4.1)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the DataFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataFile{}

// DataFile Adds support for custom fields and tags.
type DataFile struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl *string `json:"display_url,omitempty"`
	Display string `json:"display"`
	Source BriefDataSource `json:"source"`
	// File path relative to the data source's root
	Path string `json:"path"`
	LastUpdated time.Time `json:"last_updated"`
	Size int32 `json:"size"`
	// SHA256 hash of the file data
	Hash string `json:"hash"`
	AdditionalProperties map[string]interface{}
}

type _DataFile DataFile

// NewDataFile instantiates a new DataFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataFile(id int32, url string, display string, source BriefDataSource, path string, lastUpdated time.Time, size int32, hash string) *DataFile {
	this := DataFile{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Source = source
	this.Path = path
	this.LastUpdated = lastUpdated
	this.Size = size
	this.Hash = hash
	return &this
}

// NewDataFileWithDefaults instantiates a new DataFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataFileWithDefaults() *DataFile {
	this := DataFile{}
	return &this
}

// GetId returns the Id field value
func (o *DataFile) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DataFile) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DataFile) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *DataFile) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DataFile) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DataFile) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value if set, zero value otherwise.
func (o *DataFile) GetDisplayUrl() string {
	if o == nil || IsNil(o.DisplayUrl) {
		var ret string
		return ret
	}
	return *o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataFile) GetDisplayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUrl) {
		return nil, false
	}
	return o.DisplayUrl, true
}

// HasDisplayUrl returns a boolean if a field has been set.
func (o *DataFile) HasDisplayUrl() bool {
	if o != nil && !IsNil(o.DisplayUrl) {
		return true
	}

	return false
}

// SetDisplayUrl gets a reference to the given string and assigns it to the DisplayUrl field.
func (o *DataFile) SetDisplayUrl(v string) {
	o.DisplayUrl = &v
}

// GetDisplay returns the Display field value
func (o *DataFile) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DataFile) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *DataFile) SetDisplay(v string) {
	o.Display = v
}

// GetSource returns the Source field value
func (o *DataFile) GetSource() BriefDataSource {
	if o == nil {
		var ret BriefDataSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DataFile) GetSourceOk() (*BriefDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *DataFile) SetSource(v BriefDataSource) {
	o.Source = v
}

// GetPath returns the Path field value
func (o *DataFile) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DataFile) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DataFile) SetPath(v string) {
	o.Path = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *DataFile) GetLastUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *DataFile) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *DataFile) SetLastUpdated(v time.Time) {
	o.LastUpdated = v
}

// GetSize returns the Size field value
func (o *DataFile) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *DataFile) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *DataFile) SetSize(v int32) {
	o.Size = v
}

// GetHash returns the Hash field value
func (o *DataFile) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *DataFile) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *DataFile) SetHash(v string) {
	o.Hash = v
}

func (o DataFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	if !IsNil(o.DisplayUrl) {
		toSerialize["display_url"] = o.DisplayUrl
	}
	toSerialize["display"] = o.Display
	toSerialize["source"] = o.Source
	toSerialize["path"] = o.Path
	toSerialize["last_updated"] = o.LastUpdated
	toSerialize["size"] = o.Size
	toSerialize["hash"] = o.Hash

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DataFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"source",
		"path",
		"last_updated",
		"size",
		"hash",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDataFile := _DataFile{}

	err = json.Unmarshal(data, &varDataFile)

	if err != nil {
		return err
	}

	*o = DataFile(varDataFile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "source")
		delete(additionalProperties, "path")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "size")
		delete(additionalProperties, "hash")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDataFile struct {
	value *DataFile
	isSet bool
}

func (v NullableDataFile) Get() *DataFile {
	return v.value
}

func (v *NullableDataFile) Set(val *DataFile) {
	v.value = val
	v.isSet = true
}

func (v NullableDataFile) IsSet() bool {
	return v.isSet
}

func (v *NullableDataFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataFile(val *DataFile) *NullableDataFile {
	return &NullableDataFile{value: val, isSet: true}
}

func (v NullableDataFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


