package vocab

// [Object] and [Link] types use [Object] or [Link] as their base.

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
