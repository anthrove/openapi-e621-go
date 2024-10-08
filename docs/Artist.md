# Artist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**IsActive** | **bool** |  | 
**OtherNames** | **[]string** |  | 
**GroupName** | **string** |  | 
**LinkedUserId** | **NullableFloat32** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatorId** | **float32** |  | 
**IsLocked** | **bool** |  | 
**Notes** | **NullableString** |  | 

## Methods

### NewArtist

`func NewArtist(id float32, name string, updatedAt time.Time, isActive bool, otherNames []string, groupName string, linkedUserId NullableFloat32, createdAt time.Time, creatorId float32, isLocked bool, notes NullableString, ) *Artist`

NewArtist instantiates a new Artist object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtistWithDefaults

`func NewArtistWithDefaults() *Artist`

NewArtistWithDefaults instantiates a new Artist object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Artist) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Artist) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Artist) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *Artist) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Artist) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Artist) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *Artist) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Artist) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Artist) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsActive

`func (o *Artist) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Artist) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Artist) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetOtherNames

`func (o *Artist) GetOtherNames() []string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *Artist) GetOtherNamesOk() (*[]string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *Artist) SetOtherNames(v []string)`

SetOtherNames sets OtherNames field to given value.


### GetGroupName

`func (o *Artist) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *Artist) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *Artist) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetLinkedUserId

`func (o *Artist) GetLinkedUserId() float32`

GetLinkedUserId returns the LinkedUserId field if non-nil, zero value otherwise.

### GetLinkedUserIdOk

`func (o *Artist) GetLinkedUserIdOk() (*float32, bool)`

GetLinkedUserIdOk returns a tuple with the LinkedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedUserId

`func (o *Artist) SetLinkedUserId(v float32)`

SetLinkedUserId sets LinkedUserId field to given value.


### SetLinkedUserIdNil

`func (o *Artist) SetLinkedUserIdNil(b bool)`

 SetLinkedUserIdNil sets the value for LinkedUserId to be an explicit nil

### UnsetLinkedUserId
`func (o *Artist) UnsetLinkedUserId()`

UnsetLinkedUserId ensures that no value is present for LinkedUserId, not even an explicit nil
### GetCreatedAt

`func (o *Artist) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Artist) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Artist) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatorId

`func (o *Artist) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *Artist) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *Artist) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.


### GetIsLocked

`func (o *Artist) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *Artist) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *Artist) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetNotes

`func (o *Artist) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Artist) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Artist) SetNotes(v string)`

SetNotes sets Notes field to given value.


### SetNotesNil

`func (o *Artist) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *Artist) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


