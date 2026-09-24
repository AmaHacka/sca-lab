"""Агрегатор RSS — учебный сервис ЛР №4 (вариант 11).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
import bleach
import requests
from flask import Flask, jsonify, request
from vendor import markdown2

app = Flask(__name__)


@app.route("/health")
def health():
    return jsonify(status="ok")


@app.route("/feed")
def feed():
    """Загружает RSS-ленту по адресу из параметра url."""
    resp = requests.get(request.args.get("url", ""), timeout=5)
    return jsonify(status=resp.status_code, length=len(resp.content))


@app.route("/sanitize", methods=["POST"])
def sanitize():
    """Очищает HTML из ленты от опасных тегов."""
    return bleach.clean(request.get_data(as_text=True))


@app.route("/render", methods=["POST"])
def render():
    """Преобразует Markdown-описание ленты в HTML."""
    return markdown2.markdown(request.get_data(as_text=True))


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
