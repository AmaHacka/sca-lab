FROM python:3.9.7-slim-bullseye
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Библиотеки модуля аналитики переходов
ADD https://repo1.maven.org/maven2/log4j/log4j/1.2.17/log4j-1.2.17.jar /opt/analytics/lib/
ADD https://repo1.maven.org/maven2/commons-collections/commons-collections/3.2.1/commons-collections-3.2.1.jar /opt/analytics/lib/

COPY app.py ./

EXPOSE 8080
CMD ["python", "app.py"]
