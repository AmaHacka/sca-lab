"""Сервис аватаров — учебный сервис ЛР №4 (вариант 03).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
import io

from flask import Flask, jsonify, request
from PIL import Image

app = Flask(__name__)


@app.route("/health")
def health():
    return jsonify(status="ok")


@app.route("/avatar", methods=["POST"])
def avatar():
    """Принимает изображение и делает из него миниатюру 128x128."""
    img = Image.open(io.BytesIO(request.get_data()))
    img.thumbnail((128, 128))
    return jsonify(format=img.format, size=list(img.size))


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
