package orders31

import (
	"context"
	"github.com/google/uuid"
	"github.com/kirigaikabuto/common-lib31"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type ordersStore struct {
	collection *mongo.Collection
}

func NewOrdersStore(config common.MongoConfig) (OrdersStore, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.Database)
	err = db.CreateCollection(context.TODO(), config.CollectionName)
	if err != nil && !strings.Contains(err.Error(), "NamespaceExists") {
		return nil, err
	}
	collection := db.Collection(config.CollectionName)
	return &ordersStore{collection: collection}, nil
}

func (o *ordersStore) Create(order *Order) (*Order, error) {
	token := uuid.New()
	order.Id = token.String()
	_, err := o.collection.InsertOne(context.TODO(), order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *ordersStore) ListOrdersByUserId(userId string) ([]Order, error) {
	filter := bson.D{{"userid", userId}}
	orders := []Order{}
	cursor, err := o.collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
