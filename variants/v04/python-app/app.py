"""Прокси погоды — учебный сервис ЛР №4 (вариант 04).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
import requests
from flask import Flask, jsonify, request

app = Flask(__name__)


@app.route("/health")
def health():
    return jsonify(status="ok")


@app.route("/weather/<city>")
def weather(city):
    """Проксирует запрос к внешнему сервису погоды."""
    resp = requests.get("https://wttr.in/" + city, params={"format": "j1"}, timeout=5)
    return jsonify(status=resp.status_code, data=resp.json() if resp.ok else None)


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
