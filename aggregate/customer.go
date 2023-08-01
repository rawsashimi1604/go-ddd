// package aggregate holds our aggregates that combines many entities into a full object
package aggregate

import (
	"github.com/rawsashimi1604/go-ddd/entity"
	"github.com/rawsashimi1604/go-ddd/valueobject"
)

type Customer struct {
	// Person is the root entity of customer
	// which means person.ID is the main identifier for the customer

	// Small case -> external cannot use.
	// Not accessible directly to the other data. Not accessible from outside.
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}
