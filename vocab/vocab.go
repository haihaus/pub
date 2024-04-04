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
}

// Link is a reference to a resource that is identified by an URL.  When apart
// of a request links can be strings. SEE: [https://www.w3.org/ns/activitystreams#Link].
type Link struct {
	Object

	Href      string `json:"href"`
	Lang      string `json:"hreflang"`
	MediaType string `json:"mediaType"`
}

// Activity is an action that occurs, has occurred, or will occur, on behalf of
// an [Actor], or [Actor]s.  SEE: [https://www.w3.org/ns/activitystreams#Activity]
// and [https://www.w3.org/ns/activitystreams#IntransitiveActivity].
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
