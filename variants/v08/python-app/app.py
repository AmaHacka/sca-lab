"""Генератор отчётов — учебный сервис ЛР №4 (вариант 08).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
from flask import Flask, jsonify, request
from lxml import etree
from mako.template import Template

app = Flask(__name__)


@app.route("/health")
def health():
    return jsonify(status="ok")


REPORT = Template("Отчёт: ${count} элементов, корневой элемент <${root}>")


@app.route("/report", methods=["POST"])
def report():
    """Строит текстовый отчёт по XML-документу."""
    doc = etree.fromstring(request.get_data())
    return REPORT.render(count=len(doc.xpath("//*")), root=doc.tag)


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
