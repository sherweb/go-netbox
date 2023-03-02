// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewExtrasJournalEntriesListParams creates a new ExtrasJournalEntriesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJournalEntriesListParams() *ExtrasJournalEntriesListParams {
	return &ExtrasJournalEntriesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJournalEntriesListParamsWithTimeout creates a new ExtrasJournalEntriesListParams object
// with the ability to set a timeout on a request.
func NewExtrasJournalEntriesListParamsWithTimeout(timeout time.Duration) *ExtrasJournalEntriesListParams {
	return &ExtrasJournalEntriesListParams{
		timeout: timeout,
	}
}

// NewExtrasJournalEntriesListParamsWithContext creates a new ExtrasJournalEntriesListParams object
// with the ability to set a context for a request.
func NewExtrasJournalEntriesListParamsWithContext(ctx context.Context) *ExtrasJournalEntriesListParams {
	return &ExtrasJournalEntriesListParams{
		Context: ctx,
	}
}

// NewExtrasJournalEntriesListParamsWithHTTPClient creates a new ExtrasJournalEntriesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJournalEntriesListParamsWithHTTPClient(client *http.Client) *ExtrasJournalEntriesListParams {
	return &ExtrasJournalEntriesListParams{
		HTTPClient: client,
	}
}

/*
ExtrasJournalEntriesListParams contains all the parameters to send to the API endpoint

	for the extras journal entries list operation.

	Typically these are written to a http.Request.
*/
type ExtrasJournalEntriesListParams struct {

	// AssignedObjectID.
	AssignedObjectID *string

	// AssignedObjectIDGt.
	AssignedObjectIDGt *string

	// AssignedObjectIDGte.
	AssignedObjectIDGte *string

	// AssignedObjectIDLt.
	AssignedObjectIDLt *string

	// AssignedObjectIDLte.
	AssignedObjectIDLte *string

	// AssignedObjectIDn.
	AssignedObjectIDn *string

	// AssignedObjectType.
	AssignedObjectType *string

	// AssignedObjectTypen.
	AssignedObjectTypen *string

	// AssignedObjectTypeID.
	AssignedObjectTypeID *string

	// AssignedObjectTypeIDn.
	AssignedObjectTypeIDn *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// CreatedBy.
	CreatedBy *string

	// CreatedByn.
	CreatedByn *string

	// CreatedByID.
	CreatedByID *string

	// CreatedByIDn.
	CreatedByIDn *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// Kind.
	Kind *string

	// Kindn.
	Kindn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras journal entries list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJournalEntriesListParams) WithDefaults() *ExtrasJournalEntriesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras journal entries list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJournalEntriesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithTimeout(timeout time.Duration) *ExtrasJournalEntriesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithContext(ctx context.Context) *ExtrasJournalEntriesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithHTTPClient(client *http.Client) *ExtrasJournalEntriesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssignedObjectID adds the assignedObjectID to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectID(assignedObjectID *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectID(assignedObjectID)
	return o
}

// SetAssignedObjectID adds the assignedObjectId to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectID(assignedObjectID *string) {
	o.AssignedObjectID = assignedObjectID
}

// WithAssignedObjectIDGt adds the assignedObjectIDGt to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectIDGt(assignedObjectIDGt *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectIDGt(assignedObjectIDGt)
	return o
}

// SetAssignedObjectIDGt adds the assignedObjectIdGt to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectIDGt(assignedObjectIDGt *string) {
	o.AssignedObjectIDGt = assignedObjectIDGt
}

// WithAssignedObjectIDGte adds the assignedObjectIDGte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectIDGte(assignedObjectIDGte *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectIDGte(assignedObjectIDGte)
	return o
}

// SetAssignedObjectIDGte adds the assignedObjectIdGte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectIDGte(assignedObjectIDGte *string) {
	o.AssignedObjectIDGte = assignedObjectIDGte
}

// WithAssignedObjectIDLt adds the assignedObjectIDLt to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectIDLt(assignedObjectIDLt *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectIDLt(assignedObjectIDLt)
	return o
}

