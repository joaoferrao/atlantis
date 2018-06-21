package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
)

func AnyPtrToEventsCommand() *events.CommentCommand {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*events.CommentCommand))(nil)).Elem()))
	var nullValue *events.CommentCommand
	return nullValue
}

func EqPtrToEventsCommand(value *events.CommentCommand) *events.CommentCommand {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *events.CommentCommand
	return nullValue
}
