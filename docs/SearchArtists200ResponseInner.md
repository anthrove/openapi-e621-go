# SearchArtists200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**IsActive** | **bool** |  | 
**OtherNames** | **[]string** |  | 
**GroupName** | **string** |  | 
**LinkedUserId** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatorId** | **float32** |  | 
**IsLocked** | **bool** |  | 
**Notes** | **string** |  | 
**Domains** | [**[][]SearchArtists200ResponseInnerAllOfDomainsInnerInner**]([]SearchArtists200ResponseInnerAllOfDomainsInnerInner.md) |  | 
**Urls** | [**[]ArtistURL**](ArtistURL.md) |  | 

## Methods

### NewSearchArtists200ResponseInner

`func NewSearchArtists200ResponseInner(id float32, name string, updatedAt time.Time, isActive bool, otherNames []string, groupName string, linkedUserId float32, createdAt time.Time, creatorId float32, isLocked bool, notes string, domains [][]SearchArtists200ResponseInnerAllOfDomainsInnerInner, urls []ArtistURL, ) *SearchArtists200ResponseInner`

NewSearchArtists200ResponseInner instantiates a new SearchArtists200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchArtists200ResponseInnerWithDefaults

`func NewSearchArtists200ResponseInnerWithDefaults() *SearchArtists200ResponseInner`

NewSearchArtists200ResponseInnerWithDefaults instantiates a new SearchArtists200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchArtists200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchArtists200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchArtists200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *SearchArtists200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchArtists200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchArtists200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *SearchArtists200ResponseInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SearchArtists200ResponseInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SearchArtists200ResponseInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsActive

`func (o *SearchArtists200ResponseInner) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SearchArtists200ResponseInner) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SearchArtists200ResponseInner) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetOtherNames

`func (o *SearchArtists200ResponseInner) GetOtherNames() []string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *SearchArtists200ResponseInner) GetOtherNamesOk() (*[]string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *SearchArtists200ResponseInner) SetOtherNames(v []string)`

SetOtherNames sets OtherNames field to given value.


### GetGroupName

`func (o *SearchArtists200ResponseInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *SearchArtists200ResponseInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *SearchArtists200ResponseInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetLinkedUserId

`func (o *SearchArtists200ResponseInner) GetLinkedUserId() float32`

GetLinkedUserId returns the LinkedUserId field if non-nil, zero value otherwise.

### GetLinkedUserIdOk

`func (o *SearchArtists200ResponseInner) GetLinkedUserIdOk() (*float32, bool)`

GetLinkedUserIdOk returns a tuple with the LinkedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedUserId

`func (o *SearchArtists200ResponseInner) SetLinkedUserId(v float32)`

SetLinkedUserId sets LinkedUserId field to given value.


### GetCreatedAt

`func (o *SearchArtists200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SearchArtists200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SearchArtists200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatorId

`func (o *SearchArtists200ResponseInner) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *SearchArtists200ResponseInner) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *SearchArtists200ResponseInner) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.


### GetIsLocked

`func (o *SearchArtists200ResponseInner) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *SearchArtists200ResponseInner) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *SearchArtists200ResponseInner) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetNotes

`func (o *SearchArtists200ResponseInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SearchArtists200ResponseInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SearchArtists200ResponseInner) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetDomains

`func (o *SearchArtists200ResponseInner) GetDomains() [][]SearchArtists200ResponseInnerAllOfDomainsInnerInner`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *SearchArtists200ResponseInner) GetDomainsOk() (*[][]SearchArtists200ResponseInnerAllOfDomainsInnerInner, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *SearchArtists200ResponseInner) SetDomains(v [][]SearchArtists200ResponseInnerAllOfDomainsInnerInner)`

SetDomains sets Domains field to given value.


### GetUrls

`func (o *SearchArtists200ResponseInner) GetUrls() []ArtistURL`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *SearchArtists200ResponseInner) GetUrlsOk() (*[]ArtistURL, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *SearchArtists200ResponseInner) SetUrls(v []ArtistURL)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


