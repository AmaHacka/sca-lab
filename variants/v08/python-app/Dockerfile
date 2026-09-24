FROM python:3.10.12-slim-bookworm
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Библиотеки модуля экспорта в PDF
ADD https://repo1.maven.org/maven2/org/apache/commons/commons-text/1.9/commons-text-1.9.jar /opt/pdf-export/lib/
ADD https://repo1.maven.org/maven2/org/apache/commons/commons-lang3/3.11/commons-lang3-3.11.jar /opt/pdf-export/lib/
ADD https://repo1.maven.org/maven2/commons-io/commons-io/2.4/commons-io-2.4.jar /opt/pdf-export/lib/

COPY app.py ./

EXPOSE 8080
CMD ["python", "app.py"]
