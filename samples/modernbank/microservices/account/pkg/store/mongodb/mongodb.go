package mongodb

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-openapi/swag"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/tetrateio/training/samples/modernbank/microservices/account/pkg/model"
	"github.com/tetrateio/training/samples/modernbank/microservices/account/pkg/store"
)

var (
	// Enforce that MongoDB matches the Store Interface
	_ store.Interface = NewMongoDB()

	defaultAddress    = "mongodb://account-mongodb:27017"
	defaultDatabase   = "accounts"
	defaultCollection = "accounts"

	ctx = context.Background()

	randomAccountNumber = func() int64 {
		// up to 15 digit account numbers
		return rand.Int63n(999999999999999)
	}
)

func NewMongoDB() *MongoDB {
	rand.Seed(time.Now().UnixNano())
	client, _ := mongo.NewClient(defaultAddress)
	for i := 1; i < 360; i += 5 {
		time.Sleep(5 * time.Second)
		log.Printf("attempting to connect to mongodb at %v", defaultAddress)
		if err := client.Connect(ctx); err != nil {
			log.Printf("unable to connect to mongodb: %v", err)
		}
		err := client.Ping(ctx, nil)
		if err == nil {
			break
		}
		log.Printf("unable to ping mongodb: %v", err)
	}
	return &MongoDB{client: client}
}

type MongoDB struct {
	client *mongo.Client
}

func (m *MongoDB) List(owner string) ([]*model.Account, error) {
	accounts := []*model.Account{}
	res, err := m.defaultCollection().Find(ctx, bson.M{"owner": owner})
	if err != nil {
		return nil, fmt.Errorf("unable to get accounts in database: %v", err)
	}
	defer res.Close(ctx)
	for res.Next(ctx) {
		var account model.Account
		if err := res.Decode(&account); err != nil {
			return nil, fmt.Errorf("unable to decode response: %v", err)
		}
		accounts = append(accounts, &account)
	}
	if len(accounts) == 0 {
		return accounts, &store.NotFound{}
	}
	return accounts, nil
}

func (m *MongoDB) Get(owner string, number int64) (*model.Account, error) {
	var account model.Account
	res := m.defaultCollection().FindOne(ctx, bson.M{"owner": owner, "number": number})
	if res.Err().Error() == mongo.ErrNoDocuments.Error() {
		return nil, &store.NotFound{}
	} else if res.Err() != nil {
		return nil, fmt.Errorf("unable to get account in database: %v", res.Err())
	}
	return &account, res.Decode(&account)
}

func (m *MongoDB) Create(owner string) (*model.Account, error) {
	newAccountNumber, err := m.unAssignedAccountNumber()
	if err != nil {
		return nil, fmt.Errorf("error finding a vacant account number: %v", err)
	}
	newAccount := &model.Account{
		Balance: swag.Float64(0),
		Owner:   swag.String(owner),
		Number:  swag.Int64(newAccountNumber),
	}
	_, err = m.defaultCollection().InsertOne(ctx, newAccount)
	if err != nil {
		return nil, fmt.Errorf("error creating account in database: %v", err)
	}
	return newAccount, nil
}

// Not concurrency safe but close enough for a demo app
// Clashes are highly unlikely
func (m *MongoDB) unAssignedAccountNumber() (int64, error) {
	var err error
	candidate, count := int64(0), int64(0)
	for i := 0; i < 10; i++ {
		candidate = randomAccountNumber()
		count, err = m.defaultCollection().CountDocuments(ctx, bson.M{"number": candidate})
		if count == 0 {
			return candidate, nil
		}
	}
	return 0, err
}

func (m *MongoDB) Delete(owner string, number int64) error {
	res := m.defaultCollection().FindOneAndDelete(ctx, bson.M{"owner": owner, "number": number})
	if res.Err().Error() == mongo.ErrNoDocuments.Error() {
		return &store.NotFound{}
	} else if res.Err() != nil {
		return fmt.Errorf("unable to delete account in database: %v", res.Err())
	}
	return nil
}

func (m *MongoDB) UpdateBalance(number int64, deltaAmount float64) error {
	res, err := m.defaultCollection().UpdateOne(ctx, bson.M{"number": number}, bson.D{{"$inc", bson.D{{"amount", deltaAmount}}}})
	if res.ModifiedCount != 1 {
		return &store.NotFound{}
	} else if err != nil {
		return fmt.Errorf("unable to update account in database: %v", err)
	}
	return nil
}

func (m *MongoDB) defaultCollection() *mongo.Collection {
	return m.client.Database(defaultDatabase).Collection(defaultCollection)
}
