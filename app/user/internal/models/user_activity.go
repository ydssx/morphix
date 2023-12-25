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

// 获取用户活动记录
func (m *UserActivityModel) Get(userID string, startTime int64, endTime int64) ([]*UserActivity, error) {
	var activities []*UserActivity
	cursor, err := m.coll.Find(context.TODO(), bson.M{
		"user_id": userID,
		"action_time": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var activity UserActivity
		err := cursor.Decode(&activity)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &activity)
	}
	return activities, nil
}

func (m *UserActivityModel) GetByResource(userID string, startTime int64, endTime int64, resource string) ([]*UserActivity, error) {
	var activities []*UserActivity
	cursor, err := m.coll.Find(context.TODO(), bson.M{
		"user_id": userID,
		"action_time": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
		"resource": resource,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var activity UserActivity
		err := cursor.Decode(&activity)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &activity)
	}
	return activities, nil
}

func (m *UserActivityModel) GetByAction(userID string, startTime int64, endTime int64, action string) ([]*UserActivity, error) {
	var activities []*UserActivity
	cursor, err := m.coll.Find(context.TODO(), bson.M{
		"user_id": userID,
		"action_time": bson.M{
			"$gte": startTime,
			"$lte": endTime,
		},
		"action": action,
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var activity UserActivity
		err := cursor.Decode(&activity)
		if err != nil {
			return nil, err
		}
		activities = append(activities, &activity)
	}
	return activities, nil
}
