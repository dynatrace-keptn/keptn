// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package db_mock

import (
	"github.com/keptn/keptn/shipyard-controller/common"
	"github.com/keptn/keptn/shipyard-controller/models"
	"sync"
)

// EventRepoMock is a mock implementation of db.EventRepo.
//
// 	func TestSomethingThatUsesEventRepo(t *testing.T) {
//
// 		// make and configure a mocked db.EventRepo
// 		mockedEventRepo := &EventRepoMock{
// 			DeleteEventFunc: func(project string, eventID string, status common.EventStatus) error {
// 				panic("mock out the DeleteEvent method")
// 			},
// 			DeleteEventCollectionsFunc: func(project string) error {
// 				panic("mock out the DeleteEventCollections method")
// 			},
// 			GetEventsFunc: func(project string, filter common.EventFilter, status ...common.EventStatus) ([]models.Event, error) {
// 				panic("mock out the GetEvents method")
// 			},
// 			InsertEventFunc: func(project string, event models.Event, status common.EventStatus) error {
// 				panic("mock out the InsertEvent method")
// 			},
// 		}
//
// 		// use mockedEventRepo in code that requires db.EventRepo
// 		// and then make assertions.
//
// 	}
type EventRepoMock struct {
	// DeleteEventFunc mocks the DeleteEvent method.
	DeleteEventFunc func(project string, eventID string, status common.EventStatus) error

	// DeleteEventCollectionsFunc mocks the DeleteEventCollections method.
	DeleteEventCollectionsFunc func(project string) error

	// GetEventsFunc mocks the GetEvents method.
	GetEventsFunc func(project string, filter common.EventFilter, status ...common.EventStatus) ([]models.Event, error)

	// InsertEventFunc mocks the InsertEvent method.
	InsertEventFunc func(project string, event models.Event, status common.EventStatus) error

	// calls tracks calls to the methods.
	calls struct {
		// DeleteEvent holds details about calls to the DeleteEvent method.
		DeleteEvent []struct {
			// Project is the project argument value.
			Project string
			// EventID is the eventID argument value.
			EventID string
			// Status is the status argument value.
			Status common.EventStatus
		}
		// DeleteEventCollections holds details about calls to the DeleteEventCollections method.
		DeleteEventCollections []struct {
			// Project is the project argument value.
			Project string
		}
		// GetEvents holds details about calls to the GetEvents method.
		GetEvents []struct {
			// Project is the project argument value.
			Project string
			// Filter is the filter argument value.
			Filter common.EventFilter
			// Status is the status argument value.
			Status []common.EventStatus
		}
		// InsertEvent holds details about calls to the InsertEvent method.
		InsertEvent []struct {
			// Project is the project argument value.
			Project string
			// Event is the event argument value.
			Event models.Event
			// Status is the status argument value.
			Status common.EventStatus
		}
	}
	lockDeleteEvent            sync.RWMutex
	lockDeleteEventCollections sync.RWMutex
	lockGetEvents              sync.RWMutex
	lockInsertEvent            sync.RWMutex
}

// DeleteEvent calls DeleteEventFunc.
func (mock *EventRepoMock) DeleteEvent(project string, eventID string, status common.EventStatus) error {
	if mock.DeleteEventFunc == nil {
		panic("EventRepoMock.DeleteEventFunc: method is nil but EventRepo.DeleteEvent was just called")
	}
	callInfo := struct {
		Project string
		EventID string
		Status  common.EventStatus
	}{
		Project: project,
		EventID: eventID,
		Status:  status,
	}
	mock.lockDeleteEvent.Lock()
	mock.calls.DeleteEvent = append(mock.calls.DeleteEvent, callInfo)
	mock.lockDeleteEvent.Unlock()
	return mock.DeleteEventFunc(project, eventID, status)
}

