package server

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"talpidae-backend/model/game"
	"talpidae-backend/model/user"
	"talpidae-backend/storage"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGameStart(t *testing.T) {
	cases := []struct {
		name          string
		url           string
		outStatusCode int
	}{
		{
			name:          "success",
			outStatusCode: 200,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gameStorage := &storage.GameStorageMock{}
			router := mux.NewRouter()
			attachHandlers(router, gameStorage)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/game/start", nil)
			router.ServeHTTP(rec, req)

			assert.Equal(t, c.outStatusCode, rec.Result().StatusCode)
		})
	}
}

func TestGameField(t *testing.T) {
	cases := []struct {
		name              string
		beforeGameStorage storage.GameStorage
		outStatusCode     int
		outBody           string
	}{
		{
			name: "success",
			beforeGameStorage: &storage.GameStorageMock{
				GameMock: &game.GameMock{BlocksMock: [][]game.BlockType{{game.Treasure, game.SakuSaku}, {game.SakuSaku, game.SakuSaku}}},
			},
			outStatusCode: 200,
			outBody:       `{"field":[[3,0],[0,0]],"current_number_of_users":0}`,
		},
		{
			name:              "failure",
			beforeGameStorage: &storage.GameStorageMock{ErrorMock: errors.New("")},
			outStatusCode:     404,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			router := mux.NewRouter()
			attachHandlers(router, c.beforeGameStorage)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/game/field", nil)
			router.ServeHTTP(rec, req)

			if rec.Result().StatusCode != 200 {
				assert.Equal(t, c.outStatusCode, rec.Result().StatusCode)
				return
			}

			bytes, _ := ioutil.ReadAll(rec.Result().Body)
			assert.Equal(t, c.outBody, string(bytes))
		})
	}
}

func TestGameFill(t *testing.T) {
	cases := []struct {
		name              string
		beforeGameStorage storage.GameStorage
		inBody            string
		outStatusCode     int
	}{
		{
			name: "success",
			beforeGameStorage: &storage.GameStorageMock{
				GameMock: &game.GameMock{BlocksMock: [][]game.BlockType{{game.SakuSaku, game.SakuSaku}, {game.SakuSaku, game.SakuSaku}}},
			},
			inBody:        `{"h":0,"w":0,"value":3,"user_id":"1769b643-a544-3886-8504-f227ebd35aca"}`,
			outStatusCode: 200,
		},
		{
			name:          "failure: invalid body",
			inBody:        `invalid`,
			outStatusCode: 400,
		},
		{
			name:   "failure: game storage error",
			inBody: `{"h":0,"w":0,"value":3,"user_id":"1769b643-a544-3886-8504-f227ebd35aca"}`,
			beforeGameStorage: &storage.GameStorageMock{
				ErrorMock: errors.New(""),
			},
			outStatusCode: 404,
		},
		{
			name:   "failure: game error",
			inBody: `{"h":0,"w":0,"value":3,"user_id":"1769b643-a544-3886-8504-f227ebd35aca"}`,
			beforeGameStorage: &storage.GameStorageMock{
				GameMock: &game.GameMock{ErrorMock: errors.New("")},
			},
			outStatusCode: 400,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			router := mux.NewRouter()
			attachHandlers(router, c.beforeGameStorage)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/game/fill", bytes.NewBuffer([]byte(c.inBody)))
			router.ServeHTTP(rec, req)

			assert.Equal(t, c.outStatusCode, rec.Result().StatusCode)
		})
	}
}

