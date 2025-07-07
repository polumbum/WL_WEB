package service

import "errors"

// Sportsman Validation.
var (
	ErrSportsCat  = errors.New("sportsman does not meet the minimum sports category requirement")
	ErrAgeCat     = errors.New("sportsman does not meet the age category requirement")
	ErrAntidoping = errors.New("sportsman does not meet the antidoping requirement")
	ErrAccess     = errors.New("sportsman does not meet the competition access requirement")
	ErrYoung      = errors.New("sportsman does not meet the min age requirement")
)

// Organization.
var (
	ErrDateBefore = errors.New("date before today")
	ErrWrongDate  = errors.New("date before today")
)

// User.
var (
	ErrUserNotFound      = errors.New("user not found")
	ErrWrongPassword     = errors.New("wrong password")
	ErrAlreadyRegistered = errors.New("user already registered")
)

// DA Err.
var (
	//ErrNotFound        = errors.New("record not found")
	ErrADopingNotFound = errors.New("antidoping not found")
	ErrAccessNotFound  = errors.New("access not found")
	ErrConfigLoad      = errors.New("failed to load config")
	ErrNilRef          = errors.New("nil pointer to struct")
)
