# Сборка фронтенда
FROM node:14-alpine AS frontend
WORKDIR /frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci --production

FROM python:3.10.4-slim-bullseye
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Продуктивный сервер
RUN pip install --no-cache-dir gunicorn==20.1.0

# Статика фронтенда
COPY --from=frontend /frontend/node_modules ./static/node_modules

COPY app.py ./

EXPOSE 8080
CMD ["gunicorn", "-b", "0.0.0.0:8080", "app:application"]
