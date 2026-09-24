# Шпаргалка: syft, grype, jq

## syft

```bash
syft <образ>                                   # таблица пакетов (образ из локального Docker)
syft registry:alpine:3.20                      # образ прямо из реестра, без docker pull
syft dir:./python-app                          # каталог с исходниками
syft file:./app                                # один файл (например, Go-бинарь)
syft <образ> --scope all-layers                # учитывать файлы из всех слоёв, включая удалённые
syft <образ> -o json                           # syft-json в stdout
syft <образ> -o syft-json=a.json -o cyclonedx-json=a.cdx.json -o spdx-json=a.spdx.json   # несколько форматов сразу
syft convert a.json -o cyclonedx-json          # конвертация уже построенного SBOM
syft cataloger list                            # какие каталогизаторы есть
syft <образ> --select-catalogers "+go-module-binary-cataloger"   # добавить/убрать каталогизаторы
```

Полезные поля syft-json: `.artifacts[]` (`name`, `version`, `type`, `purl`, `cpes`, `licenses`, `locations[].path`,
`locations[].layerID`), `.source.metadata` (`imageID`, `layers[]`), `.distro`.

## grype

```bash
grype <образ>                                  # скан образа
grype sbom:a.json                              # скан готового SBOM (syft-json, CycloneDX, SPDX)
grype dir:./python-app                         # скан каталога
grype pkg:npm/jquery@1.12.4                    # проверить один пакет по purl
grype ... -o json > r.json                     # JSON-отчёт (также: table, cyclonedx-json, sarif, template)
grype ... --only-fixed                         # только уязвимости с доступным исправлением
grype ... --fail-on high                       # код возврата 2, если есть находки >= High
grype ... --by-cve                             # показывать CVE вместо GHSA/GO-ID, где возможно
grype ... --sort-by epss                       # сортировка: risk (по умолчанию), severity, epss, kev, package
grype ... -c .grype.yaml                       # файл политики (правила ignore)
grype ... --show-suppressed                    # показать проигнорированные находки (table)
grype explain --id CVE-XXXX-YYYY < r.json      # почему найдено совпадение
grype db status | update | check               # состояние и обновление БД
grype db search --pkg log4j-core               # записи БД по пакету
grype db search --vuln GHSA-jfh8-c2jp-5v3q     # уязвимые диапазоны версий
grype db search vuln CVE-2021-44228            # записи об уязвимости у разных поставщиков
grype db providers                             # источники данных в БД
```

Поля JSON-отчёта: `.matches[]` →
`.vulnerability` (`id`, `severity`, `cvss[]`, `epss[]`, `knownExploited[]`, `cwes[]`, `fix.state`, `fix.versions`, `risk`, `urls`),
`.relatedVulnerabilities[]` (алиасы), `.artifact` (`name`, `version`, `type`, `purl`, `locations`),
`.matchDetails[]` (как найдено). Проигнорированные находки — в `.ignoredMatches[]`.

Формат `.grype.yaml`:

```yaml
ignore:
  - vulnerability: CVE-2026-00000        # что игнорировать (любая комбинация полей)
    fix-state: wont-fix
    package:
      name: perl-base
      type: deb
      location: "/usr/lib/**"
    reason: "обоснование — обязательно"
```

## jq: частые запросы

```bash
# пакеты по типам
jq -r '.artifacts[].type' a.json | sort | uniq -c

# находки по уровням / по типам пакетов / по пакетам
jq -r '.matches[].vulnerability.severity' r.json | sort | uniq -c | sort -rn
jq -r '.matches[].artifact.type' r.json | sort | uniq -c
jq -r '.matches[] | "\(.artifact.name)@\(.artifact.version)"' r.json | sort | uniq -c | sort -rn | head

# с исправлением / без
jq '[.matches[] | select(.vulnerability.fix.state=="fixed")] | length' r.json
jq -r '.matches[] | select(.vulnerability.fix.state!="fixed") | "\(.vulnerability.fix.state)\t\(.artifact.name)\t\(.vulnerability.id)"' r.json

# уязвимости из KEV
jq -r '.matches[] | select(.vulnerability.knownExploited | length > 0) | "\(.artifact.name)@\(.artifact.version) \(.vulnerability.id)"' r.json

# исправимые High/Critical (то, что не пропустит gate)
jq -r '.matches[] | select(.vulnerability.fix.state=="fixed" and (.vulnerability.severity=="High" or .vulnerability.severity=="Critical"))
       | "\(.artifact.name)@\(.artifact.version) -> \(.vulnerability.fix.versions|join(",")) \(.vulnerability.id)"' r.json

# минимальная версия, устраняющая все уязвимости пакета
jq -r '[.matches[] | select(.artifact.name=="stdlib") | .vulnerability.fix.versions[]] | unique | .[]' r.json | sort -V | tail -1

# уникальные уязвимости
jq '[.matches[].vulnerability.id] | unique | length' r.json
```

## Слой → инструкция Dockerfile

Размер слоя в SBOM syft совпадает с размером в `docker history` в байтах. По нему слой можно сопоставить
с инструкцией:

```bash
pkg=pyyaml; sbom=sbom/py-before.syft.json; img=lab4/v00-py:before
layer=$(jq -r --arg p "$pkg" '.artifacts[] | select(.name==$p) | .locations[0].layerID' "$sbom")
size=$(jq -r --arg l "$layer" '.source.metadata.layers[] | select(.digest==$l) | .size' "$sbom")
docker history --human=false --no-trunc --format '{{.Size}}\t{{.CreatedBy}}' "$img" | awk -F'\t' -v s="$size" '$1==s'
```
```
12043267	RUN /bin/sh -c pip install --no-cache-dir -r requirements.txt # buildkit
```

## Docker

```bash
docker build -t имя:тег каталог                 # сборка
docker history --no-trunc имя:тег               # инструкции и размеры слоёв
id=$(docker create имя:тег); docker cp "$id":/путь ./файл; docker rm "$id"   # достать файл из образа
docker run --rm -d -p 8080:8080 --name t имя:тег; curl -s localhost:8080/health; docker rm -f t
docker logs t                                   # почему контейнер упал
docker image ls lab4/*                          # размеры образов
docker save имя:тег -o image.tar                # образ со всеми слоями в tar
```
