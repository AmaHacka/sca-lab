# Сборка фронтенда
FROM node:14-alpine AS frontend
WORKDIR /frontend
COPY frontend/package.json frontend/package-lock.json ./
RUN npm ci --production

FROM python:3.7-slim-buster
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Отладочные инструменты (временно, потом удаляем)
RUN pip install --no-cache-dir ipython==7.16.1
RUN pip uninstall -y ipython

# Статика фронтенда
COPY --from=frontend /frontend/node_modules ./static/node_modules

COPY app.py ./

EXPOSE 8080
CMD ["python", "app.py"]
