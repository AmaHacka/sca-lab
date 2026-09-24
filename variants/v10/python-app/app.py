"""Сервис аутентификации — учебный сервис ЛР №4 (вариант 10).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
import jwt
import rsa
from ecdsa import NIST256p, SigningKey
from flask import Flask, jsonify, request

app = Flask(__name__)


@app.route("/health")
def health():
    return jsonify(status="ok")


SECRET = "lab4-not-a-secret"


@app.route("/token")
def token():
    """Выдаёт JWT для тестового пользователя."""
    value = jwt.encode({"sub": "student"}, SECRET, algorithm="HS256")
    if isinstance(value, bytes):  # PyJWT < 2.0 возвращает bytes
        value = value.decode()
    return jsonify(token=value)


SIGNING_KEY = SigningKey.generate(curve=NIST256p)


@app.route("/sign", methods=["POST"])
def sign():
    """Подписывает тело запроса ключом ECDSA P-256."""
    return jsonify(signature=SIGNING_KEY.sign(request.get_data()).hex())


PUB_KEY, _PRIV_KEY = rsa.newkeys(512)


@app.route("/encrypt", methods=["POST"])
def encrypt():
    """Шифрует короткое сообщение открытым ключом сервиса."""
    return jsonify(data=rsa.encrypt(request.get_data()[:40], PUB_KEY).hex())


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
