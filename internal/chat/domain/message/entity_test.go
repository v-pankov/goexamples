package message

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/goexamples/internal/chat/domain/room"
	"github.com/vdrpkv/goexamples/internal/chat/domain/user"
)

func TestValidateID(t *testing.T) {
	validID := ID("123")
	require.NoError(t, validID.Validate(), "valid ID must be validated successfuly")

	emptyID := ID("")
	require.ErrorIs(t, emptyID.Validate(), ErrEmptyID, "empty ID must fail validation check")
}

func TestValidateMessage(t *testing.T) {
	var (
		validID     = ID("123")
		emptyID     = ID("")
		validUserID = user.ID("123")
		emptyUserID = user.ID("")
		validRoomID = room.ID("123")
		emptyRoomID = room.ID("")
	)

	for _, testCase := range []struct {
		name string
		give Entity
		want error
	}{
		{
			"ok",
			Entity{
				ID:     validID,
				UserID: validUserID,
				RoomID: validRoomID,
			},
			nil,
		},
		{
			"invalid id",
			Entity{
				ID:     emptyID,
				UserID: validUserID,
				RoomID: validRoomID,
			},
			ErrEmptyID,
		},
		{
			"invalid id",
			Entity{
				ID:     emptyID,
				UserID: emptyUserID,
				RoomID: validRoomID,
			},
			ErrEmptyID,
		},
		{
			"invalid id",
			Entity{
				ID:     emptyID,
				UserID: validUserID,
				RoomID: emptyRoomID,
			},
			ErrEmptyID,
		},
		{
			"invalid id",
			Entity{
				ID:     emptyID,
				UserID: emptyUserID,
				RoomID: emptyRoomID,
			},
			ErrEmptyID,
		},
		{
			"invalid user id",
			Entity{
				ID:     validID,
				UserID: emptyUserID,
				RoomID: validRoomID,
			},
			user.ErrEmptyID,
		},
		{
			"invalid user id",
			Entity{
				ID:     validID,
				UserID: emptyUserID,
				RoomID: emptyRoomID,
			},
			user.ErrEmptyID,
		},
		{
			"invalid room id",
			Entity{
				ID:     validID,
				UserID: validUserID,
				RoomID: emptyRoomID,
			},
			room.ErrEmptyID,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.give.Validate()
			if testCase.want == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, testCase.want)
			}
		})
	}
}
