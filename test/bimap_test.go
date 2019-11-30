package main

import (
	"fmt"
	bm "bimap"
	"testing"
	"github.com/stretchr/testify/assert"
) 

func TestBiMap(t *testing.T) {
	m := bm.NewBiMap()

	// Test Insert Functions
	m.Insert(1,"Vivek")
	m.Insert(2,"Mogili")
	m.Insert(3,"Sai")

	assert.Equal(t, "Vivek", m.FindByKey(1), "Pair (1,\"Vivek\") should be in BiMap.")
	assert.Equal(t, "Sai", m.FindByKey(3), "Pair (3,\"Sai\") should be in BiMap.")
	assert.Equal(t, 2, m.FindByValue("Mogili"), "Pair (2,\"Mogili\") should be in BiMap.")

	m.Insert(1,"Vaiveek") // key in map but not value so (1, "Vivek") is replaced by (1, "Vaiveek")
	m.Insert(4,"Sai") // value in map but not key so (3, "Sai") is replaced by (4, "Sai")

	assert.Nil(t, m.FindByValue("Vivek"), "Pair (1,\"Vivek\") should not be in BiMap.")
	assert.Equal(t, "Vaiveek", m.FindByKey(1), "Pair (1,\"Vaiveek\") should be in BiMap.")
	assert.Nil(t, m.FindByKey(3), "Pair (3,\"Sai\") should not be in BiMap.")
	assert.Equal(t, 4, m.FindByValue("Sai"), "Pair (4,\"Sai\") should be in BiMap.")

	// Test Remove Functions
	m.RemoveByKey(2)
	m.RemoveByValue("Sai")
	
	assert.Nil(t, m.FindByValue("Mogili"), "Pair (2,\"Mogili\") should not be in BiMap.")
	assert.Nil(t, m.FindByKey(4), "Pair (4,\"Sai\") should not be in BiMap.")
}
