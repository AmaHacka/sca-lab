"""Конвертер конфигураций — учебный сервис ЛР №4 (вариант 00).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
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


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
