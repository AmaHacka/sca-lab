"""Каталог книг — учебный сервис ЛР №4 (вариант 09).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
import requests
import yaml
from flask import Flask, jsonify, request

app = Flask(__name__)


@app.route("/health")
def health():
    return jsonify(status="ok")


@app.route("/config", methods=["POST"])
def parse_config():
    """Разбирает YAML-конфигурацию из тела запроса."""
    data = yaml.safe_load(request.get_data()) or {}
    return jsonify(keys=sorted(map(str, data)))


@app.route("/feed")
def feed():
    """Загружает RSS-ленту по адресу из параметра url."""
    resp = requests.get(request.args.get("url", ""), timeout=5)
    return jsonify(status=resp.status_code, length=len(resp.content))


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
