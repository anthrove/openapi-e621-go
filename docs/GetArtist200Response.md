# GetArtist200Response

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

### NewGetArtist200Response

`func NewGetArtist200Response(id float32, name string, updatedAt time.Time, isActive bool, otherNames []string, groupName string, linkedUserId float32, createdAt time.Time, creatorId float32, isLocked bool, notes string, domains [][]SearchArtists200ResponseInnerAllOfDomainsInnerInner, urls []ArtistURL, ) *GetArtist200Response`

NewGetArtist200Response instantiates a new GetArtist200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetArtist200ResponseWithDefaults

`func NewGetArtist200ResponseWithDefaults() *GetArtist200Response`

NewGetArtist200ResponseWithDefaults instantiates a new GetArtist200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetArtist200Response) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetArtist200Response) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetArtist200Response) SetId(v float32)`

SetId sets Id field to given value.


### GetName

`func (o *GetArtist200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetArtist200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetArtist200Response) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *GetArtist200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetArtist200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetArtist200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetIsActive

`func (o *GetArtist200Response) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GetArtist200Response) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GetArtist200Response) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetOtherNames

`func (o *GetArtist200Response) GetOtherNames() []string`

GetOtherNames returns the OtherNames field if non-nil, zero value otherwise.

### GetOtherNamesOk

`func (o *GetArtist200Response) GetOtherNamesOk() (*[]string, bool)`

GetOtherNamesOk returns a tuple with the OtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherNames

`func (o *GetArtist200Response) SetOtherNames(v []string)`

SetOtherNames sets OtherNames field to given value.


### GetGroupName

`func (o *GetArtist200Response) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GetArtist200Response) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GetArtist200Response) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetLinkedUserId

`func (o *GetArtist200Response) GetLinkedUserId() float32`

GetLinkedUserId returns the LinkedUserId field if non-nil, zero value otherwise.

### GetLinkedUserIdOk

`func (o *GetArtist200Response) GetLinkedUserIdOk() (*float32, bool)`

GetLinkedUserIdOk returns a tuple with the LinkedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedUserId

`func (o *GetArtist200Response) SetLinkedUserId(v float32)`

SetLinkedUserId sets LinkedUserId field to given value.


### GetCreatedAt

`func (o *GetArtist200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetArtist200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetArtist200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatorId

`func (o *GetArtist200Response) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *GetArtist200Response) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *GetArtist200Response) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.


### GetIsLocked

`func (o *GetArtist200Response) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *GetArtist200Response) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *GetArtist200Response) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.


### GetNotes

`func (o *GetArtist200Response) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetArtist200Response) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetArtist200Response) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetDomains

`func (o *GetArtist200Response) GetDomains() [][]SearchArtists200ResponseInnerAllOfDomainsInnerInner`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *GetArtist200Response) GetDomainsOk() (*[][]SearchArtists200ResponseInnerAllOfDomainsInnerInner, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *GetArtist200Response) SetDomains(v [][]SearchArtists200ResponseInnerAllOfDomainsInnerInner)`

SetDomains sets Domains field to given value.


### GetUrls

`func (o *GetArtist200Response) GetUrls() []ArtistURL`

GetUrls returns the Urls field if non-nil, zero value otherwise.

### GetUrlsOk

`func (o *GetArtist200Response) GetUrlsOk() (*[]ArtistURL, bool)`

GetUrlsOk returns a tuple with the Urls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrls

`func (o *GetArtist200Response) SetUrls(v []ArtistURL)`

SetUrls sets Urls field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


