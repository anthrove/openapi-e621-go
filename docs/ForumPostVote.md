# ForumPostVote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ForumPostId** | **int32** |  | 
**CreatorId** | **int32** |  | 
**Score** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CreatorName** | **string** |  | 

## Methods

### NewForumPostVote

`func NewForumPostVote(id int32, forumPostId int32, creatorId int32, score float32, createdAt time.Time, updatedAt time.Time, creatorName string, ) *ForumPostVote`

NewForumPostVote instantiates a new ForumPostVote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForumPostVoteWithDefaults

`func NewForumPostVoteWithDefaults() *ForumPostVote`

NewForumPostVoteWithDefaults instantiates a new ForumPostVote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ForumPostVote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ForumPostVote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ForumPostVote) SetId(v int32)`

SetId sets Id field to given value.


### GetForumPostId

`func (o *ForumPostVote) GetForumPostId() int32`

GetForumPostId returns the ForumPostId field if non-nil, zero value otherwise.

### GetForumPostIdOk

`func (o *ForumPostVote) GetForumPostIdOk() (*int32, bool)`

GetForumPostIdOk returns a tuple with the ForumPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumPostId

`func (o *ForumPostVote) SetForumPostId(v int32)`

SetForumPostId sets ForumPostId field to given value.


### GetCreatorId

`func (o *ForumPostVote) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ForumPostVote) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ForumPostVote) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetScore

`func (o *ForumPostVote) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ForumPostVote) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ForumPostVote) SetScore(v float32)`

SetScore sets Score field to given value.


### GetCreatedAt

`func (o *ForumPostVote) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ForumPostVote) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ForumPostVote) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ForumPostVote) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ForumPostVote) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ForumPostVote) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCreatorName

`func (o *ForumPostVote) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ForumPostVote) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ForumPostVote) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


