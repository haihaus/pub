package vocab

// [Object] and [Link] types use [Object] or [Link] as their base.

// Tombstone represents content that has been deleted.
type Tombstone struct {
  Object

  // FormerType is the type of the object that was deleted.
  // SEE: [https://www.w3.org/ns/activitystreams#formerType]
  FormerType string `json:"formerType"`
  
  // Deleted is the datetimestamp of when the object was deleted. 
  Deleted string `json:"deleted"`
}

// Relationship describes the relationship between two individuals.
// SEE: [https://www.w3.org/ns/activitystreams#Relationship]
type Relationship struct {
  Object
  
  // Relationship is the type of the relationship
  Relationship string `json:"relationship"`
}

// Profile is a content [Object] that describes another [Object], in most cases
// it is an actor.
// SEE: [https://www.w3.org/ns/activitystreams#Profile]
type Profile struct {
  Object

  // Describes is the [Object] that the Profile describes.
  // SEE: [https://www.w3.org/ns/activitystreams#describes].
  Describes Object `json:"describes"`
}

// Place represents a location that can be physical or logical.
// SEE: [https://www.w3.org/TR/activitystreams-vocabulary/#dfn-place]
type Place struct {
	Object

	// Accuracy indicates, as a percentage, how accurate the position is.
	// SEE: https://www.w3.org/ns/activitystreams#accuracy
	Accuracy float `json:"accuracy"`

	// Altitude indicates the altitude of the place.
	// SEE: [https://www.w3.org/ns/activitystreams#altitude]
	Altitude float `json:"altitude"`

	// Latitude is the latitude of the place.
	// SEE: [https://www.w3.org/ns/activitystreams#latitude]
	Latitude float `json:"latitude"`

	// Longitude is the longitude of the place.
	// SEE: [https://www.w3.org/ns/activitystreams#longitude]
	Longitude float `json:"longitude"`

	// Radius is the radius of the location.
	// SEE: [https://www.w3.org/ns/activitystreams#radius]
	Radius float `json:"radius"`
}
