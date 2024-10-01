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

// checks if the Tunnel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tunnel{}

// Tunnel Adds support for custom fields and tags.
type Tunnel struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl *string `json:"display_url,omitempty"`
	Display string `json:"display"`
	Name string `json:"name"`
	Status TunnelStatus `json:"status"`
	Group NullableBriefTunnelGroup `json:"group,omitempty"`
	Encapsulation TunnelEncapsulation `json:"encapsulation"`
	IpsecProfile NullableBriefIPSecProfile `json:"ipsec_profile,omitempty"`
	Tenant NullableBriefTenant `json:"tenant,omitempty"`
	TunnelId NullableInt64 `json:"tunnel_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	TerminationsCount int64 `json:"terminations_count"`
	AdditionalProperties map[string]interface{}
}

type _Tunnel Tunnel

// NewTunnel instantiates a new Tunnel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTunnel(id int32, url string, display string, name string, status TunnelStatus, encapsulation TunnelEncapsulation, created NullableTime, lastUpdated NullableTime, terminationsCount int64) *Tunnel {
	this := Tunnel{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Name = name
	this.Status = status
	this.Encapsulation = encapsulation
	this.Created = created
	this.LastUpdated = lastUpdated
	this.TerminationsCount = terminationsCount
	return &this
}

// NewTunnelWithDefaults instantiates a new Tunnel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTunnelWithDefaults() *Tunnel {
	this := Tunnel{}
	return &this
}

// GetId returns the Id field value
func (o *Tunnel) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tunnel) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tunnel) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *Tunnel) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Tunnel) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Tunnel) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value if set, zero value otherwise.
func (o *Tunnel) GetDisplayUrl() string {
	if o == nil || IsNil(o.DisplayUrl) {
		var ret string
		return ret
	}
	return *o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tunnel) GetDisplayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUrl) {
		return nil, false
	}
	return o.DisplayUrl, true
}

// HasDisplayUrl returns a boolean if a field has been set.
func (o *Tunnel) HasDisplayUrl() bool {
	if o != nil && !IsNil(o.DisplayUrl) {
		return true
	}

	return false
}

// SetDisplayUrl gets a reference to the given string and assigns it to the DisplayUrl field.
func (o *Tunnel) SetDisplayUrl(v string) {
	o.DisplayUrl = &v
}

// GetDisplay returns the Display field value
func (o *Tunnel) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Tunnel) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *Tunnel) SetDisplay(v string) {
	o.Display = v
}

// GetName returns the Name field value
func (o *Tunnel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tunnel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tunnel) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *Tunnel) GetStatus() TunnelStatus {
	if o == nil {
		var ret TunnelStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Tunnel) GetStatusOk() (*TunnelStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Tunnel) SetStatus(v TunnelStatus) {
	o.Status = v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tunnel) GetGroup() BriefTunnelGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefTunnelGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tunnel) GetGroupOk() (*BriefTunnelGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *Tunnel) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefTunnelGroup and assigns it to the Group field.
func (o *Tunnel) SetGroup(v BriefTunnelGroup) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *Tunnel) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *Tunnel) UnsetGroup() {
	o.Group.Unset()
}

// GetEncapsulation returns the Encapsulation field value
func (o *Tunnel) GetEncapsulation() TunnelEncapsulation {
	if o == nil {
		var ret TunnelEncapsulation
		return ret
	}

	return o.Encapsulation
}

// GetEncapsulationOk returns a tuple with the Encapsulation field value
// and a boolean to check if the value has been set.
func (o *Tunnel) GetEncapsulationOk() (*TunnelEncapsulation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encapsulation, true
}

// SetEncapsulation sets field value
func (o *Tunnel) SetEncapsulation(v TunnelEncapsulation) {
	o.Encapsulation = v
}

// GetIpsecProfile returns the IpsecProfile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tunnel) GetIpsecProfile() BriefIPSecProfile {
	if o == nil || IsNil(o.IpsecProfile.Get()) {
		var ret BriefIPSecProfile
		return ret
	}
	return *o.IpsecProfile.Get()
}

// GetIpsecProfileOk returns a tuple with the IpsecProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tunnel) GetIpsecProfileOk() (*BriefIPSecProfile, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpsecProfile.Get(), o.IpsecProfile.IsSet()
}

// HasIpsecProfile returns a boolean if a field has been set.
func (o *Tunnel) HasIpsecProfile() bool {
	if o != nil && o.IpsecProfile.IsSet() {
		return true
	}

	return false
}

// SetIpsecProfile gets a reference to the given NullableBriefIPSecProfile and assigns it to the IpsecProfile field.
func (o *Tunnel) SetIpsecProfile(v BriefIPSecProfile) {
	o.IpsecProfile.Set(&v)
}
// SetIpsecProfileNil sets the value for IpsecProfile to be an explicit nil
func (o *Tunnel) SetIpsecProfileNil() {
	o.IpsecProfile.Set(nil)
}

// UnsetIpsecProfile ensures that no value is present for IpsecProfile, not even an explicit nil
func (o *Tunnel) UnsetIpsecProfile() {
	o.IpsecProfile.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tunnel) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tunnel) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *Tunnel) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *Tunnel) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *Tunnel) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *Tunnel) UnsetTenant() {
	o.Tenant.Unset()
}

// GetTunnelId returns the TunnelId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tunnel) GetTunnelId() int64 {
	if o == nil || IsNil(o.TunnelId.Get()) {
		var ret int64
		return ret
	}
	return *o.TunnelId.Get()
}

// GetTunnelIdOk returns a tuple with the TunnelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tunnel) GetTunnelIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TunnelId.Get(), o.TunnelId.IsSet()
}

// HasTunnelId returns a boolean if a field has been set.
func (o *Tunnel) HasTunnelId() bool {
	if o != nil && o.TunnelId.IsSet() {
		return true
	}

	return false
}

// SetTunnelId gets a reference to the given NullableInt64 and assigns it to the TunnelId field.
func (o *Tunnel) SetTunnelId(v int64) {
	o.TunnelId.Set(&v)
}
// SetTunnelIdNil sets the value for TunnelId to be an explicit nil
func (o *Tunnel) SetTunnelIdNil() {
	o.TunnelId.Set(nil)
}

// UnsetTunnelId ensures that no value is present for TunnelId, not even an explicit nil
func (o *Tunnel) UnsetTunnelId() {
	o.TunnelId.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tunnel) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tunnel) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Tunnel) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tunnel) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Tunnel) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tunnel) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Tunnel) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *Tunnel) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Tunnel) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tunnel) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Tunnel) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *Tunnel) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *Tunnel) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tunnel) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *Tunnel) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *Tunnel) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Tunnel) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tunnel) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *Tunnel) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Tunnel) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tunnel) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *Tunnel) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetTerminationsCount returns the TerminationsCount field value
func (o *Tunnel) GetTerminationsCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TerminationsCount
}

// GetTerminationsCountOk returns a tuple with the TerminationsCount field value
// and a boolean to check if the value has been set.
func (o *Tunnel) GetTerminationsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TerminationsCount, true
}

// SetTerminationsCount sets field value
func (o *Tunnel) SetTerminationsCount(v int64) {
	o.TerminationsCount = v
}

func (o Tunnel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tunnel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	if !IsNil(o.DisplayUrl) {
		toSerialize["display_url"] = o.DisplayUrl
	}
	toSerialize["display"] = o.Display
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	toSerialize["encapsulation"] = o.Encapsulation
	if o.IpsecProfile.IsSet() {
		toSerialize["ipsec_profile"] = o.IpsecProfile.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if o.TunnelId.IsSet() {
		toSerialize["tunnel_id"] = o.TunnelId.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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
	toSerialize["terminations_count"] = o.TerminationsCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Tunnel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"name",
		"status",
		"encapsulation",
		"created",
		"last_updated",
		"terminations_count",
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

	varTunnel := _Tunnel{}

	err = json.Unmarshal(data, &varTunnel)

	if err != nil {
		return err
	}

	*o = Tunnel(varTunnel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "group")
		delete(additionalProperties, "encapsulation")
		delete(additionalProperties, "ipsec_profile")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "tunnel_id")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "terminations_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTunnel struct {
	value *Tunnel
	isSet bool
}

func (v NullableTunnel) Get() *Tunnel {
	return v.value
}

func (v *NullableTunnel) Set(val *Tunnel) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnel) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnel(val *Tunnel) *NullableTunnel {
	return &NullableTunnel{value: val, isSet: true}
}

func (v NullableTunnel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


