"""SSH-оркестратор — учебный сервис ЛР №4 (вариант 06).

Код намеренно простой: предмет работы — зависимости, а не логика.
"""
import io

import paramiko
import yaml
from flask import Flask, jsonify, request

app = Flask(__name__)


@app.route("/health")
def health():
    return jsonify(status="ok")


with open("inventory.yaml") as f:
    INVENTORY = yaml.load(f)


@app.route("/hosts")
def hosts():
    """Возвращает список управляемых хостов из инвентаря."""
    return jsonify(hosts=INVENTORY.get("hosts", []))


@app.route("/keys/fingerprint", methods=["POST"])
def key_fingerprint():
    """Вычисляет отпечаток закрытого RSA-ключа для инвентаризации."""
    key = paramiko.RSAKey.from_private_key(io.StringIO(request.get_data(as_text=True)))
    return jsonify(bits=key.get_bits(), fingerprint=key.get_fingerprint().hex())


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)
