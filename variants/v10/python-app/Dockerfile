FROM python:3.9-slim-bullseye
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Продуктивный сервер
RUN pip install --no-cache-dir gunicorn==19.9.0

# Библиотеки модуля аудита
ADD https://repo1.maven.org/maven2/com/h2database/h2/1.4.197/h2-1.4.197.jar /opt/audit/lib/
ADD https://repo1.maven.org/maven2/org/postgresql/postgresql/42.2.20/postgresql-42.2.20.jar /opt/audit/lib/
ADD https://repo1.maven.org/maven2/ch/qos/logback/logback-classic/1.2.3/logback-classic-1.2.3.jar /opt/audit/lib/
ADD https://repo1.maven.org/maven2/ch/qos/logback/logback-core/1.2.3/logback-core-1.2.3.jar /opt/audit/lib/

COPY app.py ./

EXPOSE 8080
CMD ["gunicorn", "-b", "0.0.0.0:8080", "app:app"]
