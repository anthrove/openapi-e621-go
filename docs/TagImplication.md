# TagImplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Reason** | **string** |  | 
**CreatorId** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**ForumPostId** | **NullableFloat32** |  | 
**AntecedentName** | **string** |  | 
**ConsequentName** | **string** |  | 
**Status** | [**TagRequestStatuses**](TagRequestStatuses.md) | Note: The \&quot;error\&quot; status will be proceeded by an error, ex: \&quot;error: Validation failed: A tag alias for tag_name already exists\&quot;  | 
**ForumTopicId** | **NullableFloat32** |  | 
**UpdatedAt** | **time.Time** |  | 
**DescendantNames** | **[]string** |  | 
**ApproverId** | **NullableFloat32** |  | 

## Methods

### NewTagImplication

`func NewTagImplication(id int32, reason string, creatorId int32, createdAt time.Time, forumPostId NullableFloat32, antecedentName string, consequentName string, status TagRequestStatuses, forumTopicId NullableFloat32, updatedAt time.Time, descendantNames []string, approverId NullableFloat32, ) *TagImplication`

NewTagImplication instantiates a new TagImplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagImplicationWithDefaults

`func NewTagImplicationWithDefaults() *TagImplication`

NewTagImplicationWithDefaults instantiates a new TagImplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TagImplication) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TagImplication) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TagImplication) SetId(v int32)`

SetId sets Id field to given value.


### GetReason

`func (o *TagImplication) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TagImplication) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TagImplication) SetReason(v string)`

SetReason sets Reason field to given value.


### GetCreatorId

`func (o *TagImplication) GetCreatorId() int32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *TagImplication) GetCreatorIdOk() (*int32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *TagImplication) SetCreatorId(v int32)`

SetCreatorId sets CreatorId field to given value.


### GetCreatedAt

`func (o *TagImplication) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TagImplication) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TagImplication) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetForumPostId

`func (o *TagImplication) GetForumPostId() float32`

GetForumPostId returns the ForumPostId field if non-nil, zero value otherwise.

### GetForumPostIdOk

`func (o *TagImplication) GetForumPostIdOk() (*float32, bool)`

GetForumPostIdOk returns a tuple with the ForumPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumPostId

`func (o *TagImplication) SetForumPostId(v float32)`

SetForumPostId sets ForumPostId field to given value.


### SetForumPostIdNil

`func (o *TagImplication) SetForumPostIdNil(b bool)`

 SetForumPostIdNil sets the value for ForumPostId to be an explicit nil

### UnsetForumPostId
`func (o *TagImplication) UnsetForumPostId()`

UnsetForumPostId ensures that no value is present for ForumPostId, not even an explicit nil
### GetAntecedentName

`func (o *TagImplication) GetAntecedentName() string`

GetAntecedentName returns the AntecedentName field if non-nil, zero value otherwise.

### GetAntecedentNameOk

`func (o *TagImplication) GetAntecedentNameOk() (*string, bool)`

GetAntecedentNameOk returns a tuple with the AntecedentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntecedentName

`func (o *TagImplication) SetAntecedentName(v string)`

SetAntecedentName sets AntecedentName field to given value.


### GetConsequentName

`func (o *TagImplication) GetConsequentName() string`

GetConsequentName returns the ConsequentName field if non-nil, zero value otherwise.

### GetConsequentNameOk

`func (o *TagImplication) GetConsequentNameOk() (*string, bool)`

GetConsequentNameOk returns a tuple with the ConsequentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsequentName

`func (o *TagImplication) SetConsequentName(v string)`

SetConsequentName sets ConsequentName field to given value.


### GetStatus

`func (o *TagImplication) GetStatus() TagRequestStatuses`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TagImplication) GetStatusOk() (*TagRequestStatuses, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TagImplication) SetStatus(v TagRequestStatuses)`

SetStatus sets Status field to given value.


### GetForumTopicId

`func (o *TagImplication) GetForumTopicId() float32`

GetForumTopicId returns the ForumTopicId field if non-nil, zero value otherwise.

### GetForumTopicIdOk

`func (o *TagImplication) GetForumTopicIdOk() (*float32, bool)`

GetForumTopicIdOk returns a tuple with the ForumTopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumTopicId

`func (o *TagImplication) SetForumTopicId(v float32)`

SetForumTopicId sets ForumTopicId field to given value.


### SetForumTopicIdNil

`func (o *TagImplication) SetForumTopicIdNil(b bool)`

 SetForumTopicIdNil sets the value for ForumTopicId to be an explicit nil

### UnsetForumTopicId
`func (o *TagImplication) UnsetForumTopicId()`

UnsetForumTopicId ensures that no value is present for ForumTopicId, not even an explicit nil
### GetUpdatedAt

`func (o *TagImplication) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TagImplication) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TagImplication) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetDescendantNames

`func (o *TagImplication) GetDescendantNames() []string`

GetDescendantNames returns the DescendantNames field if non-nil, zero value otherwise.

### GetDescendantNamesOk

`func (o *TagImplication) GetDescendantNamesOk() (*[]string, bool)`

GetDescendantNamesOk returns a tuple with the DescendantNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescendantNames

`func (o *TagImplication) SetDescendantNames(v []string)`

SetDescendantNames sets DescendantNames field to given value.


### GetApproverId

`func (o *TagImplication) GetApproverId() float32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *TagImplication) GetApproverIdOk() (*float32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *TagImplication) SetApproverId(v float32)`

SetApproverId sets ApproverId field to given value.


### SetApproverIdNil

`func (o *TagImplication) SetApproverIdNil(b bool)`

 SetApproverIdNil sets the value for ApproverId to be an explicit nil

### UnsetApproverId
`func (o *TagImplication) UnsetApproverId()`

UnsetApproverId ensures that no value is present for ApproverId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


