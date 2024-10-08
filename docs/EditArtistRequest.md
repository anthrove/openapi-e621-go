# EditArtistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtistName** | Pointer to **string** | Only usable for Janitor+ | [optional] 
**ArtistOtherNames** | Pointer to **[]string** |  | [optional] 
**ArtistOtherNamesString** | Pointer to **string** |  | [optional] 
**ArtistUrlString** | Pointer to **string** |  | [optional] 
**ArtistNotes** | Pointer to **string** |  | [optional] 
**ArtistGroupName** | Pointer to **string** |  | [optional] 
**ArtistLinkedUserId** | Pointer to **NullableFloat32** | Only usable for Janitor+ | [optional] 
**ArtistIsLocked** | Pointer to **bool** | Only usable for Janitor+ | [optional] 

## Methods

### NewEditArtistRequest

`func NewEditArtistRequest() *EditArtistRequest`

NewEditArtistRequest instantiates a new EditArtistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditArtistRequestWithDefaults

`func NewEditArtistRequestWithDefaults() *EditArtistRequest`

NewEditArtistRequestWithDefaults instantiates a new EditArtistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtistName

`func (o *EditArtistRequest) GetArtistName() string`

GetArtistName returns the ArtistName field if non-nil, zero value otherwise.

### GetArtistNameOk

`func (o *EditArtistRequest) GetArtistNameOk() (*string, bool)`

GetArtistNameOk returns a tuple with the ArtistName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistName

`func (o *EditArtistRequest) SetArtistName(v string)`

SetArtistName sets ArtistName field to given value.

### HasArtistName

`func (o *EditArtistRequest) HasArtistName() bool`

HasArtistName returns a boolean if a field has been set.

### GetArtistOtherNames

`func (o *EditArtistRequest) GetArtistOtherNames() []string`

GetArtistOtherNames returns the ArtistOtherNames field if non-nil, zero value otherwise.

### GetArtistOtherNamesOk

`func (o *EditArtistRequest) GetArtistOtherNamesOk() (*[]string, bool)`

GetArtistOtherNamesOk returns a tuple with the ArtistOtherNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistOtherNames

`func (o *EditArtistRequest) SetArtistOtherNames(v []string)`

SetArtistOtherNames sets ArtistOtherNames field to given value.

### HasArtistOtherNames

`func (o *EditArtistRequest) HasArtistOtherNames() bool`

HasArtistOtherNames returns a boolean if a field has been set.

### GetArtistOtherNamesString

`func (o *EditArtistRequest) GetArtistOtherNamesString() string`

GetArtistOtherNamesString returns the ArtistOtherNamesString field if non-nil, zero value otherwise.

### GetArtistOtherNamesStringOk

`func (o *EditArtistRequest) GetArtistOtherNamesStringOk() (*string, bool)`

GetArtistOtherNamesStringOk returns a tuple with the ArtistOtherNamesString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistOtherNamesString

`func (o *EditArtistRequest) SetArtistOtherNamesString(v string)`

SetArtistOtherNamesString sets ArtistOtherNamesString field to given value.

### HasArtistOtherNamesString

`func (o *EditArtistRequest) HasArtistOtherNamesString() bool`

HasArtistOtherNamesString returns a boolean if a field has been set.

### GetArtistUrlString

`func (o *EditArtistRequest) GetArtistUrlString() string`

GetArtistUrlString returns the ArtistUrlString field if non-nil, zero value otherwise.

### GetArtistUrlStringOk

`func (o *EditArtistRequest) GetArtistUrlStringOk() (*string, bool)`

GetArtistUrlStringOk returns a tuple with the ArtistUrlString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistUrlString

`func (o *EditArtistRequest) SetArtistUrlString(v string)`

SetArtistUrlString sets ArtistUrlString field to given value.

### HasArtistUrlString

`func (o *EditArtistRequest) HasArtistUrlString() bool`

HasArtistUrlString returns a boolean if a field has been set.

### GetArtistNotes

`func (o *EditArtistRequest) GetArtistNotes() string`

GetArtistNotes returns the ArtistNotes field if non-nil, zero value otherwise.

### GetArtistNotesOk

`func (o *EditArtistRequest) GetArtistNotesOk() (*string, bool)`

GetArtistNotesOk returns a tuple with the ArtistNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistNotes

`func (o *EditArtistRequest) SetArtistNotes(v string)`

SetArtistNotes sets ArtistNotes field to given value.

### HasArtistNotes

`func (o *EditArtistRequest) HasArtistNotes() bool`

HasArtistNotes returns a boolean if a field has been set.

### GetArtistGroupName

`func (o *EditArtistRequest) GetArtistGroupName() string`

GetArtistGroupName returns the ArtistGroupName field if non-nil, zero value otherwise.

### GetArtistGroupNameOk

`func (o *EditArtistRequest) GetArtistGroupNameOk() (*string, bool)`

GetArtistGroupNameOk returns a tuple with the ArtistGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistGroupName

`func (o *EditArtistRequest) SetArtistGroupName(v string)`

SetArtistGroupName sets ArtistGroupName field to given value.

### HasArtistGroupName

`func (o *EditArtistRequest) HasArtistGroupName() bool`

HasArtistGroupName returns a boolean if a field has been set.

### GetArtistLinkedUserId

`func (o *EditArtistRequest) GetArtistLinkedUserId() float32`

GetArtistLinkedUserId returns the ArtistLinkedUserId field if non-nil, zero value otherwise.

### GetArtistLinkedUserIdOk

`func (o *EditArtistRequest) GetArtistLinkedUserIdOk() (*float32, bool)`

GetArtistLinkedUserIdOk returns a tuple with the ArtistLinkedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistLinkedUserId

`func (o *EditArtistRequest) SetArtistLinkedUserId(v float32)`

SetArtistLinkedUserId sets ArtistLinkedUserId field to given value.

### HasArtistLinkedUserId

`func (o *EditArtistRequest) HasArtistLinkedUserId() bool`

HasArtistLinkedUserId returns a boolean if a field has been set.

### SetArtistLinkedUserIdNil

`func (o *EditArtistRequest) SetArtistLinkedUserIdNil(b bool)`

 SetArtistLinkedUserIdNil sets the value for ArtistLinkedUserId to be an explicit nil

### UnsetArtistLinkedUserId
`func (o *EditArtistRequest) UnsetArtistLinkedUserId()`

UnsetArtistLinkedUserId ensures that no value is present for ArtistLinkedUserId, not even an explicit nil
### GetArtistIsLocked

`func (o *EditArtistRequest) GetArtistIsLocked() bool`

GetArtistIsLocked returns the ArtistIsLocked field if non-nil, zero value otherwise.

### GetArtistIsLockedOk

`func (o *EditArtistRequest) GetArtistIsLockedOk() (*bool, bool)`

GetArtistIsLockedOk returns a tuple with the ArtistIsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistIsLocked

`func (o *EditArtistRequest) SetArtistIsLocked(v bool)`

SetArtistIsLocked sets ArtistIsLocked field to given value.

### HasArtistIsLocked

`func (o *EditArtistRequest) HasArtistIsLocked() bool`

HasArtistIsLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


