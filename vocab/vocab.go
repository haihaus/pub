// ActivityStreams Vocabulary in Golang
// SEE: [https://www.w3.org/TR/activitystreams-vocabulary]
package vocab

// This file outlines the core ActivityStreams vocab, SEE:
// [https://www.w3.org/TR/activitystreams-vocabulary/#types].

// Object is the base of all ActivityStream types.  SEE:
// [https://www.w3.org/ns/activitystreams#Object].
type Object struct {
	Context string `json:"@context"`
	Type    string `json:"type"`
	ID      string `json:"id"`
	Name    string `json:"name"`

	// Summary is a natural language summarization of the object.
	//
	// Can be a string or map[string]string
	// SEE: [https://www.w3.org/ns/activitystreams#summary]
	Summary *interface{} `json:"summary"`

	// Replies is an [Collection] of responses to the object.
	//
	// Can be a [Collection]
	// SEE: [https://www.w3.org/ns/activitystreams#replies]
	Replies *Collection `json:"collection"`

	// Attachment is a resource attached to, or related to, the object.
	//
	// Can be an [Object] or a [Link].
	// SEE: [https://www.w3.org/ns/activitystreams#attachment].
	Attachment *interface{} `json:"attachment"`

	// Preview identifies the entity that provides a preview of the Object.
	//
	// Can be an [Object] or [Link].
	// SEE: [https://www.w3.org/ns/activitystreams#preview]
	Preview *interface{} `json:"preview"`

	// AttributedTo identifies the actors that the object is attributed to.
	//
	// Can be an [Object] or [Link].
	// SEE: [https://www.w3.org/ns/activitystreams#attributedTo].
	AttributedTo *interface{} `json:"attributedTo"`

	// Audience identifies the entities of the population that the object is
	// relevant to.
	//
	// Can be an [Object] or [Link].
	// SEE: [https://www.w3.org/ns/activitystreams#audience]
	Audience *interface{} `json:"audience"`

	// To identifies one or more [Object]s that are the public primary audience
	// of the Object.
	//
	// Can be an [Object] or [Link], or an array of both.
	// SEE: [https://www.w3.org/ns/activitystreams#to]
	To *interface{} `json:"to"`

	// Bto identifies one or more [Object]s that are the private primary audience
	// of the Object.
	//
	// Can be an [Object] or [Link], or an array of both.
	// SEE: [https://www.w3.org/ns/activitystreams#bto]
	Bto *interface{} `json:"bto"`

	// Cc identifies the [Object] that is the public secondary audience of the/
	// Object
	//
	// Can be an [Object] or [Link].
	// SEE: [https://www.w3.org/ns/activitystreams#cc]
	Cc *interface{} `json:"cc"`

	// Bcc identifies one or more [Object]s that are the private secondary
	// audience of the Object.
	//
	// Can be an [Object] or [Link], or array of both.
	// SEE: [https://www.w3.org/ns/activitystreams#bcc]
	Bcc *interface{} `json:"bcc"`

	// Tag are the one or more "tags" tha are associated with the Object.
	//
	// Can be an [Object] or [Link], or an array of both.
	// SEE: [https://www.w3.org/ns/activitystreams#tag]
	Tag *interface{} `json:"tag"`

	// Generator identifies the entity that generates an object.
	//
	// Can be an [Object] or [Link].
	// SEE: [https://www.w3.org/ns/activitystreams#generator]
	Generator *interface{}

	// Content is the content of the Object.
	//
	// Can be an map[string]string (where keys are lang codes) or a string.
	// SEE: [https://www.w3.org/ns/activitystreams#content]
	Content *interface{} `json:"content"`

	// Icon is the Object's icon.
	//
	// Can be an [Icon], [Link], or array of both.
	// SEE: [https://www.w3.org/ns/activitystreams#icon]
	Icon *interface{} `json:"icon"`

	// Image is an image document.
	//
	// Can be an [Object], [Link], or an array of both.
	// SEE: [https://www.w3.org/ns/activitystreams#image]
	Image *interface{} `json:"image"`

	// InReplyTo indicates the entity that the object is a response to.
	//
	// Can be an [Object] or [Link].
	// SEE: [https://www.w3.org/ns/activitystreams#inReplyTo]
	InReplyTo *interface{} `json:"inReplyTo"`

	// Location is a physical or logical location(s) associated with the object.
	//
	// Can be a [Place], [Object], or [Link].
	// SEE: [https://www.w3.org/ns/activitystreams#location]
	Location *interface{} `json:"location"`

	// URL identifies one or more links to more representation of an Object.
	//
	// Can be a string, [Link], or array of both.
	// SEE: [https://www.w3.org/ns/activitystreams#url]
	URL *interface{} `json:"url"`

	// StartTime is the time that the object starts.
	//
	// SEE: [https://www.w3.org/ns/activitystreams#startTime]
	StartTime *string `json:"startTime"`

	// EndTime is the time that the object ends.
	//
	// SEE: [https://www.w3.org/ns/activitystreams#endTime]
	EndTime *string `json:"endTime"`

	// Duration is used if the object time-bound and has a approximate duration.
	//
	// Can be a string that is an xsd:duration.
	// SEE: [https://www.w3.org/ns/activitystreams#duration]
	Duration *string `json:"duration"`

	// Updated is the datetimestamp which that object was last updated at.
	//
	// Can be a string that is an xsd:dateTime.
	// SEE: [https://www.w3.org/ns/activitystreams#updated].
	Update *string `json:"update"`

	// Published is the datetimestamp that the object was published at.
	//
	// Can be a string that is an xsd:dateTime.
	// SEE: [https://www.w3.org/ns/activitystreams#published].
	Published *string `json:"published"`

	// If used as an direct [Object] MediaType identifies the MIME media type of
	// the [Content] property.  If used as apart of the [Link], it identifies the
	// MIME media type of the referenced resource.
	//
	// SEE: [https://www.w3.org/ns/activitystreams#mediaType]
	MediaType *string `json:"mediaType"`
}