func TestGameLogs(t *testing.T) {
	cases := []struct {
		name              string
		beforeGameStorage storage.GameStorage
		outStatusCode     int
		outBody           string
	}{
		{
			name: "success",
			beforeGameStorage: &storage.GameStorageMock{
				GameMock: &game.GameMock{LogsMock: []game.FillLog{{UserId: "de8a99d4-80b4-3fdf-9eb3-79bf42f5cc2e", Value: game.Treasure, Height: 0, Width: 0}}},
			},
			outStatusCode: 200,
			outBody:       `{"logs":[{"h":0,"w":0,"value":3,"user_id":"de8a99d4-80b4-3fdf-9eb3-79bf42f5cc2e"}]}`,
		},
		{
			name: "failure",
			beforeGameStorage: &storage.GameStorageMock{
				ErrorMock: errors.New(""),
			},
			outStatusCode: 404,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			router := mux.NewRouter()
			attachHandlers(router, c.beforeGameStorage)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/game/logs", nil)
			router.ServeHTTP(rec, req)

			if rec.Result().StatusCode != 200 {
				assert.Equal(t, c.outStatusCode, rec.Result().StatusCode)
				return
			}

			bytes, _ := ioutil.ReadAll(rec.Result().Body)
			assert.Equal(t, c.outBody, string(bytes))
		})
	}
}

func TestMatchJoin(t *testing.T) {
	testUUID, _ := uuid.Parse("25cc4cd0-2aff-35ed-85e7-0d36eae90966")
	cases := []struct {
		name              string
		beforeGameStorage storage.GameStorage
		inBody            string
		outStatusCode     int
		outBody           string
	}{
		{
			name: "success",
			beforeGameStorage: &storage.GameStorageMock{
				GameMock: &game.GameMock{UsersMock: []*user.User{{}, {}}},
				UUIDMock: testUUID,
			},
			inBody:        `{"id":"a442c1f7-a5b7-3bf8-b280-f3ccc031cca3","name":"Guillermo"}`,
			outStatusCode: 200,
			outBody:       "{\"game_id\":\"25cc4cd0-2aff-35ed-85e7-0d36eae90966\"}",
		},
		{
			name: "failure",
			beforeGameStorage: &storage.GameStorageMock{
				ErrorMock: errors.New(""),
			},
			inBody:        `{"id":"a442c1f7-a5b7-3bf8-b280-f3ccc031cca3","name":"Guillermo"}`,
			outStatusCode: 400,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			router := mux.NewRouter()
			attachHandlers(router, c.beforeGameStorage)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/match/join", bytes.NewBuffer([]byte(c.inBody)))
			router.ServeHTTP(rec, req)

			if rec.Result().StatusCode != 200 {
				assert.Equal(t, c.outStatusCode, rec.Result().StatusCode)
				return
			}

			bytes, _ := ioutil.ReadAll(rec.Result().Body)
			assert.Equal(t, c.outBody, string(bytes))
		})
	}
}

func TestMatchLeave(t *testing.T) {
	cases := []struct {
		name              string
		beforeGameStorage storage.GameStorage
		inBody            string
		outStatusCode     int
	}{
		{
			name: "success",
			beforeGameStorage: &storage.GameStorageMock{
				GameMock: &game.GameMock{UsersMock: []*user.User{{}, {}}},
			},
			inBody:        `{"game_id":"96b7babe-0aab-3c11-902d-db569c194aa5","user_id":"577d04f0-30b7-3074-be0a-63419fe38abb"}`,
			outStatusCode: 200,
		},
		{
			name: "failure",
			beforeGameStorage: &storage.GameStorageMock{
				ErrorMock: errors.New(""),
			},
			inBody:        `{"game_id":"96b7babe-0aab-3c11-902d-db569c194aa5","user_id":"577d04f0-30b7-3074-be0a-63419fe38abb"}`,
			outStatusCode: 400,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			router := mux.NewRouter()
			attachHandlers(router, c.beforeGameStorage)

			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/match/leave", bytes.NewBuffer([]byte(c.inBody)))
			router.ServeHTTP(rec, req)

			if rec.Result().StatusCode != 200 {
				assert.Equal(t, c.outStatusCode, rec.Result().StatusCode)
				return
			}
		})
	}
}
