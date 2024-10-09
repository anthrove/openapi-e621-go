# AvoidPosting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatorId** | **int32** |  | 
**UpdaterId** | **int32** |  | 
**ArtistId** | **int32** |  | 
**StaffNotes** | Pointer to **string** | Only visible to Janitor+ | [optional] 
**Details** | **string** |  | 
**IsActive** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAvoidPosting

`func NewAvoidPosting(id int32, creatorId int32, updaterId int32, artistId int32, details string, isActive bool, createdAt time.Time, updatedAt time.Time, ) *AvoidPosting`

NewAvoidPosting instantiates a new AvoidPosting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvoidPostingWithDefaults

`func NewAvoidPostingWithDefaults() *AvoidPosting`

NewAvoidPostingWithDefaults instantiates a new AvoidPosting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AvoidPosting) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AvoidPosting) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AvoidPosting) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatorId

`func (o *AvoidPosting) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *AvoidPosting) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *AvoidPosting) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetUpdaterId

`func (o *AvoidPosting) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *AvoidPosting) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *AvoidPosting) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.


### GetArtistId

`func (o *AvoidPosting) GetArtistId() int32`

GetArtistId returns the ArtistId field if non-nil, zero value otherwise.

### GetArtistIdOk

`func (o *AvoidPosting) GetArtistIdOk() (*int32, bool)`

GetArtistIdOk returns a tuple with the ArtistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistId

`func (o *AvoidPosting) SetArtistId(v int32)`

SetArtistId sets ArtistId field to given value.


### GetStaffNotes

`func (o *AvoidPosting) GetStaffNotes() string`

GetStaffNotes returns the StaffNotes field if non-nil, zero value otherwise.

### GetStaffNotesOk

`func (o *AvoidPosting) GetStaffNotesOk() (*string, bool)`

GetStaffNotesOk returns a tuple with the StaffNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaffNotes

`func (o *AvoidPosting) SetStaffNotes(v string)`

SetStaffNotes sets StaffNotes field to given value.

### HasStaffNotes

`func (o *AvoidPosting) HasStaffNotes() bool`

HasStaffNotes returns a boolean if a field has been set.

### GetDetails

`func (o *AvoidPosting) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AvoidPosting) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AvoidPosting) SetDetails(v string)`

SetDetails sets Details field to given value.


### GetIsActive

`func (o *AvoidPosting) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AvoidPosting) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AvoidPosting) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetCreatedAt

`func (o *AvoidPosting) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AvoidPosting) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AvoidPosting) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AvoidPosting) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AvoidPosting) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AvoidPosting) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


