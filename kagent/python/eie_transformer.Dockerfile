FROM python:3.8-slim

COPY third_party third_party

COPY kserve kserve

RUN apt-get update \
&& apt-get install -y --no-install-recommends git

RUN pip install --no-cache-dir --upgrade pip && pip install --no-cache-dir -e ./kserve

COPY transformer/eie_transformer eie_transformer
RUN pip install --no-cache-dir -e eie_transformer

ENTRYPOINT ["python", "-m", "eie_transformer"]
