# Ticket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatorId** | **int32** |  | 
**Reason** | **string** |  | 
**DispId** | **float32** |  | 
**Qtype** | [**TicketTypes**](TicketTypes.md) |  | 
**Status** | [**TicketStatuses**](TicketStatuses.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Response** | **string** |  | 
**HandlerId** | **NullableFloat32** |  | 
**ClaimantId** | Pointer to **NullableFloat32** | Only visible to Moderator+. | [optional] 
**ReportReason** | **NullableString** |  | 
**AccusedId** | **NullableFloat32** |  | 

## Methods

### NewTicket

`func NewTicket(id int32, creatorId int32, reason string, dispId float32, qtype TicketTypes, status TicketStatuses, createdAt time.Time, updatedAt time.Time, response string, handlerId NullableFloat32, reportReason NullableString, accusedId NullableFloat32, ) *Ticket`

NewTicket instantiates a new Ticket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketWithDefaults

`func NewTicketWithDefaults() *Ticket`

NewTicketWithDefaults instantiates a new Ticket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ticket) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ticket) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ticket) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *Ticket) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Ticket) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Ticket) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetReason

`func (o *Ticket) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Ticket) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Ticket) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDispId

`func (o *Ticket) GetDispId() float32`

GetDispId returns the DispId field if non-nil, zero value otherwise.

### GetDispIdOk

`func (o *Ticket) GetDispIdOk() (*float32, bool)`

GetDispIdOk returns a tuple with the DispId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispId

`func (o *Ticket) SetDispId(v float32)`

SetDispId sets DispId field to given value.


### GetQtype

`func (o *Ticket) GetQtype() TicketTypes`

GetQtype returns the Qtype field if non-nil, zero value otherwise.

### GetQtypeOk

`func (o *Ticket) GetQtypeOk() (*TicketTypes, bool)`

GetQtypeOk returns a tuple with the Qtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQtype

`func (o *Ticket) SetQtype(v TicketTypes)`

SetQtype sets Qtype field to given value.


### GetStatus

`func (o *Ticket) GetStatus() TicketStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ticket) GetStatusOk() (*TicketStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ticket) SetStatus(v TicketStatuses)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Ticket) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Ticket) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Ticket) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Ticket) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Ticket) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Ticket) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetResponse

`func (o *Ticket) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Ticket) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Ticket) SetResponse(v string)`

SetResponse sets Response field to given value.


### GetHandlerId

`func (o *Ticket) GetHandlerId() float32`

GetHandlerId returns the HandlerId field if non-nil, zero value otherwise.

### GetHandlerIdOk

`func (o *Ticket) GetHandlerIdOk() (*float32, bool)`

GetHandlerIdOk returns a tuple with the HandlerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerId

`func (o *Ticket) SetHandlerId(v float32)`

SetHandlerId sets HandlerId field to given value.


### SetHandlerIdNil

`func (o *Ticket) SetHandlerIdNil(b bool)`

 SetHandlerIdNil sets the value for HandlerId to be an explicit nil

### UnsetHandlerId
`func (o *Ticket) UnsetHandlerId()`

UnsetHandlerId ensures that no value is present for HandlerId, not even an explicit nil
### GetClaimantId

`func (o *Ticket) GetClaimantId() float32`

GetClaimantId returns the ClaimantId field if non-nil, zero value otherwise.

### GetClaimantIdOk

`func (o *Ticket) GetClaimantIdOk() (*float32, bool)`

GetClaimantIdOk returns a tuple with the ClaimantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimantId

`func (o *Ticket) SetClaimantId(v float32)`

SetClaimantId sets ClaimantId field to given value.

### HasClaimantId

`func (o *Ticket) HasClaimantId() bool`

HasClaimantId returns a boolean if a field has been set.

### SetClaimantIdNil

`func (o *Ticket) SetClaimantIdNil(b bool)`

 SetClaimantIdNil sets the value for ClaimantId to be an explicit nil

### UnsetClaimantId
`func (o *Ticket) UnsetClaimantId()`

UnsetClaimantId ensures that no value is present for ClaimantId, not even an explicit nil
### GetReportReason

`func (o *Ticket) GetReportReason() string`

GetReportReason returns the ReportReason field if non-nil, zero value otherwise.

### GetReportReasonOk

`func (o *Ticket) GetReportReasonOk() (*string, bool)`

GetReportReasonOk returns a tuple with the ReportReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportReason

`func (o *Ticket) SetReportReason(v string)`

SetReportReason sets ReportReason field to given value.


### SetReportReasonNil

`func (o *Ticket) SetReportReasonNil(b bool)`

 SetReportReasonNil sets the value for ReportReason to be an explicit nil

### UnsetReportReason
`func (o *Ticket) UnsetReportReason()`

UnsetReportReason ensures that no value is present for ReportReason, not even an explicit nil
### GetAccusedId

`func (o *Ticket) GetAccusedId() float32`

GetAccusedId returns the AccusedId field if non-nil, zero value otherwise.

### GetAccusedIdOk

`func (o *Ticket) GetAccusedIdOk() (*float32, bool)`

GetAccusedIdOk returns a tuple with the AccusedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccusedId

`func (o *Ticket) SetAccusedId(v float32)`

SetAccusedId sets AccusedId field to given value.


### SetAccusedIdNil

`func (o *Ticket) SetAccusedIdNil(b bool)`

 SetAccusedIdNil sets the value for AccusedId to be an explicit nil

### UnsetAccusedId
`func (o *Ticket) UnsetAccusedId()`

UnsetAccusedId ensures that no value is present for AccusedId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


