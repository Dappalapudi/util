package docid

import "github.com/rs/xid"

// New creates new uids.
func New() string {
	return xid.New().String()
}
