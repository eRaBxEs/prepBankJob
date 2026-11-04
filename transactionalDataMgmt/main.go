package main

import (
	"fmt"
)

/*GlobalStore holds the (global) variables*/
var GlobalStore = make(map[string]string)

/* Transaction points to a key-value store*/
type Transaction struct {
	store map[string]string // every transaction has it's own local store
	next  *Transaction
}

/*TransactionStack maintains a list of active/suspended transactions */
type TransactionStack struct {
	top  *Transaction
	size int // more meta data can be saved like Stack limit etc.
}

/*PushTransaction creates a new active transaction*/
func (ts *TransactionStack) PushTransaction() {
	// push a new Transaction, this is the current active transaction
	temp := Transaction{store: make(map[string]string)}
	temp.next = ts.top
	ts.top = &temp
	ts.size++
}

/*PopTransaction deletes a transaction from stack*/
func (ts *TransactionStack) PopTransaction() {
	// pop the Transaction from stack, no longer active
	if ts.top == nil {
		// basically stack underflow
		fmt.Printf("ERROR: No Active Transactions\n")
	} else {
		node := &Transaction{}
		ts.top = ts.top.next
		node.next = nil
		ts.size--
	}
}

/*Peek returns the active transaction*/
func (ts *TransactionStack) Peek() *Transaction {
	return ts.top
}
