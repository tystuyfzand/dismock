package dismock

import (
	"testing"

	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMocker_React(t *testing.T) {
	m, s := NewSession(t)

	var (
		channelID discord.ChannelID = 123
		messageID discord.MessageID = 456
		emoji                       = "🍆"
	)

	m.React(channelID, messageID, emoji)

	err := s.React(channelID, messageID, emoji)
	require.NoError(t, err)

	m.Eval()
}

func TestMocker_Unreact(t *testing.T) {
	m, s := NewSession(t)

	var (
		channelID discord.ChannelID = 123
		messageID discord.MessageID = 456
		emoji                       = "🍆"
	)

	m.Unreact(channelID, messageID, emoji)

	err := s.Unreact(channelID, messageID, emoji)
	require.NoError(t, err)

	m.Eval()
}

func TestMocker_Reactions(t *testing.T) {
	successCases := []struct {
		name      string
		reactions int
		limit     uint
	}{
		{
			name:      "limited",
			reactions: 130,
			limit:     199,
		},
		{
			name:      "unlimited",
			reactions: 200,
			limit:     0,
		},
	}

	t.Run("success", func(t *testing.T) {
		for _, c := range successCases {
			t.Run(c.name, func(t *testing.T) {
				m, s := NewSession(t)

				var (
					channelID discord.ChannelID = 123
					messageID discord.MessageID = 456
					emoji                       = "🍆"
				)

				expect := make([]discord.User, c.reactions)

				for i := 1; i < c.reactions+1; i++ {
					expect[i-1] = discord.User{ID: discord.UserID(i)}
				}

				m.Reactions(channelID, messageID, c.limit, emoji, expect)

				actual, err := s.Reactions(channelID, messageID, emoji, c.limit)
				require.NoError(t, err)

				assert.Equal(t, expect, actual)

				m.Eval()
			})
		}
	})

	t.Run("nil users", func(t *testing.T) {
		m, s := NewSession(t)

		var (
			channelID discord.ChannelID = 123
			messageID discord.MessageID = 456
			emoji                       = "🍆"
		)

		m.Reactions(channelID, messageID, 100, emoji, nil)

		actual, err := s.Reactions(channelID, messageID, emoji, 100)
		require.NoError(t, err)

		assert.Len(t, actual, 0)

		m.Eval()
	})

	t.Run("limit smaller than guilds", func(t *testing.T) {
		m, _ := NewSession(t)

		assert.Panics(t, func() {
			m.Reactions(123, 456, 1, "abc", []discord.User{{}, {}})
		})
	})
}

func TestMocker_ReactionsBefore(t *testing.T) {
	successCases := []struct {
		name      string
		reactions int
		limit     uint
	}{
		{
			name:      "limited",
			reactions: 130,
			limit:     199,
		},
		{
			name:      "unlimited",
			reactions: 200,
			limit:     0,
		},
	}

	t.Run("success", func(t *testing.T) {
		for _, c := range successCases {
			t.Run(c.name, func(t *testing.T) {
				m, s := NewSession(t)

				var (
					channelID discord.ChannelID = 123
					messageID discord.MessageID = 456
					emoji                       = "🍆"

					before discord.UserID = 999999999999
				)

				expect := make([]discord.User, c.reactions)

				for i := 1; i < c.reactions+1; i++ {
					expect[i-1] = discord.User{ID: discord.UserID(i)}
				}

				m.ReactionsBefore(channelID, messageID, before, c.limit, emoji, expect)

				actual, err := s.ReactionsBefore(channelID, messageID, before, emoji, c.limit)
				require.NoError(t, err)

				assert.Equal(t, expect, actual)

				m.Eval()
			})
		}
	})

	t.Run("nil users", func(t *testing.T) {
		m, s := NewSession(t)

		var (
			channelID discord.ChannelID = 123
			messageID discord.MessageID = 456
			emoji                       = "🍆"
		)

		m.ReactionsBefore(channelID, messageID, 0, 100, emoji, nil)

		actual, err := s.ReactionsBefore(channelID, messageID, 0, emoji, 100)
		require.NoError(t, err)

		assert.Len(t, actual, 0)

		m.Eval()
	})

	t.Run("failure", func(t *testing.T) {
		tMock := new(testing.T)

		m, s := NewSession(tMock)

		var (
			channelID discord.ChannelID = 123
			messageID discord.MessageID = 456
			emoji                       = "abc"
		)

		expect := []discord.User{
			{
				ID: 123,
			},
			{
				ID: 456,
			},
		}

		m.ReactionsBefore(channelID, messageID, 890, 100, emoji, expect)

		actual, err := s.ReactionsBefore(channelID, messageID, 789, emoji, 100)
		require.NoError(t, err)

		assert.Equal(t, expect, actual)
		assert.True(t, tMock.Failed())
	})

	t.Run("limit smaller than guilds", func(t *testing.T) {
		m, _ := NewSession(t)

		assert.Panics(t, func() {
			m.ReactionsBefore(123, 456, 0, 1, "abc", []discord.User{{}, {}})
		})
	})
}

func TestMocker_ReactionsAfter(t *testing.T) {
	successCases := []struct {
		name      string
		reactions int
		limit     uint
	}{
		{
			name:      "limited",
			reactions: 130,
			limit:     199,
		},
		{
			name:      "unlimited",
			reactions: 200,
			limit:     0,
		},
	}

	t.Run("success", func(t *testing.T) {
		for _, c := range successCases {
			t.Run(c.name, func(t *testing.T) {
				m, s := NewSession(t)

				var (
					channelID discord.ChannelID = 123
					messageID discord.MessageID = 456
					emoji                       = "🍆"

					after discord.UserID = 123
				)

				expect := make([]discord.User, c.reactions)

				for i := int(after) + 1; i < c.reactions+int(after)+1; i++ {
					expect[i-int(after)-1] = discord.User{ID: discord.UserID(i)}
				}

				m.ReactionsAfter(channelID, messageID, after, c.limit, emoji, expect)

				actual, err := s.ReactionsAfter(channelID, messageID, after, emoji, c.limit)
				require.NoError(t, err)

				assert.Equal(t, expect, actual)

				m.Eval()
			})
		}
	})

	t.Run("nil users", func(t *testing.T) {
		m, s := NewSession(t)

		var (
			channelID discord.ChannelID = 123
			messageID discord.MessageID = 456
			emoji                       = "🍆"
		)

		m.ReactionsAfter(channelID, messageID, 0, 100, emoji, nil)

		actual, err := s.ReactionsAfter(channelID, messageID, 0, emoji, 100)
		require.NoError(t, err)

		assert.Len(t, actual, 0)

		m.Eval()
	})

	t.Run("failure", func(t *testing.T) {
		tMock := new(testing.T)

		m, s := NewSession(tMock)

		var (
			channelID discord.ChannelID = 123
			messageID discord.MessageID = 456
			emoji                       = "🍆"
		)

		expect := []discord.User{
			{
				ID: 456,
			},
			{
				ID: 789,
			},
		}

		m.ReactionsAfter(channelID, messageID, 123, 100, emoji, expect)

		actual, err := s.ReactionsAfter(channelID, messageID, 321, emoji, 100)
		require.NoError(t, err)

		assert.Equal(t, expect, actual)
		assert.True(t, tMock.Failed())
	})

	t.Run("limit smaller than guilds", func(t *testing.T) {
		m, _ := NewSession(t)

		assert.Panics(t, func() {
			m.ReactionsAfter(123, 456, 0, 1, "abc", []discord.User{{}, {}})
		})
	})
}

func TestMocker_DeleteUserReaction(t *testing.T) {
	m, s := NewSession(t)

	var (
		channelID discord.ChannelID = 123
		messageID discord.MessageID = 456
		userID    discord.UserID    = 789
		emoji                       = "🍆"
	)

	m.DeleteUserReaction(channelID, messageID, userID, emoji)

	err := s.DeleteUserReaction(channelID, messageID, userID, emoji)
	require.NoError(t, err)

	m.Eval()
}

func TestMocker_DeleteReactions(t *testing.T) {
	m, s := NewSession(t)

	var (
		channelID discord.ChannelID = 123
		messageID discord.MessageID = 456
		emoji                       = "🍆"
	)

	m.DeleteReactions(channelID, messageID, emoji)

	err := s.DeleteReactions(channelID, messageID, emoji)
	require.NoError(t, err)

	m.Eval()
}

func TestMocker_DeleteAllReactions(t *testing.T) {
	m, s := NewSession(t)

	var (
		channelID discord.ChannelID = 123
		messageID discord.MessageID = 456
	)

	m.DeleteAllReactions(channelID, messageID)

	err := s.DeleteAllReactions(channelID, messageID)
	require.NoError(t, err)

	m.Eval()
}
