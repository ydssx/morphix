package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserActivity struct {
	ID         primitive.ObjectID `bson:"_id"`
	UserID     int64              `bson:"user_id"`
	Action     string             `bson:"action"`
	ActionTime int64              `bson:"action_time"`
	ActionIp   string             `bson:"action_ip"`
	Resource   string             `bson:"resource"`
}

type UserActivityModel struct {
	coll *mongo.Collection
}

func NewUserActivityModel(coll *mongo.Collection) *UserActivityModel {
	return &UserActivityModel{coll: coll}
}

// 创建用户活动记录
func (m *UserActivityModel) Create(activity *UserActivity) error {
	_, err := m.coll.InsertOne(context.TODO(), activity)
	return err
}

type UserActivityCond struct {
	StartTime int64
	EndTime   int64
	Resource  string
	Action    string
}

// Get retrieves user activity records based on the given conditions.
func (m *UserActivityModel) Get(userID int64, cond *UserActivityCond) ([]*UserActivity, error) {
	var activities []*UserActivity

	filter := bson.M{"user_id": userID}

	if cond.StartTime != 0 {
		filter["action_time"] = bson.M{
			"$gte": cond.StartTime,
		}
	}

	if cond.EndTime != 0 {
		if filter["action_time"] == nil {
			filter["action_time"] = bson.M{
				"$lte": cond.EndTime,
			}
		} else {
			filter["action_time"].(bson.M)["$lte"] = cond.EndTime
		}
	}

	if cond.Resource != "" {
		filter["resource"] = cond.Resource
	}

	if cond.Action != "" {
		filter["action"] = cond.Action
	}

	cursor, err := m.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var activity UserActivity
		if err := cursor.Decode(&activity); err != nil {
			return nil, err
		}
		activities = append(activities, &activity)
	}

	return activities, nil
}

func (m *UserActivityModel) DeleteByUser(userID string) error {
	_, err := m.coll.DeleteMany(context.TODO(), bson.M{"user_id": userID})
	return err
}

func (m *UserActivityModel) DeleteByResource(userID string, resource string) error {
	_, err := m.coll.DeleteMany(context.TODO(), bson.M{"user_id": userID, "resource": resource})
	return err
}

func (m *UserActivityModel) DeleteByAction(userID string, action string) error {
	_, err := m.coll.DeleteMany(context.TODO(), bson.M{"user_id": userID, "action": action})
	return err
}
