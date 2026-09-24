FROM python:3.8-slim-bullseye
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Библиотеки старого модуля кэширования
ADD https://repo1.maven.org/maven2/org/springframework/spring-beans/5.3.17/spring-beans-5.3.17.jar /opt/legacy-cache/lib/
ADD https://repo1.maven.org/maven2/org/springframework/spring-core/5.3.17/spring-core-5.3.17.jar /opt/legacy-cache/lib/
ADD https://repo1.maven.org/maven2/org/springframework/spring-jcl/5.3.17/spring-jcl-5.3.17.jar /opt/legacy-cache/lib/

COPY app.py ./

EXPOSE 8080
CMD ["python", "app.py"]
