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

// checks if the IKEProposal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IKEProposal{}

// IKEProposal Adds support for custom fields and tags.
type IKEProposal struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl *string `json:"display_url,omitempty"`
	Display string `json:"display"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	AuthenticationMethod IKEProposalAuthenticationMethod `json:"authentication_method"`
	EncryptionAlgorithm IKEProposalEncryptionAlgorithm `json:"encryption_algorithm"`
	AuthenticationAlgorithm *IKEProposalAuthenticationAlgorithm `json:"authentication_algorithm,omitempty"`
	Group IKEProposalGroup `json:"group"`
	// Security association lifetime (in seconds)
	SaLifetime NullableInt32 `json:"sa_lifetime,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _IKEProposal IKEProposal

// NewIKEProposal instantiates a new IKEProposal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIKEProposal(id int32, url string, display string, name string, authenticationMethod IKEProposalAuthenticationMethod, encryptionAlgorithm IKEProposalEncryptionAlgorithm, group IKEProposalGroup, created NullableTime, lastUpdated NullableTime) *IKEProposal {
	this := IKEProposal{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.AuthenticationMethod = authenticationMethod
	this.EncryptionAlgorithm = encryptionAlgorithm
	this.Group = group
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewIKEProposalWithDefaults instantiates a new IKEProposal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIKEProposalWithDefaults() *IKEProposal {
	this := IKEProposal{}
	return &this
}

// GetId returns the Id field value
func (o *IKEProposal) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IKEProposal) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *IKEProposal) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IKEProposal) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value if set, zero value otherwise.
func (o *IKEProposal) GetDisplayUrl() string {
	if o == nil || IsNil(o.DisplayUrl) {
		var ret string
		return ret
	}
	return *o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetDisplayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUrl) {
		return nil, false
	}
	return o.DisplayUrl, true
}

// HasDisplayUrl returns a boolean if a field has been set.
func (o *IKEProposal) HasDisplayUrl() bool {
	if o != nil && !IsNil(o.DisplayUrl) {
		return true
	}

	return false
}

// SetDisplayUrl gets a reference to the given string and assigns it to the DisplayUrl field.
func (o *IKEProposal) SetDisplayUrl(v string) {
	o.DisplayUrl = &v
}

// GetDisplay returns the Display field value
func (o *IKEProposal) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *IKEProposal) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *IKEProposal) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IKEProposal) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IKEProposal) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IKEProposal) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IKEProposal) SetDescription(v string) {
	o.Description = &v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *IKEProposal) GetAuthenticationMethod() IKEProposalAuthenticationMethod {
	if o == nil {
		var ret IKEProposalAuthenticationMethod
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetAuthenticationMethodOk() (*IKEProposalAuthenticationMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *IKEProposal) SetAuthenticationMethod(v IKEProposalAuthenticationMethod) {
	o.AuthenticationMethod = v
}

// GetEncryptionAlgorithm returns the EncryptionAlgorithm field value
func (o *IKEProposal) GetEncryptionAlgorithm() IKEProposalEncryptionAlgorithm {
	if o == nil {
		var ret IKEProposalEncryptionAlgorithm
		return ret
	}

	return o.EncryptionAlgorithm
}

// GetEncryptionAlgorithmOk returns a tuple with the EncryptionAlgorithm field value
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetEncryptionAlgorithmOk() (*IKEProposalEncryptionAlgorithm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncryptionAlgorithm, true
}

// SetEncryptionAlgorithm sets field value
func (o *IKEProposal) SetEncryptionAlgorithm(v IKEProposalEncryptionAlgorithm) {
	o.EncryptionAlgorithm = v
}

// GetAuthenticationAlgorithm returns the AuthenticationAlgorithm field value if set, zero value otherwise.
func (o *IKEProposal) GetAuthenticationAlgorithm() IKEProposalAuthenticationAlgorithm {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		var ret IKEProposalAuthenticationAlgorithm
		return ret
	}
	return *o.AuthenticationAlgorithm
}

// GetAuthenticationAlgorithmOk returns a tuple with the AuthenticationAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetAuthenticationAlgorithmOk() (*IKEProposalAuthenticationAlgorithm, bool) {
	if o == nil || IsNil(o.AuthenticationAlgorithm) {
		return nil, false
	}
	return o.AuthenticationAlgorithm, true
}

