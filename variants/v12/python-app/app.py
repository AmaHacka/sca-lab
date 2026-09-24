"""Файловый обменник — учебный сервис ЛР №4 (вариант 12).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
import os

from flask import Flask, jsonify, request
from werkzeug.utils import secure_filename

app = Flask(__name__)


@app.route("/health")
def health():
    return jsonify(status="ok")


UPLOAD_DIR = "/tmp/uploads"


@app.route("/upload", methods=["POST"])
def upload():
    """Сохраняет загруженный файл во временный каталог."""
    os.makedirs(UPLOAD_DIR, exist_ok=True)
    f = request.files["file"]
    name = secure_filename(f.filename)
    f.save(os.path.join(UPLOAD_DIR, name))
    return jsonify(saved=name)


if __name__ == "__main__":
    from waitress import serve

    serve(app, host="0.0.0.0", port=8080)
