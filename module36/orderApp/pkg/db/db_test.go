package db

import (
	"orderApp/pkg/model"
	"testing"
)

func TestCRUDOperations(t *testing.T) {
	db := NewDB()

	// CREATE
	order := model.Order{
		IsOpen:          true,
		DeliveryTime:    1627847285,
		DeliveryAddress: "123 Main St",
		Products: []model.Product{
			{ID: 1, Name: "Product 1", Price: 10.0},
			{ID: 2, Name: "Product 2", Price: 20.0},
		},
	}
	createdID := db.AddOrder(order)

	if createdID != 1 {
		t.Fatalf("Expected created order ID to be 1, got %d", createdID)
	}

	// READ
	retrievedOrders := db.GetOrders()
	if len(retrievedOrders) != 1 {
		t.Fatalf("Expected 1 order, got %d", len(retrievedOrders))
	}

	retrievedOrder := retrievedOrders[0]
	if retrievedOrder.ID != createdID || retrievedOrder.DeliveryAddress != "123 Main St" {
		t.Errorf("Expected retrieved order to have ID %d and address '123 Main St', got ID %d and address '%s'", createdID, retrievedOrder.ID, retrievedOrder.DeliveryAddress)
	}

	// UPDATE
	updatedOrder := model.Order{
		ID:              createdID,
		IsOpen:          false,
		DeliveryTime:    1627847285,
		DeliveryAddress: "456 Elm St",
		Products: []model.Product{
			{ID: 2, Name: "Product 3", Price: 30.0},
		},
	}
	db.UpdateOrder(updatedOrder)

	updatedRetrievedOrder := db.store[createdID]
	if updatedRetrievedOrder.DeliveryAddress != "456 Elm St" {
		t.Errorf("Expected updated order to have address '456 Elm St' and be closed, got address '%s' and IsOpen=%t", updatedRetrievedOrder.DeliveryAddress, updatedRetrievedOrder.IsOpen)
	}

	// DELETE
	db.DeleteOrder(createdID)
	if len(db.store) != 0 {
		t.Fatalf("Expected store to be empty after deletion, got %d orders", len(db.store))
	}
	if _, exists := db.store[createdID]; exists {
		t.Errorf("Expected order with ID %d to be deleted, but it still exists", createdID)
	}
}