// DeleteEventCalls gets all the calls that were made to DeleteEvent.
// Check the length with:
//     len(mockedEventRepo.DeleteEventCalls())
func (mock *EventRepoMock) DeleteEventCalls() []struct {
	Project string
	EventID string
	Status  common.EventStatus
} {
	var calls []struct {
		Project string
		EventID string
		Status  common.EventStatus
	}
	mock.lockDeleteEvent.RLock()
	calls = mock.calls.DeleteEvent
	mock.lockDeleteEvent.RUnlock()
	return calls
}

// DeleteEventCollections calls DeleteEventCollectionsFunc.
func (mock *EventRepoMock) DeleteEventCollections(project string) error {
	if mock.DeleteEventCollectionsFunc == nil {
		panic("EventRepoMock.DeleteEventCollectionsFunc: method is nil but EventRepo.DeleteEventCollections was just called")
	}
	callInfo := struct {
		Project string
	}{
		Project: project,
	}
	mock.lockDeleteEventCollections.Lock()
	mock.calls.DeleteEventCollections = append(mock.calls.DeleteEventCollections, callInfo)
	mock.lockDeleteEventCollections.Unlock()
	return mock.DeleteEventCollectionsFunc(project)
}

// DeleteEventCollectionsCalls gets all the calls that were made to DeleteEventCollections.
// Check the length with:
//     len(mockedEventRepo.DeleteEventCollectionsCalls())
func (mock *EventRepoMock) DeleteEventCollectionsCalls() []struct {
	Project string
} {
	var calls []struct {
		Project string
	}
	mock.lockDeleteEventCollections.RLock()
	calls = mock.calls.DeleteEventCollections
	mock.lockDeleteEventCollections.RUnlock()
	return calls
}

// GetEvents calls GetEventsFunc.
func (mock *EventRepoMock) GetEvents(project string, filter common.EventFilter, status ...common.EventStatus) ([]models.Event, error) {
	if mock.GetEventsFunc == nil {
		panic("EventRepoMock.GetEventsFunc: method is nil but EventRepo.GetEvents was just called")
	}
	callInfo := struct {
		Project string
		Filter  common.EventFilter
		Status  []common.EventStatus
	}{
		Project: project,
		Filter:  filter,
		Status:  status,
	}
	mock.lockGetEvents.Lock()
	mock.calls.GetEvents = append(mock.calls.GetEvents, callInfo)
	mock.lockGetEvents.Unlock()
	return mock.GetEventsFunc(project, filter, status...)
}

// GetEventsCalls gets all the calls that were made to GetEvents.
// Check the length with:
//     len(mockedEventRepo.GetEventsCalls())
func (mock *EventRepoMock) GetEventsCalls() []struct {
	Project string
	Filter  common.EventFilter
	Status  []common.EventStatus
} {
	var calls []struct {
		Project string
		Filter  common.EventFilter
		Status  []common.EventStatus
	}
	mock.lockGetEvents.RLock()
	calls = mock.calls.GetEvents
	mock.lockGetEvents.RUnlock()
	return calls
}

// InsertEvent calls InsertEventFunc.
func (mock *EventRepoMock) InsertEvent(project string, event models.Event, status common.EventStatus) error {
	if mock.InsertEventFunc == nil {
		panic("EventRepoMock.InsertEventFunc: method is nil but EventRepo.InsertEvent was just called")
	}
	callInfo := struct {
		Project string
		Event   models.Event
		Status  common.EventStatus
	}{
		Project: project,
		Event:   event,
		Status:  status,
	}
	mock.lockInsertEvent.Lock()
	mock.calls.InsertEvent = append(mock.calls.InsertEvent, callInfo)
	mock.lockInsertEvent.Unlock()
	return mock.InsertEventFunc(project, event, status)
}

// InsertEventCalls gets all the calls that were made to InsertEvent.
// Check the length with:
//     len(mockedEventRepo.InsertEventCalls())
func (mock *EventRepoMock) InsertEventCalls() []struct {
	Project string
	Event   models.Event
	Status  common.EventStatus
} {
	var calls []struct {
		Project string
		Event   models.Event
		Status  common.EventStatus
	}
	mock.lockInsertEvent.RLock()
	calls = mock.calls.InsertEvent
	mock.lockInsertEvent.RUnlock()
	return calls
}
