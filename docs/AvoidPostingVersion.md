# AvoidPostingVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**UpdaterId** | **int32** |  | 
**AvoidPostingId** | **int32** |  | 
**Details** | **string** |  | 
**StaffNotes** | Pointer to **string** | Only visible to Janitor+ | [optional] 
**IsActive** | **bool** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAvoidPostingVersion

`func NewAvoidPostingVersion(id int32, updaterId int32, avoidPostingId int32, details string, isActive bool, updatedAt time.Time, ) *AvoidPostingVersion`

NewAvoidPostingVersion instantiates a new AvoidPostingVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvoidPostingVersionWithDefaults

`func NewAvoidPostingVersionWithDefaults() *AvoidPostingVersion`

NewAvoidPostingVersionWithDefaults instantiates a new AvoidPostingVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AvoidPostingVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvoidPostingVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvoidPostingVersion) SetId(v int32)`

SetId sets Id field to given value.


### GetUpdaterId

`func (o *AvoidPostingVersion) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *AvoidPostingVersion) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *AvoidPostingVersion) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.


### GetAvoidPostingId

`func (o *AvoidPostingVersion) GetAvoidPostingId() int32`

GetAvoidPostingId returns the AvoidPostingId field if non-nil, zero value otherwise.

### GetAvoidPostingIdOk

`func (o *AvoidPostingVersion) GetAvoidPostingIdOk() (*int32, bool)`

GetAvoidPostingIdOk returns a tuple with the AvoidPostingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvoidPostingId

`func (o *AvoidPostingVersion) SetAvoidPostingId(v int32)`

SetAvoidPostingId sets AvoidPostingId field to given value.


### GetDetails

`func (o *AvoidPostingVersion) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AvoidPostingVersion) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AvoidPostingVersion) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetStaffNotes

`func (o *AvoidPostingVersion) GetStaffNotes() string`

GetStaffNotes returns the StaffNotes field if non-nil, zero value otherwise.

### GetStaffNotesOk

`func (o *AvoidPostingVersion) GetStaffNotesOk() (*string, bool)`

GetStaffNotesOk returns a tuple with the StaffNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaffNotes

`func (o *AvoidPostingVersion) SetStaffNotes(v string)`

SetStaffNotes sets StaffNotes field to given value.

### HasStaffNotes

`func (o *AvoidPostingVersion) HasStaffNotes() bool`

HasStaffNotes returns a boolean if a field has been set.

### GetIsActive

`func (o *AvoidPostingVersion) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AvoidPostingVersion) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AvoidPostingVersion) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetUpdatedAt

`func (o *AvoidPostingVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AvoidPostingVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AvoidPostingVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


