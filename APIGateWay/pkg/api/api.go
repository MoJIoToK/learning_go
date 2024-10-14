package api

import (
	"APIGateWay/pkg/config"
	"APIGateWay/pkg/logger"
	"APIGateWay/pkg/middleware"
	"APIGateWay/pkg/model"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

type API struct {
	r     *mux.Router
	proxy map[string]string
	cl    *http.Client
}

const (
	news     = "news"
	comments = "comments"
	censor   = "censor"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrBadRequset = errors.New("bad request")
)

func New(cfg *config.Config) *API {
	p := make(map[string]string)
	p[news] = cfg.News
	p[comments] = cfg.Comments
	p[censor] = cfg.Censor

	api := API{
		r:     mux.NewRouter(),
		proxy: p,
		cl:    &http.Client{},
	}
	api.endpoints()
	return &api
}

func (api *API) Router() *mux.Router {
	return api.r
}

func (api *API) endpoints() {
	api.r.Use(middleware.RequestID)
	api.r.Use(middleware.Logger)
	api.r.Use(middleware.RealIP)
	api.r.HandleFunc("/news", api.News).Methods(http.MethodGet, http.MethodOptions)
	api.r.HandleFunc("/news/id/{id}", api.Detailed).Methods(http.MethodGet, http.MethodOptions)
	api.r.HandleFunc("/news/comments/new", api.AddComment).Methods(http.MethodPost, http.MethodOptions)
}

func (api *API) News(w http.ResponseWriter, r *http.Request) {
	const operation = "APIGateWay.api.News"

	log := slog.Default().With(
		slog.String("op", operation),
		slog.String("request_id", middleware.GetRequestID(r.Context())),
	)

	log.Info("Request to receive posts")

	resp, err := request(api.proxy[news], r, api.cl)
	if err != nil {
		log.Error("Failed to receive posts", logger.Err(err))
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	log.Debug("News received successfully")

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Error("Failed to copy response body", logger.Err(err))
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	log.Info("request served successfully")
}

func (api *API) Detailed(w http.ResponseWriter, r *http.Request) {
	const operation = "APIGateWay.api.Detailed"

	log := slog.Default().With(
		slog.String("op", operation),
		slog.String("request_id", middleware.GetRequestID(r.Context())),
	)

	log.Info("Request to receive post by ID with comments")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Объявим функцию, которая вызывает request и декодирует
	// ответ в структуру.
	fn := func(host string, req *http.Request, data any) error {
		resp, err := request(host, req, api.cl)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			return ErrNotFound
		}
		if resp.StatusCode == http.StatusBadRequest {
			return ErrBadRequset
		}

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("request status = %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("readAll err = %w", err)
		}

		err = json.Unmarshal(body, data)
		if err != nil {
			return fmt.Errorf("unmarshal err = %w", err)
		}
		return nil
	}

	var post model.NewsShortDetailed
	var comment []model.FullComment
	var errProxy error
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	rNews := r.Clone(r.Context())
	rComm := r.Clone(ctx)

	// Сделаем запросы в сервис новостей и сервис комментариев
	// в отдельных горутинах.
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := fn(api.proxy[news], rNews, &post)
		// Если при получении новости возникла ошибка, то сохраняем ее
		// для дальнейшей обработки, затем прерываем запрос комментариев.
		if err != nil {
			log.Error("Failed to receive post", logger.Err(err))
			errProxy = err
			cancel()
			return
		}
		log.Debug("Post received successfully")
	}()

	go func() {
		defer wg.Done()
		uri := rComm.URL.Path
		uri = strings.ReplaceAll(uri, "post/id", "comments")
		rComm.URL.Path = uri
		err := fn(api.proxy[comments], rComm, &comment)
		// Если при получении комментариев возникла ошибка, то не обрабатываем
		// ее как в горутине получения новости, так как ошибка получения новости
		// критична, а получения комментариев - нет.
		if err != nil {
			log.Error("Failed to receive comments", logger.Err(err))
			return
		}
		log.Debug("Comments received successfully")
	}()

	wg.Wait()

	if errProxy != nil {
		log.Error("failed to find post by ID", logger.Err(errProxy))
		if errors.Is(errProxy, ErrNotFound) {
			http.Error(w, "post not found", http.StatusNotFound)
			return
		}
		if errors.Is(errProxy, ErrBadRequset) {
			http.Error(w, "incorrect post id", http.StatusBadRequest)
			return
		}
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}

	fullPost := model.NewsFullDetailed{
		News:     post,
		Comments: comment,
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "\t")
	err := enc.Encode(fullPost)
	if err != nil {
		log.Error("Failed to encode post and comments", logger.Err(err))
		http.Error(w, "Failed to encode post and comments", http.StatusInternalServerError)
		return
	}
	log.Info("Request served successfully")
}

func (api *API) AddComment(w http.ResponseWriter, r *http.Request) {
	const operation = "APIGateWay.api.AddComment"

	log := slog.Default().With(
		slog.String("op", operation),
		slog.String("request_id", middleware.GetRequestID(r.Context())),
	)

	log.Info("Request to add new comment")

	// Создаем копии тела запроса.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Failed to read request body", logger.Err(err))
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	rc1 := io.NopCloser(bytes.NewBuffer(body))
	rc2 := io.NopCloser(bytes.NewBuffer(body))

	// Клонируем запрос с немодифицированнымм телом и меняем
	// путь запроса.
	rCens := r.Clone(r.Context())
	rCens.Body = rc1
	rCens.URL.Path = ""

	log.Debug("Checking new comment")

	respCensor, err := request(api.proxy[censor], rCens, api.cl)
	if err != nil {
		log.Error("Failed to check comment", logger.Err(err))
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer respCensor.Body.Close()
	io.Copy(io.Discard, respCensor.Body)

	if respCensor.StatusCode != http.StatusOK {
		log.Error("Comment contains inappropriate words", slog.Int("code", respCensor.StatusCode))
		http.Error(w, "Comment contains inappropriate words", http.StatusBadRequest)
		return
	}
	log.Debug("Comment checked successfully")

	// Клонируем запрос с немодифицированнымм телом.
	rComm := r.Clone(r.Context())
	rComm.Body = rc2

	respComm, err := request(api.proxy[comments], rComm, api.cl)
	if err != nil {
		log.Error("failed to add new comment", logger.Err(err))
		http.Error(w, "service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer respComm.Body.Close()

	// Полностью копируем ответ сервиса комментариев в ResponseWriter.
	copyHeader(w.Header(), respComm.Header)
	w.WriteHeader(respComm.StatusCode)
	_, err = io.Copy(w, respComm.Body)
	if err != nil {
		log.Error("failed to copy response body", logger.Err(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	log.Info("request served successfully")
}
