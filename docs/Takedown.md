# Takedown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Status** | **string** |  | 
**ApproverId** | **NullableFloat32** |  | 
**ReasonHidden** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**PostCount** | **int32** |  | 

## Methods

### NewTakedown

`func NewTakedown(id int32, status string, approverId NullableFloat32, reasonHidden bool, createdAt time.Time, updatedAt time.Time, postCount int32, ) *Takedown`

NewTakedown instantiates a new Takedown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTakedownWithDefaults

`func NewTakedownWithDefaults() *Takedown`

NewTakedownWithDefaults instantiates a new Takedown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Takedown) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Takedown) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Takedown) SetId(v int32)`

SetId sets Id field to given value.


### GetStatus

`func (o *Takedown) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Takedown) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Takedown) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetApproverId

`func (o *Takedown) GetApproverId() float32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *Takedown) GetApproverIdOk() (*float32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *Takedown) SetApproverId(v float32)`

SetApproverId sets ApproverId field to given value.


### SetApproverIdNil

`func (o *Takedown) SetApproverIdNil(b bool)`

 SetApproverIdNil sets the value for ApproverId to be an explicit nil

### UnsetApproverId
`func (o *Takedown) UnsetApproverId()`

UnsetApproverId ensures that no value is present for ApproverId, not even an explicit nil
### GetReasonHidden

`func (o *Takedown) GetReasonHidden() bool`

GetReasonHidden returns the ReasonHidden field if non-nil, zero value otherwise.

### GetReasonHiddenOk

`func (o *Takedown) GetReasonHiddenOk() (*bool, bool)`

GetReasonHiddenOk returns a tuple with the ReasonHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonHidden

`func (o *Takedown) SetReasonHidden(v bool)`

SetReasonHidden sets ReasonHidden field to given value.


### GetCreatedAt

`func (o *Takedown) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Takedown) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Takedown) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Takedown) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Takedown) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Takedown) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPostCount

`func (o *Takedown) GetPostCount() int32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *Takedown) GetPostCountOk() (*int32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *Takedown) SetPostCount(v int32)`

SetPostCount sets PostCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