// SetAssignedObjectIDLt adds the assignedObjectIdLt to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectIDLt(assignedObjectIDLt *string) {
	o.AssignedObjectIDLt = assignedObjectIDLt
}

// WithAssignedObjectIDLte adds the assignedObjectIDLte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectIDLte(assignedObjectIDLte *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectIDLte(assignedObjectIDLte)
	return o
}

// SetAssignedObjectIDLte adds the assignedObjectIdLte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectIDLte(assignedObjectIDLte *string) {
	o.AssignedObjectIDLte = assignedObjectIDLte
}

// WithAssignedObjectIDn adds the assignedObjectIDn to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectIDn(assignedObjectIDn *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectIDn(assignedObjectIDn)
	return o
}

// SetAssignedObjectIDn adds the assignedObjectIdN to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectIDn(assignedObjectIDn *string) {
	o.AssignedObjectIDn = assignedObjectIDn
}

// WithAssignedObjectType adds the assignedObjectType to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectType(assignedObjectType *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectType(assignedObjectType)
	return o
}

// SetAssignedObjectType adds the assignedObjectType to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectType(assignedObjectType *string) {
	o.AssignedObjectType = assignedObjectType
}

// WithAssignedObjectTypen adds the assignedObjectTypen to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectTypen(assignedObjectTypen *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectTypen(assignedObjectTypen)
	return o
}

// SetAssignedObjectTypen adds the assignedObjectTypeN to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectTypen(assignedObjectTypen *string) {
	o.AssignedObjectTypen = assignedObjectTypen
}

// WithAssignedObjectTypeID adds the assignedObjectTypeID to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectTypeID(assignedObjectTypeID *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectTypeID(assignedObjectTypeID)
	return o
}

// SetAssignedObjectTypeID adds the assignedObjectTypeId to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectTypeID(assignedObjectTypeID *string) {
	o.AssignedObjectTypeID = assignedObjectTypeID
}

// WithAssignedObjectTypeIDn adds the assignedObjectTypeIDn to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithAssignedObjectTypeIDn(assignedObjectTypeIDn *string) *ExtrasJournalEntriesListParams {
	o.SetAssignedObjectTypeIDn(assignedObjectTypeIDn)
	return o
}

// SetAssignedObjectTypeIDn adds the assignedObjectTypeIdN to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetAssignedObjectTypeIDn(assignedObjectTypeIDn *string) {
	o.AssignedObjectTypeIDn = assignedObjectTypeIDn
}

// WithCreated adds the created to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithCreated(created *string) *ExtrasJournalEntriesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithCreatedGte(createdGte *string) *ExtrasJournalEntriesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithCreatedLte(createdLte *string) *ExtrasJournalEntriesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithCreatedBy adds the createdBy to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithCreatedBy(createdBy *string) *ExtrasJournalEntriesListParams {
	o.SetCreatedBy(createdBy)
	return o
}

// SetCreatedBy adds the createdBy to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetCreatedBy(createdBy *string) {
	o.CreatedBy = createdBy
}

// WithCreatedByn adds the createdByn to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithCreatedByn(createdByn *string) *ExtrasJournalEntriesListParams {
	o.SetCreatedByn(createdByn)
	return o
}

// SetCreatedByn adds the createdByN to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetCreatedByn(createdByn *string) {
	o.CreatedByn = createdByn
}

// WithCreatedByID adds the createdByID to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithCreatedByID(createdByID *string) *ExtrasJournalEntriesListParams {
	o.SetCreatedByID(createdByID)
	return o
}

// SetCreatedByID adds the createdById to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetCreatedByID(createdByID *string) {
	o.CreatedByID = createdByID
}

// WithCreatedByIDn adds the createdByIDn to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithCreatedByIDn(createdByIDn *string) *ExtrasJournalEntriesListParams {
	o.SetCreatedByIDn(createdByIDn)
	return o
}

// SetCreatedByIDn adds the createdByIdN to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetCreatedByIDn(createdByIDn *string) {
	o.CreatedByIDn = createdByIDn
}

