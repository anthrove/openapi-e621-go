# PostVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**PostId** | **int32** |  | 
**Tags** | **string** |  | 
**UpdaterId** | **int32** |  | 
**UpdatedAt** | **time.Time** |  | 
**Rating** | [**Ratings**](Ratings.md) |  | 
**ParentId** | **NullableInt32** |  | 
**Source** | **string** |  | 
**Description** | **string** |  | 
**Reason** | **NullableString** |  | 
**LockedTags** | **NullableString** |  | 
**AddedTags** | **[]string** |  | 
**RemovedTags** | **[]string** |  | 
**AddedLockedTags** | **[]string** |  | 
**RemovedLockedTags** | **[]string** |  | 
**RatingChanged** | **bool** |  | 
**ParentChanged** | **bool** |  | 
**SourceChanged** | **bool** |  | 
**DescriptionChanged** | **bool** |  | 
**Version** | **float32** |  | 
**ObsoleteAddedTags** | **string** |  | 
**ObsoleteRemovedTags** | **string** |  | 
**UnchangedTags** | **string** |  | 
**UpdaterName** | **string** |  | 

## Methods

### NewPostVersion

`func NewPostVersion(id int32, postId int32, tags string, updaterId int32, updatedAt time.Time, rating Ratings, parentId NullableInt32, source string, description string, reason NullableString, lockedTags NullableString, addedTags []string, removedTags []string, addedLockedTags []string, removedLockedTags []string, ratingChanged bool, parentChanged bool, sourceChanged bool, descriptionChanged bool, version float32, obsoleteAddedTags string, obsoleteRemovedTags string, unchangedTags string, updaterName string, ) *PostVersion`

NewPostVersion instantiates a new PostVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostVersionWithDefaults

`func NewPostVersionWithDefaults() *PostVersion`

NewPostVersionWithDefaults instantiates a new PostVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostVersion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostVersion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostVersion) SetId(v int32)`

SetId sets Id field to given value.


### GetPostId

`func (o *PostVersion) GetPostId() int32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *PostVersion) GetPostIdOk() (*int32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *PostVersion) SetPostId(v int32)`

SetPostId sets PostId field to given value.


### GetTags

`func (o *PostVersion) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostVersion) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostVersion) SetTags(v string)`

SetTags sets Tags field to given value.


### GetUpdaterId

`func (o *PostVersion) GetUpdaterId() int32`

GetUpdaterId returns the UpdaterId field if non-nil, zero value otherwise.

### GetUpdaterIdOk

`func (o *PostVersion) GetUpdaterIdOk() (*int32, bool)`

GetUpdaterIdOk returns a tuple with the UpdaterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterId

`func (o *PostVersion) SetUpdaterId(v int32)`

SetUpdaterId sets UpdaterId field to given value.


### GetUpdatedAt

`func (o *PostVersion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostVersion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostVersion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetRating

`func (o *PostVersion) GetRating() Ratings`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PostVersion) GetRatingOk() (*Ratings, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PostVersion) SetRating(v Ratings)`

SetRating sets Rating field to given value.


### GetParentId

`func (o *PostVersion) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *PostVersion) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *PostVersion) SetParentId(v int32)`

SetParentId sets ParentId field to given value.


### SetParentIdNil

`func (o *PostVersion) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *PostVersion) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetSource

`func (o *PostVersion) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PostVersion) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PostVersion) SetSource(v string)`

SetSource sets Source field to given value.


### GetDescription

`func (o *PostVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostVersion) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetReason

`func (o *PostVersion) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostVersion) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostVersion) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *PostVersion) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *PostVersion) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetLockedTags

`func (o *PostVersion) GetLockedTags() string`

GetLockedTags returns the LockedTags field if non-nil, zero value otherwise.

### GetLockedTagsOk

`func (o *PostVersion) GetLockedTagsOk() (*string, bool)`

GetLockedTagsOk returns a tuple with the LockedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedTags

`func (o *PostVersion) SetLockedTags(v string)`

SetLockedTags sets LockedTags field to given value.


### SetLockedTagsNil

`func (o *PostVersion) SetLockedTagsNil(b bool)`

 SetLockedTagsNil sets the value for LockedTags to be an explicit nil

### UnsetLockedTags
`func (o *PostVersion) UnsetLockedTags()`

UnsetLockedTags ensures that no value is present for LockedTags, not even an explicit nil
### GetAddedTags

`func (o *PostVersion) GetAddedTags() []string`

GetAddedTags returns the AddedTags field if non-nil, zero value otherwise.

### GetAddedTagsOk

`func (o *PostVersion) GetAddedTagsOk() (*[]string, bool)`

GetAddedTagsOk returns a tuple with the AddedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedTags

`func (o *PostVersion) SetAddedTags(v []string)`

SetAddedTags sets AddedTags field to given value.


### GetRemovedTags

`func (o *PostVersion) GetRemovedTags() []string`

GetRemovedTags returns the RemovedTags field if non-nil, zero value otherwise.

### GetRemovedTagsOk

`func (o *PostVersion) GetRemovedTagsOk() (*[]string, bool)`

GetRemovedTagsOk returns a tuple with the RemovedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedTags

`func (o *PostVersion) SetRemovedTags(v []string)`

SetRemovedTags sets RemovedTags field to given value.


