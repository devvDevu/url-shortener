package url_usecase

import (
	"context"
	"testing"
	"url-shortener/internal/common/types/error_with_codes"
	"url-shortener/internal/common/types/url_types"
	"url-shortener/internal/model/url_model"
	url_service_mock "url-shortener/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetUrlByCode(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем моки для зависимостей
	mockUrlService := url_service_mock.NewMockserviceI(ctrl)

	// Создаем экземпляр usecase с моками
	usecase := NewUrlUsecase(mockUrlService)

	ctx := context.Background()

	t.Run("Success url found", func(t *testing.T) {
		// Подготовка тестовых данных
		urlId := url_types.UrlId(1)
		urlCode := url_types.UrlCode("test")
		urlOriginal := url_types.UrlOriginal("https://www.google.com")
		expectedItem := &url_model.Url{
			Id:       urlId,
			Code:     urlCode,
			Original: urlOriginal,
		}

		// Ожидаемые вызовы
		mockUrlService.EXPECT().
			GetUrlByCode(ctx, urlCode).
			Return(expectedItem, true, nil)

		// Вызов тестируемого метода
		respOriginalUrl, err := usecase.GetUrlByCode(ctx, urlCode)

		// Проверки
		assert.NoError(t, err)
		assert.Equal(t, urlOriginal, *respOriginalUrl)
	})

	t.Run("Error url not found", func(t *testing.T) {
		urlCode := url_types.UrlCode("12345")

		// Ожидаемые вызовы
		mockUrlService.EXPECT().
			GetUrlByCode(ctx, urlCode).
			Return(nil, false, error_with_codes.ErrorFailedToGetUrlByCode)

		// Вызов тестируемого метода
		_, err := usecase.GetUrlByCode(ctx, urlCode)

		// Проверки
		assert.Error(t, err)
		assert.Equal(t, error_with_codes.ErrorFailedToGetUrlByCode, err)
	})

}

func TestGetUrlList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем моки для зависимостей
	mockUrlService := url_service_mock.NewMockserviceI(ctrl)

	// Создаем экземпляр usecase с моками
	usecase := NewUrlUsecase(mockUrlService)

	ctx := context.Background()

	t.Run("Success url list found", func(t *testing.T) {
		// Подготовка тестовых данных
		urlId := url_types.UrlId(1)
		urlCode := url_types.UrlCode("test")
		urlOriginal := url_types.UrlOriginal("https://www.google.com")
		expectedItem := []*url_model.Url{
			{
				Id:       urlId,
				Code:     urlCode,
				Original: urlOriginal,
			},
		}

		// Ожидаемые вызовы
		mockUrlService.EXPECT().
			GetUrlList(ctx).
			Return(expectedItem, true, nil)

		// Вызов тестируемого метода
		respMap, err := usecase.GetUrlList(ctx)

		// Проверки
		assert.NoError(t, err)
		assert.Equal(t, urlOriginal, respMap[urlId].Original)
		assert.Equal(t, urlCode, respMap[urlId].Code)
	})

	t.Run("Error url list not found", func(t *testing.T) {

		// Ожидаемые вызовы
		mockUrlService.EXPECT().
			GetUrlList(ctx).
			Return(nil, false, error_with_codes.ErrorFailedToGetUrlList)

		// Вызов тестируемого метода
		_, err := usecase.GetUrlList(ctx)

		// Проверки
		assert.Error(t, err)
		assert.Equal(t, error_with_codes.ErrorFailedToGetUrlList, err)
	})

	t.Run("Empty url list", func(t *testing.T) {
		// Ожидаемые вызовы
		mockUrlService.EXPECT().
			GetUrlList(ctx).
			Return(nil, true, nil)

		// Вызов тестируемого метода
		respOriginalUrl, err := usecase.GetUrlList(ctx)

		// Проверки
		assert.NoError(t, err)
		assert.Equal(t, 0, len(respOriginalUrl))
	})

}
func TestMakeShortUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем моки для зависимостей
	mockUrlService := url_service_mock.NewMockserviceI(ctrl)

	// Создаем экземпляр usecase с моками
	usecase := NewUrlUsecase(mockUrlService)

	ctx := context.Background()

	t.Run("Success make short url", func(t *testing.T) {
		urlOriginal := url_types.UrlOriginal("https://www.google.com")

		// Ожидаемые вызовы
		mockUrlService.EXPECT().
			CreateUrl(ctx, gomock.Any()).
			Do(func(ctx context.Context, url *url_model.Url) {
				assert.Equal(t, urlOriginal, url.Original)
				assert.Len(t, url.Code, codeLength)
			}).
			Return(nil)

		// Вызов тестируемого метода
		ok, err := usecase.MakeShortUrl(ctx, urlOriginal)

		// Проверки
		assert.NoError(t, err)
		assert.True(t, ok)
	})

	t.Run("Error failed to create url", func(t *testing.T) {
		urlOriginal := url_types.UrlOriginal("https://www.google.com")

		// Ожидаемые вызовы
		mockUrlService.EXPECT().
			CreateUrl(ctx, gomock.Any()).
			Do(func(ctx context.Context, url *url_model.Url) {
				assert.Equal(t, urlOriginal, url.Original)
				assert.Len(t, url.Code, codeLength)
			}).
			Return(error_with_codes.ErrorFailedToCreateUrl)

		// Вызов тестируемого метода
		_, err := usecase.MakeShortUrl(ctx, urlOriginal)

		// Проверки
		assert.Error(t, err)
		assert.Equal(t, error_with_codes.ErrorFailedToCreateUrl, err)
	})
}
