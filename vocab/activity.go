package vocab

// Activity Types use [Activity] as their base.

// Question is a question that is being asked
type Question struct {
	Object

	// OneOf identifies an exclusive set of options.
	OneOf []interface{} `json:"oneOf"`

	// AnyOf identifies an inclusive set of options.
	AnyOf []interface{} `json:"anyOf"`

	// Closed that the question is no longer accepting answers.  Can be an
	// [Object], [Link], dateTime, or boolean.
	Closed interface{} `json:"closed"`
}
