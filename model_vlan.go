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

// checks if the VLAN type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VLAN{}

// VLAN Adds support for custom fields and tags.
type VLAN struct {
	Id int32 `json:"id"`
	Url string `json:"url"`
	DisplayUrl string `json:"display_url"`
	Display string `json:"display"`
	Site NullableBriefSite `json:"site,omitempty"`
	Group NullableBriefVLANGroup `json:"group,omitempty"`
	// Numeric VLAN ID (1-4094)
	Vid int32 `json:"vid"`
	Name string `json:"name"`
	Tenant NullableBriefTenant `json:"tenant,omitempty"`
	Status *IPRangeStatus `json:"status,omitempty"`
	Role NullableBriefRole `json:"role,omitempty"`
	Description *string `json:"description,omitempty"`
	Comments *string `json:"comments,omitempty"`
	L2vpnTermination NullableBriefL2VPNTermination `json:"l2vpn_termination"`
	Tags []NestedTag `json:"tags,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Created NullableTime `json:"created"`
	LastUpdated NullableTime `json:"last_updated"`
	PrefixCount *int64 `json:"prefix_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VLAN VLAN

// NewVLAN instantiates a new VLAN object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVLAN(id int32, url string, displayUrl string, display string, vid int32, name string, l2vpnTermination NullableBriefL2VPNTermination, created NullableTime, lastUpdated NullableTime) *VLAN {
	this := VLAN{}
	this.Id = id
	this.Url = url
	this.DisplayUrl = displayUrl
	this.Display = display
	this.Vid = vid
	this.Name = name
	this.L2vpnTermination = l2vpnTermination
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewVLANWithDefaults instantiates a new VLAN object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVLANWithDefaults() *VLAN {
	this := VLAN{}
	return &this
}

// GetId returns the Id field value
func (o *VLAN) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VLAN) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *VLAN) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *VLAN) SetUrl(v string) {
	o.Url = v
}

// GetDisplayUrl returns the DisplayUrl field value
func (o *VLAN) GetDisplayUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayUrl
}

// GetDisplayUrlOk returns a tuple with the DisplayUrl field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetDisplayUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayUrl, true
}

// SetDisplayUrl sets field value
func (o *VLAN) SetDisplayUrl(v string) {
	o.DisplayUrl = v
}

// GetDisplay returns the Display field value
func (o *VLAN) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *VLAN) SetDisplay(v string) {
	o.Display = v
}

// GetSite returns the Site field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLAN) GetSite() BriefSite {
	if o == nil || IsNil(o.Site.Get()) {
		var ret BriefSite
		return ret
	}
	return *o.Site.Get()
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetSiteOk() (*BriefSite, bool) {
	if o == nil {
		return nil, false
	}
	return o.Site.Get(), o.Site.IsSet()
}

// HasSite returns a boolean if a field has been set.
func (o *VLAN) HasSite() bool {
	if o != nil && o.Site.IsSet() {
		return true
	}

	return false
}

// SetSite gets a reference to the given NullableBriefSite and assigns it to the Site field.
func (o *VLAN) SetSite(v BriefSite) {
	o.Site.Set(&v)
}
// SetSiteNil sets the value for Site to be an explicit nil
func (o *VLAN) SetSiteNil() {
	o.Site.Set(nil)
}

// UnsetSite ensures that no value is present for Site, not even an explicit nil
func (o *VLAN) UnsetSite() {
	o.Site.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLAN) GetGroup() BriefVLANGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret BriefVLANGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetGroupOk() (*BriefVLANGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *VLAN) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableBriefVLANGroup and assigns it to the Group field.
func (o *VLAN) SetGroup(v BriefVLANGroup) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *VLAN) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *VLAN) UnsetGroup() {
	o.Group.Unset()
}

// GetVid returns the Vid field value
func (o *VLAN) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *VLAN) SetVid(v int32) {
	o.Vid = v
}

// GetName returns the Name field value
func (o *VLAN) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VLAN) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VLAN) SetName(v string) {
	o.Name = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLAN) GetTenant() BriefTenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret BriefTenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetTenantOk() (*BriefTenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *VLAN) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableBriefTenant and assigns it to the Tenant field.
