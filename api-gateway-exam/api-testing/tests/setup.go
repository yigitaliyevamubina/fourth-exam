package tests

import (
	"bytes"
	"exam/api-gateway/api-testing/storage/kv"
	"io"
	"net/http"
	"net/http/httptest"
	"os"

	_ "github.com/lib/pq" //postgres drivers
	// "github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const testHost = "http://localhost:9191"

func SetupMinimumInstance(path string) error {
	_ = path
	conf := viper.New()
	conf.Set("mode", "test")

	// client := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })
	// kv.Init(kv.NewRedisClient(client))

	// psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	"localhost",
	// 	5432,
	// 	"postgres",
	// 	"mubina2007",
	// 	"postgres",
	// )

	// db, err := sql.Open("postgres", psqlString)
	// if err != nil {
	// 	return err
	// }

	// kv.Init(kv.NewPostgres(db))
	kv.Init(kv.NewInMemoryInst())

	return nil
}

func Serve(handler func(c *gin.Context), method, uri string, body []byte) (*httptest.ResponseRecorder, error) {
	r := gin.New()

	gin.SetMode(gin.TestMode)

	switch method {
	case http.MethodPost:
		r.POST(uri, handler)
	case http.MethodGet:
		r.GET(uri, handler)
	case http.MethodDelete:
		r.DELETE(uri, handler)
	case http.MethodPatch:
		r.PATCH(uri, handler)
	}

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	recorder := httptest.NewRecorder()

	r.ServeHTTP(recorder, req)

	return recorder, nil
}

func NewResponse() *http.Response {
	return &http.Response{}
}

func NewRequest(method string, uri string, body []byte) *http.Request {
	req, err := http.NewRequest(method, testHost+uri, nil)
	if err != nil {
		return nil
	}

	req.Header.Set("Host", "localhost")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	req.Header.Set("X-Forwarded-For", "79.104.42.249")

	if body != nil {
		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	return req
}

func OpenFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}
