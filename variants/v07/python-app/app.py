"""Чат-бот — учебный сервис ЛР №4 (вариант 07).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
import jwt
from aiohttp import web

SECRET = "lab4-not-a-secret"


async def health(request):
    return web.json_response({"status": "ok"})


async def token(request):
    """Выдаёт JWT для подключения клиента к чату."""
    value = jwt.encode({"sub": request.query.get("user", "guest")}, SECRET, algorithm="HS256")
    if isinstance(value, bytes):  # PyJWT < 2.0 возвращает bytes
        value = value.decode()
    return web.json_response({"token": value})


async def message(request):
    """Отвечает на сообщение пользователя."""
    data = await request.json()
    return web.json_response({"reply": "Эхо: %s" % data.get("text", "")})


app = web.Application()
app.add_routes([
    web.get("/health", health),
    web.get("/token", token),
    web.post("/message", message),
])

if __name__ == "__main__":
    web.run_app(app, host="0.0.0.0", port=8080)
