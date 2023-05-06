package membertypes

import "time"

type SchoolType int8

const (
	Elementary SchoolType = iota + 1
)

type VirtualStudent struct {
	DOB        time.Time  `json:"dob"`
	School     string     `json:"school"`
	SchoolType SchoolType `json:"school_type"`
}

func (st SchoolType) String() string {
	switch st {
	case Elementary:
		return "elementary"
	}
	return "unknown school type"
}
