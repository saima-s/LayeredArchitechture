package service
type mockStore struct{}

type MockDataStore struct {
	store mockStore
}

func NewS(customer mockStore) MockDataStore {
	return MockDataStore{store: customer}
}
