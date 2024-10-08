# WikiPageVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Title** | **string** |  | 
**Body** | **string** |  | 
**UpdaterId** | **float32** |  | 
**WikiPageId** | **float32** |  | 
**IsLocked** | **bool** |  | 
**OtherNames** | **[]string** |  | 
**IsDeleted** | **bool** |  | 
**Reason** | **NullableString** |  | 
**Parent** | **NullableString** |  | 

## Methods

### NewWikiPageVersion

`func NewWikiPageVersion(id float32, createdAt time.Time, updatedAt time.Time, title string, body string, updaterId float32, wikiPageId float32, isLocked bool, otherNames []string, isDeleted bool, reason NullableString, parent NullableString, ) *WikiPageVersion`

NewWikiPageVersion instantiates a new WikiPageVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWikiPageVersionWithDefaults

`func NewWikiPageVersionWithDefaults() *WikiPageVersion`

NewWikiPageVersionWithDefaults instantiates a new WikiPageVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WikiPageVersion) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WikiPageVersion) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WikiPageVersion) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *WikiPageVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WikiPageVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WikiPageVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *WikiPageVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WikiPageVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WikiPageVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetTitle

`func (o *WikiPageVersion) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WikiPageVersion) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WikiPageVersion) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBody

`func (o *WikiPageVersion) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WikiPageVersion) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WikiPageVersion) SetBody(v string)`

SetBody sets Body field to given value.


### GetUpdaterId

`func (o *WikiPageVersion) GetUpdaterId() float32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *WikiPageVersion) GetUpdaterIdOk() (*float32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *WikiPageVersion) SetUpdaterId(v float32)`

SetUpdaterId sets UpdaterId field to given value.


### GetWikiPageId

`func (o *WikiPageVersion) GetWikiPageId() float32`

GetWikiPageId returns the WikiPageId field if non-nil, zero value otherwise.

### GetWikiPageIdOk

`func (o *WikiPageVersion) GetWikiPageIdOk() (*float32, bool)`

GetWikiPageIdOk returns a tuple with the WikiPageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiPageId

`func (o *WikiPageVersion) SetWikiPageId(v float32)`

SetWikiPageId sets WikiPageId field to given value.


### GetIsLocked

`func (o *WikiPageVersion) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *WikiPageVersion) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *WikiPageVersion) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetOtherNames

`func (o *WikiPageVersion) GetOtherNames() []string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *WikiPageVersion) GetOtherNamesOk() (*[]string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *WikiPageVersion) SetOtherNames(v []string)`

SetOtherNames sets OtherNames field to given value.


### GetIsDeleted

`func (o *WikiPageVersion) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *WikiPageVersion) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *WikiPageVersion) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetReason

`func (o *WikiPageVersion) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WikiPageVersion) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WikiPageVersion) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *WikiPageVersion) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *WikiPageVersion) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetParent

`func (o *WikiPageVersion) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WikiPageVersion) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WikiPageVersion) SetParent(v string)`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *WikiPageVersion) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *WikiPageVersion) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


