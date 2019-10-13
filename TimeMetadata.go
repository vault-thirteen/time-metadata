package tm

import (
	"time"
)

type TimeMetadata struct {

	// UNIX Timestamp of Creation Time.
	creationTime int64

	// UNIX Timestamp of Update Time.
	updateTime int64
}

// Creates a new Time Meta-Data with 'Creation Time' Field filled with the
// current Time.
func New() *TimeMetadata {
	return &TimeMetadata{
		creationTime: time.Now().Unix(),
	}
}

// Updates the 'Update Time' Field of the Time Meta-Data Object using the
// current Time.
func (this *TimeMetadata) Update() {
	this.updateTime = time.Now().Unix()
}

// Returns the 'Creation Time' Field of the Time Meta-Data Object.
func (this TimeMetadata) GetCreationTime() int64 {
	return this.creationTime
}

// Returns the 'Update Time' Field of the Time Meta-Data Object.
func (this TimeMetadata) GetUpdateTime() int64 {
	return this.updateTime
}
