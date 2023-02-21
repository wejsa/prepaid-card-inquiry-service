package dto

type IsSkippabler interface {
	IsSkippable() bool
}

type IsFollowingSkippabler interface {
	IsFollowingSkippable() bool
}
