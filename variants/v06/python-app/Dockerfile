FROM python:3.7-slim-bullseye
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Библиотеки модуля инвентаризации
ADD https://repo1.maven.org/maven2/com/fasterxml/jackson/core/jackson-databind/2.9.8/jackson-databind-2.9.8.jar /opt/inventory/lib/
ADD https://repo1.maven.org/maven2/com/thoughtworks/xstream/xstream/1.4.17/xstream-1.4.17.jar /opt/inventory/lib/

COPY app.py inventory.yaml ./

EXPOSE 8080
CMD ["python", "app.py"]