// WithID adds the id to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithID(id *string) *ExtrasJournalEntriesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithIDGt(iDGt *string) *ExtrasJournalEntriesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithIDGte(iDGte *string) *ExtrasJournalEntriesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithIDLt(iDLt *string) *ExtrasJournalEntriesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithIDLte(iDLte *string) *ExtrasJournalEntriesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithIDn(iDn *string) *ExtrasJournalEntriesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithKind adds the kind to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithKind(kind *string) *ExtrasJournalEntriesListParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithKindn adds the kindn to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithKindn(kindn *string) *ExtrasJournalEntriesListParams {
	o.SetKindn(kindn)
	return o
}

// SetKindn adds the kindN to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetKindn(kindn *string) {
	o.Kindn = kindn
}

// WithLastUpdated adds the lastUpdated to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithLastUpdated(lastUpdated *string) *ExtrasJournalEntriesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *ExtrasJournalEntriesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *ExtrasJournalEntriesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithLimit(limit *int64) *ExtrasJournalEntriesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithOffset(offset *int64) *ExtrasJournalEntriesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) WithQ(q *string) *ExtrasJournalEntriesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras journal entries list params
func (o *ExtrasJournalEntriesListParams) SetQ(q *string) {
	o.Q = q
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJournalEntriesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssignedObjectID != nil {

		// query param assigned_object_id
		var qrAssignedObjectID string

		if o.AssignedObjectID != nil {
			qrAssignedObjectID = *o.AssignedObjectID
		}
		qAssignedObjectID := qrAssignedObjectID
		if qAssignedObjectID != "" {

			if err := r.SetQueryParam("assigned_object_id", qAssignedObjectID); err != nil {
				return err
			}
		}
	}

	if o.AssignedObjectIDGt != nil {

		// query param assigned_object_id__gt
		var qrAssignedObjectIDGt string

		if o.AssignedObjectIDGt != nil {
			qrAssignedObjectIDGt = *o.AssignedObjectIDGt
		}
		qAssignedObjectIDGt := qrAssignedObjectIDGt
		if qAssignedObjectIDGt != "" {

			if err := r.SetQueryParam("assigned_object_id__gt", qAssignedObjectIDGt); err != nil {
				return err
			}
		}
	}

	if o.AssignedObjectIDGte != nil {

		// query param assigned_object_id__gte
		var qrAssignedObjectIDGte string

		if o.AssignedObjectIDGte != nil {
			qrAssignedObjectIDGte = *o.AssignedObjectIDGte
		}
		qAssignedObjectIDGte := qrAssignedObjectIDGte
		if qAssignedObjectIDGte != "" {

			if err := r.SetQueryParam("assigned_object_id__gte", qAssignedObjectIDGte); err != nil {
				return err
			}
		}
	}

	if o.AssignedObjectIDLt != nil {

		// query param assigned_object_id__lt
		var qrAssignedObjectIDLt string

		if o.AssignedObjectIDLt != nil {
			qrAssignedObjectIDLt = *o.AssignedObjectIDLt
		}
		qAssignedObjectIDLt := qrAssignedObjectIDLt
		if qAssignedObjectIDLt != "" {

			if err := r.SetQueryParam("assigned_object_id__lt", qAssignedObjectIDLt); err != nil {
				return err
			}
		}
	}

	if o.AssignedObjectIDLte != nil {

		// query param assigned_object_id__lte
		var qrAssignedObjectIDLte string

		if o.AssignedObjectIDLte != nil {
			qrAssignedObjectIDLte = *o.AssignedObjectIDLte
		}
		qAssignedObjectIDLte := qrAssignedObjectIDLte
		if qAssignedObjectIDLte != "" {

			if err := r.SetQueryParam("assigned_object_id__lte", qAssignedObjectIDLte); err != nil {
				return err
			}
		}
	}

	if o.AssignedObjectIDn != nil {

		// query param assigned_object_id__n
		var qrAssignedObjectIDn string

		if o.AssignedObjectIDn != nil {
			qrAssignedObjectIDn = *o.AssignedObjectIDn
		}
		qAssignedObjectIDn := qrAssignedObjectIDn
		if qAssignedObjectIDn != "" {

			if err := r.SetQueryParam("assigned_object_id__n", qAssignedObjectIDn); err != nil {
				return err
			}
		}
	}

	if o.AssignedObjectType != nil {

		// query param assigned_object_type
		var qrAssignedObjectType string

		if o.AssignedObjectType != nil {
			qrAssignedObjectType = *o.AssignedObjectType
		}
		qAssignedObjectType := qrAssignedObjectType
		if qAssignedObjectType != "" {

			if err := r.SetQueryParam("assigned_object_type", qAssignedObjectType); err != nil {
				return err
			}
		}
	}

	if o.AssignedObjectTypen != nil {

		// query param assigned_object_type__n
		var qrAssignedObjectTypen string

		if o.AssignedObjectTypen != nil {
			qrAssignedObjectTypen = *o.AssignedObjectTypen
		}
		qAssignedObjectTypen := qrAssignedObjectTypen
		if qAssignedObjectTypen != "" {

			if err := r.SetQueryParam("assigned_object_type__n", qAssignedObjectTypen); err != nil {
				return err
			}
		}
	}

	if o.AssignedObjectTypeID != nil {

		// query param assigned_object_type_id
		var qrAssignedObjectTypeID string

		if o.AssignedObjectTypeID != nil {
			qrAssignedObjectTypeID = *o.AssignedObjectTypeID
		}
		qAssignedObjectTypeID := qrAssignedObjectTypeID
		if qAssignedObjectTypeID != "" {

			if err := r.SetQueryParam("assigned_object_type_id", qAssignedObjectTypeID); err != nil {
				return err
			}
		}
	}

	if o.AssignedObjectTypeIDn != nil {

		// query param assigned_object_type_id__n
		var qrAssignedObjectTypeIDn string

		if o.AssignedObjectTypeIDn != nil {
			qrAssignedObjectTypeIDn = *o.AssignedObjectTypeIDn
		}
		qAssignedObjectTypeIDn := qrAssignedObjectTypeIDn
		if qAssignedObjectTypeIDn != "" {

			if err := r.SetQueryParam("assigned_object_type_id__n", qAssignedObjectTypeIDn); err != nil {
				return err
			}
		}
	}

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.CreatedBy != nil {

		// query param created_by
		var qrCreatedBy string

		if o.CreatedBy != nil {
			qrCreatedBy = *o.CreatedBy
		}
		qCreatedBy := qrCreatedBy
		if qCreatedBy != "" {

			if err := r.SetQueryParam("created_by", qCreatedBy); err != nil {
				return err
			}
		}
	}

	if o.CreatedByn != nil {

		// query param created_by__n
		var qrCreatedByn string

		if o.CreatedByn != nil {
			qrCreatedByn = *o.CreatedByn
		}
		qCreatedByn := qrCreatedByn
		if qCreatedByn != "" {

			if err := r.SetQueryParam("created_by__n", qCreatedByn); err != nil {
				return err
			}
		}
	}

	if o.CreatedByID != nil {

		// query param created_by_id
		var qrCreatedByID string

		if o.CreatedByID != nil {
			qrCreatedByID = *o.CreatedByID
		}
		qCreatedByID := qrCreatedByID
		if qCreatedByID != "" {

			if err := r.SetQueryParam("created_by_id", qCreatedByID); err != nil {
				return err
			}
		}
	}

	if o.CreatedByIDn != nil {

		// query param created_by_id__n
		var qrCreatedByIDn string

		if o.CreatedByIDn != nil {
			qrCreatedByIDn = *o.CreatedByIDn
		}
		qCreatedByIDn := qrCreatedByIDn
		if qCreatedByIDn != "" {

			if err := r.SetQueryParam("created_by_id__n", qCreatedByIDn); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.Kind != nil {

		// query param kind
		var qrKind string

		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {

			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}
	}

	if o.Kindn != nil {

		// query param kind__n
		var qrKindn string

		if o.Kindn != nil {
			qrKindn = *o.Kindn
		}
		qKindn := qrKindn
		if qKindn != "" {

			if err := r.SetQueryParam("kind__n", qKindn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
