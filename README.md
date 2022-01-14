# The demo for usage of [ent proto](https://github.com/pigfall/contrib)
We will create three schema for demo

# Init ent schema
```
go run entgo.io/ent/cmd/ent init User
go run entgo.io/ent/cmd/ent init Dept
go run entgo.io/ent/cmd/ent init Car

the schema releations:
user m2m dept
user o2m car
```

Replace go mod  change go.mod file
```
# Before we use entproto, we need to change to the forked repo
# in go.mod

replace entgo.io/contrib => github.com/pigfall/contrib v0.1.13
```


# Tag proto info to schema

look at the source file which located at file ent/user.go. Below  is the main code snip in ent/user.go file. please go to the source file ent/user.go to see full code
```
func (User) Fields() []ent.Field {
	return []ent.Field {
		field.String("id").// the primary key
		StorageKey(USER_ID). // override the primary key from id to  user_id, the primary column in db will be user_id
			NotEmpty(). 
			MaxLen(100).
			Annotations(
				entproto.Field(
					1, // will map to proto file 1
					entproto.FieldName(USER_ID), // will map to proto field with name user_id
				),
			), // add out proto annotation

			field.String("phone").
			NotEmpty().
			MaxLen(100).
			Optional(). // optional field
			Comment("手机号"). // comment
			Annotations(
				entproto.Field(
					2, // if without the entproto.FieldName option, the mapped proto field name will be the same as schema field. 
				),
			),

			// 人种
			field.Enum("race").
				Values( // white, black,yellow
						RACE_YELLOW,
						RACE_WHITE,
						RACE_BLACK,
				).
				Annotations(
					entproto.Field(
						3,
					),
					entproto.Enum( // map to proto enum
						map[string]int32{
							RACE_YELLOW:1,
							RACE_WHITE:2,
							RACE_BLACK:3,
						},
					),
				),
	}
}
```


# Modify ent/generate.go file, add the following code
```
//go:generate bash -c "OPTION_GO_PACKAGE=github.com/pigfall/pf_ent_proto_demo/api  go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema -targetPath ../api"

// the OPTION_GO_PACKAGE is the proto package name 
// the targetPath will be the proto file saved path. the working directory is the ent , so the releative path will be ent/../api
```

# Generate the schema and proto
```
go generate ./...
```
check the generated file in ent and api folder

