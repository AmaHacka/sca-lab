# Сборка фронтенда
FROM node:14-alpine AS frontend
WORKDIR /frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci --production

FROM python:3.9.1-slim-buster
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Статика фронтенда
COPY --from=frontend /frontend/node_modules ./static/node_modules

COPY app.py ./
COPY vendor/ ./vendor/

EXPOSE 8080
CMD ["python", "app.py"]
