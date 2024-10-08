# UserNameChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**ApproverId** | **float32** |  | 
**UserId** | **float32** |  | 
**OriginalName** | **string** |  | 
**DesiredName** | **string** |  | 
**ChangeReason** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Status** | **string** |  | 

## Methods

### NewUserNameChangeRequest

`func NewUserNameChangeRequest(id float32, approverId float32, userId float32, originalName string, desiredName string, createdAt time.Time, updatedAt time.Time, status string, ) *UserNameChangeRequest`

NewUserNameChangeRequest instantiates a new UserNameChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNameChangeRequestWithDefaults

`func NewUserNameChangeRequestWithDefaults() *UserNameChangeRequest`

NewUserNameChangeRequestWithDefaults instantiates a new UserNameChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserNameChangeRequest) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserNameChangeRequest) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserNameChangeRequest) SetId(v float32)`

SetId sets Id field to given value.


### GetApproverId

`func (o *UserNameChangeRequest) GetApproverId() float32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *UserNameChangeRequest) GetApproverIdOk() (*float32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *UserNameChangeRequest) SetApproverId(v float32)`

SetApproverId sets ApproverId field to given value.


### GetUserId

`func (o *UserNameChangeRequest) GetUserId() float32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserNameChangeRequest) GetUserIdOk() (*float32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserNameChangeRequest) SetUserId(v float32)`

SetUserId sets UserId field to given value.


### GetOriginalName

`func (o *UserNameChangeRequest) GetOriginalName() string`

GetOriginalName returns the OriginalName field if non-nil, zero value otherwise.

### GetOriginalNameOk

`func (o *UserNameChangeRequest) GetOriginalNameOk() (*string, bool)`

GetOriginalNameOk returns a tuple with the OriginalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalName

`func (o *UserNameChangeRequest) SetOriginalName(v string)`

SetOriginalName sets OriginalName field to given value.


### GetDesiredName

`func (o *UserNameChangeRequest) GetDesiredName() string`

GetDesiredName returns the DesiredName field if non-nil, zero value otherwise.

### GetDesiredNameOk

`func (o *UserNameChangeRequest) GetDesiredNameOk() (*string, bool)`

GetDesiredNameOk returns a tuple with the DesiredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredName

`func (o *UserNameChangeRequest) SetDesiredName(v string)`

SetDesiredName sets DesiredName field to given value.


### GetChangeReason

`func (o *UserNameChangeRequest) GetChangeReason() string`

GetChangeReason returns the ChangeReason field if non-nil, zero value otherwise.

### GetChangeReasonOk

`func (o *UserNameChangeRequest) GetChangeReasonOk() (*string, bool)`

GetChangeReasonOk returns a tuple with the ChangeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeReason

`func (o *UserNameChangeRequest) SetChangeReason(v string)`

SetChangeReason sets ChangeReason field to given value.

### HasChangeReason

`func (o *UserNameChangeRequest) HasChangeReason() bool`

HasChangeReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserNameChangeRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserNameChangeRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserNameChangeRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserNameChangeRequest) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserNameChangeRequest) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserNameChangeRequest) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetStatus

`func (o *UserNameChangeRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserNameChangeRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserNameChangeRequest) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


