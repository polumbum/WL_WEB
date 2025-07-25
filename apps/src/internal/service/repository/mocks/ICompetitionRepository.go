// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	domain "src/internal/domain"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ICompetitionRepository is an autogenerated mock type for the ICompetitionRepository type
type ICompetitionRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: competition
func (_m *ICompetitionRepository) Create(competition *domain.Competition) error {
	ret := _m.Called(competition)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Competition) error); ok {
		r0 = rf(competition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *ICompetitionRepository) Delete(id uuid.UUID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRegistration provides a mock function with given fields: smID, compID
func (_m *ICompetitionRepository) DeleteRegistration(smID uuid.UUID, compID uuid.UUID) error {
	ret := _m.Called(smID, compID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRegistration")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(smID, compID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCompetitionByID provides a mock function with given fields: competitionID
func (_m *ICompetitionRepository) GetCompetitionByID(competitionID uuid.UUID) (*domain.Competition, error) {
	ret := _m.Called(competitionID)

	if len(ret) == 0 {
		panic("no return value specified for GetCompetitionByID")
	}

	var r0 *domain.Competition
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*domain.Competition, error)); ok {
		return rf(competitionID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *domain.Competition); ok {
		r0 = rf(competitionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Competition)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(competitionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListByOrgID provides a mock function with given fields: id
func (_m *ICompetitionRepository) ListByOrgID(id uuid.UUID) ([]*domain.Competition, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for ListByOrgID")
	}

	var r0 []*domain.Competition
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]*domain.Competition, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []*domain.Competition); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Competition)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCompetitions provides a mock function with given fields: page, batch, sort, filter
func (_m *ICompetitionRepository) ListCompetitions(page int, batch int, sort string, filter string) ([]*domain.Competition, error) {
	ret := _m.Called(page, batch, sort, filter)

	if len(ret) == 0 {
		panic("no return value specified for ListCompetitions")
	}

	var r0 []*domain.Competition
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]*domain.Competition, error)); ok {
		return rf(page, batch, sort, filter)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []*domain.Competition); ok {
		r0 = rf(page, batch, sort, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Competition)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) error); ok {
		r1 = rf(page, batch, sort, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterSportsman provides a mock function with given fields: compApplication
func (_m *ICompetitionRepository) RegisterSportsman(compApplication *domain.CompApplication) error {
	ret := _m.Called(compApplication)

	if len(ret) == 0 {
		panic("no return value specified for RegisterSportsman")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.CompApplication) error); ok {
		r0 = rf(compApplication)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: competition
func (_m *ICompetitionRepository) Update(competition *domain.Competition) error {
	ret := _m.Called(competition)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Competition) error); ok {
		r0 = rf(competition)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewICompetitionRepository creates a new instance of ICompetitionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICompetitionRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICompetitionRepository {
	mock := &ICompetitionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
