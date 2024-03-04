package terminals

import (
	"math/rand"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Terminal struct {
	Terminal string `gorm:"not null" json:"terminal"`
	Addr     string `gorm:"not null" json:"addr"`
}

type Terminals struct {
	repo Repository
}

type Repository interface {
	Get(terminal string) (addr string, ok bool)
	Set(params SetParams) (response SetResponse, err error)
	RandomInsert(n int) error
}

func New(repo Repository) *Terminals {
	return &Terminals{
		repo: repo,
	}
}

type GetParams struct {
	Terminal string `json:"terminal"`
}

type GetResponse struct {
	Addr string `json:"addr"`
}

func (t *Terminals) Get(c echo.Context) error {
	params := GetParams{}
	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	addr, ok := t.repo.Get(params.Terminal)
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "Not found")
	}

	return c.JSON(http.StatusCreated, GetResponse{
		Addr: addr,
	})
}

type SetParams struct {
	Terminal string `json:"terminal"`
	Addr     string `json:"addr"`
}

type SetResponse struct {
	Terminal string `json:"terminal"`
	Addr     string `json:"addr"`
}

func (t *Terminals) Set(c echo.Context) error {
	params := SetParams{}
	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if params.Terminal == "" || params.Addr == "" {
		params.Terminal = randStringBytes(10)
		params.Addr = randStringBytes(10)
	}

	response, err := t.repo.Set(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	return c.JSON(http.StatusCreated, response)
}

type RandomInsertParams struct {
	N int `json:"n"`
}

func (t *Terminals) RandomInsert(c echo.Context) error {
	params := RandomInsertParams{}
	if err := c.Bind(&params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	err := t.repo.RandomInsert(params.N)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal error")
	}

	return c.NoContent(http.StatusNoContent)
}

func (t *Terminals) AddRoutes(g *echo.Group) {
	g.POST("/get", t.Get)
	g.POST("/set", t.Set)
	g.POST("/random-insert", t.RandomInsert)
}

var letterBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
