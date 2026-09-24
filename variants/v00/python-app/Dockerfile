FROM python:3.8-slim-buster
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Библиотеки модуля отчётов (перенесены со старого сервера)
ADD https://repo1.maven.org/maven2/org/apache/logging/log4j/log4j-core/2.14.1/log4j-core-2.14.1.jar /opt/reporting/lib/
ADD https://repo1.maven.org/maven2/org/apache/logging/log4j/log4j-api/2.14.1/log4j-api-2.14.1.jar /opt/reporting/lib/

COPY app.py ./

EXPOSE 8080
CMD ["python", "app.py"]