// HasAuthenticationAlgorithm returns a boolean if a field has been set.
func (o *IKEProposal) HasAuthenticationAlgorithm() bool {
	if o != nil && !IsNil(o.AuthenticationAlgorithm) {
		return true
	}

	return false
}

// SetAuthenticationAlgorithm gets a reference to the given IKEProposalAuthenticationAlgorithm and assigns it to the AuthenticationAlgorithm field.
func (o *IKEProposal) SetAuthenticationAlgorithm(v IKEProposalAuthenticationAlgorithm) {
	o.AuthenticationAlgorithm = &v
}

// GetGroup returns the Group field value
func (o *IKEProposal) GetGroup() IKEProposalGroup {
	if o == nil {
		var ret IKEProposalGroup
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetGroupOk() (*IKEProposalGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *IKEProposal) SetGroup(v IKEProposalGroup) {
	o.Group = v
}

// GetSaLifetime returns the SaLifetime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IKEProposal) GetSaLifetime() int32 {
	if o == nil || IsNil(o.SaLifetime.Get()) {
		var ret int32
		return ret
	}
	return *o.SaLifetime.Get()
}

// GetSaLifetimeOk returns a tuple with the SaLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IKEProposal) GetSaLifetimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SaLifetime.Get(), o.SaLifetime.IsSet()
}

// HasSaLifetime returns a boolean if a field has been set.
func (o *IKEProposal) HasSaLifetime() bool {
	if o != nil && o.SaLifetime.IsSet() {
		return true
	}

	return false
}

// SetSaLifetime gets a reference to the given NullableInt32 and assigns it to the SaLifetime field.
func (o *IKEProposal) SetSaLifetime(v int32) {
	o.SaLifetime.Set(&v)
}
// SetSaLifetimeNil sets the value for SaLifetime to be an explicit nil
func (o *IKEProposal) SetSaLifetimeNil() {
	o.SaLifetime.Set(nil)
}

// UnsetSaLifetime ensures that no value is present for SaLifetime, not even an explicit nil
func (o *IKEProposal) UnsetSaLifetime() {
	o.SaLifetime.Unset()
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IKEProposal) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IKEProposal) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *IKEProposal) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IKEProposal) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IKEProposal) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *IKEProposal) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IKEProposal) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IKEProposal) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IKEProposal) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *IKEProposal) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IKEProposal) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IKEProposal) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *IKEProposal) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *IKEProposal) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IKEProposal) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *IKEProposal) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o IKEProposal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IKEProposal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	if !IsNil(o.DisplayUrl) {
		toSerialize["display_url"] = o.DisplayUrl
	}
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["authentication_method"] = o.AuthenticationMethod
	toSerialize["encryption_algorithm"] = o.EncryptionAlgorithm
	if !IsNil(o.AuthenticationAlgorithm) {
		toSerialize["authentication_algorithm"] = o.AuthenticationAlgorithm
	}
	toSerialize["group"] = o.Group
	if o.SaLifetime.IsSet() {
		toSerialize["sa_lifetime"] = o.SaLifetime.Get()
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IKEProposal) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"authentication_method",
		"encryption_algorithm",
		"group",
		"created",
		"last_updated",
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

	varIKEProposal := _IKEProposal{}

	err = json.Unmarshal(data, &varIKEProposal)

	if err != nil {
		return err
	}

	*o = IKEProposal(varIKEProposal)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "authentication_method")
		delete(additionalProperties, "encryption_algorithm")
		delete(additionalProperties, "authentication_algorithm")
		delete(additionalProperties, "group")
		delete(additionalProperties, "sa_lifetime")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIKEProposal struct {
	value *IKEProposal
	isSet bool
}

func (v NullableIKEProposal) Get() *IKEProposal {
	return v.value
}

func (v *NullableIKEProposal) Set(val *IKEProposal) {
	v.value = val
	v.isSet = true
}

func (v NullableIKEProposal) IsSet() bool {
	return v.isSet
}

func (v *NullableIKEProposal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIKEProposal(val *IKEProposal) *NullableIKEProposal {
	return &NullableIKEProposal{value: val, isSet: true}
}

func (v NullableIKEProposal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIKEProposal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


