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

// checks if the ASN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ASN{}

// ASN Adds support for custom fields and tags.
type ASN struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	// 16- or 32-bit autonomous system number
	Asn int64 `json:"asn"`
	Rir NullableBriefRIR `json:"rir,omitempty"`
	Tenant NullableBriefTenant `json:"tenant,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	SiteCount int64 `json:"site_count"`
	ProviderCount int64 `json:"provider_count"`
	AdditionalProperties map[string]interface{}
}

type _ASN ASN

// NewASN instantiates a new ASN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewASN(id int32, url string, displayUrl string, display string, asn int64, created NullableTime, lastUpdated NullableTime, siteCount int64, providerCount int64) *ASN {
	this := ASN{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Asn = asn
	this.Created = created
	this.LastUpdated = lastUpdated
	this.SiteCount = siteCount
	this.ProviderCount = providerCount
	return &this
}

// NewASNWithDefaults instantiates a new ASN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewASNWithDefaults() *ASN {
	this := ASN{}
	return &this
}

// GetId returns the Id field value
func (o *ASN) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ASN) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ASN) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ASN) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ASN) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ASN) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *ASN) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *ASN) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *ASN) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *ASN) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *ASN) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *ASN) SetDisplay(v string) {
	o.Display = v
}

// GetAsn returns the Asn field value
func (o *ASN) GetAsn() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Asn
}

// GetAsnOk returns a tuple with the Asn field value
// and a boolean to check if the value has been set.
func (o *ASN) GetAsnOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asn, true
}

// SetAsn sets field value
func (o *ASN) SetAsn(v int64) {
	o.Asn = v
}

// GetRir returns the Rir field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ASN) GetRir() BriefRIR {
	if o == nil || IsNil(o.Rir.Get()) {
		var ret BriefRIR
		return ret
	}
	return *o.Rir.Get()
}

// GetRirOk returns a tuple with the Rir field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ASN) GetRirOk() (*BriefRIR, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rir.Get(), o.Rir.IsSet()
}

// HasRir returns a boolean if a field has been set.
func (o *ASN) HasRir() bool {
	if o != nil && o.Rir.IsSet() {
		return true
	}

	return false
}

// SetRir gets a reference to the given NullableBriefRIR and assigns it to the Rir field.
func (o *ASN) SetRir(v BriefRIR) {
	o.Rir.Set(&v)
}
// SetRirNil sets the value for Rir to be an explicit nil
func (o *ASN) SetRirNil() {
	o.Rir.Set(nil)
}

// UnsetRir ensures that no value is present for Rir, not even an explicit nil
func (o *ASN) UnsetRir() {
	o.Rir.Unset()
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ASN) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ASN) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *ASN) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *ASN) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *ASN) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *ASN) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ASN) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASN) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ASN) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ASN) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ASN) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASN) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ASN) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ASN) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ASN) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASN) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ASN) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *ASN) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *ASN) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASN) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *ASN) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *ASN) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ASN) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ASN) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *ASN) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *ASN) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ASN) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *ASN) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetSiteCount returns the SiteCount field value
func (o *ASN) GetSiteCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SiteCount
}

// GetSiteCountOk returns a tuple with the SiteCount field value
// and a boolean to check if the value has been set.
func (o *ASN) GetSiteCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteCount, true
}

// SetSiteCount sets field value
func (o *ASN) SetSiteCount(v int64) {
	o.SiteCount = v
}

// GetProviderCount returns the ProviderCount field value
func (o *ASN) GetProviderCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ProviderCount
}

// GetProviderCountOk returns a tuple with the ProviderCount field value
// and a boolean to check if the value has been set.
func (o *ASN) GetProviderCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderCount, true
}

// SetProviderCount sets field value
func (o *ASN) SetProviderCount(v int64) {
	o.ProviderCount = v
}

func (o ASN) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ASN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	toSerialize["asn"] = o.Asn
	if o.Rir.IsSet() {
		toSerialize["rir"] = o.Rir.Get()
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
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
	toSerialize["site_count"] = o.SiteCount
	toSerialize["provider_count"] = o.ProviderCount

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ASN) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"asn",
		"created",
		"last_updated",
		"site_count",
		"provider_count",
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

	varASN := _ASN{}

	err = json.Unmarshal(data, &varASN)

	if err != nil {
		return err
	}

	*o = ASN(varASN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "asn")
		delete(additionalProperties, "rir")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "site_count")
		delete(additionalProperties, "provider_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableASN struct {
	value *ASN
	isSet bool
}

func (v NullableASN) Get() *ASN {
	return v.value
}

func (v *NullableASN) Set(val *ASN) {
	v.value = val
	v.isSet = true
}

func (v NullableASN) IsSet() bool {
	return v.isSet
}

func (v *NullableASN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableASN(val *ASN) *NullableASN {
	return &NullableASN{value: val, isSet: true}
}

func (v NullableASN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableASN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


