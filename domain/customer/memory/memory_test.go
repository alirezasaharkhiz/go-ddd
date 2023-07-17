package memory

import (
	"ddd/aggregator"
	"ddd/domain/customer"
	"errors"
	"github.com/google/uuid"
	"testing"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cust, err := aggregator.NewCustomer("mamad")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetId()

	mr := MemoryRepository{
		customers: map[uuid.UUID]aggregator.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("32d6f671-45af-4307-b137-4e06fa9003d3"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := mr.Get(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("expected error %v but we got %v", tc.expectedErr, err)
			}
		})
	}
}
