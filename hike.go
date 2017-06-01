package hikestr

import (
	"net/url"
)

// MetaElevation lists a maximum and minimum elevation present on a hike.
type MetaElevation struct {
	// Gain is the number of feet that will be ascended while hiking this trail.
	Gain int `json:"gain"`
	// Loss is the number of feet that will be descended while hiking this trail.
	Loss int `json:"loss"`
	// Max is the highest point in feet of a hike
	Max int `json:"max"`
	// Min is the lowest point in feet of a hike
	Min int `json:"min"`
}

// Hike represents a single entry
type Hike struct {
	// Details is a link to further information about this hike.
	Details url.URL `json:"details"`
	// DogFriendly denotes whether or not dogs are permitted on the trail.
	DogFriendly bool `json:"dogs"`
	//MetaElevation has information about the maximum and minimum elevation of a hike.
	MetaElevation `json:"elevation"`
	// Length is the number of miles that will be hiked while on this trail.
	Length float32 `json:"length"`
	// TrailHead is position of the trail head.
	TrailHead Location `json:"location"`
	// Permits lists the permits that are accepted to access this hike.
	Permits []string `json:"permits"`
	// Rating is the score out of a hundred of how enjoyable this hike is.
	Rating int8 `json:"rating"`
}