// Icon represents an icon.
type Icon struct {
	Object

	Width  int `json:"width"`
	Height int `json:"height"`
}

// Link is a reference to a resource that is identified by an URL.  When apart
// of a request links can be strings. SEE: [https://www.w3.org/ns/activitystreams#Link].
type Link struct {
	Object

	Href string `json:"href"`
	Lang string `json:"hreflang"`
}

// Activity is an action that occurs, has occurred, or will occur, on behalf of
// an [Actor], or [Actor]s.  SEE: [https://www.w3.org/ns/activitystreams#Activity],
// in some cases Activity can represent extended Activity types.
type Activity struct {
	Object

	Actor Actor `json:"actor"`

	// RelatedObject is the [Object] that the [Actor] is acting on behalf.  It
	// can be an [Link], [Object], or string that represents an URL to an ID.
	RelatedObject *interface{} `json:"object"`

	// Target is mostly used of IntransitiveActivities, which are activities that
	// are related to the English proposition "to."  It can be a [Link],
	// [Object], or string that represents an URL to an ID.
	Target *interface{} `json:"target"`

	// Origin is the is the source of the Activity.
	Origin *interface{} `json:"origin"`
}

// Collection is an (un)ordered set of [Object]s.  It is required to check the
// type, as an "OrderedCollection" will *always* be strictly ordered.  SEE:
// [https://www.w3.org/ns/activitystreams#Collection] and
// [https://www.w3.org/ns/activitystreams#OrderedCollection].
type Collection struct {
	Object

	// TotalItems the total number of items in the set.  When apart of a
	// [CollectionPage] it will be nil.
	TotalItems *int `json:"totalItems"`

	// The items that are stored by the set, it can be an [Object] or [Link].
	Items interface{} `json:"items"`
}

// CollectionPage is a subset of a [Collection].  It is required to check the
// type, as an "OrderedCollectionPage" will relate to an "OrderedCollection".
// SEE: [https://www.w3.org/ns/activitystreams#OrderedCollectionPage] and
// [https://www.w3.org/ns/activitystreams#CollectionPage].
type CollectionPage struct {
	Collection

	// Whatever [Collection] the subset is apart of.  It can be a [Link] or
	// [Collection].
	PartOf interface{} `json:"partOf"`
}
