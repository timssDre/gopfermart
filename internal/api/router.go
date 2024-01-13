package api

import "github.com/gin-gonic/gin"

func (s *RestAPI) setRoutes(r *gin.Engine) {
	r.POST("/api/user/register", s.Registration) //— регистрация пользователя;
	//POST /api/users/login — аутентификация пользователя;
	//POST /api/users/orders — загрузка пользователем номера заказа для расчёта;
	//GET /api/users/orders — получение списка загруженных пользователем номеров заказов, статусов их обработки и информации о начислениях;
	//GET /api/users/balance — получение текущего баланса счёта баллов лояльности пользователя;
	//POST /api/users/balance/withdraw — запрос на списание баллов с накопительного счёта в счёт оплаты нового заказа;
	//GET /api/users/withdrawals — получение информации о выводе средств с накопительного счёта пользователем.
}
