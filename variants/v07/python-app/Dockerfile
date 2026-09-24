# Сборка фронтенда
FROM node:14-alpine AS frontend
WORKDIR /frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci --production

FROM python:3.10-slim-bullseye
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Утилита для запуска процесса от непривилегированного пользователя
ADD https://github.com/tianon/gosu/releases/download/1.14/gosu-amd64 /usr/local/bin/gosu
RUN chmod +x /usr/local/bin/gosu

# Статика фронтенда
COPY --from=frontend /frontend/node_modules ./static/node_modules

COPY app.py ./

EXPOSE 8080
ENTRYPOINT ["gosu", "nobody"]
CMD ["python", "app.py"]
