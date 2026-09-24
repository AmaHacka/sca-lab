# Старый сканер файлов (Spring Boot-подобный толстый JAR): сборка архива
FROM alpine:3.17 AS legacy
RUN apk add --no-cache zip
WORKDIR /build
ADD https://repo1.maven.org/maven2/commons-fileupload/commons-fileupload/1.3.1/commons-fileupload-1.3.1.jar BOOT-INF/lib/
ADD https://repo1.maven.org/maven2/commons-io/commons-io/2.4/commons-io-2.4.jar BOOT-INF/lib/
ADD https://repo1.maven.org/maven2/commons-beanutils/commons-beanutils/1.9.3/commons-beanutils-1.9.3.jar BOOT-INF/lib/
RUN mkdir -p META-INF \
    && printf 'Manifest-Version: 1.0\nMain-Class: org.springframework.boot.loader.JarLauncher\n' > META-INF/MANIFEST.MF \
    && zip -r0 /build/file-scanner.jar META-INF BOOT-INF

FROM python:3.8.10-slim-buster
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Старый сканер файлов (Spring Boot-подобный толстый JAR)
COPY --from=legacy /build/file-scanner.jar /opt/file-scanner/file-scanner.jar

COPY app.py ./

EXPOSE 8080
CMD ["python", "app.py"]
