"""Сервис заметок — учебный сервис ЛР №4 (вариант 01).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
from flask import Flask, jsonify, request
from pygments import highlight
from pygments.formatters import HtmlFormatter
from pygments.lexers import guess_lexer

app = Flask(__name__)


@app.route("/health")
def health():
    return jsonify(status="ok")


@app.route("/highlight", methods=["POST"])
def highlight_snippet():
    """Подсвечивает синтаксис фрагмента кода из заметки."""
    code = request.get_data(as_text=True)
    return highlight(code, guess_lexer(code), HtmlFormatter())


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
