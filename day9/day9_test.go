package day9

import (
	"testing"
)

func Test_NewGameStr(t *testing.T) {
	input := "10 players; last marble is worth 1618 points"
	game := NewGameStr(input)

	if len(game.players) != 10 {
		t.Errorf("Expected players = 10, got %d", len(game.players))
	}

	if game.LastMarble != 1618 {
		t.Errorf("Expected players = 1618, got %d", game.LastMarble)
	}
}

func Test_getIndex(t *testing.T) {
	circle := []int{0, 2, 1, 3,}
	place := getIndex(len(circle), 3, 2)

	if place != 1 {
		t.Errorf("Expected place = 1, got %d", place)
	}
}

// GetHighScore tests
func testGetHighScore(game Game, highScore int, t *testing.T) {
	hs := GetHighScore(game)

	if hs != highScore {
		t.Errorf("Expected highScore = %d, got %d", highScore, hs)
	}
}

func Test_GetHighScore32(t *testing.T) {
	testGetHighScore(NewGameParams(9, 32), 32, t)
}

func Test_GetHighScore1618(t *testing.T) {
	testGetHighScore(NewGameParams(10, 1618), 8317, t)
}

func Test_GetHighScore7999(t *testing.T) {
	testGetHighScore(NewGameParams(13, 7999), 146373, t)
}

func Test_GetHighScore1104(t *testing.T) {
	testGetHighScore(NewGameParams(17, 1104), 2764, t)
}

func Test_GetHighScore6111(t *testing.T) {
	testGetHighScore(NewGameParams(21, 6111), 54718, t)
}

func Test_GetHighScore5807(t *testing.T) {
	testGetHighScore(NewGameParams(30, 5807), 37305, t)
}

// GetHighScoreFast tests
func testGetHighScoreFast(game Game, highScore int, t *testing.T) {
	hs := GetHighScoreFast(game)

	if hs != highScore {
		t.Errorf("Expected highScore = %d, got %d", highScore, hs)
	}
}

func Test_GetHighScoreFast32(t *testing.T) {
	testGetHighScoreFast(NewGameParams(9, 32), 32, t)
}

func Test_GetHighScoreFast1618(t *testing.T) {
	testGetHighScoreFast(NewGameParams(10, 1618), 8317, t)
}

func Test_GetHighScoreFast7999(t *testing.T) {
	testGetHighScoreFast(NewGameParams(13, 7999), 146373, t)
}

func Test_GetHighScoreFast1104(t *testing.T) {
	testGetHighScoreFast(NewGameParams(17, 1104), 2764, t)
}

func Test_GetHighScoreFast6111(t *testing.T) {
	testGetHighScoreFast(NewGameParams(21, 6111), 54718, t)
}

func Test_GetHighScoreFast5807(t *testing.T) {
	testGetHighScoreFast(NewGameParams(30, 5807), 37305, t)
}

func Test_GetHighScoreFastInput(t *testing.T) {
	testGetHighScoreFast(NewGameStr(Input), 408679, t)
}

func Test_GetHighScoreFast100Input(t *testing.T) {
	game := NewGameStr(Input)
	game.LastMarble *= 100

	testGetHighScoreFast(game, 3443939356, t)
}

// Benchmarks
func benchmarkGetHighScore(game Game, b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetHighScore(game)
	}
}

func benchmarkGetHighScoreFast(game Game, b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetHighScoreFast(game)
	}
}

func BenchmarkGetHighScoreInput(b *testing.B) {
	benchmarkGetHighScore(NewGameStr(Input), b)
}

func BenchmarkGetHighScoreFastInput(b *testing.B) {
	benchmarkGetHighScoreFast(NewGameStr(Input), b)
}

func BenchmarkGetHighScore2Input(b *testing.B) {
	game := NewGameStr(Input)
	game.LastMarble *= 2

	benchmarkGetHighScore(game, b)
}

func BenchmarkGetHighScoreFast2Input(b *testing.B) {
	game := NewGameStr(Input)
	game.LastMarble *= 2

	benchmarkGetHighScoreFast(game, b)
}

func BenchmarkGetHighScoreFast100Input(b *testing.B) {
	game := NewGameStr(Input)
	game.LastMarble *= 100

	benchmarkGetHighScoreFast(game, b)
}
