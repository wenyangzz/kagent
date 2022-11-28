FROM python:3.9-slim

COPY third_party third_party

COPY kserve kserve
RUN pip install --no-cache-dir --upgrade pip && pip install --no-cache-dir -e ./kserve

RUN apt-get update && apt-get install -y --no-install-recommends \
    libgomp1 && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY model_server/sklearnserver sklearnserver
RUN sed -i 's/scikit-learn == 1.1.2/scikit-learn == 1.0.2/g' sklearnserver/setup.py
RUN pip install --no-cache-dir -e ./sklearnserver

RUN useradd kserve -m -u 1000 -d /home/kserve
USER 1000
ENTRYPOINT ["python", "-m", "sklearnserver"]
