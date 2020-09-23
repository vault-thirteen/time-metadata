// TimeMetadata.go.

////////////////////////////////////////////////////////////////////////////////
//
// Copyright © 2019..2020 by Vault Thirteen.
//
// All rights reserved. No part of this publication may be reproduced,
// distributed, or transmitted in any form or by any means, including
// photocopying, recording, or other electronic or mechanical methods,
// without the prior written permission of the publisher, except in the case
// of brief quotations embodied in critical reviews and certain other
// noncommercial uses permitted by copyright law. For permission requests,
// write to the publisher, addressed “Copyright Protected Material” at the
// address below.
//
////////////////////////////////////////////////////////////////////////////////
//
// Web Site Address:	https://github.com/vault-thirteen.
//
////////////////////////////////////////////////////////////////////////////////

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
