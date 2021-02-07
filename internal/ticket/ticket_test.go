package ticket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"nipatest/main/db"
	"nipatest/main/internal/model"
	"nipatest/main/internal/repository"
	"nipatest/main/router"
	"strings"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func dropTestDB() {
	if err := db.DropTestDB(); err != nil {
		log.Fatal(err)
	}
}

func newContainer() (*echo.Echo, *gorm.DB, *Ticket) {
	var e *echo.Echo
	e = router.New()
	v1 := e.Group("/api")
	d := db.TestDB()
	db.AutoMigrate(d)

	tr := repository.NewTicketRepo(d)
	loadFixtures(tr)
	tk := NewTicket(tr)

	tk.Register(v1)
	return e, d, tk
}

func loadFixtures(tr *repository.TicketRepo) {
	t1 := model.Ticket{
		Name:        "test1",
		Description: "test1",
		ContactInfo: "test1",
		Status:      "PENDING",
		CreatedAt:   time.Date(2021, 2, 7, 14, 6, 44, 0, time.Local),
	}
	if _, err := tr.Create(&t1); err != nil {
		panic(err)
	}
	t2 := model.Ticket{
		Name:        "test2",
		Description: "test2",
		ContactInfo: "test2",
		Status:      "REJECTED",
		CreatedAt:   time.Date(2021, 2, 7, 14, 7, 44, 0, time.Local),
	}
	if _, err := tr.Create(&t2); err != nil {
		panic(err)
	}
	t3 := model.Ticket{
		Name:        "test3",
		Description: "test3",
		ContactInfo: "test3",
		Status:      "PENDING",
		CreatedAt:   time.Date(2021, 2, 7, 14, 8, 44, 0, time.Local),
	}
	if _, err := tr.Create(&t3); err != nil {
		panic(err)
	}
}

func TestGetTicketCase(t *testing.T) {
	e, d, tkc := newContainer()
	defer dropTestDB()
	defer d.Close()

	t.Run("unable to get ticket wrong filter by status", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/api/ticket?status=MOCK", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Get(c))
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("able to get ticket filter by status", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/api/ticket?status=REJECTED", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Get(c))
		if assert.Equal(t, http.StatusOK, rec.Code) {
			var dat []model.Ticket
			if err := json.Unmarshal(rec.Body.Bytes(), &dat); err != nil {
				panic(err)
			}
			assert.Len(t, dat, 1)
			assert.Equal(t, "test2", dat[0].Name)
			assert.Equal(t, model.TicketStatusRejected, dat[0].Status)
		}
	})

	t.Run("unable to get ticket wrong filter by range time", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/api/ticket?from=MOCK&to=MOCK", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Get(c))
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("able to get ticket filter by range time", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, fmt.Sprintf("/api/ticket?from=%s&to=%s", "1612681604", "1612681666"), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Get(c))
		if assert.Equal(t, http.StatusOK, rec.Code) {
			var dat []model.Ticket
			if err := json.Unmarshal(rec.Body.Bytes(), &dat); err != nil {
				panic(err)
			}
			assert.Len(t, dat, 2)
			assert.Equal(t, "test1", dat[0].Name)
			assert.Equal(t, "test2", dat[1].Name)
		}
	})

	t.Run("able to get ticket filter by range time from only", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, fmt.Sprintf("/api/ticket?from=%s", "1612681604"), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Get(c))
		if assert.Equal(t, http.StatusOK, rec.Code) {
			var dat []model.Ticket
			if err := json.Unmarshal(rec.Body.Bytes(), &dat); err != nil {
				panic(err)
			}
			assert.Len(t, dat, 3)
		}
	})

	t.Run("able to get ticket filter by range time", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, fmt.Sprintf("/api/ticket?to=%s", "1612681666"), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Get(c))
		if assert.Equal(t, http.StatusOK, rec.Code) {
			var dat []model.Ticket
			if err := json.Unmarshal(rec.Body.Bytes(), &dat); err != nil {
				panic(err)
			}
			assert.Len(t, dat, 2)
		}
	})

	t.Run("unable to get ticket wrong pagination", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, "/api/ticket?page=MOCK&page_size=MOCK", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Get(c))
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})

	t.Run("able to get ticket pagination", func(t *testing.T) {
		req := httptest.NewRequest(echo.GET, fmt.Sprintf("/api/ticket?page=%d&page_size=%d", 1, 2), nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Get(c))
		if assert.Equal(t, http.StatusOK, rec.Code) {
			var dat []model.Ticket
			if err := json.Unmarshal(rec.Body.Bytes(), &dat); err != nil {
				panic(err)
			}
			assert.Len(t, dat, 2)
			assert.Equal(t, "test1", dat[0].Name)
			assert.Equal(t, "test2", dat[1].Name)
		}
	})
}

func TestPostTicketCase(t *testing.T) {
	e, d, tkc := newContainer()
	defer dropTestDB()
	defer d.Close()
	var (
		reqJSON = `{"name":"testpost","desc":"testpost","contactInfo":"testpost"}`
	)

	t.Run("able to post ticket", func(t *testing.T) {
		req := httptest.NewRequest(echo.POST, "/api/ticket", strings.NewReader(reqJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Create(c))
		if assert.Equal(t, http.StatusCreated, rec.Code) {
			var dat model.Ticket
			if err := json.Unmarshal(rec.Body.Bytes(), &dat); err != nil {
				panic(err)
			}
			assert.Equal(t, "testpost", dat.Name)
			assert.Equal(t, "testpost", dat.Description)
			assert.Equal(t, "testpost", dat.ContactInfo)
		}
	})

	t.Run("unable to post ticket", func(t *testing.T) {
		req := httptest.NewRequest(echo.POST, "/api/ticket", strings.NewReader(`{}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Create(c))
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})
}

func TestPutTicketCase(t *testing.T) {
	e, d, tkc := newContainer()
	defer dropTestDB()
	defer d.Close()
	var (
		reqJSON = `{"status":"ACCEPTED"}`
	)

	t.Run("able to update ticket", func(t *testing.T) {
		req := httptest.NewRequest(echo.PUT, "/api/ticket/:id", strings.NewReader(reqJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/api/ticket/:id")
		c.SetParamNames("id")
		c.SetParamValues("1")
		assert.NoError(t, tkc.Update(c))
		if assert.Equal(t, http.StatusOK, rec.Code) {
			assert.Contains(t, rec.Body.String(), "Update ticket successfully")
		}
	})

	t.Run("unable to update ticket", func(t *testing.T) {
		req := httptest.NewRequest(echo.POST, "/api/ticket", strings.NewReader(`{}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		assert.NoError(t, tkc.Update(c))
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	})
}