func (o *VLAN) SetTenant(v BriefTenant) {
	o.Tenant.Set(&v)
}
// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *VLAN) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *VLAN) UnsetTenant() {
	o.Tenant.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VLAN) GetStatus() IPRangeStatus {
	if o == nil || IsNil(o.Status) {
		var ret IPRangeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetStatusOk() (*IPRangeStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VLAN) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IPRangeStatus and assigns it to the Status field.
func (o *VLAN) SetStatus(v IPRangeStatus) {
	o.Status = &v
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VLAN) GetRole() BriefRole {
	if o == nil || IsNil(o.Role.Get()) {
		var ret BriefRole
		return ret
	}
	return *o.Role.Get()
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetRoleOk() (*BriefRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.Role.Get(), o.Role.IsSet()
}

// HasRole returns a boolean if a field has been set.
func (o *VLAN) HasRole() bool {
	if o != nil && o.Role.IsSet() {
		return true
	}

	return false
}

// SetRole gets a reference to the given NullableBriefRole and assigns it to the Role field.
func (o *VLAN) SetRole(v BriefRole) {
	o.Role.Set(&v)
}
// SetRoleNil sets the value for Role to be an explicit nil
func (o *VLAN) SetRoleNil() {
	o.Role.Set(nil)
}

// UnsetRole ensures that no value is present for Role, not even an explicit nil
func (o *VLAN) UnsetRole() {
	o.Role.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VLAN) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VLAN) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VLAN) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *VLAN) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *VLAN) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *VLAN) SetComments(v string) {
	o.Comments = &v
}

// GetL2vpnTermination returns the L2vpnTermination field value
// If the value is explicit nil, the zero value for BriefL2VPNTermination will be returned
func (o *VLAN) GetL2vpnTermination() BriefL2VPNTermination {
	if o == nil || o.L2vpnTermination.Get() == nil {
		var ret BriefL2VPNTermination
		return ret
	}

	return *o.L2vpnTermination.Get()
}

// GetL2vpnTerminationOk returns a tuple with the L2vpnTermination field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetL2vpnTerminationOk() (*BriefL2VPNTermination, bool) {
	if o == nil {
		return nil, false
	}
	return o.L2vpnTermination.Get(), o.L2vpnTermination.IsSet()
}

// SetL2vpnTermination sets field value
func (o *VLAN) SetL2vpnTermination(v BriefL2VPNTermination) {
	o.L2vpnTermination.Set(&v)
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VLAN) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VLAN) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *VLAN) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *VLAN) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *VLAN) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *VLAN) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VLAN) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *VLAN) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *VLAN) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VLAN) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *VLAN) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

// GetPrefixCount returns the PrefixCount field value if set, zero value otherwise.
func (o *VLAN) GetPrefixCount() int64 {
	if o == nil || IsNil(o.PrefixCount) {
		var ret int64
		return ret
	}
	return *o.PrefixCount
}

// GetPrefixCountOk returns a tuple with the PrefixCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VLAN) GetPrefixCountOk() (*int64, bool) {
	if o == nil || IsNil(o.PrefixCount) {
		return nil, false
	}
	return o.PrefixCount, true
}

// HasPrefixCount returns a boolean if a field has been set.
func (o *VLAN) HasPrefixCount() bool {
	if o != nil && !IsNil(o.PrefixCount) {
		return true
	}

	return false
}

// SetPrefixCount gets a reference to the given int64 and assigns it to the PrefixCount field.
func (o *VLAN) SetPrefixCount(v int64) {
	o.PrefixCount = &v
}

func (o VLAN) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VLAN) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display_url"] = o.DisplayUrl
	toSerialize["display"] = o.Display
	if o.Site.IsSet() {
		toSerialize["site"] = o.Site.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	toSerialize["vid"] = o.Vid
	toSerialize["name"] = o.Name
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Role.IsSet() {
		toSerialize["role"] = o.Role.Get()
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	toSerialize["l2vpn_termination"] = o.L2vpnTermination.Get()
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()
	if !IsNil(o.PrefixCount) {
		toSerialize["prefix_count"] = o.PrefixCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VLAN) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display_url",
		"display",
		"vid",
		"name",
		"l2vpn_termination",
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

	varVLAN := _VLAN{}

	err = json.Unmarshal(data, &varVLAN)

	if err != nil {
		return err
	}

	*o = VLAN(varVLAN)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display_url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "site")
		delete(additionalProperties, "group")
		delete(additionalProperties, "vid")
		delete(additionalProperties, "name")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "status")
		delete(additionalProperties, "role")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "l2vpn_termination")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		delete(additionalProperties, "prefix_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVLAN struct {
	value *VLAN
	isSet bool
}

func (v NullableVLAN) Get() *VLAN {
	return v.value
}

func (v *NullableVLAN) Set(val *VLAN) {
	v.value = val
	v.isSet = true
}

func (v NullableVLAN) IsSet() bool {
	return v.isSet
}

func (v *NullableVLAN) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVLAN(val *VLAN) *NullableVLAN {
	return &NullableVLAN{value: val, isSet: true}
}

func (v NullableVLAN) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVLAN) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


