package server

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"talpidae-backend/model/game"
	"talpidae-backend/storage"
	"testing"

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
			req := httptest.NewRequest(http.MethodGet, "/start", nil)
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
			outBody:       `{"field":[[3,0],[0,0]]}`,
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
			req := httptest.NewRequest(http.MethodGet, "/field", nil)
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
			req := httptest.NewRequest(http.MethodPost, "/fill", bytes.NewBuffer([]byte(c.inBody)))
			router.ServeHTTP(rec, req)

			assert.Equal(t, c.outStatusCode, rec.Result().StatusCode)
		})
	}
}
