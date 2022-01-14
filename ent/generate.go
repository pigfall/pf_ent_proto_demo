package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
//go:generate bash -c "OPTION_GO_PACKAGE=github.com/pigfall/pf_ent_proto_demo/api  go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema -targetPath ../api"
