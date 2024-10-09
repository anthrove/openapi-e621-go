# PostSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatorId** | **int32** |  | 
**IsPublic** | **bool** |  | 
**Name** | **string** |  | 
**Shortname** | **string** |  | 
**Description** | **string** |  | 
**PostCount** | **int32** |  | 
**TransferOnDelete** | **bool** |  | 
**PostIds** | **[]int32** |  | 

## Methods

### NewPostSet

`func NewPostSet(id int32, createdAt time.Time, updatedAt time.Time, creatorId int32, isPublic bool, name string, shortname string, description string, postCount int32, transferOnDelete bool, postIds []int32, ) *PostSet`

NewPostSet instantiates a new PostSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSetWithDefaults

`func NewPostSetWithDefaults() *PostSet`

NewPostSetWithDefaults instantiates a new PostSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostSet) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostSet) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostSet) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *PostSet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostSet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostSet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PostSet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostSet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostSet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatorId

`func (o *PostSet) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PostSet) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PostSet) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetIsPublic

`func (o *PostSet) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *PostSet) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *PostSet) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetName

`func (o *PostSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostSet) SetName(v string)`

SetName sets Name field to given value.


### GetShortname

`func (o *PostSet) GetShortname() string`

GetShortname returns the Shortname field if non-nil, zero value otherwise.

### GetShortnameOk

`func (o *PostSet) GetShortnameOk() (*string, bool)`

GetShortnameOk returns a tuple with the Shortname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortname

`func (o *PostSet) SetShortname(v string)`

SetShortname sets Shortname field to given value.


### GetDescription

`func (o *PostSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostSet) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetPostCount

`func (o *PostSet) GetPostCount() int32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *PostSet) GetPostCountOk() (*int32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *PostSet) SetPostCount(v int32)`

SetPostCount sets PostCount field to given value.


### GetTransferOnDelete

`func (o *PostSet) GetTransferOnDelete() bool`

GetTransferOnDelete returns the TransferOnDelete field if non-nil, zero value otherwise.

### GetTransferOnDeleteOk

`func (o *PostSet) GetTransferOnDeleteOk() (*bool, bool)`

GetTransferOnDeleteOk returns a tuple with the TransferOnDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferOnDelete

`func (o *PostSet) SetTransferOnDelete(v bool)`

SetTransferOnDelete sets TransferOnDelete field to given value.


### GetPostIds

`func (o *PostSet) GetPostIds() []int32`

GetPostIds returns the PostIds field if non-nil, zero value otherwise.

### GetPostIdsOk

`func (o *PostSet) GetPostIdsOk() (*[]int32, bool)`

GetPostIdsOk returns a tuple with the PostIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostIds

`func (o *PostSet) SetPostIds(v []int32)`

SetPostIds sets PostIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