### GetAddedLockedTags

`func (o *PostVersion) GetAddedLockedTags() []string`

GetAddedLockedTags returns the AddedLockedTags field if non-nil, zero value otherwise.

### GetAddedLockedTagsOk

`func (o *PostVersion) GetAddedLockedTagsOk() (*[]string, bool)`

GetAddedLockedTagsOk returns a tuple with the AddedLockedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedLockedTags

`func (o *PostVersion) SetAddedLockedTags(v []string)`

SetAddedLockedTags sets AddedLockedTags field to given value.


### GetRemovedLockedTags

`func (o *PostVersion) GetRemovedLockedTags() []string`

GetRemovedLockedTags returns the RemovedLockedTags field if non-nil, zero value otherwise.

### GetRemovedLockedTagsOk

`func (o *PostVersion) GetRemovedLockedTagsOk() (*[]string, bool)`

GetRemovedLockedTagsOk returns a tuple with the RemovedLockedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedLockedTags

`func (o *PostVersion) SetRemovedLockedTags(v []string)`

SetRemovedLockedTags sets RemovedLockedTags field to given value.


### GetRatingChanged

`func (o *PostVersion) GetRatingChanged() bool`

GetRatingChanged returns the RatingChanged field if non-nil, zero value otherwise.

### GetRatingChangedOk

`func (o *PostVersion) GetRatingChangedOk() (*bool, bool)`

GetRatingChangedOk returns a tuple with the RatingChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingChanged

`func (o *PostVersion) SetRatingChanged(v bool)`

SetRatingChanged sets RatingChanged field to given value.


### GetParentChanged

`func (o *PostVersion) GetParentChanged() bool`

GetParentChanged returns the ParentChanged field if non-nil, zero value otherwise.

### GetParentChangedOk

`func (o *PostVersion) GetParentChangedOk() (*bool, bool)`

GetParentChangedOk returns a tuple with the ParentChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentChanged

`func (o *PostVersion) SetParentChanged(v bool)`

SetParentChanged sets ParentChanged field to given value.


### GetSourceChanged

`func (o *PostVersion) GetSourceChanged() bool`

GetSourceChanged returns the SourceChanged field if non-nil, zero value otherwise.

### GetSourceChangedOk

`func (o *PostVersion) GetSourceChangedOk() (*bool, bool)`

GetSourceChangedOk returns a tuple with the SourceChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceChanged

`func (o *PostVersion) SetSourceChanged(v bool)`

SetSourceChanged sets SourceChanged field to given value.


### GetDescriptionChanged

`func (o *PostVersion) GetDescriptionChanged() bool`

GetDescriptionChanged returns the DescriptionChanged field if non-nil, zero value otherwise.

### GetDescriptionChangedOk

`func (o *PostVersion) GetDescriptionChangedOk() (*bool, bool)`

GetDescriptionChangedOk returns a tuple with the DescriptionChanged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionChanged

`func (o *PostVersion) SetDescriptionChanged(v bool)`

SetDescriptionChanged sets DescriptionChanged field to given value.


### GetVersion

`func (o *PostVersion) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostVersion) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostVersion) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetObsoleteAddedTags

`func (o *PostVersion) GetObsoleteAddedTags() string`

GetObsoleteAddedTags returns the ObsoleteAddedTags field if non-nil, zero value otherwise.

### GetObsoleteAddedTagsOk

`func (o *PostVersion) GetObsoleteAddedTagsOk() (*string, bool)`

GetObsoleteAddedTagsOk returns a tuple with the ObsoleteAddedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoleteAddedTags

`func (o *PostVersion) SetObsoleteAddedTags(v string)`

SetObsoleteAddedTags sets ObsoleteAddedTags field to given value.


### GetObsoleteRemovedTags

`func (o *PostVersion) GetObsoleteRemovedTags() string`

GetObsoleteRemovedTags returns the ObsoleteRemovedTags field if non-nil, zero value otherwise.

### GetObsoleteRemovedTagsOk

`func (o *PostVersion) GetObsoleteRemovedTagsOk() (*string, bool)`

GetObsoleteRemovedTagsOk returns a tuple with the ObsoleteRemovedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObsoleteRemovedTags

`func (o *PostVersion) SetObsoleteRemovedTags(v string)`

SetObsoleteRemovedTags sets ObsoleteRemovedTags field to given value.


### GetUnchangedTags

`func (o *PostVersion) GetUnchangedTags() string`

GetUnchangedTags returns the UnchangedTags field if non-nil, zero value otherwise.

### GetUnchangedTagsOk

`func (o *PostVersion) GetUnchangedTagsOk() (*string, bool)`

GetUnchangedTagsOk returns a tuple with the UnchangedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnchangedTags

`func (o *PostVersion) SetUnchangedTags(v string)`

SetUnchangedTags sets UnchangedTags field to given value.


### GetUpdaterName

`func (o *PostVersion) GetUpdaterName() string`

GetUpdaterName returns the UpdaterName field if non-nil, zero value otherwise.

### GetUpdaterNameOk

`func (o *PostVersion) GetUpdaterNameOk() (*string, bool)`

GetUpdaterNameOk returns a tuple with the UpdaterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdaterName

`func (o *PostVersion) SetUpdaterName(v string)`

SetUpdaterName sets UpdaterName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


