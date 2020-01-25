package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}

	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/players/Pepper", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatusCode(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/players/Floyd", nil)
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatusCode(t, resp.Code, http.StatusOK)
		assertResponseBody(t, resp.Body.String(), "10")

	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		req := newGetScoreRequest("Apollo")
		resp := httptest.NewRecorder()

		server.ServeHTTP(resp, req)

		assertStatusCode(t, resp.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
	}

	server := &PlayerServer{&store}

	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"

		req := newPostWinRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatusCode(t, res.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
		}
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatusCode(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest("POST", "/players/"+name, nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest("GET", "/players/"+name, nil)
	return req
}

func assertStatusCode(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got status code %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
