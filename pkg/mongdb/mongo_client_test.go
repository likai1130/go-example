package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

const (
	MONGODB_DATABASE = "nftdisk"
	MONGODB_DATABASE_COLLECT = "users"
)

type UserData struct {
	Id         string `bson:"_id,omitempty" json:"id"`
	Name       string `bson:"name" json:"name"`
	Number     int    `bson:"number" json:"number"`
	Age        int    `bson:"age" json:"age"`
	BirthMonth int    `bson:"birthMonth" json:"birthMonth"`
}

var cli *mongo.Client

func init() {
	clientOptions := options.Client().SetAuth(options.Credential{Username: "mongoadmin", Password: "secret"})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("init databases fail!")
	}
	cli = client
}

//增
func TestInsertUser(t *testing.T){
	//cli, err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	userModel := UserData{
		Name:       "kli",
		Number:     152664,
		Age:        18,
		BirthMonth: 1852,
	}
	insertOneResult, err := coll.InsertOne(ctx, userModel)
	if err != nil {
		log.Fatalf("Insert user data is error : %s",err.Error())
		return
	}
	log.Printf("Insert user data is success! id = %v\n",insertOneResult.InsertedID)
	docId := insertOneResult.InsertedID.(primitive.ObjectID)
	log.Printf("insert one ID str is : %s\n", docId.String())
}

func TestInsertMany(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli,err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)
	userDatas := []UserData{
		UserData{
			Name:       "kli1",
			Number:     152664,
			Age:        17,
			BirthMonth: 1852,
		},
		UserData{
			Name:       "kli2",
			Number:     152665,
			Age:        18,
			BirthMonth: 1853,
		},
		UserData{
			Name:       "kli3",
			Number:     152666,
			Age:        19,
			BirthMonth: 1854,
		},
	}
	var list []interface{}
	marshal, _ := json.Marshal(userDatas)
	json.Unmarshal(marshal,&list)
	many, err := coll.InsertMany(ctx, list)
	if err != nil {
		log.Fatalf("Insert many user data is error : %s",err.Error())
		return
	}
	log.Printf("Insert many user data is success! ids = %v\n",many.InsertedIDs)

}

//删
func TestDeleteUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli,err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)
	filter := bson.M{"birthMonth": bson.M{"$lte": 3}}
	one, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatalf("Delete one user fail. err is %s",err.Error())
		return
	}
	log.Printf("Delete user success ! count is %d \n",one.DeletedCount)
}

func TestDeleteUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli,err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)
	filter := bson.M{"age": bson.M{"$lte": 18}}
	one, err := coll.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatalf("Delete many users fail. err is %s",err.Error())
		return
	}
	log.Printf("Delete users success ! count is %d \n",one.DeletedCount)
}

//改
func TestUpdateUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli,err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)

	filter := bson.M{"name": "kli"}
	value := bson.M{"$set": bson.M{"number": 1024}} //操作

	_, err = coll.UpdateOne(ctx, filter, value)
	if err != nil {
		log.Fatalf("Update one user is fail,err is %s",err.Error())
		return
	}
	return
}

func TestUpdateUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli,err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)

	names := []string{"kli", "kli1"}
	filter := bson.M{"name": bson.M{"$in": names}}
	value := bson.M{"$set": bson.M{"birthMonth":3}}
	result, err := coll.UpdateMany(ctx, filter, value)
	if err != nil {
		log.Fatalf("Update many users is fail,err is %s",err.Error())
		return
	}
	log.Println("Update many user success !, result is ", result)
	return
}

//查
func TestFindUser(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli,err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)
	//过滤条件
	filter := bson.M{"name": "kli"}
	singleResult := coll.FindOne(ctx, filter)
	if singleResult == nil || singleResult.Err() != nil {
		log.Fatalf("Find one user is error")
		if singleResult.Err() != nil {
			log.Fatalf(singleResult.Err().Error())
		}
		return
	}
	userData := &UserData{}
	err = singleResult.Decode(userData)
	if err != nil {
		log.Fatalf("Find one user is error:%s\n",err.Error())
		return
	}
	marshal, _ := json.Marshal(userData)
	log.Printf("Find one user is success!\n %s",string(marshal))
}

func TestFindUsers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli,err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)

	//过滤条件
	filter := bson.M{"birthMonth": bson.M{"$gte": 0}}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		log.Fatalf("Find users is error: %s",err.Error())
		return
	}
	defer cursor.Close(ctx)

	list := []UserData{}
	if err = cursor.All(ctx, &list); err != nil{
		log.Fatalf("Find users is error: %s",err.Error())
		return
	}
	marshal, _ := json.Marshal(list)
	log.Printf("Find users is success!\n %s",string(marshal))
}

//分组查询
func TestFindGroup(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli,err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)

	//复杂查询，先匹配后分组
	pipline := bson.A{
		bson.M{
			"$match": bson.M{"birthMonth": 1854},
		},
		bson.M{"$group":bson.M{
			"_id":        bson.M{"birthMonthUid": "$birthMonth"},
			"totalCount": bson.M{"$sum": 1},
			"nameG":      bson.M{"$min": "$name"},
			"ageG":       bson.M{"$min": "$age"},
		}},
	}
	log.Printf("pipline is %v\n",pipline)
	cursor, err := coll.Aggregate(ctx, pipline)
	if err != nil {
		log.Fatalf("dao.findGroup collection.Aggregate() error=[%s]\n", err.Error())
		return
	}

	for cursor.Next(context.Background()) {
		doc := cursor.Current

		totalCount, err_2 := doc.LookupErr("totalCount")
		if err_2 != nil {
			fmt.Printf("dao.findGroup totalCount err_2=[%s]\n", err_2)
			return
		}

		nameData, err_4 := doc.LookupErr("nameG")
		if err_4 != nil {
			fmt.Printf("dao.findGroup insertDateG err_4=[%s]\n", err_4)
			return
		}

		ageData, err_5 := doc.LookupErr("ageG")
		if err_5 != nil {
			fmt.Printf("dao.findGroup ageG err_5=[%s]\n", err_5)
			continue
		}
		fmt.Println("totalCount is ", totalCount)
		fmt.Println("nameData is ", nameData)
		fmt.Println("ageData is ", ageData)
	}

}

//分页查询
func TestLimitPage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), MONGODB_CONNECT_TIMEOUT)
	defer cancel()

	cli,err := NewMongoCliInstance()
	coll := cli.Database(MONGODB_DATABASE).Collection(MONGODB_DATABASE_COLLECT)
	filter := bson.M{"age": bson.M{"$gte": 0}}

	SORT := bson.D{{"number", -1}}
	findOptions := options.Find().SetSort(SORT)
	skipTmp := int64((1 - 1) * 10)
	limitTmp := int64(10)
	findOptions.Skip = &skipTmp
	findOptions.Limit = &limitTmp

	cursor, err := coll.Find(ctx, filter, findOptions)
	defer cursor.Close(ctx)
	if err != nil {
		log.Fatalf("Find limit page error is %s", err.Error())
		return
	}

	list := []UserData{}
	err = cursor.All(ctx, &list)
	if err != nil {
		log.Fatalf("Find limit page error is %s", err.Error())
		return
	}
	marshal, _ := json.Marshal(list)
	log.Printf("Find limit is success!\n %s",string(marshal))
}