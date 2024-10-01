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

// checks if the WirelessLAN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelessLAN{}

// WirelessLAN Adds support for custom fields and tags.
type WirelessLAN struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl *string `json:"display_url,omitempty"`
	Display string `json:"display"`
	Ssid string `json:"ssid"`
	Description *string `json:"description,omitempty"`
	Group NullableBriefWirelessLANGroup `json:"group,omitempty"`
	Status *WirelessLANStatus `json:"status,omitempty"`
	Vlan NullableBriefVLAN `json:"vlan,omitempty"`
	Tenant NullableBriefTenant `json:"tenant,omitempty"`
	AuthType *WirelessLANAuthType `json:"auth_type,omitempty"`
	AuthCipher *WirelessLANAuthCipher `json:"auth_cipher,omitempty"`
	AuthPsk *string `json:"auth_psk,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _WirelessLAN WirelessLAN

// NewWirelessLAN instantiates a new WirelessLAN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelessLAN(id int32, url string, display string, ssid string, created NullableTime, lastUpdated NullableTime) *WirelessLAN {
	this := WirelessLAN{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.Ssid = ssid
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewWirelessLANWithDefaults instantiates a new WirelessLAN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelessLANWithDefaults() *WirelessLAN {
	this := WirelessLAN{}
	return &this
}

// GetId returns the Id field value
func (o *WirelessLAN) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WirelessLAN) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WirelessLAN) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WirelessLAN) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value if set, zero value otherwise.
func (o *WirelessLAN) GetDisplayUrl() string {
	if o == nil || IsNil(o.DisplayUrl) {
		var ret string
		return ret
	}
	return *o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetDisplayUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayUrl) {
		return nil, false
	}
	return o.DisplayUrl, true
}

// HasDisplayUrl returns a boolean if a field has been set.
func (o *WirelessLAN) HasDisplayUrl() bool {
	if o != nil && !IsNil(o.DisplayUrl) {
		return true
	}

	return false
}

// SetDisplayUrl gets a reference to the given string and assigns it to the DisplayUrl field.
func (o *WirelessLAN) SetDisplayUrl(v string) {
	o.DisplayUrl = &v
}

// GetDisplay returns the Display field value
func (o *WirelessLAN) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WirelessLAN) SetDisplay(v string) {
	o.Display = v
}

// GetSsid returns the Ssid field value
func (o *WirelessLAN) GetSsid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetSsidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ssid, true
}

// SetSsid sets field value
func (o *WirelessLAN) SetSsid(v string) {
	o.Ssid = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WirelessLAN) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WirelessLAN) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WirelessLAN) SetDescription(v string) {
	o.Description = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WirelessLAN) GetGroup() BriefWirelessLANGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefWirelessLANGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessLAN) GetGroupOk() (*BriefWirelessLANGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *WirelessLAN) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefWirelessLANGroup and assigns it to the Group field.
func (o *WirelessLAN) SetGroup(v BriefWirelessLANGroup) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *WirelessLAN) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *WirelessLAN) UnsetGroup() {
	o.Group.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WirelessLAN) GetStatus() WirelessLANStatus {
	if o == nil || IsNil(o.Status) {
		var ret WirelessLANStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetStatusOk() (*WirelessLANStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WirelessLAN) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WirelessLANStatus and assigns it to the Status field.
func (o *WirelessLAN) SetStatus(v WirelessLANStatus) {
	o.Status = &v
}

// GetVlan returns the Vlan field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WirelessLAN) GetVlan() BriefVLAN {
	if o == nil || IsNil(o.Vlan.Get()) {
		var ret BriefVLAN
		return ret
	}
	return *o.Vlan.Get()
}

// GetVlanOk returns a tuple with the Vlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessLAN) GetVlanOk() (*BriefVLAN, bool) {
	if o == nil {
		return nil, false
	}
	return o.Vlan.Get(), o.Vlan.IsSet()
}

// HasVlan returns a boolean if a field has been set.
func (o *WirelessLAN) HasVlan() bool {
	if o != nil && o.Vlan.IsSet() {
		return true
	}

	return false
}

// SetVlan gets a reference to the given NullableBriefVLAN and assigns it to the Vlan field.
func (o *WirelessLAN) SetVlan(v BriefVLAN) {
	o.Vlan.Set(&v)
}
// SetVlanNil sets the value for Vlan to be an explicit nil
func (o *WirelessLAN) SetVlanNil() {
	o.Vlan.Set(nil)
}

// UnsetVlan ensures that no value is present for Vlan, not even an explicit nil
func (o *WirelessLAN) UnsetVlan() {
	o.Vlan.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WirelessLAN) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessLAN) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WirelessLAN) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *WirelessLAN) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WirelessLAN) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WirelessLAN) UnsetTenant() {
	o.Tenant.Unset()
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *WirelessLAN) GetAuthType() WirelessLANAuthType {
	if o == nil || IsNil(o.AuthType) {
		var ret WirelessLANAuthType
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetAuthTypeOk() (*WirelessLANAuthType, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *WirelessLAN) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given WirelessLANAuthType and assigns it to the AuthType field.
func (o *WirelessLAN) SetAuthType(v WirelessLANAuthType) {
	o.AuthType = &v
}

// GetAuthCipher returns the AuthCipher field value if set, zero value otherwise.
func (o *WirelessLAN) GetAuthCipher() WirelessLANAuthCipher {
	if o == nil || IsNil(o.AuthCipher) {
		var ret WirelessLANAuthCipher
		return ret
	}
	return *o.AuthCipher
}

// GetAuthCipherOk returns a tuple with the AuthCipher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetAuthCipherOk() (*WirelessLANAuthCipher, bool) {
	if o == nil || IsNil(o.AuthCipher) {
		return nil, false
	}
	return o.AuthCipher, true
}

// HasAuthCipher returns a boolean if a field has been set.
func (o *WirelessLAN) HasAuthCipher() bool {
	if o != nil && !IsNil(o.AuthCipher) {
		return true
	}

	return false
}

// SetAuthCipher gets a reference to the given WirelessLANAuthCipher and assigns it to the AuthCipher field.
func (o *WirelessLAN) SetAuthCipher(v WirelessLANAuthCipher) {
	o.AuthCipher = &v
}

// GetAuthPsk returns the AuthPsk field value if set, zero value otherwise.
func (o *WirelessLAN) GetAuthPsk() string {
	if o == nil || IsNil(o.AuthPsk) {
		var ret string
		return ret
	}
	return *o.AuthPsk
}

// GetAuthPskOk returns a tuple with the AuthPsk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetAuthPskOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPsk) {
		return nil, false
	}
	return o.AuthPsk, true
}

// HasAuthPsk returns a boolean if a field has been set.
func (o *WirelessLAN) HasAuthPsk() bool {
	if o != nil && !IsNil(o.AuthPsk) {
		return true
	}

	return false
}

// SetAuthPsk gets a reference to the given string and assigns it to the AuthPsk field.
func (o *WirelessLAN) SetAuthPsk(v string) {
	o.AuthPsk = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WirelessLAN) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WirelessLAN) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WirelessLAN) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WirelessLAN) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WirelessLAN) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *WirelessLAN) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WirelessLAN) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLAN) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WirelessLAN) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WirelessLAN) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WirelessLAN) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessLAN) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *WirelessLAN) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WirelessLAN) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessLAN) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *WirelessLAN) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o WirelessLAN) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelessLAN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	if !IsNil(o.DisplayUrl) {
		toSerialize["display_url"] = o.DisplayUrl
	}
	toSerialize["display"] = o.Display
	toSerialize["ssid"] = o.Ssid
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Vlan.IsSet() {
		toSerialize["vlan"] = o.Vlan.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.AuthCipher) {
		toSerialize["auth_cipher"] = o.AuthCipher
	}
	if !IsNil(o.AuthPsk) {
		toSerialize["auth_psk"] = o.AuthPsk
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

func (o *WirelessLAN) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"ssid",
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

	varWirelessLAN := _WirelessLAN{}

	err = json.Unmarshal(data, &varWirelessLAN)

	if err != nil {
		return err
	}

	*o = WirelessLAN(varWirelessLAN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "description")
		delete(additionalProperties, "group")
		delete(additionalProperties, "status")
		delete(additionalProperties, "vlan")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "auth_cipher")
		delete(additionalProperties, "auth_psk")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWirelessLAN struct {
	value *WirelessLAN
	isSet bool
}

func (v NullableWirelessLAN) Get() *WirelessLAN {
	return v.value
}

func (v *NullableWirelessLAN) Set(val *WirelessLAN) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLAN) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLAN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLAN(val *WirelessLAN) *NullableWirelessLAN {
	return &NullableWirelessLAN{value: val, isSet: true}
}

func (v NullableWirelessLAN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLAN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


