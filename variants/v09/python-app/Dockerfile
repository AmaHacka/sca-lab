# Сборка фронтенда
FROM node:14-alpine AS frontend
WORKDIR /frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci --production

FROM python:3.8-alpine3.12
WORKDIR /app

# Компилятор нужен для сборки C-расширений
RUN apk add --no-cache gcc musl-dev

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Уменьшаем образ: компилятор больше не нужен
RUN apk del gcc musl-dev

# Статика фронтенда
COPY --from=frontend /frontend/node_modules ./static/node_modules

COPY app.py ./

EXPOSE 8080
CMD ["python", "app.py"]
