package main

import (
	"context"
	fields "github.com/gbaptista/requested-fields"
	"fmt"
	"log"
)

type Query struct {
	Field fields.Field `graphql:"query"`
}

type UserResolver struct {
	Field fields.Field `graphql:"user"`
}

type AddressResolver struct {
	Field fields.Field `graphql:"address"`
}

func (queryResolver *Query) User(ctx context.Context) *UserResolver {
	userResolver := &UserResolver{}
	userResolver.Field.SetParent(queryResolver)
	log.Println(fmt.Sprintf(
		"Query.User Fields: %v", fields.RequestedFor(ctx, userResolver)))

	return userResolver
}

func (userResolver *UserResolver) Name(ctx context.Context) *string {
	name := "Harry Potter"

	return &name
}

func (userResolver *UserResolver) Address(ctx context.Context) *AddressResolver {
	addressResolver := &AddressResolver{}
	addressResolver.Field.SetParent(userResolver)

	log.Println(fmt.Sprintf(
		"User.Address Fields: %v", fields.RequestedFor(ctx, addressResolver)))

	return addressResolver
}

func (addressResolver *AddressResolver) City(ctx context.Context) *string {
	city := "Little Whinging"

	return &city
}

func (addressResolver *AddressResolver) Street(ctx context.Context) *string {
	street := "4 Privet Drive"

	return &street
}
