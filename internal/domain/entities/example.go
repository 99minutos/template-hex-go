package entities

type SubExample struct {
	SubExampleId   int    `json:"sub_example_id" bson:"sub_example_id"`
	SubExampleName string `json:"sub_example_name" bson:"sub_example_name"`
}
type Example struct {
	Id         string     `json:"id" bson:"_id,omitempty"`
	FirstName  string     `json:"first_name" bson:"first_name"`
	LastName   string     `json:"last_name" bson:"last_name"`
	SubExample SubExample `json:"sub_example" bson:"sub_example"`
}
